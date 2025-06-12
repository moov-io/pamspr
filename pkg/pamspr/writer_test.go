package pamspr

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// mockWriter simulates write errors
type mockWriter struct {
	err error
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	if m.err != nil {
		return 0, m.err
	}
	return len(p), nil
}

// TestNewWriter tests the NewWriter constructor
func TestNewWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if writer == nil {
		t.Fatal("NewWriter returned nil")
	}
	if writer.w == nil {
		t.Error("writer has nil io.Writer")
	}
	if writer.validator == nil {
		t.Error("writer has nil validator")
	}
	if writer.errors == nil {
		t.Error("writer has nil errors slice")
	}
}

// TestWriter_Write_CompleteACHFile tests writing a complete ACH file with all record types
func TestWriter_Write_CompleteACHFile(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TREASURY SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:              "01",
					AgencyACHText:           "TREA",
					ScheduleNumber:          "00000000123456",
					PaymentTypeCode:         "Vendor",
					StandardEntryClassCode:  "CCD",
					AgencyLocationCode:      "12345678",
					FederalEmployerIDNumber: "123456789 ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{
							RecordCode:                   "02",
							AgencyAccountIdentifier:      "ACC123456789012 ",
							Amount:                       150000, // $1,500.00
							AgencyPaymentTypeCode:        "V",
							IsTOP_Offset:                 "0",
							PayeeName:                    "ACME CORPORATION",
							PayeeAddressLine1:            "123 BUSINESS STREET",
							PayeeAddressLine2:            "SUITE 100",
							CityName:                     "NEW YORK",
							StateName:                    "NEW YORK",
							StateCodeText:                "NY",
							PostalCode:                   "10001",
							PostalCodeExtension:          "1234 ",
							CountryCodeText:              "US",
							RoutingNumber:                "021000021",
							AccountNumber:                "123456789012345  ",
							ACH_TransactionCode:          "22",
							PayeeIdentifierAdditional:    "         ",
							PayeeNameAdditional:          "",
							PaymentID:                    "PAY000000001        ",
							Reconcilement:                "INVOICE #12345",
							TIN:                          "123456789",
							PaymentRecipientTINIndicator: "2",
							AdditionalPayeeTINIndicator:  "0",
							AmountEligibleForOffset:      "0000150000",
							PayeeAddressLine3:            "",
							PayeeAddressLine4:            "",
							CountryName:                  "UNITED STATES",
							ConsularCode:                 "   ",
							SubPaymentTypeCode:           "VENDOR PAYMENT",
							PayerMechanism:               "ACH CREDIT",
							PaymentDescriptionCode:       "  ",
							Addenda: []*ACHAddendum{
								{
									RecordCode:         "03",
									PaymentID:          "PAY000000001        ",
									AddendaInformation: "INVOICE 12345 DATED 2024-01-15",
								},
							},
							CARSTASBETC: []*CARSTASBETC{
								{
									RecordCode:                    "G ",
									PaymentID:                     "PAY000000001        ",
									SubLevelPrefixCode:            "20",
									AllocationTransferAgencyID:    "020",
									AgencyIdentifier:              "020",
									BeginningPeriodOfAvailability: "2024",
									EndingPeriodOfAvailability:    "2024",
									AvailabilityTypeCode:          "X",
									MainAccountCode:               "1000",
									SubAccountCode:                "000",
									BusinessEventTypeCode:         "DISB0001",
									AccountClassificationAmount:   150000,
									IsCredit:                      "0",
								},
							},
							DNP: &DNPRecord{
								RecordCode: "DD",
								PaymentID:  "PAY000000001        ",
								DNPDetail:  "DNP TEST DETAIL INFORMATION",
							},
						},
						// Second payment without optional records
						&ACHPayment{
							RecordCode:                   "02",
							AgencyAccountIdentifier:      "ACC987654321098 ",
							Amount:                       75000, // $750.00
							AgencyPaymentTypeCode:        "S",
							IsTOP_Offset:                 "1",
							PayeeName:                    "JOHN DOE",
							PayeeAddressLine1:            "456 MAIN STREET",
							PayeeAddressLine2:            "",
							CityName:                     "LOS ANGELES",
							StateName:                    "CALIFORNIA",
							StateCodeText:                "CA",
							PostalCode:                   "90001",
							PostalCodeExtension:          "     ",
							CountryCodeText:              "US",
							RoutingNumber:                "122000247",
							AccountNumber:                "987654321       ",
							ACH_TransactionCode:          "27",
							PayeeIdentifierAdditional:    "         ",
							PayeeNameAdditional:          "",
							PaymentID:                    "PAY000000002        ",
							Reconcilement:                "",
							TIN:                          "987654321",
							PaymentRecipientTINIndicator: "1",
							AdditionalPayeeTINIndicator:  "0",
							AmountEligibleForOffset:      "0000075000",
							PayeeAddressLine3:            "",
							PayeeAddressLine4:            "",
							CountryName:                  "UNITED STATES",
							ConsularCode:                 "   ",
							SubPaymentTypeCode:           "SALARY PAYMENT",
							PayerMechanism:               "ACH DEBIT",
							PaymentDescriptionCode:       "  ",
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  2,
						ScheduleAmount: 225000, // $2,250.00 total
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   9, // H + 01 + 02 + 03 + G + DD + 02 + T + E = 9
			TotalCountPayments:  2,
			TotalAmountPayments: 225000,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	// Verify we have the correct number of lines
	expectedLines := 9 // H + 01 + 02 + 03 + G + DD + 02 + T + E
	if len(lines) != expectedLines {
		t.Errorf("expected %d lines, got %d", expectedLines, len(lines))
	}

	// Verify each line has correct length
	for i, line := range lines {
		if len(line) != RecordLength {
			t.Errorf("line %d has incorrect length: expected %d, got %d", i+1, RecordLength, len(line))
		}
	}

	// Verify record codes are in correct order
	expectedCodes := []string{"H ", "01", "02", "03", "G ", "DD", "02", "T ", "E "}
	for i, code := range expectedCodes {
		if !strings.HasPrefix(lines[i], code) {
			t.Errorf("line %d: expected record code %q, got %q", i+1, code, lines[i][:2])
		}
	}

	// Verify specific content
	if !strings.Contains(output, "TREASURY SYSTEM") {
		t.Error("output missing input system")
	}
	if !strings.Contains(output, "ACME CORPORATION") {
		t.Error("output missing first payee name")
	}
	if !strings.Contains(output, "JOHN DOE") {
		t.Error("output missing second payee name")
	}
	if !strings.Contains(output, "INVOICE 12345 DATED 2024-01-15") {
		t.Error("output missing addenda information")
	}
	if !strings.Contains(output, "DNP TEST DETAIL INFORMATION") {
		t.Error("output missing DNP detail")
	}
}

