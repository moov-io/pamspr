package pamspr

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// WriterConfig configures writer behavior
type WriterConfig struct {
	// BufferSize sets the output buffer size (default: 64KB)
	BufferSize int

	// EnableValidation enables validation during writing (default: true)
	EnableValidation bool

	// FlushInterval controls how often to flush buffer (default: every 100 records)
	FlushInterval int

	// ChecksumValidation enables checksum calculation during writing
	ChecksumValidation bool
}

// DefaultWriterConfig returns sensible defaults
func DefaultWriterConfig() *WriterConfig {
	return &WriterConfig{
		BufferSize:         64 * 1024, // 64KB buffer
		EnableValidation:   true,
		FlushInterval:      100,
		ChecksumValidation: false,
	}
}

// Writer writes PAM SPR files in a memory-efficient streaming fashion
type Writer struct {
	writer    io.Writer
	buffer    *bufio.Writer
	validator *Validator
	config    *WriterConfig

	// State tracking
	recordCount   int64
	paymentCount  int64
	scheduleCount int64
	totalAmount   int64
	errors        []error
	isFinalized   bool

	// Memory-efficient formatting buffer
	lineBuffer strings.Builder
}

// NewWriter creates a new PAM SPR writer with streaming capability
// This provides memory-efficient writing for large files
func NewWriter(w io.Writer) *Writer {
	return NewWriterWithConfig(w, DefaultWriterConfig())
}

// NewWriterWithConfig creates a writer with custom configuration
func NewWriterWithConfig(w io.Writer, config *WriterConfig) *Writer {

	var buffer *bufio.Writer
	if config.BufferSize > 0 {
		buffer = bufio.NewWriterSize(w, config.BufferSize)
	} else {
		buffer = bufio.NewWriter(w)
	}

	writer := &Writer{
		writer:    w,
		buffer:    buffer,
		validator: NewValidator(),
		config:    config,
		errors:    make([]error, 0),
	}

	// Pre-allocate line buffer to avoid reallocations
	writer.lineBuffer.Grow(RecordLength + 10)

	return writer
}

// WriteFileHeader writes the file header record
func (w *Writer) WriteFileHeader(header *FileHeader) error {
	if w.recordCount > 0 {
		return fmt.Errorf("file header must be written first")
	}

	if w.config.EnableValidation {
		if err := w.validator.ValidateFileHeader(header); err != nil {
			return fmt.Errorf("validating file header: %w", err)
		}
	}

	line, err := w.formatFileHeader(header)
	if err != nil {
		return fmt.Errorf("formatting file header: %w", err)
	}

	return w.writeLine(line)
}

// WriteScheduleHeader writes a schedule header record
func (w *Writer) WriteScheduleHeader(schedule Schedule) error {
	if w.recordCount == 0 {
		return fmt.Errorf("file header must be written before schedule header")
	}

	var line string
	var err error

	switch s := schedule.(type) {
	case *ACHSchedule:
		line, err = w.formatACHScheduleHeader(s.Header)
	case *CheckSchedule:
		line, err = w.formatCheckScheduleHeader(s.Header)
	default:
		return fmt.Errorf("unknown schedule type")
	}

	if err != nil {
		return fmt.Errorf("formatting schedule header: %w", err)
	}

	w.scheduleCount++
	return w.writeLine(line)
}

