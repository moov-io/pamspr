package pamspr

import (
	"fmt"
)

// CommonParser handles parsing of shared record types (CARS/TAS/BETC, DNP, Schedule Trailer)
type CommonParser struct {
	validator *Validator
}

// NewCommonParser creates a new common parser
func NewCommonParser(validator *Validator) *CommonParser {
	return &CommonParser{
		validator: validator,
	}
}

// ParseCARSTASBETC parses a CARS/TAS/BETC record ("G ")
func (p *CommonParser) ParseCARSTASBETC(line string) (*CARSTASBETC, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("G ")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for CARS/TAS/BETC record")
	}

	record := &CARSTASBETC{
		RecordCode:                    extractField(line, fields["RecordCode"]),
		PaymentID:                     extractFieldTrimmed(line, fields["PaymentID"]),
		SubLevelPrefixCode:            extractFieldTrimmed(line, fields["SubLevelPrefixCode"]),
		AllocationTransferAgencyID:    extractFieldTrimmed(line, fields["AllocationTransferAgencyID"]),
		AgencyIdentifier:              extractFieldTrimmed(line, fields["AgencyIdentifier"]),
		BeginningPeriodOfAvailability: extractFieldTrimmed(line, fields["BeginningPeriodOfAvailability"]),
		EndingPeriodOfAvailability:    extractFieldTrimmed(line, fields["EndingPeriodOfAvailability"]),
		AvailabilityTypeCode:          extractFieldTrimmed(line, fields["AvailabilityTypeCode"]),
		MainAccountCode:               extractFieldTrimmed(line, fields["MainAccountCode"]),
		SubAccountCode:                extractFieldTrimmed(line, fields["SubAccountCode"]),
		BusinessEventTypeCode:         extractFieldTrimmed(line, fields["BusinessEventTypeCode"]),
		AccountClassificationAmount:   parseAmount(extractField(line, fields["AccountClassificationAmount"])),
		IsCredit:                      extractFieldTrimmed(line, fields["IsCredit"]),
	}

	return record, nil
}

// ParseDNP parses a DNP record ("DD")
func (p *CommonParser) ParseDNP(line string) (*DNPRecord, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("DD")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for DNP record")
	}

	record := &DNPRecord{
		RecordCode: extractField(line, fields["RecordCode"]),
		PaymentID:  extractFieldTrimmed(line, fields["PaymentID"]),
		DNPDetail:  extractFieldTrimmed(line, fields["DNPDetail"]),
	}

	return record, nil
}

// ParseScheduleTrailer parses a schedule trailer record ("T ")
func (p *CommonParser) ParseScheduleTrailer(line string) (*ScheduleTrailer, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("T ")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for schedule trailer")
	}

	trailer := &ScheduleTrailer{
		RecordCode:     extractField(line, fields["RecordCode"]),
		ScheduleCount:  parseAmount(extractField(line, fields["ScheduleCount"])),
		ScheduleAmount: parseAmount(extractField(line, fields["ScheduleAmount"])),
	}

	return trailer, nil
}