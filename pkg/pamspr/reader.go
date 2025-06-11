package pamspr

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Reader reads PAM SPR files
type Reader struct {
	scanner      *bufio.Scanner
	validator    *Validator
	lineNum      int
	errors       []error
	nextLine     *string // For line pushback
	fileParser   *FileParser
	achParser    *ACHParser
	checkParser  *CheckParser
	commonParser *CommonParser
}

// NewReader creates a new PAM SPR reader
func NewReader(r io.Reader) *Reader {
	validator := NewValidator()
	return &Reader{
		scanner:      bufio.NewScanner(r),
		validator:    validator,
		errors:       make([]error, 0),
		nextLine:     nil,
		fileParser:   NewFileParser(validator),
		achParser:    NewACHParser(validator),
		checkParser:  NewCheckParser(validator),
		commonParser: NewCommonParser(validator),
	}
}

// scanLine reads the next line, handling pushback
func (r *Reader) scanLine() (string, bool) {
	if r.nextLine != nil {
		line := *r.nextLine
		r.nextLine = nil
		r.lineNum++
		return line, true
	}

	if r.scanner.Scan() {
		r.lineNum++
		return r.scanner.Text(), true
	}

	return "", false
}

// pushBackLine pushes a line back to be read again
func (r *Reader) pushBackLine(line string) {
	r.nextLine = &line
	r.lineNum-- // Decrement since it will be incremented again
}

// Read reads a complete PAM SPR file
func (r *Reader) Read() (*File, error) {
	file := &File{
		Schedules: make([]Schedule, 0),
	}

	// Read file header
	line, ok := r.scanLine()
	if !ok {
		return nil, fmt.Errorf("missing file header")
	}

	header, err := r.fileParser.ParseFileHeader(line)
	if err != nil {
		return nil, fmt.Errorf("line %d: %w", r.lineNum, err)
	}
	file.Header = header

	// Read schedules
	for {
		line, ok := r.scanLine()
		if !ok {
			break
		}

		// Check for file trailer
		if strings.HasPrefix(line, "E ") {
			trailer, err := r.fileParser.ParseFileTrailer(line)
			if err != nil {
				return nil, fmt.Errorf("line %d: %w", r.lineNum, err)
			}
			file.Trailer = trailer
			break
		}

		// Parse schedule
		schedule, err := r.parseSchedule(line)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", r.lineNum, err)
		}
		if schedule != nil {
			file.Schedules = append(file.Schedules, schedule)
		}
	}

	if err := r.scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	// Validate file structure
	if err := r.validator.ValidateFileStructure(file); err != nil {
		return nil, fmt.Errorf("file validation: %w", err)
	}

	// Validate balancing
	if err := r.validator.ValidateBalancing(file); err != nil {
		return nil, fmt.Errorf("balancing validation: %w", err)
	}

	return file, nil
}

// parseSchedule parses a schedule starting with the given line
func (r *Reader) parseSchedule(firstLine string) (Schedule, error) {
	recordCode := r.extractField(firstLine, 1, 2)

	switch recordCode {
	case "01":
		return r.parseACHSchedule(firstLine)
	case "11":
		return r.parseCheckSchedule(firstLine)
	default:
		return nil, fmt.Errorf("unexpected record code: %s", recordCode)
	}
}