// WritePayment writes a payment record (and optionally associated records)
func (w *Writer) WritePayment(payment Payment) error {
	var line string
	var err error

	switch p := payment.(type) {
	case *ACHPayment:
		line, err = w.formatACHPayment(p)
		if err != nil {
			return fmt.Errorf("formatting ACH payment: %w", err)
		}

		// Write main payment record
		if err := w.writeLine(line); err != nil {
			return err
		}

		// Write associated records
		for _, addendum := range p.Addenda {
			addendumLine, err := w.formatACHAddendum(addendum)
			if err != nil {
				return fmt.Errorf("formatting ACH addendum: %w", err)
			}
			if err := w.writeLine(addendumLine); err != nil {
				return err
			}
		}

		for _, cars := range p.CARSTASBETC {
			carsLine, err := w.formatCARSTASBETC(cars)
			if err != nil {
				return fmt.Errorf("formatting CARS record: %w", err)
			}
			if err := w.writeLine(carsLine); err != nil {
				return err
			}
		}

		if p.DNP != nil {
			dnpLine, err := w.formatDNP(p.DNP)
			if err != nil {
				return fmt.Errorf("formatting DNP record: %w", err)
			}
			if err := w.writeLine(dnpLine); err != nil {
				return err
			}
		}

		w.totalAmount += p.Amount

	case *CheckPayment:
		line, err = w.formatCheckPayment(p)
		if err != nil {
			return fmt.Errorf("formatting check payment: %w", err)
		}

		// Write main payment record
		if err := w.writeLine(line); err != nil {
			return err
		}

		// Write associated records
		if p.Stub != nil {
			stubLine, err := w.formatCheckStub(p.Stub)
			if err != nil {
				return fmt.Errorf("formatting check stub: %w", err)
			}
			if err := w.writeLine(stubLine); err != nil {
				return err
			}
		}

		for _, cars := range p.CARSTASBETC {
			carsLine, err := w.formatCARSTASBETC(cars)
			if err != nil {
				return fmt.Errorf("formatting CARS record: %w", err)
			}
			if err := w.writeLine(carsLine); err != nil {
				return err
			}
		}

		if p.DNP != nil {
			dnpLine, err := w.formatDNP(p.DNP)
			if err != nil {
				return fmt.Errorf("formatting DNP record: %w", err)
			}
			if err := w.writeLine(dnpLine); err != nil {
				return err
			}
		}

		w.totalAmount += p.Amount

	default:
		return fmt.Errorf("unknown payment type")
	}

	w.paymentCount++

	// Periodic flush to control memory usage
	if w.config.FlushInterval > 0 && w.recordCount%int64(w.config.FlushInterval) == 0 {
		if err := w.buffer.Flush(); err != nil {
			return fmt.Errorf("flushing buffer: %w", err)
		}
	}

	return nil
}

// WriteScheduleTrailer writes a schedule trailer record
func (w *Writer) WriteScheduleTrailer(trailer *ScheduleTrailer) error {
	line, err := w.formatScheduleTrailer(trailer)
	if err != nil {
		return fmt.Errorf("formatting schedule trailer: %w", err)
	}

	return w.writeLine(line)
}

// WriteFileTrailer writes the file trailer record and finalizes the file
func (w *Writer) WriteFileTrailer(trailer *FileTrailer) error {
	if w.isFinalized {
		return fmt.Errorf("file already finalized")
	}

	// Update trailer with actual counts if not provided
	if trailer.TotalCountRecords == 0 {
		trailer.TotalCountRecords = w.recordCount + 1 // +1 for the trailer itself
	}
	if trailer.TotalCountPayments == 0 {
		trailer.TotalCountPayments = w.paymentCount
	}
	if trailer.TotalAmountPayments == 0 {
		trailer.TotalAmountPayments = w.totalAmount
	}

	if w.config.EnableValidation {
		// Validate trailer totals match what we've written
		if trailer.TotalCountPayments != w.paymentCount {
			return fmt.Errorf("trailer payment count %d doesn't match written payments %d",
				trailer.TotalCountPayments, w.paymentCount)
		}
		if trailer.TotalAmountPayments != w.totalAmount {
			return fmt.Errorf("trailer amount %d doesn't match written amount %d",
				trailer.TotalAmountPayments, w.totalAmount)
		}
	}

	line, err := w.formatFileTrailer(trailer)
	if err != nil {
		return fmt.Errorf("formatting file trailer: %w", err)
	}

	if err := w.writeLine(line); err != nil {
		return err
	}

	// Final flush
	if err := w.buffer.Flush(); err != nil {
		return fmt.Errorf("final flush: %w", err)
	}

	w.isFinalized = true
	return nil
}