// TestWriter_Write_CompleteCheckFile tests writing a complete check file
func TestWriter_Write_CompleteCheckFile(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "CHECK SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&CheckSchedule{
				Header: &CheckScheduleHeader{
					RecordCode:                "11",
					ScheduleNumber:            "00000000456789",
					PaymentTypeCode:           "Refund",
					AgencyLocationCode:        "87654321",
					CheckPaymentEnclosureCode: "stub      ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&CheckPayment{
							RecordCode:                   "12",
							AgencyAccountIdentifier:      "CHK123456789012 ",
							Amount:                       250000, // $2,500.00
							AgencyPaymentTypeCode:        "R",
							IsTOP_Offset:                 "0",
							PayeeName:                    "JANE SMITH",
							PayeeAddressLine1:            "789 OAK AVENUE",
							PayeeAddressLine2:            "APT 5B",
							PayeeAddressLine3:            "BUILDING C",
							PayeeAddressLine4:            "",
							CityName:                     "CHICAGO",
							StateName:                    "ILLINOIS",
							StateCodeText:                "IL",
							PostalCode:                   "60601",
							PostalCodeExtension:          "9876 ",
							PostNetBarcodeDeliveryPoint:  "123",
							CountryName:                  "UNITED STATES",
							ConsularCode:                 "   ",
							CheckLegendText1:             "TAX REFUND FOR 2023",
							CheckLegendText2:             "THANK YOU FOR YOUR PATIENCE",
							PayeeIdentifier_Secondary:    "123456789",
							PartyName_Secondary:          "JOINT ACCOUNT HOLDER",
							PaymentID:                    "CHK000000001        ",
							Reconcilement:                "REFUND FOR OVERPAYMENT",
							SpecialHandling:              "CERTIFIED MAIL",
							TIN:                          "123456789",
							USPSIntelligentMailBarcode:   "12345678901234567890123456789012345678901234567890",
							PaymentRecipientTINIndicator: "1",
							SecondaryPayeeTINIndicator:   "1",
							AmountEligibleForOffset:      "0000250000",
							SubPaymentTypeCode:           "TAX REFUND",
							PayerMechanism:               "PAPER CHECK",
							PaymentDescriptionCode:       "TR",
							Stub: &CheckStub{
								RecordCode: "13",
								PaymentID:  "CHK000000001        ",
								PaymentIdentificationLines: [14]string{
									"TAX YEAR: 2023",
									"FORM: 1040",
									"FILING STATUS: MARRIED FILING JOINTLY",
									"ADJUSTED GROSS INCOME: $75,000",
									"TOTAL TAX: $12,500",
									"TOTAL PAYMENTS: $15,000",
									"REFUND AMOUNT: $2,500",
									"INTEREST: $0.00",
									"",
									"KEEP THIS STUB FOR YOUR RECORDS",
									"",
									"",
									"",
									"",
								},
							},
							CARSTASBETC: []*CARSTASBETC{
								{
									RecordCode:                    "G ",
									PaymentID:                     "CHK000000001        ",
									SubLevelPrefixCode:            "20",
									AllocationTransferAgencyID:    "015",
									AgencyIdentifier:              "015",
									BeginningPeriodOfAvailability: "2024",
									EndingPeriodOfAvailability:    "2024",
									AvailabilityTypeCode:          "X",
									MainAccountCode:               "2000",
									SubAccountCode:                "100",
									BusinessEventTypeCode:         "REFUND01",
									AccountClassificationAmount:   250000,
									IsCredit:                      "1",
								},
							},
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 250000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   6, // H + 11 + 12 + 13 + G + T + E = 7
			TotalCountPayments:  1,
			TotalAmountPayments: 250000,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	// Verify we have the correct number of lines
	expectedLines := 7 // H + 11 + 12 + 13 + G + T + E
	if len(lines) != expectedLines {
		t.Errorf("expected %d lines, got %d", expectedLines, len(lines))
	}

	// Verify record codes are in correct order
	expectedCodes := []string{"H ", "11", "12", "13", "G ", "T ", "E "}
	for i, code := range expectedCodes {
		if !strings.HasPrefix(lines[i], code) {
			t.Errorf("line %d: expected record code %q, got %q", i+1, code, lines[i][:2])
		}
	}

	// Verify specific content
	if !strings.Contains(output, "CHECK SYSTEM") {
		t.Error("output missing input system")
	}
	if !strings.Contains(output, "JANE SMITH") {
		t.Error("output missing payee name")
	}
	if !strings.Contains(output, "TAX REFUND FOR 2023") {
		t.Error("output missing check legend text")
	}
	if !strings.Contains(output, "TAX YEAR: 2023") {
		t.Error("output missing stub line")
	}
}

// TestWriter_Write_CTXAddendum tests writing CTX (04) addendum records
func TestWriter_Write_CTXAddendum(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Create a long CTX addendum string (800 chars)
	ctxData := strings.Repeat("CTX DATA ", 88) + "CTX DATA" // 800 chars exactly

	addendum := &ACHAddendum{
		RecordCode:         "04",
		PaymentID:          "CTXPAYMENT12345678  ",
		AddendaInformation: ctxData,
	}

	err := writer.writeACHAddendum(addendum)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	line := strings.TrimRight(output, "\n")

	if len(line) != RecordLength {
		t.Errorf("CTX addendum line has incorrect length: expected %d, got %d", RecordLength, len(line))
	}

	if !strings.HasPrefix(line, "04") {
		t.Error("CTX addendum missing record code")
	}

	if !strings.Contains(line, "CTXPAYMENT12345678") {
		t.Error("CTX addendum missing payment ID")
	}
}

// TestWriter_Write_ValidationFailure tests file validation failure
func TestWriter_Write_ValidationFailure(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Create a file that will fail validation
	file := &File{
		Header: &FileHeader{
			RecordCode: "XX", // Invalid record code
		},
	}

	err := writer.Write(file)
	if err == nil {
		t.Error("expected validation error, got nil")
	}
	if !strings.Contains(err.Error(), "file validation") {
		t.Errorf("expected file validation error, got: %v", err)
	}
}

// TestWriter_Write_WriteError tests handling of write errors
func TestWriter_Write_WriteError(t *testing.T) {
	mockErr := errors.New("write failed")
	writer := NewWriter(&mockWriter{err: mockErr})

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TEST",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   2,
			TotalCountPayments:  0,
			TotalAmountPayments: 0,
		},
	}

	err := writer.Write(file)
	if err == nil {
		t.Error("expected write error, got nil")
	}
	if !strings.Contains(err.Error(), "write failed") {
		t.Errorf("expected write failed error, got: %v", err)
	}
}

