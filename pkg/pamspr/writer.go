package pamspr

import (
	"fmt"
	"io"
	"strings"
)

// Writer writes PAM SPR files
type Writer struct {
	w         io.Writer
	validator *Validator
	formatter *FieldFormatter
	errors    []error
}

// NewWriter creates a new PAM SPR writer
func NewWriter(w io.Writer) *Writer {
	validator := NewValidator()
	return &Writer{
		w:         w,
		validator: validator,
		formatter: NewFieldFormatter(validator),
		errors:    make([]error, 0),
	}
}

// Write writes a complete PAM SPR file
func (w *Writer) Write(file *File) error {
	// Validate file structure before writing
	if err := w.validator.ValidateFileStructure(file); err != nil {
		return fmt.Errorf("file validation: %w", err)
	}

	// Write file header
	if err := w.writeFileHeader(file.Header); err != nil {
		return fmt.Errorf("writing file header: %w", err)
	}

	// Write schedules
	for i, schedule := range file.Schedules {
		if err := w.writeSchedule(schedule); err != nil {
			return fmt.Errorf("writing schedule %d: %w", i, err)
		}
	}

	// Write file trailer
	if err := w.writeFileTrailer(file.Trailer); err != nil {
		return fmt.Errorf("writing file trailer: %w", err)
	}

	return nil
}

// writeFileHeader writes the file header record
func (w *Writer) writeFileHeader(header *FileHeader) error {
	line := w.formatField(header.RecordCode, 2) +
		w.formatField(header.InputSystem, 40) +
		w.formatField(header.StandardPaymentVersion, 3) +
		w.formatField(header.IsRequestedForSameDayACH, 1) +
		w.formatField("", 804) // Filler

	return w.writeLine(line)
}

// writeSchedule writes a complete schedule
func (w *Writer) writeSchedule(schedule Schedule) error {
	switch s := schedule.(type) {
	case *ACHSchedule:
		return w.writeACHSchedule(s)
	case *CheckSchedule:
		return w.writeCheckSchedule(s)
	default:
		return fmt.Errorf("unknown schedule type")
	}
}

// writeACHSchedule writes an ACH schedule
func (w *Writer) writeACHSchedule(schedule *ACHSchedule) error {
	// Write schedule header
	if err := w.writeACHScheduleHeader(schedule.Header); err != nil {
		return err
	}

	// Write payments
	for _, payment := range schedule.Payments {
		if achPayment, ok := payment.(*ACHPayment); ok {
			// Write payment record
			if err := w.writeACHPayment(achPayment); err != nil {
				return err
			}

			// Write associated records
			for _, addendum := range achPayment.Addenda {
				if err := w.writeACHAddendum(addendum); err != nil {
					return err
				}
			}

			for _, cars := range achPayment.CARSTASBETC {
				if err := w.writeCARSTASBETC(cars); err != nil {
					return err
				}
			}

			if achPayment.DNP != nil {
				if err := w.writeDNP(achPayment.DNP); err != nil {
					return err
				}
			}
		}
	}

	// Write schedule trailer
	return w.writeScheduleTrailer(schedule.Trailer)
}

// writeACHScheduleHeader writes an ACH schedule header record
func (w *Writer) writeACHScheduleHeader(header *ACHScheduleHeader) error {
	line := w.formatField(header.RecordCode, 2) +
		w.formatField(header.AgencyACHText, 4) +
		w.formatFieldRightJustified(header.ScheduleNumber, 14, '0') +
		w.formatField(header.PaymentTypeCode, 25) +
		w.formatField(header.StandardEntryClassCode, 3) +
		w.formatNumeric(header.AgencyLocationCode, 8) +
		w.formatField("", 1) + // Filler
		w.formatField(header.FederalEmployerIDNumber, 10) +
		w.formatField("", 783) // Filler

	return w.writeLine(line)
}

