package pamspr

import (
	"fmt"
	"strings"
)

// Note: RecordLength constant is defined in file.go

// FieldDefinition defines a field's position and formatting rules
type FieldDefinition struct {
	Start     int                    // 1-based start position
	End       int                    // 1-based end position (inclusive)
	Length    int                    // Calculated field length
	Required  bool                   // Whether field is required
	Validator func(string) error     // Optional validation function
	Formatter func(string, int) string // Optional custom formatter
}

// NewFieldDef creates a field definition with calculated length
func NewFieldDef(start, end int, required bool) FieldDefinition {
	return FieldDefinition{
		Start:    start,
		End:      end,
		Length:   end - start + 1,
		Required: required,
	}
}

// File Header Record ("H ") Field Positions
var FileHeaderFields = map[string]FieldDefinition{
	"RecordCode":               NewFieldDef(1, 2, true),    // "H "
	"InputSystem":              NewFieldDef(3, 42, true),   // 40 chars
	"StandardPaymentVersion":   NewFieldDef(43, 45, true),  // "502"
	"IsRequestedForSameDayACH": NewFieldDef(46, 46, false), // "0", "1", or blank
	"Filler":                   NewFieldDef(47, 850, false), // Remaining space
}

// ACH Schedule Header Record ("01") Field Positions
var ACHScheduleHeaderFields = map[string]FieldDefinition{
	"RecordCode":              NewFieldDef(1, 2, true),     // "01"
	"AgencyACHText":           NewFieldDef(3, 6, true),     // 4 chars
	"ScheduleNumber":          NewFieldDef(7, 20, true),    // 14 chars, right-justified
	"PaymentTypeCode":         NewFieldDef(21, 45, true),   // 25 chars
	"StandardEntryClassCode":  NewFieldDef(46, 48, true),   // 3 chars
	"AgencyLocationCode":      NewFieldDef(49, 56, true),   // 8 chars, numeric
	"Filler1":                 NewFieldDef(57, 57, false),  // 1 char
	"FederalEmployerIDNumber": NewFieldDef(58, 67, true),   // 10 chars
	"Filler2":                 NewFieldDef(68, 850, false), // Remaining space
}

// ACH Payment Record ("02") Field Positions
var ACHPaymentFields = map[string]FieldDefinition{
	"RecordCode":                   NewFieldDef(1, 2, true),     // "02"
	"AgencyAccountIdentifier":      NewFieldDef(3, 18, true),    // 16 chars
	"Amount":                       NewFieldDef(19, 28, true),   // 10 digits, amount in cents
	"AgencyPaymentTypeCode":        NewFieldDef(29, 29, true),   // 1 char
	"IsTOPOffset":                  NewFieldDef(30, 30, true),   // "0" or "1"
	"PayeeName":                    NewFieldDef(31, 65, true),   // 35 chars
	"PayeeAddressLine1":            NewFieldDef(66, 100, false), // 35 chars
	"PayeeAddressLine2":            NewFieldDef(101, 135, false), // 35 chars
	"CityName":                     NewFieldDef(136, 162, false), // 27 chars
	"StateName":                    NewFieldDef(163, 172, false), // 10 chars
	"StateCodeText":                NewFieldDef(173, 174, false), // 2 chars
	"PostalCode":                   NewFieldDef(175, 179, false), // 5 chars
	"PostalCodeExtension":          NewFieldDef(180, 184, false), // 5 chars
	"CountryCodeText":              NewFieldDef(185, 186, false), // 2 chars
	"RoutingNumber":                NewFieldDef(187, 195, true),  // 9 chars
	"AccountNumber":                NewFieldDef(196, 212, true),  // 17 chars
	"ACHTransactionCode":           NewFieldDef(213, 214, true),  // 2 chars
	"PayeeIdentifierAdditional":    NewFieldDef(215, 223, false), // 9 chars
	"PayeeNameAdditional":          NewFieldDef(224, 258, false), // 35 chars
	"PaymentID":                    NewFieldDef(259, 278, true),  // 20 chars
	"Reconcilement":                NewFieldDef(279, 378, false), // 100 chars
	"TIN":                          NewFieldDef(379, 387, false), // 9 chars
	"PaymentRecipientTINIndicator": NewFieldDef(388, 388, false), // 1 char
	"AdditionalPayeeTINIndicator":  NewFieldDef(389, 389, false), // 1 char
	"AmountEligibleForOffset":      NewFieldDef(390, 399, false), // 10 digits
	"PayeeAddressLine3":            NewFieldDef(400, 434, false), // 35 chars
	"PayeeAddressLine4":            NewFieldDef(435, 469, false), // 35 chars
	"CountryName":                  NewFieldDef(470, 509, false), // 40 chars
	"ConsularCode":                 NewFieldDef(510, 512, false), // 3 chars
	"SubPaymentTypeCode":           NewFieldDef(513, 544, false), // 32 chars
	"PayerMechanism":               NewFieldDef(545, 564, false), // 20 chars
	"PaymentDescriptionCode":       NewFieldDef(565, 566, false), // 2 chars
	"Filler":                       NewFieldDef(567, 850, false), // Remaining space
}