// TestWriter_writeACHAddendum_InvalidRecordCode tests invalid addendum record code
func TestWriter_writeACHAddendum_InvalidRecordCode(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	addendum := &ACHAddendum{
		RecordCode: "99", // Invalid code
	}

	err := writer.writeACHAddendum(addendum)
	if err == nil {
		t.Error("expected error for invalid record code, got nil")
	}
	if !strings.Contains(err.Error(), "invalid addendum record code") {
		t.Errorf("expected invalid addendum record code error, got: %v", err)
	}
}

// TestWriter_formatField tests field formatting functions
func TestWriter_formatField(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	tests := []struct {
		name        string
		value       string
		length      int
		expected    string
		shouldPanic bool
	}{
		{"short value padded", "ABC", 5, "ABC  ", false},
		{"exact length", "ABCDE", 5, "ABCDE", false},
		{"truncated", "ABCDEFGH", 5, "ABCDE", false}, // Returns truncated value but accumulates error
		{"empty value", "", 3, "   ", false},
		{"spaces preserved", "A B", 5, "A B  ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("formatField(%q, %d) should have panicked but didn't", tt.value, tt.length)
					}
				}()
			}

			result := writer.formatField(tt.value, tt.length)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("formatField(%q, %d) = %q, want %q", tt.value, tt.length, result, tt.expected)
			}
		})
	}
}