// writeACHPayment writes an ACH payment data record
func (w *Writer) writeACHPayment(payment *ACHPayment) error {
	line := w.formatField(payment.RecordCode, 2) +
		w.formatField(payment.AgencyAccountIdentifier, 16) +
		w.formatAmount(payment.Amount, 10) +
		w.formatField(payment.AgencyPaymentTypeCode, 1) +
		w.formatField(payment.IsTOP_Offset, 1) +
		w.formatField(payment.PayeeName, 35) +
		w.formatField(payment.PayeeAddressLine1, 35) +
		w.formatField(payment.PayeeAddressLine2, 35) +
		w.formatField(payment.CityName, 27) +
		w.formatField(payment.StateName, 10) +
		w.formatField(payment.StateCodeText, 2) +
		w.formatField(payment.PostalCode, 5) +
		w.formatField(payment.PostalCodeExtension, 5) +
		w.formatField(payment.CountryCodeText, 2) +
		w.formatField(payment.RoutingNumber, 9) +
		w.formatField(payment.AccountNumber, 17) +
		w.formatField(payment.ACH_TransactionCode, 2) +
		w.formatField(payment.PayeeIdentifierAdditional, 9) +
		w.formatField(payment.PayeeNameAdditional, 35) +
		w.formatField(payment.PaymentID, 20) +
		w.formatFieldNoJustify(payment.Reconcilement, 100) + // Don't justify reconcilement
		w.formatField(payment.TIN, 9) +
		w.formatField(payment.PaymentRecipientTINIndicator, 1) +
		w.formatField(payment.AdditionalPayeeTINIndicator, 1) +
		w.formatField(payment.AmountEligibleForOffset, 10) +
		w.formatField(payment.PayeeAddressLine3, 35) +
		w.formatField(payment.PayeeAddressLine4, 35) +
		w.formatField(payment.CountryName, 40) +
		w.formatField(payment.ConsularCode, 3) +
		w.formatField(payment.SubPaymentTypeCode, 32) +
		w.formatField(payment.PayerMechanism, 20) +
		w.formatField(payment.PaymentDescriptionCode, 2) +
		w.formatField("", 284) // Filler

	return w.writeLine(line)
}

// writeCheckSchedule writes a check schedule
func (w *Writer) writeCheckSchedule(schedule *CheckSchedule) error {
	// Write schedule header
	if err := w.writeCheckScheduleHeader(schedule.Header); err != nil {
		return err
	}

	// Write payments
	for _, payment := range schedule.Payments {
		if checkPayment, ok := payment.(*CheckPayment); ok {
			// Write payment record
			if err := w.writeCheckPayment(checkPayment); err != nil {
				return err
			}

			// Write associated records
			if checkPayment.Stub != nil {
				if err := w.writeCheckStub(checkPayment.Stub); err != nil {
					return err
				}
			}

			for _, cars := range checkPayment.CARSTASBETC {
				if err := w.writeCARSTASBETC(cars); err != nil {
					return err
				}
			}

			if checkPayment.DNP != nil {
				if err := w.writeDNP(checkPayment.DNP); err != nil {
					return err
				}
			}
		}
	}

	// Write schedule trailer
	return w.writeScheduleTrailer(schedule.Trailer)
}

// writeCheckScheduleHeader writes a check schedule header record
func (w *Writer) writeCheckScheduleHeader(header *CheckScheduleHeader) error {
	line := w.formatField(header.RecordCode, 2) +
		w.formatFieldRightJustified(header.ScheduleNumber, 14, '0') +
		w.formatField(header.PaymentTypeCode, 25) +
		w.formatNumeric(header.AgencyLocationCode, 8) +
		w.formatField("", 9) + // Filler
		w.formatField(header.CheckPaymentEnclosureCode, 10) +
		w.formatField("", 782) // Filler

	return w.writeLine(line)
}