// Check Schedule Header Record ("11") Field Positions
var CheckScheduleHeaderFields = map[string]FieldDefinition{
	"RecordCode":                NewFieldDef(1, 2, true),     // "11"
	"ScheduleNumber":            NewFieldDef(3, 16, true),    // 14 chars, right-justified
	"PaymentTypeCode":           NewFieldDef(17, 41, true),   // 25 chars
	"AgencyLocationCode":        NewFieldDef(42, 49, true),   // 8 chars, numeric
	"Filler1":                   NewFieldDef(50, 58, false),  // 9 chars
	"CheckPaymentEnclosureCode": NewFieldDef(59, 68, false),  // 10 chars
	"Filler2":                   NewFieldDef(69, 850, false), // Remaining space
}

// Check Payment Record ("12") Field Positions
var CheckPaymentFields = map[string]FieldDefinition{
	"RecordCode":                   NewFieldDef(1, 2, true),     // "12"
	"AgencyAccountIdentifier":      NewFieldDef(3, 18, true),    // 16 chars
	"Amount":                       NewFieldDef(19, 28, true),   // 10 digits, amount in cents
	"AgencyPaymentTypeCode":        NewFieldDef(29, 29, true),   // 1 char
	"IsTOPOffset":                  NewFieldDef(30, 30, true),   // "0" or "1"
	"PayeeName":                    NewFieldDef(31, 65, true),   // 35 chars
	"PayeeAddressLine1":            NewFieldDef(66, 100, false), // 35 chars
	"PayeeAddressLine2":            NewFieldDef(101, 135, false), // 35 chars
	"PayeeAddressLine3":            NewFieldDef(136, 170, false), // 35 chars
	"PayeeAddressLine4":            NewFieldDef(171, 205, false), // 35 chars
	"CityName":                     NewFieldDef(206, 232, false), // 27 chars
	"StateName":                    NewFieldDef(233, 242, false), // 10 chars
	"StateCodeText":                NewFieldDef(243, 244, false), // 2 chars
	"PostalCode":                   NewFieldDef(245, 249, false), // 5 chars
	"PostalCodeExtension":          NewFieldDef(250, 254, false), // 5 chars
	"PostNetBarcodeDeliveryPoint":  NewFieldDef(255, 257, false), // 3 chars
	"Filler1":                      NewFieldDef(258, 271, false), // 14 chars
	"CountryName":                  NewFieldDef(272, 311, false), // 40 chars
	"ConsularCode":                 NewFieldDef(312, 314, false), // 3 chars
	"CheckLegendText1":             NewFieldDef(315, 369, false), // 55 chars
	"CheckLegendText2":             NewFieldDef(370, 424, false), // 55 chars
	"PayeeIdentifierSecondary":     NewFieldDef(425, 433, false), // 9 chars
	"PartyNameSecondary":           NewFieldDef(434, 468, false), // 35 chars
	"PaymentID":                    NewFieldDef(469, 488, true),  // 20 chars
	"Reconcilement":                NewFieldDef(489, 588, false), // 100 chars
	"SpecialHandling":              NewFieldDef(589, 638, false), // 50 chars
	"TIN":                          NewFieldDef(639, 647, false), // 9 chars
	"USPSIntelligentMailBarcode":   NewFieldDef(648, 697, false), // 50 chars
	"PaymentRecipientTINIndicator": NewFieldDef(698, 698, false), // 1 char
	"SecondaryPayeeTINIndicator":   NewFieldDef(699, 699, false), // 1 char
	"AmountEligibleForOffset":      NewFieldDef(700, 709, false), // 10 digits
	"SubPaymentTypeCode":           NewFieldDef(710, 741, false), // 32 chars
	"PayerMechanism":               NewFieldDef(742, 761, false), // 20 chars
	"PaymentDescriptionCode":       NewFieldDef(762, 763, false), // 2 chars
	"Filler":                       NewFieldDef(764, 850, false), // Remaining space
}