// Flush forces a buffer flush
func (w *Writer) Flush() error {
	return w.buffer.Flush()
}

// GetStats returns current writing statistics
func (w *Writer) GetStats() (recordCount, paymentCount, scheduleCount int64, totalAmount int64) {
	return w.recordCount, w.paymentCount, w.scheduleCount, w.totalAmount
}

// Memory-efficient formatting methods using string builder

func (w *Writer) formatFileHeader(header *FileHeader) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(header.RecordCode, 2)
	w.appendField(header.InputSystem, 40)
	w.appendField(header.StandardPaymentVersion, 3)
	w.appendField(header.IsRequestedForSameDayACH, 1)
	w.appendFiller(804)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatACHScheduleHeader(header *ACHScheduleHeader) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(header.RecordCode, 2)
	w.appendField(header.AgencyACHText, 4)
	w.appendFieldRightJustified(header.ScheduleNumber, 14, '0')
	w.appendField(header.PaymentTypeCode, 25)
	w.appendField(header.StandardEntryClassCode, 3)
	w.appendNumeric(header.AgencyLocationCode, 8)
	w.appendFiller(1)
	w.appendField(header.FederalEmployerIDNumber, 10)
	w.appendFiller(783)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatCheckScheduleHeader(header *CheckScheduleHeader) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(header.RecordCode, 2)
	w.appendFieldRightJustified(header.ScheduleNumber, 14, '0')
	w.appendField(header.PaymentTypeCode, 25)
	w.appendNumeric(header.AgencyLocationCode, 8)
	w.appendFiller(9)
	w.appendField(header.CheckPaymentEnclosureCode, 10)
	w.appendFiller(782)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatACHPayment(payment *ACHPayment) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(payment.RecordCode, 2)
	w.appendField(payment.AgencyAccountIdentifier, 16)
	w.appendAmount(payment.Amount, 10)
	w.appendField(payment.AgencyPaymentTypeCode, 1)
	w.appendField(payment.IsTOP_Offset, 1)
	w.appendField(payment.PayeeName, 35)
	w.appendField(payment.PayeeAddressLine1, 35)
	w.appendField(payment.PayeeAddressLine2, 35)
	w.appendField(payment.CityName, 27)
	w.appendField(payment.StateName, 10)
	w.appendField(payment.StateCodeText, 2)
	w.appendField(payment.PostalCode, 5)
	w.appendField(payment.PostalCodeExtension, 5)
	w.appendField(payment.CountryCodeText, 2)
	w.appendField(payment.RoutingNumber, 9)
	w.appendField(payment.AccountNumber, 17)
	w.appendField(payment.ACH_TransactionCode, 2)
	w.appendField(payment.PayeeIdentifierAdditional, 9)
	w.appendField(payment.PayeeNameAdditional, 35)
	w.appendField(payment.PaymentID, 20)
	w.appendFieldNoJustify(payment.Reconcilement, 100)
	w.appendField(payment.TIN, 9)
	w.appendField(payment.PaymentRecipientTINIndicator, 1)
	w.appendField(payment.AdditionalPayeeTINIndicator, 1)
	w.appendField(payment.AmountEligibleForOffset, 10)
	w.appendField(payment.PayeeAddressLine3, 35)
	w.appendField(payment.PayeeAddressLine4, 35)
	w.appendField(payment.CountryName, 40)
	w.appendField(payment.ConsularCode, 3)
	w.appendField(payment.SubPaymentTypeCode, 32)
	w.appendField(payment.PayerMechanism, 20)
	w.appendField(payment.PaymentDescriptionCode, 2)
	w.appendFiller(284)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatCheckPayment(payment *CheckPayment) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(payment.RecordCode, 2)
	w.appendField(payment.AgencyAccountIdentifier, 16)
	w.appendAmount(payment.Amount, 10)
	w.appendField(payment.AgencyPaymentTypeCode, 1)
	w.appendField(payment.IsTOP_Offset, 1)
	w.appendField(payment.PayeeName, 35)
	w.appendField(payment.PayeeAddressLine1, 35)
	w.appendField(payment.PayeeAddressLine2, 35)
	w.appendField(payment.PayeeAddressLine3, 35)
	w.appendField(payment.PayeeAddressLine4, 35)
	w.appendField(payment.CityName, 27)
	w.appendField(payment.StateName, 10)
	w.appendField(payment.StateCodeText, 2)
	w.appendField(payment.PostalCode, 5)
	w.appendField(payment.PostalCodeExtension, 5)
	w.appendField(payment.PostNetBarcodeDeliveryPoint, 3)
	w.appendFiller(14)
	w.appendField(payment.CountryName, 40)
	w.appendField(payment.ConsularCode, 3)
	w.appendField(payment.CheckLegendText1, 55)
	w.appendField(payment.CheckLegendText2, 55)
	w.appendField(payment.PayeeIdentifier_Secondary, 9)
	w.appendField(payment.PartyName_Secondary, 35)
	w.appendField(payment.PaymentID, 20)
	w.appendFieldNoJustify(payment.Reconcilement, 100)
	w.appendField(payment.SpecialHandling, 50)
	w.appendField(payment.TIN, 9)
	w.appendField(payment.USPSIntelligentMailBarcode, 50)
	w.appendField(payment.PaymentRecipientTINIndicator, 1)
	w.appendField(payment.SecondaryPayeeTINIndicator, 1)
	w.appendField(payment.AmountEligibleForOffset, 10)
	w.appendField(payment.SubPaymentTypeCode, 32)
	w.appendField(payment.PayerMechanism, 20)
	w.appendField(payment.PaymentDescriptionCode, 2)
	w.appendFiller(87)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatACHAddendum(addendum *ACHAddendum) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(addendum.RecordCode, 2)
	w.appendField(addendum.PaymentID, 20)

	if addendum.RecordCode == "03" {
		w.appendField(addendum.AddendaInformation, 80)
		w.appendFiller(748)
	} else if addendum.RecordCode == "04" { // CTX
		w.appendField(addendum.AddendaInformation, 800)
		w.appendFiller(28)
	} else {
		return "", fmt.Errorf("invalid addendum record code: %s", addendum.RecordCode)
	}

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatCARSTASBETC(cars *CARSTASBETC) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(cars.RecordCode, 2)
	w.appendField(cars.PaymentID, 20)
	w.appendField(cars.SubLevelPrefixCode, 2)
	w.appendField(cars.AllocationTransferAgencyID, 3)
	w.appendField(cars.AgencyIdentifier, 3)
	w.appendField(cars.BeginningPeriodOfAvailability, 4)
	w.appendField(cars.EndingPeriodOfAvailability, 4)
	w.appendField(cars.AvailabilityTypeCode, 1)
	w.appendField(cars.MainAccountCode, 4)
	w.appendField(cars.SubAccountCode, 3)
	w.appendField(cars.BusinessEventTypeCode, 8)
	w.appendAmount(cars.AccountClassificationAmount, 10)
	w.appendField(cars.IsCredit, 1)
	w.appendFiller(785)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatCheckStub(stub *CheckStub) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(stub.RecordCode, 2)
	w.appendField(stub.PaymentID, 20)

	for i := 0; i < 14; i++ {
		w.appendField(stub.PaymentIdentificationLines[i], 55)
	}

	w.appendFiller(58)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatDNP(dnp *DNPRecord) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(dnp.RecordCode, 2)
	w.appendField(dnp.PaymentID, 20)
	w.appendFieldNoJustify(dnp.DNPDetail, 766)
	w.appendFiller(62)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatScheduleTrailer(trailer *ScheduleTrailer) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(trailer.RecordCode, 2)
	w.appendFiller(10)
	w.appendNumeric(fmt.Sprintf("%d", trailer.ScheduleCount), 8)
	w.appendFiller(3)
	w.appendAmount(trailer.ScheduleAmount, 15)
	w.appendFiller(812)

	return w.lineBuffer.String(), nil
}

