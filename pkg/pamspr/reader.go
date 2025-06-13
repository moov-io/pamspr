package pamspr

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// PaymentCallback is called for each payment as it's parsed
// Return false to stop processing, true to continue
type PaymentCallback func(payment Payment, scheduleIndex int, paymentIndex int) bool

// ScheduleCallback is called when a schedule header is encountered
type ScheduleCallback func(schedule Schedule, scheduleIndex int) bool

// RecordCallback is called for each raw record (for debugging/monitoring)
type RecordCallback func(recordType string, lineNumber int, line string)

// ReaderConfig configures the reader behavior
type ReaderConfig struct {
	// BufferSize sets the scanner buffer size (default: 64KB)
	BufferSize int

	// EnableValidation enables validation during streaming (default: true)
	EnableValidation bool

	// CollectErrors determines if errors should be accumulated (default: true)
	CollectErrors bool

	// MaxErrors limits error collection to prevent memory growth (0 = unlimited)
	MaxErrors int

	// SkipInvalidRecords continues processing on invalid records (default: false)
	SkipInvalidRecords bool
}

// DefaultConfig returns sensible defaults for the reader
func DefaultConfig() *ReaderConfig {
	return &ReaderConfig{
		BufferSize:         64 * 1024, // 64KB buffer
		EnableValidation:   true,
		CollectErrors:      true,
		MaxErrors:          1000, // Limit to 1000 errors
		SkipInvalidRecords: false,
	}
}

// Reader processes PAM SPR files in a streaming fashion
// without loading the entire file into memory
type Reader struct {
	scanner     *bufio.Scanner
	validator   *Validator
	config      *ReaderConfig
	lineNum     int
	errors      []error
	nextLine    *string
	currentFile *FileHeader

	// Parsers
	fileParser   *FileParser
	achParser    *ACHParser
	checkParser  *CheckParser
	commonParser *CommonParser

	// Statistics
	stats Stats
}

// Stats tracks processing statistics
type Stats struct {
	LinesProcessed     int64
	PaymentsProcessed  int64
	SchedulesProcessed int64
	ErrorsEncountered  int64
	BytesProcessed     int64
}

// NewReader creates a new PAM SPR reader with streaming capability
// This provides memory-efficient processing for large files
func NewReader(r io.Reader) *Reader {
	return NewReaderWithConfig(r, DefaultConfig())
}

// NewReaderWithConfig creates a reader with custom configuration
func NewReaderWithConfig(r io.Reader, config *ReaderConfig) *Reader {

	validator := NewValidator()
	scanner := bufio.NewScanner(r)

	// Set custom buffer size if specified
	if config.BufferSize > 0 {
		buffer := make([]byte, config.BufferSize)
		scanner.Buffer(buffer, config.BufferSize)
	}

	return &Reader{
		scanner:      scanner,
		validator:    validator,
		config:       config,
		errors:       make([]error, 0),
		fileParser:   NewFileParser(validator),
		achParser:    NewACHParser(validator),
		checkParser:  NewCheckParser(validator),
		commonParser: NewCommonParser(validator),
		stats:        Stats{},
	}
}

// GetStats returns current processing statistics
func (r *Reader) GetStats() Stats {
	return r.stats
}

// GetErrors returns accumulated errors (if enabled)
func (r *Reader) GetErrors() []error {
	return r.errors
}