// ACH Addendum Record ("03") Field Positions
var ACHAddendum03Fields = map[string]FieldDefinition{
	"RecordCode":         NewFieldDef(1, 2, true),     // "03"
	"PaymentID":          NewFieldDef(3, 22, true),    // 20 chars
	"AddendaInformation": NewFieldDef(23, 102, false), // 80 chars
	"Filler":             NewFieldDef(103, 850, false), // Remaining space
}

// CTX Addendum Record ("04") Field Positions
var ACHAddendum04Fields = map[string]FieldDefinition{
	"RecordCode":         NewFieldDef(1, 2, true),     // "04"
	"PaymentID":          NewFieldDef(3, 22, true),    // 20 chars
	"AddendaInformation": NewFieldDef(23, 822, false), // 800 chars
	"Filler":             NewFieldDef(823, 850, false), // Remaining space
}

// Check Stub Record ("13") Field Positions
var CheckStubFields = map[string]FieldDefinition{
	"RecordCode": NewFieldDef(1, 2, true),   // "13"
	"PaymentID":  NewFieldDef(3, 22, true),  // 20 chars
	// PaymentIdentificationLines: 14 lines of 55 chars each (positions 23-792)
	"Line1":  NewFieldDef(23, 77, false),   // 55 chars
	"Line2":  NewFieldDef(78, 132, false),  // 55 chars
	"Line3":  NewFieldDef(133, 187, false), // 55 chars
	"Line4":  NewFieldDef(188, 242, false), // 55 chars
	"Line5":  NewFieldDef(243, 297, false), // 55 chars
	"Line6":  NewFieldDef(298, 352, false), // 55 chars
	"Line7":  NewFieldDef(353, 407, false), // 55 chars
	"Line8":  NewFieldDef(408, 462, false), // 55 chars
	"Line9":  NewFieldDef(463, 517, false), // 55 chars
	"Line10": NewFieldDef(518, 572, false), // 55 chars
	"Line11": NewFieldDef(573, 627, false), // 55 chars
	"Line12": NewFieldDef(628, 682, false), // 55 chars
	"Line13": NewFieldDef(683, 737, false), // 55 chars
	"Line14": NewFieldDef(738, 792, false), // 55 chars
	"Filler": NewFieldDef(793, 850, false), // 58 chars
}

// CARS TAS/BETC Record ("G ") Field Positions
var CARSTASBETCFields = map[string]FieldDefinition{
	"RecordCode":                    NewFieldDef(1, 2, true),   // "G "
	"PaymentID":                     NewFieldDef(3, 22, true),  // 20 chars
	"SubLevelPrefixCode":            NewFieldDef(23, 24, true), // 2 chars
	"AllocationTransferAgencyID":    NewFieldDef(25, 27, true), // 3 chars
	"AgencyIdentifier":              NewFieldDef(28, 30, true), // 3 chars
	"BeginningPeriodOfAvailability": NewFieldDef(31, 34, true), // 4 chars
	"EndingPeriodOfAvailability":    NewFieldDef(35, 38, true), // 4 chars
	"AvailabilityTypeCode":          NewFieldDef(39, 39, true), // 1 char
	"MainAccountCode":               NewFieldDef(40, 43, true), // 4 chars
	"SubAccountCode":                NewFieldDef(44, 46, true), // 3 chars
	"BusinessEventTypeCode":         NewFieldDef(47, 54, true), // 8 chars
	"AccountClassificationAmount":   NewFieldDef(55, 64, true), // 10 digits
	"IsCredit":                      NewFieldDef(65, 65, true), // 1 char: "0" or "1"
	"Filler":                        NewFieldDef(66, 850, false), // Remaining space
}