// TestWriter_formatFieldRightJustified tests right-justified field formatting
func TestWriter_formatFieldRightJustified(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	tests := []struct {
		name        string
		value       string
		length      int
		padChar     rune
		expected    string
		shouldPanic bool
	}{
		{"zeros padding", "123", 5, '0', "00123", false},
		{"spaces padding", "ABC", 5, ' ', "  ABC", false},
		{"exact length", "12345", 5, '0', "12345", false},
		{"truncated", "123456", 5, '0', "12345", false}, // Returns truncated value but accumulates error
		{"empty value", "", 3, '0', "000", false},
		{"trim spaces", "  123  ", 5, '0', "00123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("formatFieldRightJustified(%q, %d, %q) should have panicked but didn't",
							tt.value, tt.length, tt.padChar)
					}
				}()
			}

			result := writer.formatFieldRightJustified(tt.value, tt.length, tt.padChar)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("formatFieldRightJustified(%q, %d, %q) = %q, want %q",
					tt.value, tt.length, tt.padChar, result, tt.expected)
			}
		})
	}
}

// TestWriter_formatFieldNoJustify tests no-justify field formatting
func TestWriter_formatFieldNoJustify(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	tests := []struct {
		name        string
		value       string
		length      int
		expected    string
		shouldPanic bool
	}{
		{"short value padded", "ABC", 5, "ABC  ", false},
		{"exact length", "ABCDE", 5, "ABCDE", false},
		{"truncated", "ABCDEFGH", 5, "ABCDE", false}, // Returns truncated value but accumulates error
		{"empty value", "", 3, "   ", false},
		{"preserves internal spaces", "A  B", 5, "A  B ", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("formatFieldNoJustify(%q, %d) should have panicked but didn't", tt.value, tt.length)
					}
				}()
			}

			result := writer.formatFieldNoJustify(tt.value, tt.length)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("formatFieldNoJustify(%q, %d) = %q, want %q", tt.value, tt.length, result, tt.expected)
			}
		})
	}
}

// TestWriter_formatNumeric tests numeric field formatting
func TestWriter_formatNumeric(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	tests := []struct {
		name        string
		value       string
		length      int
		expected    string
		shouldPanic bool
	}{
		{"simple number", "123", 5, "00123", false},
		{"with non-numeric", "12-34", 6, "001234", false},
		{"letters removed", "ABC123DEF", 5, "00123", false},
		{"empty string", "", 3, "000", false},
		{"all non-numeric", "ABC", 3, "000", false},
		{"truncated", "123456789", 5, "12345", false}, // Returns truncated value but accumulates error
		{"spaces removed", "1 2 3", 5, "00123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("formatNumeric(%q, %d) should have panicked but didn't", tt.value, tt.length)
					}
				}()
			}

			result := writer.formatNumeric(tt.value, tt.length)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("formatNumeric(%q, %d) = %q, want %q", tt.value, tt.length, result, tt.expected)
			}
		})
	}
}

