package pamspr

import (
	"fmt"
)

// ACHParser handles parsing of ACH schedule and payment records
type ACHParser struct {
	validator *Validator
}

// NewACHParser creates a new ACH parser
func NewACHParser(validator *Validator) *ACHParser {
	return &ACHParser{
		validator: validator,
	}
}

// ParseACHScheduleHeader parses an ACH schedule header record ("01")
func (p *ACHParser) ParseACHScheduleHeader(line string) (*ACHScheduleHeader, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("01")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for ACH schedule header")
	}

	header := &ACHScheduleHeader{
		RecordCode:              extractField(line, fields["RecordCode"]),
		AgencyACHText:           extractField(line, fields["AgencyACHText"]),
		ScheduleNumber:          extractField(line, fields["ScheduleNumber"]),
		PaymentTypeCode:         extractField(line, fields["PaymentTypeCode"]),
		StandardEntryClassCode:  extractField(line, fields["StandardEntryClassCode"]),
		AgencyLocationCode:      extractField(line, fields["AgencyLocationCode"]),
		FederalEmployerIDNumber: extractField(line, fields["FederalEmployerIDNumber"]),
	}

	return header, nil
}

// ParseACHPayment parses an ACH payment record ("02")
func (p *ACHParser) ParseACHPayment(line string) (*ACHPayment, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("02")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for ACH payment")
	}

	payment := &ACHPayment{
		RecordCode:                   extractField(line, fields["RecordCode"]),
		AgencyAccountIdentifier:      extractField(line, fields["AgencyAccountIdentifier"]),
		Amount:                       parseAmount(extractField(line, fields["Amount"])),
		AgencyPaymentTypeCode:        extractField(line, fields["AgencyPaymentTypeCode"]),
		IsTOP_Offset:                 extractField(line, fields["IsTOP_Offset"]),
		PayeeName:                    extractField(line, fields["PayeeName"]),
		PayeeAddressLine1:            extractField(line, fields["PayeeAddressLine1"]),
		PayeeAddressLine2:            extractField(line, fields["PayeeAddressLine2"]),
		CityName:                     extractField(line, fields["CityName"]),
		StateName:                    extractField(line, fields["StateName"]),
		StateCodeText:                extractField(line, fields["StateCodeText"]),
		PostalCode:                   extractField(line, fields["PostalCode"]),
		PostalCodeExtension:          extractField(line, fields["PostalCodeExtension"]),
		CountryCodeText:              extractField(line, fields["CountryCodeText"]),
		RoutingNumber:                extractField(line, fields["RoutingNumber"]),
		AccountNumber:                extractField(line, fields["AccountNumber"]),
		ACH_TransactionCode:          extractField(line, fields["ACH_TransactionCode"]),
		PayeeIdentifierAdditional:    extractField(line, fields["PayeeIdentifierAdditional"]),
		PayeeNameAdditional:          extractField(line, fields["PayeeNameAdditional"]),
		PaymentID:                    extractField(line, fields["PaymentID"]),
		Reconcilement:                extractField(line, fields["Reconcilement"]),
		TIN:                          extractField(line, fields["TIN"]),
		PaymentRecipientTINIndicator: extractField(line, fields["PaymentRecipientTINIndicator"]),
		AdditionalPayeeTINIndicator:  extractField(line, fields["AdditionalPayeeTINIndicator"]),
		AmountEligibleForOffset:      extractField(line, fields["AmountEligibleForOffset"]),
		PayeeAddressLine3:            extractField(line, fields["PayeeAddressLine3"]),
		PayeeAddressLine4:            extractField(line, fields["PayeeAddressLine4"]),
		CountryName:                  extractField(line, fields["CountryName"]),
		ConsularCode:                 extractField(line, fields["ConsularCode"]),
		SubPaymentTypeCode:           extractField(line, fields["SubPaymentTypeCode"]),
		PayerMechanism:               extractField(line, fields["PayerMechanism"]),
		PaymentDescriptionCode:       extractField(line, fields["PaymentDescriptionCode"]),
		Addenda:                      make([]*ACHAddendum, 0),
		CARSTASBETC:                  make([]*CARSTASBETC, 0),
	}

	// Note: Validation is handled at the Reader level, not in individual parsers
	// This maintains backward compatibility with the original parsing behavior

	return payment, nil
}

// ParseACHAddendum parses an ACH addendum record ("03" or "04")
func (p *ACHParser) ParseACHAddendum(line string) (*ACHAddendum, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	recordCode := extractFieldByPosition(line, 1, 2)

	var fields map[string]FieldDefinition
	switch recordCode {
	case "03":
		fields = GetFieldDefinitions("03")
	case "04":
		fields = GetFieldDefinitions("04")
	default:
		return nil, fmt.Errorf("invalid addendum record code: %s", recordCode)
	}

	if fields == nil {
		return nil, fmt.Errorf("no field definitions for addendum record %s", recordCode)
	}

	addendum := &ACHAddendum{
		RecordCode:         extractField(line, fields["RecordCode"]),
		PaymentID:          extractField(line, fields["PaymentID"]),
		AddendaInformation: extractField(line, fields["AddendaInformation"]),
	}

	return addendum, nil
}

// Helper function for extracting fields by position (legacy support)
func extractFieldByPosition(line string, start, end int) string {
	field := FieldDefinition{Start: start, End: end}
	value, err := SecureExtractField(line, field, "field", DefaultSecurityConfig())
	if err != nil {
		// TODO: Add proper error handling/logging
		return ""
	}
	return value
}
