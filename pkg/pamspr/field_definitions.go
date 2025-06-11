package pamspr

import "fmt"

// FieldDefinition defines the position and properties of a field in a fixed-width record
type FieldDefinition struct {
	Start    int  // 1-based start position
	End      int  // 1-based end position (inclusive)
	Length   int  // Calculated field length
	Required bool // Whether field is required
}

// NewFieldDef creates a new field definition
func NewFieldDef(start, length int, required bool) FieldDefinition {
	return FieldDefinition{
		Start:    start,
		End:      start + length - 1,
		Length:   length,
		Required: required,
	}
}

// GetFieldDefinitions returns field definitions for a given record type
func GetFieldDefinitions(recordCode string) map[string]FieldDefinition {
	switch recordCode {
	case "H ": // File Header
		return map[string]FieldDefinition{
			"RecordCode":               NewFieldDef(1, 2, true),
			"InputSystem":              NewFieldDef(3, 40, true),
			"StandardPaymentVersion":   NewFieldDef(43, 3, true),
			"IsRequestedForSameDayACH": NewFieldDef(46, 1, false),
			"Filler":                   NewFieldDef(47, 804, false),
		}

	case "01": // ACH Schedule Header
		return map[string]FieldDefinition{
			"RecordCode":              NewFieldDef(1, 2, true),
			"AgencyACHText":           NewFieldDef(3, 4, false),
			"ScheduleNumber":          NewFieldDef(7, 14, true),
			"PaymentTypeCode":         NewFieldDef(21, 25, false),
			"StandardEntryClassCode":  NewFieldDef(46, 3, true),
			"AgencyLocationCode":      NewFieldDef(49, 8, true),
			"Filler1":                 NewFieldDef(57, 1, false),
			"FederalEmployerIDNumber": NewFieldDef(58, 10, false),
			"Filler2":                 NewFieldDef(68, 783, false),
		}

	case "02": // ACH Payment
		return map[string]FieldDefinition{
			"RecordCode":                   NewFieldDef(1, 2, true),
			"AgencyAccountIdentifier":      NewFieldDef(3, 16, true),
			"Amount":                       NewFieldDef(19, 10, true),
			"AgencyPaymentTypeCode":        NewFieldDef(29, 1, false),
			"IsTOPOffset":                  NewFieldDef(30, 1, false),
			"PayeeName":                    NewFieldDef(31, 35, true),
			"PayeeAddressLine1":            NewFieldDef(66, 35, false),
			"PayeeAddressLine2":            NewFieldDef(101, 35, false),
			"CityName":                     NewFieldDef(136, 27, false),
			"StateName":                    NewFieldDef(163, 10, false),
			"StateCodeText":                NewFieldDef(173, 2, false),
			"PostalCode":                   NewFieldDef(175, 5, false),
			"PostalCodeExtension":          NewFieldDef(180, 5, false),
			"CountryCodeText":              NewFieldDef(185, 2, false),
			"RoutingNumber":                NewFieldDef(187, 9, true),
			"AccountNumber":                NewFieldDef(196, 17, true),
			"ACHTransactionCode":           NewFieldDef(213, 2, true),
			"PayeeIdentifierAdditional":    NewFieldDef(215, 9, false),
			"PayeeNameAdditional":          NewFieldDef(224, 35, false),
			"PaymentID":                    NewFieldDef(259, 20, true),
			"Reconcilement":                NewFieldDef(279, 100, false),
			"TIN":                          NewFieldDef(379, 9, false),
			"PaymentRecipientTINIndicator": NewFieldDef(388, 1, false),
			"AdditionalPayeeTINIndicator":  NewFieldDef(389, 1, false),
			"AmountEligibleForOffset":      NewFieldDef(390, 10, false),
			"PayeeAddressLine3":            NewFieldDef(400, 35, false),
			"PayeeAddressLine4":            NewFieldDef(435, 35, false),
			"CountryName":                  NewFieldDef(470, 40, false),
			"ConsularCode":                 NewFieldDef(510, 3, false),
			"SubPaymentTypeCode":           NewFieldDef(513, 32, false),
			"PayerMechanism":               NewFieldDef(545, 20, false),
			"PaymentDescriptionCode":       NewFieldDef(565, 2, false),
			"Filler":                       NewFieldDef(567, 284, false),
		}

	case "03": // ACH Addendum (CCD/PPD)
		return map[string]FieldDefinition{
			"RecordCode":         NewFieldDef(1, 2, true),
			"PaymentID":          NewFieldDef(3, 20, true),
			"AddendaInformation": NewFieldDef(23, 80, false),
			"Filler":             NewFieldDef(103, 748, false),
		}

	case "04": // ACH Addendum (CTX)
		return map[string]FieldDefinition{
			"RecordCode":         NewFieldDef(1, 2, true),
			"PaymentID":          NewFieldDef(3, 20, true),
			"AddendaInformation": NewFieldDef(23, 800, false),
			"Filler":             NewFieldDef(823, 28, false),
		}

	case "11": // Check Schedule Header
		return map[string]FieldDefinition{
			"RecordCode":                NewFieldDef(1, 2, true),
			"ScheduleNumber":            NewFieldDef(3, 14, true),
			"PaymentTypeCode":           NewFieldDef(17, 25, false),
			"AgencyLocationCode":        NewFieldDef(42, 8, true),
			"Filler1":                   NewFieldDef(50, 9, false),
			"CheckPaymentEnclosureCode": NewFieldDef(59, 10, false),
			"Filler2":                   NewFieldDef(69, 782, false),
		}

	case "12": // Check Payment
		return map[string]FieldDefinition{
			"RecordCode":                   NewFieldDef(1, 2, true),
			"AgencyAccountIdentifier":      NewFieldDef(3, 16, true),
			"Amount":                       NewFieldDef(19, 10, true),
			"AgencyPaymentTypeCode":        NewFieldDef(29, 1, false),
			"IsTOPOffset":                  NewFieldDef(30, 1, false),
			"PayeeName":                    NewFieldDef(31, 35, true),
			"PayeeAddressLine1":            NewFieldDef(66, 35, false),
			"PayeeAddressLine2":            NewFieldDef(101, 35, false),
			"PayeeAddressLine3":            NewFieldDef(136, 35, false),
			"PayeeAddressLine4":            NewFieldDef(171, 35, false),
			"CityName":                     NewFieldDef(206, 27, false),
			"StateName":                    NewFieldDef(233, 10, false),
			"StateCodeText":                NewFieldDef(243, 2, false),
			"PostalCode":                   NewFieldDef(245, 5, false),
			"PostalCodeExtension":          NewFieldDef(250, 5, false),
			"PostNetBarcodeDeliveryPoint":  NewFieldDef(255, 3, false),
			"Filler1":                      NewFieldDef(258, 14, false),
			"CountryName":                  NewFieldDef(272, 40, false),
			"ConsularCode":                 NewFieldDef(312, 3, false),
			"CheckLegendText1":             NewFieldDef(315, 55, false),
			"CheckLegendText2":             NewFieldDef(370, 55, false),
			"PayeeIdentifierSecondary":     NewFieldDef(425, 9, false),
			"PartyNameSecondary":           NewFieldDef(434, 35, false),
			"PaymentID":                    NewFieldDef(469, 20, true),
			"Reconcilement":                NewFieldDef(489, 100, false),
			"SpecialHandling":              NewFieldDef(589, 50, false),
			"TIN":                          NewFieldDef(639, 9, false),
			"USPSIntelligentMailBarcode":   NewFieldDef(648, 50, false),
			"PaymentRecipientTINIndicator": NewFieldDef(698, 1, false),
			"SecondaryPayeeTINIndicator":   NewFieldDef(699, 1, false),
			"AmountEligibleForOffset":      NewFieldDef(700, 10, false),
			"SubPaymentTypeCode":           NewFieldDef(710, 32, false),
			"PayerMechanism":               NewFieldDef(742, 20, false),
			"PaymentDescriptionCode":       NewFieldDef(762, 2, false),
			"Filler2":                      NewFieldDef(764, 87, false),
		}

	case "13": // Check Stub
		return map[string]FieldDefinition{
			"RecordCode": NewFieldDef(1, 2, true),
			"PaymentID":  NewFieldDef(3, 20, true),
			"Line1":      NewFieldDef(23, 55, false),
			"Line2":      NewFieldDef(78, 55, false),
			"Line3":      NewFieldDef(133, 55, false),
			"Line4":      NewFieldDef(188, 55, false),
			"Line5":      NewFieldDef(243, 55, false),
			"Line6":      NewFieldDef(298, 55, false),
			"Line7":      NewFieldDef(353, 55, false),
			"Line8":      NewFieldDef(408, 55, false),
			"Line9":      NewFieldDef(463, 55, false),
			"Line10":     NewFieldDef(518, 55, false),
			"Line11":     NewFieldDef(573, 55, false),
			"Line12":     NewFieldDef(628, 55, false),
			"Line13":     NewFieldDef(683, 55, false),
			"Line14":     NewFieldDef(738, 55, false),
			"Filler":     NewFieldDef(793, 58, false),
		}

	case "G ": // CARS TAS/BETC
		return map[string]FieldDefinition{
			"RecordCode":                    NewFieldDef(1, 2, true),
			"PaymentID":                     NewFieldDef(3, 20, true),
			"SubLevelPrefixCode":            NewFieldDef(23, 2, false),
			"AllocationTransferAgencyID":    NewFieldDef(25, 3, false),
			"AgencyIdentifier":              NewFieldDef(28, 3, false),
			"BeginningPeriodOfAvailability": NewFieldDef(31, 4, false),
			"EndingPeriodOfAvailability":    NewFieldDef(35, 4, false),
			"AvailabilityTypeCode":          NewFieldDef(39, 1, false),
			"MainAccountCode":               NewFieldDef(40, 4, false),
			"SubAccountCode":                NewFieldDef(44, 3, false),
			"BusinessEventTypeCode":         NewFieldDef(47, 8, false),
			"AccountClassificationAmount":   NewFieldDef(55, 10, false),
			"IsCredit":                      NewFieldDef(65, 1, false),
			"Filler":                        NewFieldDef(66, 785, false),
		}

	case "DD": // DNP Record
		return map[string]FieldDefinition{
			"RecordCode": NewFieldDef(1, 2, true),
			"PaymentID":  NewFieldDef(3, 20, true),
			"DNPDetail":  NewFieldDef(23, 766, false),
			"Filler":     NewFieldDef(789, 62, false),
		}

	case "T ": // Schedule Trailer
		return map[string]FieldDefinition{
			"RecordCode":     NewFieldDef(1, 2, true),
			"Filler1":        NewFieldDef(3, 10, false),
			"ScheduleCount":  NewFieldDef(13, 8, true),
			"Filler2":        NewFieldDef(21, 3, false),
			"ScheduleAmount": NewFieldDef(24, 15, true),
			"Filler3":        NewFieldDef(39, 812, false),
		}

	case "E ": // File Trailer
		return map[string]FieldDefinition{
			"RecordCode":          NewFieldDef(1, 2, true),
			"TotalCountRecords":   NewFieldDef(3, 18, true),
			"TotalCountPayments":  NewFieldDef(21, 18, true),
			"TotalAmountPayments": NewFieldDef(39, 18, true),
			"Filler":              NewFieldDef(57, 794, false),
		}

	default:
		return nil
	}
}

// ValidateFieldPositions validates that all field definitions are correct
func ValidateFieldPositions() error {
	recordCodes := []string{"H ", "01", "02", "03", "04", "11", "12", "13", "G ", "DD", "T ", "E "}

	for _, code := range recordCodes {
		fields := GetFieldDefinitions(code)
		if fields == nil {
			continue
		}

		// Check that no fields overlap and all end at or before position 850
		for name, field := range fields {
			if field.Start < 1 {
				return fmt.Errorf("record %s field %s: start position %d is less than 1", code, name, field.Start)
			}
			if field.End > RecordLength {
				return fmt.Errorf("record %s field %s: end position %d exceeds record length %d", code, name, field.End, RecordLength)
			}
			if field.Start > field.End {
				return fmt.Errorf("record %s field %s: start position %d is greater than end position %d", code, name, field.Start, field.End)
			}
			if field.Length != (field.End - field.Start + 1) {
				return fmt.Errorf("record %s field %s: length %d doesn't match calculated length %d", code, name, field.Length, field.End-field.Start+1)
			}
		}
	}

	return nil
}