// DNP Record ("DD") Field Positions
var DNPFields = map[string]FieldDefinition{
	"RecordCode": NewFieldDef(1, 2, true),     // "DD"
	"PaymentID":  NewFieldDef(3, 22, true),    // 20 chars
	"DNPDetail":  NewFieldDef(23, 788, false), // 766 chars
	"Filler":     NewFieldDef(789, 850, false), // 62 chars
}

// Schedule Trailer Record ("T ") Field Positions
var ScheduleTrailerFields = map[string]FieldDefinition{
	"RecordCode":     NewFieldDef(1, 2, true),     // "T "
	"Filler1":        NewFieldDef(3, 12, false),   // 10 chars
	"ScheduleCount":  NewFieldDef(13, 20, true),   // 8 digits
	"Filler2":        NewFieldDef(21, 23, false),  // 3 chars
	"ScheduleAmount": NewFieldDef(24, 38, true),   // 15 digits
	"Filler3":        NewFieldDef(39, 850, false), // Remaining space
}

// File Trailer Record ("E ") Field Positions
var FileTrailerFields = map[string]FieldDefinition{
	"RecordCode":          NewFieldDef(1, 2, true),     // "E "
	"TotalCountRecords":   NewFieldDef(3, 20, true),    // 18 digits
	"TotalCountPayments":  NewFieldDef(21, 38, true),   // 18 digits
	"TotalAmountPayments": NewFieldDef(39, 56, true),   // 18 digits
	"Filler":              NewFieldDef(57, 850, false), // Remaining space
}

// GetFieldDefinitions returns the field map for a given record type
func GetFieldDefinitions(recordCode string) map[string]FieldDefinition {
	switch recordCode {
	case "H ":
		return FileHeaderFields
	case "01":
		return ACHScheduleHeaderFields
	case "02":
		return ACHPaymentFields
	case "03":
		return ACHAddendum03Fields
	case "04":
		return ACHAddendum04Fields
	case "11":
		return CheckScheduleHeaderFields
	case "12":
		return CheckPaymentFields
	case "13":
		return CheckStubFields
	case "G ":
		return CARSTASBETCFields
	case "DD":
		return DNPFields
	case "T ":
		return ScheduleTrailerFields
	case "E ":
		return FileTrailerFields
	default:
		return nil
	}
}

// ValidateFieldPositions validates that all field definitions are consistent
func ValidateFieldPositions() error {
	allRecordTypes := []string{"H ", "01", "02", "03", "04", "11", "12", "13", "G ", "DD", "T ", "E "}
	
	for _, recordCode := range allRecordTypes {
		fields := GetFieldDefinitions(recordCode)
		if fields == nil {
			continue
		}
		
		// Check that all fields fit within record length
		for fieldName, field := range fields {
			if field.End > RecordLength {
				return fmt.Errorf("record %s field %s ends at position %d, exceeds record length %d", 
					recordCode, fieldName, field.End, RecordLength)
			}
			if field.Start < 1 {
				return fmt.Errorf("record %s field %s starts at position %d, must be >= 1", 
					recordCode, fieldName, field.Start)
			}
			if field.Start > field.End {
				return fmt.Errorf("record %s field %s has start %d > end %d", 
					recordCode, fieldName, field.Start, field.End)
			}
			if field.Length != (field.End-field.Start+1) {
				return fmt.Errorf("record %s field %s has incorrect length calculation", 
					recordCode, fieldName)
			}
		}
		
		// Check for overlapping fields (exclude filler fields)
		var positions []int
		for fieldName, field := range fields {
			if !strings.Contains(fieldName, "Filler") {
				for pos := field.Start; pos <= field.End; pos++ {
					positions = append(positions, pos)
				}
			}
		}
		
		// Check for duplicates
		posMap := make(map[int]bool)
		for _, pos := range positions {
			if posMap[pos] {
				return fmt.Errorf("record %s has overlapping fields at position %d", recordCode, pos)
			}
			posMap[pos] = true
		}
	}
	
	return nil
}