// parseACHSchedule parses an ACH schedule
func (r *Reader) parseACHSchedule(headerLine string) (*ACHSchedule, error) {
	// Parse header using ACH parser
	header, err := r.achParser.ParseACHScheduleHeader(headerLine)
	if err != nil {
		return nil, err
	}

	schedule := &ACHSchedule{
		Header: header,
		BaseSchedule: BaseSchedule{
			ScheduleNumber: strings.TrimSpace(header.ScheduleNumber),
			PaymentType:    strings.TrimSpace(header.PaymentTypeCode),
			ALC:            strings.TrimSpace(header.AgencyLocationCode),
			Payments:       make([]Payment, 0),
		},
	}

	// Read payments until schedule trailer
	var currentPayment *ACHPayment

	for {
		line, ok := r.scanLine()
		if !ok {
			return nil, fmt.Errorf("unexpected end of file in ACH schedule")
		}

		recordCode := r.extractField(line, 1, 2)

		switch recordCode {
		case "02": // ACH Payment
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			payment, err := r.achParser.ParseACHPayment(line)
			if err != nil {
				return nil, err
			}
			payment.StandardEntryClassCode = header.StandardEntryClassCode
			currentPayment = payment

		case "03", "04": // ACH Addendum
			if currentPayment == nil {
				return nil, fmt.Errorf("addendum without payment")
			}
			addendum, err := r.achParser.ParseACHAddendum(line)
			if err != nil {
				return nil, err
			}
			currentPayment.Addenda = append(currentPayment.Addenda, addendum)

		case "G ": // CARS TAS/BETC
			if currentPayment == nil {
				return nil, fmt.Errorf("CARS record without payment")
			}
			cars, err := r.commonParser.ParseCARSTASBETC(line)
			if err != nil {
				return nil, err
			}
			currentPayment.CARSTASBETC = append(currentPayment.CARSTASBETC, cars)

		case "DD": // DNP
			if currentPayment == nil {
				return nil, fmt.Errorf("DNP record without payment")
			}
			dnp, err := r.commonParser.ParseDNP(line)
			if err != nil {
				return nil, err
			}
			currentPayment.DNP = dnp

		case "T ": // Schedule Trailer
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			trailer, err := r.commonParser.ParseScheduleTrailer(line)
			if err != nil {
				return nil, err
			}
			schedule.Trailer = trailer
			return schedule, nil

		default:
			// Put back for next schedule/file trailer
			r.pushBackLine(line)
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			return schedule, nil
		}
	}
}

// parseCheckSchedule parses a check schedule
func (r *Reader) parseCheckSchedule(headerLine string) (*CheckSchedule, error) {
	// Parse header using Check parser
	header, err := r.checkParser.ParseCheckScheduleHeader(headerLine)
	if err != nil {
		return nil, err
	}

	schedule := &CheckSchedule{
		Header: header,
		BaseSchedule: BaseSchedule{
			ScheduleNumber: strings.TrimSpace(header.ScheduleNumber),
			PaymentType:    strings.TrimSpace(header.PaymentTypeCode),
			ALC:            strings.TrimSpace(header.AgencyLocationCode),
			Payments:       make([]Payment, 0),
		},
	}

	// Read payments until schedule trailer
	var currentPayment *CheckPayment

	for {
		line, ok := r.scanLine()
		if !ok {
			return nil, fmt.Errorf("unexpected end of file in check schedule")
		}

		recordCode := r.extractField(line, 1, 2)

		switch recordCode {
		case "12": // Check Payment
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			payment, err := r.checkParser.ParseCheckPayment(line)
			if err != nil {
				return nil, err
			}
			currentPayment = payment

		case "13": // Check Stub
			if currentPayment == nil {
				return nil, fmt.Errorf("check stub without payment")
			}
			stub, err := r.checkParser.ParseCheckStub(line)
			if err != nil {
				return nil, err
			}
			currentPayment.Stub = stub

		case "G ": // CARS TAS/BETC
			if currentPayment == nil {
				return nil, fmt.Errorf("CARS record without payment")
			}
			cars, err := r.commonParser.ParseCARSTASBETC(line)
			if err != nil {
				return nil, err
			}
			currentPayment.CARSTASBETC = append(currentPayment.CARSTASBETC, cars)

		case "DD": // DNP
			if currentPayment == nil {
				return nil, fmt.Errorf("DNP record without payment")
			}
			dnp, err := r.commonParser.ParseDNP(line)
			if err != nil {
				return nil, err
			}
			currentPayment.DNP = dnp

		case "T ": // Schedule Trailer
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			trailer, err := r.commonParser.ParseScheduleTrailer(line)
			if err != nil {
				return nil, err
			}
			schedule.Trailer = trailer
			return schedule, nil

		default:
			// Put back for next schedule/file trailer
			r.pushBackLine(line)
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			return schedule, nil
		}
	}
}

// Helper methods (keep only essential ones used by this file)
func (r *Reader) extractField(line string, start, end int) string {
	if start > len(line) || end > len(line) {
		return ""
	}
	return line[start-1 : end]
}

func (r *Reader) parseAmount(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	amount, _ := strconv.ParseInt(s, 10, 64)
	return amount
}