func (w *Writer) formatFileTrailer(trailer *FileTrailer) (string, error) {
	w.lineBuffer.Reset()
	w.lineBuffer.Grow(RecordLength)

	w.appendField(trailer.RecordCode, 2)
	w.appendNumeric(fmt.Sprintf("%d", trailer.TotalCountRecords), 18)
	w.appendNumeric(fmt.Sprintf("%d", trailer.TotalCountPayments), 18)
	w.appendAmount(trailer.TotalAmountPayments, 18)
	w.appendFiller(794)

	return w.lineBuffer.String(), nil
}

// Memory-efficient field formatting helpers using string builder

func (w *Writer) appendField(value string, length int) {
	if len(value) >= length {
		w.lineBuffer.WriteString(value[:length])
	} else {
		w.lineBuffer.WriteString(value)
		w.appendSpaces(length - len(value))
	}
}

func (w *Writer) appendFieldRightJustified(value string, length int, padChar rune) {
	value = strings.TrimSpace(value)
	if len(value) >= length {
		w.lineBuffer.WriteString(value[:length])
	} else {
		padCount := length - len(value)
		for i := 0; i < padCount; i++ {
			w.lineBuffer.WriteRune(padChar)
		}
		w.lineBuffer.WriteString(value)
	}
}

func (w *Writer) appendFieldNoJustify(value string, length int) {
	w.appendField(value, length) // Same as left-justified
}