// TestWriter_formatAmount tests amount field formatting
func TestWriter_formatAmount(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	tests := []struct {
		name        string
		cents       int64
		length      int
		expected    string
		shouldPanic bool
	}{
		{"simple amount", 12345, 10, "0000012345", false},
		{"zero amount", 0, 8, "00000000", false},
		{"large amount", 9999999999, 10, "9999999999", false},
		{"negative amount", -12345, 10, "0000012345", false}, // Note: negatives become positive
		{"truncated", 123456789012, 10, "1234567890", false}, // Returns truncated value but accumulates error
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("formatAmount(%d, %d) should have panicked but didn't", tt.cents, tt.length)
					}
				}()
			}

			result := writer.formatAmount(tt.cents, tt.length)
			if !tt.shouldPanic && result != tt.expected {
				t.Errorf("formatAmount(%d, %d) = %q, want %q", tt.cents, tt.length, result, tt.expected)
			}
		})
	}
}

// TestWriter_writeLine tests line writing with validation
func TestWriter_writeLine(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		wantErr bool
	}{
		{"valid line", strings.Repeat("X", RecordLength), false},
		{"too short", strings.Repeat("X", RecordLength-1), true},
		{"too long", strings.Repeat("X", RecordLength+1), true},
		{"empty line", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			writer := NewWriter(&buf)
			err := writer.writeLine(tt.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeLine() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && !strings.Contains(buf.String(), tt.line) {
				t.Error("line not written to buffer")
			}
		})
	}
}

// TestWriter_Write_MixedSchedules tests writing a file with both ACH and Check schedules
func TestWriter_Write_MixedSchedules(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "MIXED SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0", // Test mixed ACH and Check schedules
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:              "01",
					AgencyACHText:           "AGCY",
					ScheduleNumber:          "00000000000001",
					PaymentTypeCode:         "Salary",
					StandardEntryClassCode:  "PPD",
					AgencyLocationCode:      "11111111",
					FederalEmployerIDNumber: "111111111 ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{
							RecordCode:              "02",
							AgencyAccountIdentifier: "SALARY001       ",
							Amount:                  500000,
							PaymentID:               "ACH001              ",
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 500000,
					},
				},
			},
			&CheckSchedule{
				Header: &CheckScheduleHeader{
					RecordCode:                "11",
					ScheduleNumber:            "00000000000002",
					PaymentTypeCode:           "Vendor",
					AgencyLocationCode:        "22222222",
					CheckPaymentEnclosureCode: "letter    ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&CheckPayment{
							RecordCode:              "12",
							AgencyAccountIdentifier: "VENDOR001       ",
							Amount:                  300000,
							PaymentID:               "CHK001              ",
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 300000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   8, // H + 01 + 02 + T + 11 + 12 + T + E
			TotalCountPayments:  2,
			TotalAmountPayments: 800000,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}

	// Verify both ACH and Check records are present
	if !strings.Contains(output, "ACH001") {
		t.Error("missing ACH payment ID")
	}
	if !strings.Contains(output, "CHK001") {
		t.Error("missing Check payment ID")
	}
}

// TestWriter_Write_EmptySchedules tests writing a file with no schedules
func TestWriter_Write_EmptySchedules(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "EMPTY SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   2,
			TotalCountPayments:  0,
			TotalAmountPayments: 0,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	if len(lines) != 2 {
		t.Errorf("expected 2 lines (header and trailer), got %d", len(lines))
	}
}

// TestWriter_FieldPaddingEdgeCases tests edge cases in field padding
func TestWriter_FieldPaddingEdgeCases(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Test with maximum length values
	maxString := strings.Repeat("A", 1000) // Much longer than any field

	// Test that oversized input returns truncated value and accumulates error
	t.Run("formatField with oversized input", func(t *testing.T) {
		result := writer.formatField(maxString, 10)
		if len(result) != 10 {
			t.Errorf("formatField should return truncated value of length 10, got %d", len(result))
		}
		if len(writer.errors) == 0 {
			t.Errorf("formatField should accumulate error for oversized input")
		}
	})

	t.Run("formatFieldRightJustified with oversized input", func(t *testing.T) {
		// Clear previous errors
		writer.errors = writer.errors[:0]
		result := writer.formatFieldRightJustified(maxString, 10, '0')
		if len(result) != 10 {
			t.Errorf("formatFieldRightJustified should return truncated value of length 10, got %d", len(result))
		}
		if len(writer.errors) == 0 {
			t.Errorf("formatFieldRightJustified should accumulate error for oversized input")
		}
	})

	// Test exact length (should not panic)
	t.Run("exact length", func(t *testing.T) {
		exactString := strings.Repeat("B", 10)
		result := writer.formatField(exactString, 10)
		if len(result) != 10 {
			t.Errorf("formatField with exact length should return exact length, got %d", len(result))
		}
	})
}