// writeCheckPayment writes a check payment data record
func (w *Writer) writeCheckPayment(payment *CheckPayment) error {
	line := w.formatField(payment.RecordCode, 2) +
		w.formatField(payment.AgencyAccountIdentifier, 16) +
		w.formatAmount(payment.Amount, 10) +
		w.formatField(payment.AgencyPaymentTypeCode, 1) +
		w.formatField(payment.IsTOP_Offset, 1) +
		w.formatField(payment.PayeeName, 35) +
		w.formatField(payment.PayeeAddressLine1, 35) +
		w.formatField(payment.PayeeAddressLine2, 35) +
		w.formatField(payment.PayeeAddressLine3, 35) +
		w.formatField(payment.PayeeAddressLine4, 35) +
		w.formatField(payment.CityName, 27) +
		w.formatField(payment.StateName, 10) +
		w.formatField(payment.StateCodeText, 2) +
		w.formatField(payment.PostalCode, 5) +
		w.formatField(payment.PostalCodeExtension, 5) +
		w.formatField(payment.PostNetBarcodeDeliveryPoint, 3) +
		w.formatField("", 14) + // Filler
		w.formatField(payment.CountryName, 40) +
		w.formatField(payment.ConsularCode, 3) +
		w.formatField(payment.CheckLegendText1, 55) +
		w.formatField(payment.CheckLegendText2, 55) +
		w.formatField(payment.PayeeIdentifier_Secondary, 9) +
		w.formatField(payment.PartyName_Secondary, 35) +
		w.formatField(payment.PaymentID, 20) +
		w.formatFieldNoJustify(payment.Reconcilement, 100) + // Don't justify reconcilement
		w.formatField(payment.SpecialHandling, 50) +
		w.formatField(payment.TIN, 9) +
		w.formatField(payment.USPSIntelligentMailBarcode, 50) +
		w.formatField(payment.PaymentRecipientTINIndicator, 1) +
		w.formatField(payment.SecondaryPayeeTINIndicator, 1) +
		w.formatField(payment.AmountEligibleForOffset, 10) +
		w.formatField(payment.SubPaymentTypeCode, 32) +
		w.formatField(payment.PayerMechanism, 20) +
		w.formatField(payment.PaymentDescriptionCode, 2) +
		w.formatField("", 87) // Filler

	return w.writeLine(line)
}

// writeACHAddendum writes an ACH addendum record
func (w *Writer) writeACHAddendum(addendum *ACHAddendum) error {
	var line string

	if addendum.RecordCode == "03" {
		line = w.formatField(addendum.RecordCode, 2) +
			w.formatField(addendum.PaymentID, 20) +
			w.formatField(addendum.AddendaInformation, 80) +
			w.formatField("", 748) // Filler
	} else if addendum.RecordCode == "04" { // CTX
		line = w.formatField(addendum.RecordCode, 2) +
			w.formatField(addendum.PaymentID, 20) +
			w.formatField(addendum.AddendaInformation, 800) +
			w.formatField("", 28) // Filler
	} else {
		return fmt.Errorf("invalid addendum record code: %s", addendum.RecordCode)
	}

	return w.writeLine(line)
}

// writeCARSTASBETC writes a CARS TAS/BETC record
func (w *Writer) writeCARSTASBETC(cars *CARSTASBETC) error {
	line := w.formatField(cars.RecordCode, 2) +
		w.formatField(cars.PaymentID, 20) +
		w.formatField(cars.SubLevelPrefixCode, 2) +
		w.formatField(cars.AllocationTransferAgencyID, 3) +
		w.formatField(cars.AgencyIdentifier, 3) +
		w.formatField(cars.BeginningPeriodOfAvailability, 4) +
		w.formatField(cars.EndingPeriodOfAvailability, 4) +
		w.formatField(cars.AvailabilityTypeCode, 1) +
		w.formatField(cars.MainAccountCode, 4) +
		w.formatField(cars.SubAccountCode, 3) +
		w.formatField(cars.BusinessEventTypeCode, 8) +
		w.formatAmount(cars.AccountClassificationAmount, 10) +
		w.formatField(cars.IsCredit, 1) +
		w.formatField("", 785) // Filler

	return w.writeLine(line)
}

