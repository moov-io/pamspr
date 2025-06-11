package pamspr

import (
	"fmt"
	"strings"
)

// CheckParser handles parsing of check schedule and payment records
type CheckParser struct {
	validator *Validator
}

// NewCheckParser creates a new check parser
func NewCheckParser(validator *Validator) *CheckParser {
	return &CheckParser{
		validator: validator,
	}
}

// ParseCheckScheduleHeader parses a check schedule header record ("11")
func (p *CheckParser) ParseCheckScheduleHeader(line string) (*CheckScheduleHeader, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("11")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for check schedule header")
	}

	header := &CheckScheduleHeader{
		RecordCode:                extractField(line, fields["RecordCode"]),
		ScheduleNumber:            extractField(line, fields["ScheduleNumber"]),
		PaymentTypeCode:           extractField(line, fields["PaymentTypeCode"]),
		AgencyLocationCode:        extractField(line, fields["AgencyLocationCode"]),
		CheckPaymentEnclosureCode: strings.TrimSpace(extractField(line, fields["CheckPaymentEnclosureCode"])),
	}

	return header, nil
}

// ParseCheckPayment parses a check payment record ("12")
func (p *CheckParser) ParseCheckPayment(line string) (*CheckPayment, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("12")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for check payment")
	}

	payment := &CheckPayment{
		RecordCode:                   extractField(line, fields["RecordCode"]),
		AgencyAccountIdentifier:      extractField(line, fields["AgencyAccountIdentifier"]),
		Amount:                       parseAmount(extractField(line, fields["Amount"])),
		AgencyPaymentTypeCode:        extractField(line, fields["AgencyPaymentTypeCode"]),
		IsTOPOffset:                  extractField(line, fields["IsTOPOffset"]),
		PayeeName:                    extractField(line, fields["PayeeName"]),
		PayeeAddressLine1:            extractField(line, fields["PayeeAddressLine1"]),
		PayeeAddressLine2:            extractField(line, fields["PayeeAddressLine2"]),
		PayeeAddressLine3:            extractField(line, fields["PayeeAddressLine3"]),
		PayeeAddressLine4:            extractField(line, fields["PayeeAddressLine4"]),
		CityName:                     extractField(line, fields["CityName"]),
		StateName:                    extractField(line, fields["StateName"]),
		StateCodeText:                extractField(line, fields["StateCodeText"]),
		PostalCode:                   extractField(line, fields["PostalCode"]),
		PostalCodeExtension:          extractField(line, fields["PostalCodeExtension"]),
		PostNetBarcodeDeliveryPoint:  extractField(line, fields["PostNetBarcodeDeliveryPoint"]),
		CountryName:                  extractField(line, fields["CountryName"]),
		ConsularCode:                 extractField(line, fields["ConsularCode"]),
		CheckLegendText1:             extractField(line, fields["CheckLegendText1"]),
		CheckLegendText2:             extractField(line, fields["CheckLegendText2"]),
		PayeeIdentifierSecondary:     extractField(line, fields["PayeeIdentifierSecondary"]),
		PartyNameSecondary:           extractField(line, fields["PartyNameSecondary"]),
		PaymentID:                    extractField(line, fields["PaymentID"]),
		Reconcilement:                extractField(line, fields["Reconcilement"]),
		SpecialHandling:              extractField(line, fields["SpecialHandling"]),
		TIN:                          extractField(line, fields["TIN"]),
		USPSIntelligentMailBarcode:   extractField(line, fields["USPSIntelligentMailBarcode"]),
		PaymentRecipientTINIndicator: extractField(line, fields["PaymentRecipientTINIndicator"]),
		SecondaryPayeeTINIndicator:   extractField(line, fields["SecondaryPayeeTINIndicator"]),
		AmountEligibleForOffset:      extractField(line, fields["AmountEligibleForOffset"]),
		SubPaymentTypeCode:           extractField(line, fields["SubPaymentTypeCode"]),
		PayerMechanism:               extractField(line, fields["PayerMechanism"]),
		PaymentDescriptionCode:       extractField(line, fields["PaymentDescriptionCode"]),
		CARSTASBETC:                  make([]*CARSTASBETC, 0),
	}

	// Note: Validation is handled at the Reader level, not in individual parsers
	// This maintains backward compatibility with the original parsing behavior

	return payment, nil
}

// ParseCheckStub parses a check stub record ("13")
func (p *CheckParser) ParseCheckStub(line string) (*CheckStub, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("13")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for check stub")
	}

	stub := &CheckStub{
		RecordCode: extractField(line, fields["RecordCode"]),
		PaymentID:  extractField(line, fields["PaymentID"]),
	}

	// Parse 14 payment identification lines
	lineFields := []string{"Line1", "Line2", "Line3", "Line4", "Line5", "Line6", "Line7",
		"Line8", "Line9", "Line10", "Line11", "Line12", "Line13", "Line14"}

	for i, lineField := range lineFields {
		if field, exists := fields[lineField]; exists {
			stub.PaymentIdentificationLines[i] = extractField(line, field)
		}
	}

	return stub, nil
}
