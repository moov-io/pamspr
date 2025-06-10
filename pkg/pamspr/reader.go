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
	scanner   *bufio.Scanner
	validator *Validator
	lineNum   int
	errors    []error
	nextLine  *string // For line pushback
}

// NewReader creates a new PAM SPR reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		scanner:   bufio.NewScanner(r),
		validator: NewValidator(),
		errors:    make([]error, 0),
		nextLine:  nil,
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

	header, err := r.parseFileHeader(line)
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
			trailer, err := r.parseFileTrailer(line)
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

// parseFileHeader parses a file header record
func (r *Reader) parseFileHeader(line string) (*FileHeader, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	header := &FileHeader{
		RecordCode:               r.extractField(line, 1, 2),
		InputSystem:              r.extractField(line, 3, 42),
		StandardPaymentVersion:   r.extractField(line, 43, 45),
		IsRequestedForSameDayACH: r.extractField(line, 46, 46),
	}

	// Validate header
	if err := r.validator.ValidateFileHeader(header); err != nil {
		return nil, err
	}

	return header, nil
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
	// Parse header
	header := &ACHScheduleHeader{
		RecordCode:              r.extractField(headerLine, 1, 2),
		AgencyACHText:           r.extractField(headerLine, 3, 6),
		ScheduleNumber:          r.extractField(headerLine, 7, 20),
		PaymentTypeCode:         r.extractField(headerLine, 21, 45),
		StandardEntryClassCode:  r.extractField(headerLine, 46, 48),
		AgencyLocationCode:      r.extractField(headerLine, 49, 56),
		FederalEmployerIDNumber: r.extractField(headerLine, 58, 67),
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
			payment, err := r.parseACHPayment(line)
			if err != nil {
				return nil, err
			}
			payment.StandardEntryClassCode = header.StandardEntryClassCode
			currentPayment = payment

		case "03", "04": // ACH Addendum
			if currentPayment == nil {
				return nil, fmt.Errorf("addendum without payment")
			}
			addendum, err := r.parseACHAddendum(line)
			if err != nil {
				return nil, err
			}
			currentPayment.Addenda = append(currentPayment.Addenda, addendum)

		case "G ": // CARS TAS/BETC
			if currentPayment == nil {
				return nil, fmt.Errorf("CARS record without payment")
			}
			cars, err := r.parseCARSTASBETC(line)
			if err != nil {
				return nil, err
			}
			currentPayment.CARSTASBETC = append(currentPayment.CARSTASBETC, cars)

		case "DD": // DNP
			if currentPayment == nil {
				return nil, fmt.Errorf("DNP record without payment")
			}
			dnp, err := r.parseDNP(line)
			if err != nil {
				return nil, err
			}
			currentPayment.DNP = dnp

		case "T ": // Schedule Trailer
			if currentPayment != nil {
				schedule.Payments = append(schedule.Payments, currentPayment)
			}
			trailer, err := r.parseScheduleTrailer(line)
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

// parseACHPayment parses an ACH payment record
func (r *Reader) parseACHPayment(line string) (*ACHPayment, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length")
	}

	payment := &ACHPayment{
		RecordCode:                   r.extractField(line, 1, 2),
		AgencyAccountIdentifier:      r.extractField(line, 3, 18),
		Amount:                       r.parseAmount(r.extractField(line, 19, 28)),
		AgencyPaymentTypeCode:        r.extractField(line, 29, 29),
		IsTOPOffset:                  r.extractField(line, 30, 30),
		PayeeName:                    r.extractField(line, 31, 65),
		PayeeAddressLine1:            r.extractField(line, 66, 100),
		PayeeAddressLine2:            r.extractField(line, 101, 135),
		CityName:                     r.extractField(line, 136, 162),
		StateName:                    r.extractField(line, 163, 172),
		StateCodeText:                r.extractField(line, 173, 174),
		PostalCode:                   r.extractField(line, 175, 179),
		PostalCodeExtension:          r.extractField(line, 180, 184),
		CountryCodeText:              r.extractField(line, 185, 186),
		RoutingNumber:                r.extractField(line, 187, 195),
		AccountNumber:                r.extractField(line, 196, 212),
		ACHTransactionCode:           r.extractField(line, 213, 214),
		PayeeIdentifierAdditional:    r.extractField(line, 215, 223),
		PayeeNameAdditional:          r.extractField(line, 224, 258),
		PaymentID:                    r.extractField(line, 259, 278),
		Reconcilement:                r.extractField(line, 279, 378),
		TIN:                          r.extractField(line, 379, 387),
		PaymentRecipientTINIndicator: r.extractField(line, 388, 388),
		AdditionalPayeeTINIndicator:  r.extractField(line, 389, 389),
		AmountEligibleForOffset:      r.extractField(line, 390, 399),
		PayeeAddressLine3:            r.extractField(line, 400, 434),
		PayeeAddressLine4:            r.extractField(line, 435, 469),
		CountryName:                  r.extractField(line, 470, 509),
		ConsularCode:                 r.extractField(line, 510, 512),
		SubPaymentTypeCode:           r.extractField(line, 513, 544),
		PayerMechanism:               r.extractField(line, 545, 564),
		PaymentDescriptionCode:       r.extractField(line, 565, 566),
		Addenda:                      make([]*ACHAddendum, 0),
		CARSTASBETC:                  make([]*CARSTASBETC, 0),
	}

	// Validate payment
	if err := r.validator.ValidateACHPayment(payment); err != nil {
		r.errors = append(r.errors, fmt.Errorf("line %d: %w", r.lineNum, err))
	}

	return payment, nil
}

// parseCheckSchedule parses a check schedule
func (r *Reader) parseCheckSchedule(headerLine string) (*CheckSchedule, error) {
	// Similar to parseACHSchedule but for checks
	header := &CheckScheduleHeader{
		RecordCode:                r.extractField(headerLine, 1, 2),
		ScheduleNumber:            r.extractField(headerLine, 3, 16),
		PaymentTypeCode:           r.extractField(headerLine, 17, 41),
		AgencyLocationCode:        r.extractField(headerLine, 42, 49),
		CheckPaymentEnclosureCode: r.extractField(headerLine, 59, 68),
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

	// Continue reading check payments...
	// Implementation similar to ACH schedule parsing

	return schedule, nil
}

// Helper methods
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

// Additional parsing methods for other record types...
func (r *Reader) parseACHAddendum(line string) (*ACHAddendum, error) {
	recordCode := r.extractField(line, 1, 2)

	if recordCode == "03" {
		return &ACHAddendum{
			RecordCode:         recordCode,
			PaymentID:          r.extractField(line, 3, 22),
			AddendaInformation: r.extractField(line, 23, 102),
		}, nil
	} else if recordCode == "04" {
		return &ACHAddendum{
			RecordCode:         recordCode,
			PaymentID:          r.extractField(line, 3, 22),
			AddendaInformation: r.extractField(line, 23, 822),
		}, nil
	}

	return nil, fmt.Errorf("invalid addendum record code: %s", recordCode)
}

func (r *Reader) parseCARSTASBETC(line string) (*CARSTASBETC, error) {
	return &CARSTASBETC{
		RecordCode:                    r.extractField(line, 1, 2),
		PaymentID:                     r.extractField(line, 3, 22),
		SubLevelPrefixCode:            r.extractField(line, 23, 24),
		AllocationTransferAgencyID:    r.extractField(line, 25, 27),
		AgencyIdentifier:              r.extractField(line, 28, 30),
		BeginningPeriodOfAvailability: r.extractField(line, 31, 34),
		EndingPeriodOfAvailability:    r.extractField(line, 35, 38),
		AvailabilityTypeCode:          r.extractField(line, 39, 39),
		MainAccountCode:               r.extractField(line, 40, 43),
		SubAccountCode:                r.extractField(line, 44, 46),
		BusinessEventTypeCode:         r.extractField(line, 47, 54),
		AccountClassificationAmount:   r.parseAmount(r.extractField(line, 55, 64)),
		IsCredit:                      r.extractField(line, 65, 65),
	}, nil
}

func (r *Reader) parseDNP(line string) (*DNPRecord, error) {
	return &DNPRecord{
		RecordCode: r.extractField(line, 1, 2),
		PaymentID:  r.extractField(line, 3, 22),
		DNPDetail:  r.extractField(line, 23, 788),
	}, nil
}

func (r *Reader) parseScheduleTrailer(line string) (*ScheduleTrailer, error) {
	return &ScheduleTrailer{
		RecordCode:     r.extractField(line, 1, 2),
		ScheduleCount:  r.parseAmount(r.extractField(line, 13, 20)),
		ScheduleAmount: r.parseAmount(r.extractField(line, 24, 38)),
	}, nil
}

func (r *Reader) parseFileTrailer(line string) (*FileTrailer, error) {
	return &FileTrailer{
		RecordCode:          r.extractField(line, 1, 2),
		TotalCountRecords:   r.parseAmount(r.extractField(line, 3, 20)),
		TotalCountPayments:  r.parseAmount(r.extractField(line, 21, 38)),
		TotalAmountPayments: r.parseAmount(r.extractField(line, 39, 56)),
	}, nil
}