func (w *Writer) appendNumeric(value string, length int) {
	// Extract only numeric characters
	numericOnly := strings.Builder{}
	for _, r := range value {
		if r >= '0' && r <= '9' {
			numericOnly.WriteRune(r)
		}
	}

	numeric := numericOnly.String()
	if len(numeric) >= length {
		w.lineBuffer.WriteString(numeric[:length])
	} else {
		// Right justify with zeros
		padCount := length - len(numeric)
		for i := 0; i < padCount; i++ {
			w.lineBuffer.WriteByte('0')
		}
		w.lineBuffer.WriteString(numeric)
	}
}

func (w *Writer) appendAmount(cents int64, length int) {
	// Handle negative amounts
	amount := cents
	if amount < 0 {
		amount = -amount
	}

	amountStr := fmt.Sprintf("%d", amount)
	w.appendNumeric(amountStr, length)
}

func (w *Writer) appendFiller(length int) {
	w.appendSpaces(length)
}

func (w *Writer) appendSpaces(count int) {
	for i := 0; i < count; i++ {
		w.lineBuffer.WriteByte(' ')
	}
}

func (w *Writer) writeLine(line string) error {
	if len(line) != RecordLength {
		return fmt.Errorf("invalid line length: expected %d, got %d", RecordLength, len(line))
	}

	if _, err := w.buffer.WriteString(line); err != nil {
		return err
	}
	if err := w.buffer.WriteByte('\n'); err != nil {
		return err
	}

	w.recordCount++
	return nil
}

// Write writes a complete PAM SPR file in streaming fashion
// This method provides compatibility with the traditional Writer.Write() API
// while using streaming internally for memory efficiency
func (w *Writer) Write(file *File) error {
	// Write file header
	if err := w.WriteFileHeader(file.Header); err != nil {
		return fmt.Errorf("writing file header: %w", err)
	}

	// Write schedules
	for i, schedule := range file.Schedules {
		if err := w.WriteScheduleHeader(schedule); err != nil {
			return fmt.Errorf("writing schedule %d header: %w", i, err)
		}

		// Write payments
		for _, payment := range schedule.GetPayments() {
			if err := w.WritePayment(payment); err != nil {
				return fmt.Errorf("writing payment in schedule %d: %w", i, err)
			}
		}

		// Write schedule trailer
		if err := w.WriteScheduleTrailer(schedule.GetTrailer()); err != nil {
			return fmt.Errorf("writing schedule %d trailer: %w", i, err)
		}
	}

	// Write file trailer
	if err := w.WriteFileTrailer(file.Trailer); err != nil {
		return fmt.Errorf("writing file trailer: %w", err)
	}

	return nil
}