// TestWriter_AllRecordTypes tests that all record types can be written
func TestWriter_AllRecordTypes(t *testing.T) {
	tests := []struct {
		name   string
		record interface{}
		writer func(*Writer, interface{}) error
	}{
		{
			name: "FileHeader",
			record: &FileHeader{
				RecordCode:               "H ",
				InputSystem:              "TEST",
				StandardPaymentVersion:   "502",
				IsRequestedForSameDayACH: "0",
			},
			writer: func(w *Writer, r interface{}) error {
				return w.writeFileHeader(r.(*FileHeader))
			},
		},
		{
			name: "ACHScheduleHeader",
			record: &ACHScheduleHeader{
				RecordCode:              "01",
				AgencyACHText:           "TEST",
				ScheduleNumber:          "12345",
				PaymentTypeCode:         "Vendor",
				StandardEntryClassCode:  "CCD",
				AgencyLocationCode:      "12345678",
				FederalEmployerIDNumber: "123456789",
			},
			writer: func(w *Writer, r interface{}) error {
				return w.writeACHScheduleHeader(r.(*ACHScheduleHeader))
			},
		},
		{
			name: "CheckScheduleHeader",
			record: &CheckScheduleHeader{
				RecordCode:                "11",
				ScheduleNumber:            "12345",
				PaymentTypeCode:           "Refund",
				AgencyLocationCode:        "87654321",
				CheckPaymentEnclosureCode: "stub",
			},
			writer: func(w *Writer, r interface{}) error {
				return w.writeCheckScheduleHeader(r.(*CheckScheduleHeader))
			},
		},
		{
			name: "ScheduleTrailer",
			record: &ScheduleTrailer{
				RecordCode:     "T ",
				ScheduleCount:  5,
				ScheduleAmount: 100000,
			},
			writer: func(w *Writer, r interface{}) error {
				return w.writeScheduleTrailer(r.(*ScheduleTrailer))
			},
		},
		{
			name: "FileTrailer",
			record: &FileTrailer{
				RecordCode:          "E ",
				TotalCountRecords:   100,
				TotalCountPayments:  50,
				TotalAmountPayments: 5000000,
			},
			writer: func(w *Writer, r interface{}) error {
				return w.writeFileTrailer(r.(*FileTrailer))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			writer := NewWriter(&buf)

			err := tt.writer(writer, tt.record)
			if err != nil {
				t.Errorf("unexpected error writing %s: %v", tt.name, err)
			}

			output := buf.String()
			if len(strings.TrimSpace(output)) == 0 {
				t.Errorf("no output for %s", tt.name)
			}

			// Check line length (don't trim spaces as padding is required)
			lines := strings.Split(output, "\n")
			if len(lines) > 0 && lines[0] != "" {
				line := lines[0]
				if len(line) != RecordLength {
					t.Errorf("%s: expected line length %d, got %d", tt.name, RecordLength, len(line))
				}
			}
		})
	}
}

// BenchmarkWriter_Write benchmarks writing performance
func BenchmarkWriter_Write(b *testing.B) {
	// Create a file with multiple payments
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "BENCHMARK TEST",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:              "01",
					AgencyACHText:           "BNCH",
					ScheduleNumber:          "00000000000001",
					PaymentTypeCode:         "Vendor",
					StandardEntryClassCode:  "CCD",
					AgencyLocationCode:      "12345678",
					FederalEmployerIDNumber: "123456789 ",
				},
				BaseSchedule: BaseSchedule{
					Payments: make([]Payment, 100), // 100 payments
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  100,
						ScheduleAmount: 10000000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   103,
			TotalCountPayments:  100,
			TotalAmountPayments: 10000000,
		},
	}

	// Initialize payments
	for i := 0; i < 100; i++ {
		file.Schedules[0].(*ACHSchedule).Payments[i] = &ACHPayment{
			RecordCode:              "02",
			AgencyAccountIdentifier: fmt.Sprintf("ACC%013d ", i),
			Amount:                  100000,
			PaymentID:               fmt.Sprintf("PAY%017d ", i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		writer := NewWriter(&buf)
		if err := writer.Write(file); err != nil {
			b.Fatal(err)
		}
	}
}