// writeCheckStub writes a check stub record
func (w *Writer) writeCheckStub(stub *CheckStub) error {
	line := w.formatField(stub.RecordCode, 2) +
		w.formatField(stub.PaymentID, 20)

	for i := 0; i < 14; i++ {
		line += w.formatField(stub.PaymentIdentificationLines[i], 55)
	}

	line += w.formatField("", 58) // Filler

	return w.writeLine(line)
}

// writeDNP writes a DNP record
func (w *Writer) writeDNP(dnp *DNPRecord) error {
	line := w.formatField(dnp.RecordCode, 2) +
		w.formatField(dnp.PaymentID, 20) +
		w.formatFieldNoJustify(dnp.DNPDetail, 766) + // Don't justify DNP detail
		w.formatField("", 62) // Filler

	return w.writeLine(line)
}

// writeScheduleTrailer writes a schedule trailer record
func (w *Writer) writeScheduleTrailer(trailer *ScheduleTrailer) error {
	line := w.formatField(trailer.RecordCode, 2) +
		w.formatField("", 10) + // Filler
		w.formatNumeric(fmt.Sprintf("%d", trailer.ScheduleCount), 8) +
		w.formatField("", 3) + // Filler
		w.formatAmount(trailer.ScheduleAmount, 15) +
		w.formatField("", 812) // Filler

	return w.writeLine(line)
}

// writeFileTrailer writes the file trailer record
func (w *Writer) writeFileTrailer(trailer *FileTrailer) error {
	line := w.formatField(trailer.RecordCode, 2) +
		w.formatNumeric(fmt.Sprintf("%d", trailer.TotalCountRecords), 18) +
		w.formatNumeric(fmt.Sprintf("%d", trailer.TotalCountPayments), 18) +
		w.formatAmount(trailer.TotalAmountPayments, 18) +
		w.formatField("", 794) // Filler

	return w.writeLine(line)
}

// Helper methods for field formatting

// formatField formats a field with left justification and blank padding
func (w *Writer) formatField(value string, length int) string {
	if len(value) > length {
		return value[:length]
	}
	return value + strings.Repeat(" ", length-len(value))
}

// formatFieldRightJustified formats a field with right justification
func (w *Writer) formatFieldRightJustified(value string, length int, padChar rune) string {
	value = strings.TrimSpace(value)
	if len(value) > length {
		return value[:length]
	}
	return strings.Repeat(string(padChar), length-len(value)) + value
}

// formatFieldNoJustify returns the field value as-is, padded to length
func (w *Writer) formatFieldNoJustify(value string, length int) string {
	if len(value) > length {
		return value[:length]
	}
	if len(value) < length {
		return value + strings.Repeat(" ", length-len(value))
	}
	return value
}

// formatNumeric formats a numeric field with right justification and zero padding
func (w *Writer) formatNumeric(value string, length int) string {
	// Remove non-numeric characters using strings.Builder for efficiency
	var builder strings.Builder
	builder.Grow(len(value)) // Pre-allocate capacity
	for _, r := range value {
		if r >= '0' && r <= '9' {
			builder.WriteRune(r)
		}
	}
	numeric := builder.String()

	if len(numeric) > length {
		return numeric[:length]
	}
	return strings.Repeat("0", length-len(numeric)) + numeric
}

// formatAmount formats an amount field (in cents) with right justification and zero padding
func (w *Writer) formatAmount(cents int64, length int) string {
	return w.formatNumeric(fmt.Sprintf("%d", cents), length)
}

// writeLine writes a line to the output
func (w *Writer) writeLine(line string) error {
	if len(line) != RecordLength {
		return fmt.Errorf("invalid line length: expected %d, got %d", RecordLength, len(line))
	}

	_, err := fmt.Fprintln(w.w, line)
	return err
}