// ProcessFile streams through the entire file calling callbacks for each component
func (r *Reader) ProcessFile(
	scheduleCallback ScheduleCallback,
	paymentCallback PaymentCallback,
	recordCallback RecordCallback,
) error {
	// Read file header and notify callback
	line, ok := r.scanLine()
	if !ok {
		return fmt.Errorf("file is empty")
	}

	if len(line) != RecordLength {
		return fmt.Errorf("line 1: invalid record length %d, expected %d", len(line), RecordLength)
	}

	if !strings.HasPrefix(line, "H ") {
		return fmt.Errorf("line 1: expected file header, got record code '%s'", line[:2])
	}

	header, err := r.fileParser.ParseFileHeader(line)
	if err != nil {
		return fmt.Errorf("parsing file header: %w", err)
	}

	if r.config.EnableValidation {
		if err := r.validator.ValidateFileHeader(header); err != nil {
			return fmt.Errorf("validating file header: %w", err)
		}
	}

	r.currentFile = header

	// Call record callback for header
	if recordCallback != nil {
		recordCallback("H ", r.lineNum, line)
	}

	scheduleIndex := 0

	// Process schedules
	for {
		line, ok := r.scanLine()
		if !ok {
			break
		}

		if len(line) < 2 {
			r.addError(fmt.Errorf("line %d: line too short", r.lineNum))
			continue
		}

		recordCode := line[:2]

		// Call record callback if provided
		if recordCallback != nil {
			recordCallback(recordCode, r.lineNum, line)
		}

		switch recordCode {
		case "01", "11": // Schedule headers
			schedule, err := r.parseScheduleHeader(line)
			if err != nil {
				if !r.config.SkipInvalidRecords {
					return fmt.Errorf("parsing schedule header at line %d: %w", r.lineNum, err)
				}
				r.addError(err)
				continue
			}

			r.stats.SchedulesProcessed++

			// Call schedule callback
			if scheduleCallback != nil && !scheduleCallback(schedule, scheduleIndex) {
				return nil // Stop processing if callback returns false
			}

			// Process payments within this schedule
			err = r.processSchedulePayments(schedule, scheduleIndex, paymentCallback, recordCallback)
			if err != nil {
				return err
			}

			scheduleIndex++

		case "E ": // File trailer
			// Call record callback for trailer
			if recordCallback != nil {
				recordCallback("E ", r.lineNum, line)
			}
			return nil // End of schedules

		default:
			if !r.config.SkipInvalidRecords {
				return fmt.Errorf("unexpected record code '%s' at line %d", recordCode, r.lineNum)
			}
			r.addError(fmt.Errorf("line %d: unexpected record code '%s'", r.lineNum, recordCode))
		}
	}

	return nil
}

// ProcessPaymentsOnly streams through file calling callback only for payments
// This is optimized for payment-only processing without building schedule objects
func (r *Reader) ProcessPaymentsOnly(callback PaymentCallback) error {
	// Read file header (but don't store full structure)
	_, err := r.readFileHeader()
	if err != nil {
		return fmt.Errorf("reading file header: %w", err)
	}

	scheduleIndex := 0
	currentSchedule := (*BaseSchedule)(nil)

	for {
		line, ok := r.scanLine()
		if !ok {
			break
		}

		if len(line) < 2 {
			continue
		}

		recordCode := line[:2]

		switch recordCode {
		case "01", "11": // Schedule headers
			scheduleIndex++
			// Parse minimal schedule info but don't store full object
			currentSchedule = &BaseSchedule{}

		case "02": // ACH Payment
			if currentSchedule == nil {
				return fmt.Errorf("payment record without schedule header at line %d", r.lineNum)
			}

			payment, err := r.achParser.ParseACHPayment(line)
			if err != nil {
				if !r.config.SkipInvalidRecords {
					return fmt.Errorf("parsing ACH payment at line %d: %w", r.lineNum, err)
				}
				r.addError(err)
				continue
			}

			r.stats.PaymentsProcessed++

			// Call payment callback
			if !callback(payment, scheduleIndex-1, int(r.stats.PaymentsProcessed-1)) {
				return nil // Stop processing
			}

		case "12": // Check Payment
			if currentSchedule == nil {
				return fmt.Errorf("payment record without schedule header at line %d", r.lineNum)
			}

			payment, err := r.checkParser.ParseCheckPayment(line)
			if err != nil {
				if !r.config.SkipInvalidRecords {
					return fmt.Errorf("parsing check payment at line %d: %w", r.lineNum, err)
				}
				r.addError(err)
				continue
			}

			r.stats.PaymentsProcessed++

			// Call payment callback
			if !callback(payment, scheduleIndex-1, int(r.stats.PaymentsProcessed-1)) {
				return nil // Stop processing
			}

		case "E ": // File trailer
			return nil

		// Skip associated records and schedule trailers in payment-only mode
		case "03", "04", "13", "G ", "DD", "T ":
			continue

		default:
			if !r.config.SkipInvalidRecords {
				return fmt.Errorf("unexpected record code '%s' at line %d", recordCode, r.lineNum)
			}
		}
	}

	return nil
}

// ValidateFileStructureOnly performs streaming validation without building objects
func (r *Reader) ValidateFileStructureOnly() error {
	var hasHeader, hasTrailer bool
	var scheduleCount int

	for {
		line, ok := r.scanLine()
		if !ok {
			break
		}

		if len(line) != RecordLength {
			return fmt.Errorf("line %d: invalid record length %d, expected %d", r.lineNum, len(line), RecordLength)
		}

		recordCode := line[:2]

		switch recordCode {
		case "H ":
			if hasHeader {
				return fmt.Errorf("line %d: multiple file headers", r.lineNum)
			}
			hasHeader = true
			if r.config.EnableValidation {
				header, err := r.fileParser.ParseFileHeader(line)
				if err != nil {
					return fmt.Errorf("line %d: invalid file header: %w", r.lineNum, err)
				}
				if err := r.validator.ValidateFileHeader(header); err != nil {
					return fmt.Errorf("line %d: file header validation: %w", r.lineNum, err)
				}
			}

		case "01", "11":
			scheduleCount++

		case "E ":
			hasTrailer = true
			// Should be last record
			if r.scanner.Scan() {
				return fmt.Errorf("line %d: records found after file trailer", r.lineNum+1)
			}

		case "02", "12", "03", "04", "13", "G ", "DD", "T ":
			// Valid record types, just count them

		default:
			return fmt.Errorf("line %d: invalid record code '%s'", r.lineNum, recordCode)
		}

		r.stats.LinesProcessed++
	}

	if !hasHeader {
		return fmt.Errorf("file missing header record")
	}
	if !hasTrailer {
		return fmt.Errorf("file missing trailer record")
	}
	if scheduleCount == 0 {
		return fmt.Errorf("file contains no schedules")
	}

	return nil
}

// Helper methods

func (r *Reader) scanLine() (string, bool) {
	if r.nextLine != nil {
		line := *r.nextLine
		r.nextLine = nil
		r.lineNum++
		r.stats.BytesProcessed += int64(len(line))
		return line, true
	}

	if r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++
		r.stats.LinesProcessed++
		r.stats.BytesProcessed += int64(len(line))
		return line, true
	}

	return "", false
}

func (r *Reader) pushBackLine(line string) {
	r.nextLine = &line
	r.lineNum--
	r.stats.LinesProcessed--
}

func (r *Reader) addError(err error) {
	if !r.config.CollectErrors {
		return
	}

	r.stats.ErrorsEncountered++

	if r.config.MaxErrors > 0 && len(r.errors) >= r.config.MaxErrors {
		return // Don't exceed max errors
	}

	r.errors = append(r.errors, err)
}

func (r *Reader) readFileHeader() (*FileHeader, error) {
	line, ok := r.scanLine()
	if !ok {
		return nil, fmt.Errorf("file is empty")
	}

	if len(line) != RecordLength {
		return nil, fmt.Errorf("line 1: invalid record length %d, expected %d", len(line), RecordLength)
	}

	if !strings.HasPrefix(line, "H ") {
		return nil, fmt.Errorf("line 1: expected file header, got record code '%s'", line[:2])
	}

	header, err := r.fileParser.ParseFileHeader(line)
	if err != nil {
		return nil, fmt.Errorf("parsing file header: %w", err)
	}

	if r.config.EnableValidation {
		if err := r.validator.ValidateFileHeader(header); err != nil {
			return nil, fmt.Errorf("validating file header: %w", err)
		}
	}

	return header, nil
}

func (r *Reader) parseScheduleHeader(line string) (Schedule, error) {
	recordCode := line[:2]

	switch recordCode {
	case "01":
		header, err := r.achParser.ParseACHScheduleHeader(line)
		if err != nil {
			return nil, err
		}
		return &ACHSchedule{Header: header, BaseSchedule: BaseSchedule{}}, nil

	case "11":
		header, err := r.checkParser.ParseCheckScheduleHeader(line)
		if err != nil {
			return nil, err
		}
		return &CheckSchedule{Header: header, BaseSchedule: BaseSchedule{}}, nil

	default:
		return nil, fmt.Errorf("invalid schedule header record code: %s", recordCode)
	}
}

func (r *Reader) processSchedulePayments(
	schedule Schedule,
	scheduleIndex int,
	paymentCallback PaymentCallback,
	recordCallback RecordCallback,
) error {
	paymentIndex := 0

	for {
		line, ok := r.scanLine()
		if !ok {
			return fmt.Errorf("unexpected end of file in schedule")
		}

		if len(line) < 2 {
			continue
		}

		recordCode := line[:2]

		// Call record callback if provided
		if recordCallback != nil {
			recordCallback(recordCode, r.lineNum, line)
		}

		switch recordCode {
		case "02": // ACH Payment
			payment, err := r.achParser.ParseACHPayment(line)
			if err != nil {
				if !r.config.SkipInvalidRecords {
					return fmt.Errorf("parsing ACH payment at line %d: %w", r.lineNum, err)
				}
				r.addError(err)
				continue
			}

			r.stats.PaymentsProcessed++

			if paymentCallback != nil && !paymentCallback(payment, scheduleIndex, paymentIndex) {
				return nil // Stop processing
			}
			paymentIndex++

		case "12": // Check Payment
			payment, err := r.checkParser.ParseCheckPayment(line)
			if err != nil {
				if !r.config.SkipInvalidRecords {
					return fmt.Errorf("parsing check payment at line %d: %w", r.lineNum, err)
				}
				r.addError(err)
				continue
			}

			r.stats.PaymentsProcessed++

			if paymentCallback != nil && !paymentCallback(payment, scheduleIndex, paymentIndex) {
				return nil // Stop processing
			}
			paymentIndex++

		case "03", "04", "13", "G ", "DD": // Associated records - skip in streaming mode
			continue

		case "T ": // Schedule trailer
			return nil // End of schedule

		case "01", "11": // Next schedule
			r.pushBackLine(line)
			return nil

		case "E ": // File trailer
			r.pushBackLine(line)
			return nil

		default:
			if !r.config.SkipInvalidRecords {
				return fmt.Errorf("unexpected record code '%s' at line %d", recordCode, r.lineNum)
			}
			r.addError(fmt.Errorf("line %d: unexpected record code '%s'", r.lineNum, recordCode))
		}
	}
}

// ReadAll reads a complete PAM SPR file in streaming fashion
// This method provides compatibility with the traditional Reader.Read() API
// while using streaming internally for memory efficiency
func (r *Reader) ReadAll() (*File, error) {
	return r.readAllLegacyCompatible()
}

// ReadPayments reads only payments from the file in streaming fashion
// Returns a slice of all payments found in the file
func (r *Reader) ReadPayments() ([]Payment, error) {
	var payments []Payment

	err := r.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
		payments = append(payments, payment)
		return true
	})

	if err != nil {
		return nil, err
	}

	return payments, nil
}

// Read is an alias for ReadAll to maintain compatibility
func (r *Reader) Read() (*File, error) {
	return r.ReadAll()
}

// For compatibility, implement a simple ReadAll that works like the legacy reader
// This is less efficient but maintains full compatibility
func (r *Reader) readAllLegacyCompatible() (*File, error) {
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

	// Check that we found a file trailer
	if file.Trailer == nil {
		return nil, fmt.Errorf("missing file trailer")
	}

	return file, nil
}

// parseSchedule parses a schedule starting with the given line
func (r *Reader) parseSchedule(firstLine string) (Schedule, error) {
	recordCode := firstLine[:2]

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

		recordCode := line[:2]

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

		recordCode := line[:2]

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
