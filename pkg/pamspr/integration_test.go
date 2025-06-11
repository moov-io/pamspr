package pamspr

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TestIntegrationFullFileRoundTrip tests complete file read/write cycle
func TestIntegrationFullFileRoundTrip(t *testing.T) {
	tests := []struct {
		name        string
		createFile  func() *File
		description string
	}{
		{
			name:        "ACH only file",
			description: "File with single ACH schedule and multiple payments",
			createFile:  createTestACHFile,
		},
		{
			name:        "Check only file",
			description: "File with single check schedule and payments with stubs",
			createFile:  createTestCheckFile,
		},
		{
			name:        "Mixed ACH and Check file",
			description: "File with both ACH and check schedules",
			createFile:  createTestMixedFile,
		},
		{
			name:        "ACH with all addenda types",
			description: "ACH payments with 03 and 04 (CTX) addenda",
			createFile:  createTestACHWithAddenda,
		},
		{
			name:        "File with CARS/TAS/BETC records",
			description: "Payments with accounting classification records",
			createFile:  createTestFileWithCARS,
		},
		{
			name:        "File with DNP records",
			description: "Payments with DNP detail records",
			createFile:  createTestFileWithDNP,
		},
		{
			name:        "Same Day ACH file",
			description: "File marked for same day ACH processing",
			createFile:  createTestSameDayACHFile,
		},
		{
			name:        "Large file performance test",
			description: "File with 1000 payments",
			createFile:  createTestLargeFile,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test file
			originalFile := tt.createFile()

			// Write to buffer
			var buf bytes.Buffer
			writer := NewWriter(&buf)
			if err := writer.Write(originalFile); err != nil {
				t.Fatalf("Failed to write file: %v", err)
			}

			// Read back from buffer (make a copy since reading consumes the buffer)
			bufCopy := bytes.NewReader(buf.Bytes())
			reader := NewReader(bufCopy)
			readFile, err := reader.Read()
			if err != nil {
				t.Fatalf("Failed to read file: %v", err)
			}

			// Validate structure matches
			validateFileStructure(t, originalFile, readFile)

			// Write again and compare output
			var buf2 bytes.Buffer
			writer2 := NewWriter(&buf2)
			if err := writer2.Write(readFile); err != nil {
				t.Fatalf("Failed to write file second time: %v", err)
			}

			// Compare outputs should be identical
			if buf.String() != buf2.String() {
				t.Error("Round-trip produced different output")
				t.Logf("First write length: %d", len(buf.String()))
				t.Logf("Second write length: %d", len(buf2.String()))
				lines1 := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
				lines2 := strings.Split(strings.TrimRight(buf2.String(), "\n"), "\n")
				t.Logf("First write lines: %d", len(lines1))
				t.Logf("Second write lines: %d", len(lines2))
				if len(lines2) > 0 {
					t.Logf("First line second write: %q (len=%d)", lines2[0], len(lines2[0]))
				}
			}

			// Validate all records are 850 characters
			lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
			for i, line := range lines {
				if len(line) != RecordLength {
					t.Errorf("Line %d has incorrect length: expected %d, got %d",
						i+1, RecordLength, len(line))
				}
			}
		})
	}
}

// TestIntegrationValidation tests that validation catches all error cases
func TestIntegrationValidation(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name       string
		createFile func() *File
		wantErr    bool
		errType    string
	}{
		{
			name:       "Invalid routing number",
			createFile: createFileWithInvalidRoutingNumber,
			wantErr:    true,
			errType:    "routing_number",
		},
		{
			name:       "Unbalanced file totals",
			createFile: createUnbalancedFile,
			wantErr:    true,
			errType:    "balance",
		},
		{
			name:       "SDA with check payments",
			createFile: createInvalidSDAFile,
			wantErr:    true,
			errType:    "sda_ach_only",
		},
		{
			name:       "CTX without addendum",
			createFile: createCTXWithoutAddendum,
			wantErr:    true,
			errType:    "ctx_required",
		},
		{
			name:       "Mixed payment types in schedule",
			createFile: createMixedPaymentSchedule,
			wantErr:    true,
			errType:    "payment_type_consistency",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := tt.createFile()

			var err error
			// Use appropriate validation method based on error type
			if tt.errType == "balance" {
				err = validator.ValidateBalancing(file)
			} else {
				err = validator.ValidateFileStructure(file)
			}

			if tt.wantErr && err == nil {
				t.Error("Expected validation error but got none")
			}
			if !tt.wantErr && err != nil {
				t.Errorf("Unexpected validation error: %v", err)
			}
		})
	}
}

// TestIntegrationFieldPositions verifies all field positions are correct
func TestIntegrationFieldPositions(t *testing.T) {
	// Test ACH payment field positions - create a valid 850-character line with real data
	// Use field positions to build the line properly
	line := make([]byte, 850)
	for i := range line {
		line[i] = ' ' // Fill with spaces first
	}

	// Set specific fields at their correct positions
	// RecordCode (1-2)
	copy(line[0:2], "02")
	// AgencyAccountIdentifier (3-18)
	copy(line[2:18], "ACC123456789012 ")
	// Amount (19-28)
	copy(line[18:28], "0000150000")
	// PayeeName (31-65)
	copy(line[30:65], "PAYEE NAME HERE                    ")
	// RoutingNumber (187-195)
	copy(line[186:195], "021000021")

	achLine := string(line)

	parser := NewACHParser(nil) // Disable validation for field position test
	payment, err := parser.ParseACHPayment(achLine)
	if err != nil {
		t.Fatalf("Failed to parse ACH payment: %v", err)
	}

	// Verify critical field extractions
	if payment.AgencyAccountIdentifier != "ACC123456789012 " {
		t.Errorf("Incorrect agency account: %q", payment.AgencyAccountIdentifier)
	}
	if payment.Amount != 150000 {
		t.Errorf("Incorrect amount: %d", payment.Amount)
	}
	if payment.PayeeName != "PAYEE NAME HERE                    " {
		t.Errorf("Incorrect payee name: %q", payment.PayeeName)
	}
	if payment.RoutingNumber != "021000021" {
		t.Errorf("Incorrect routing number: %q", payment.RoutingNumber)
	}
}

// Helper functions to create test files
func createTestACHFile() *File {
	return &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TEST SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:              "01",
					AgencyACHText:           "TEST",
					ScheduleNumber:          "00000000000001",
					PaymentTypeCode:         "Vendor",
					StandardEntryClassCode:  "PPD",
					AgencyLocationCode:      "12345678",
					FederalEmployerIDNumber: "123456789 ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{
							RecordCode:              "02",
							AgencyAccountIdentifier: "ACC001          ",
							Amount:                  100000, // $1,000.00
							PayeeName:               "TEST PAYEE 1",
							RoutingNumber:           "021000021",
							AccountNumber:           "1234567890",
							ACHTransactionCode:      "22",
							PaymentID:               "PAY001              ",
						},
						&ACHPayment{
							RecordCode:              "02",
							AgencyAccountIdentifier: "ACC002          ",
							Amount:                  200000, // $2,000.00
							PayeeName:               "TEST PAYEE 2",
							RoutingNumber:           "122000247",
							AccountNumber:           "9876543210",
							ACHTransactionCode:      "27",
							PaymentID:               "PAY002              ",
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  2,
						ScheduleAmount: 300000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   6, // H + 01 + 02 + 02 + T + E
			TotalCountPayments:  2,
			TotalAmountPayments: 300000,
		},
	}
}

func createTestCheckFile() *File {
	return &File{
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
					ScheduleNumber:            "00000000000001",
					PaymentTypeCode:           "Refund",
					AgencyLocationCode:        "87654321",
					CheckPaymentEnclosureCode: "stub",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&CheckPayment{
							RecordCode:              "12",
							AgencyAccountIdentifier: "CHK001          ",
							Amount:                  150000, // $1,500.00
							PayeeName:               "CHECK PAYEE",
							PaymentID:               "CHK001              ",
							Stub: &CheckStub{
								RecordCode: "13",
								PaymentID:  "CHK001              ",
								PaymentIdentificationLines: [14]string{
									"INVOICE #12345",
									"DATE: 2024-01-01",
								},
							},
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 150000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   6, // H + 11 + 12 + 13 + T + E
			TotalCountPayments:  1,
			TotalAmountPayments: 150000,
		},
	}
}

func createTestMixedFile() *File {
	achFile := createTestACHFile()
	checkSchedule := createTestCheckFile().Schedules[0]

	// Add check schedule to ACH file
	achFile.Schedules = append(achFile.Schedules, checkSchedule)

	// Update totals
	achFile.Trailer.TotalCountRecords = 10 // H + 01 + 02 + 02 + T + 11 + 12 + 13 + T + E
	achFile.Trailer.TotalCountPayments = 3
	achFile.Trailer.TotalAmountPayments = 450000

	return achFile
}

func createTestACHWithAddenda() *File {
	file := createTestACHFile()

	// Add standard addendum
	if achSchedule, ok := AsACHSchedule(file.Schedules[0]); ok {
		if achPayment, ok := AsACHPayment(achSchedule.GetPayments()[0]); ok {
			achPayment.SetAddenda([]*ACHAddendum{
				{
					RecordCode:         "03",
					PaymentID:          "PAY001              ",
					AddendaInformation: "INVOICE 12345 DATED 2024-01-01",
				},
			})
		}
	}

	// Add CTX addendum to second payment
	if achSchedule, ok := AsACHSchedule(file.Schedules[0]); ok {
		if achPayment, ok := AsACHPayment(achSchedule.GetPayments()[1]); ok {
			achPayment.SetStandardEntryClassCode("CTX")
			achPayment.SetAddenda([]*ACHAddendum{
				{
					RecordCode:         "04",
					PaymentID:          "PAY002              ",
					AddendaInformation: "ISA*00*          *00*          " + strings.Repeat("X", 768),
				},
			})
		}
	}

	// Update record count
	file.Trailer.TotalCountRecords = 8 // Added 2 addenda records

	return file
}

func createTestFileWithCARS() *File {
	file := createTestACHFile()

	// Add CARS record
	if achSchedule, ok := AsACHSchedule(file.Schedules[0]); ok {
		if achPayment, ok := AsACHPayment(achSchedule.GetPayments()[0]); ok {
			achPayment.SetCARSTASBETC([]*CARSTASBETC{
				{
					RecordCode:                    "G ",
					PaymentID:                     "PAY001              ",
					SubLevelPrefixCode:            "20",
					AllocationTransferAgencyID:    "020",
					AgencyIdentifier:              "020",
					BeginningPeriodOfAvailability: "2024",
					EndingPeriodOfAvailability:    "2024",
					AvailabilityTypeCode:          "X",
					MainAccountCode:               "1000",
					SubAccountCode:                "000",
					BusinessEventTypeCode:         "DISB0001",
					AccountClassificationAmount:   100000,
					IsCredit:                      "0",
				},
			})
		}
	}

	// Update record count
	file.Trailer.TotalCountRecords = 7 // Added 1 CARS record

	return file
}

func createTestFileWithDNP() *File {
	file := createTestACHFile()

	// Add DNP record
	if achSchedule, ok := AsACHSchedule(file.Schedules[0]); ok {
		if achPayment, ok := AsACHPayment(achSchedule.GetPayments()[0]); ok {
			achPayment.SetDNP(&DNPRecord{
				RecordCode: "DD",
				PaymentID:  "PAY001              ",
				DNPDetail:  "DNP TEST DETAIL INFORMATION",
			})
		}
	}

	// Update record count
	file.Trailer.TotalCountRecords = 7 // Added 1 DNP record

	return file
}

func createTestSameDayACHFile() *File {
	file := createTestACHFile()
	file.Header.IsRequestedForSameDayACH = "1"

	// Ensure amounts are under SDA limit
	if achSchedule, ok := AsACHSchedule(file.Schedules[0]); ok {
		for _, payment := range achSchedule.GetPayments() {
			if payment.GetAmount() > 100000000 { // $1M limit
				payment.SetAmount(99999999) // Clean interface usage!
			}
		}
	}

	return file
}

func createTestLargeFile() *File {
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "PERFORMANCE TEST",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   1004, // H + 01 + 1000 payments + T + E
			TotalCountPayments:  1000,
			TotalAmountPayments: 100000000, // $1M total
		},
	}

	// Create schedule with 1000 payments
	schedule := &ACHSchedule{
		Header: &ACHScheduleHeader{
			RecordCode:              "01",
			AgencyACHText:           "PERF",
			ScheduleNumber:          "00000000000001",
			PaymentTypeCode:         "Vendor",
			StandardEntryClassCode:  "PPD",
			AgencyLocationCode:      "12345678",
			FederalEmployerIDNumber: "123456789 ",
		},
		BaseSchedule: BaseSchedule{
			Payments: make([]Payment, 1000),
			Trailer: &ScheduleTrailer{
				RecordCode:     "T ",
				ScheduleCount:  1000,
				ScheduleAmount: 100000000,
			},
		},
	}

	// Add 1000 payments
	for i := 0; i < 1000; i++ {
		schedule.Payments[i] = &ACHPayment{
			RecordCode:              "02",
			AgencyAccountIdentifier: "ACC" + padLeft(fmt.Sprintf("%d", i), 13, '0'),
			Amount:                  100000, // $1,000 each
			PayeeName:               "PAYEE " + fmt.Sprintf("%d", i),
			RoutingNumber:           "021000021",
			AccountNumber:           "1234567890",
			ACHTransactionCode:      "22",
			PaymentID:               "PAY" + padLeft(fmt.Sprintf("%d", i), 17, '0'),
		}
	}

	file.Schedules = append(file.Schedules, schedule)
	file.Trailer.TotalCountRecords = 1004 // H + 01 + 1000*02 + T + E

	return file
}

// Validation test helper functions
func createFileWithInvalidRoutingNumber() *File {
	file := createTestACHFile()
	if achSchedule, ok := AsACHSchedule(file.Schedules[0]); ok {
		if achPayment, ok := AsACHPayment(achSchedule.GetPayments()[0]); ok {
			achPayment.SetRoutingNumber("123456789") // Invalid
		}
	}
	return file
}

func createUnbalancedFile() *File {
	file := createTestACHFile()
	file.Trailer.TotalAmountPayments = 999999 // Doesn't match sum
	return file
}

func createInvalidSDAFile() *File {
	file := createTestCheckFile()
	file.Header.IsRequestedForSameDayACH = "1" // Check payments not allowed for SDA
	return file
}

func createCTXWithoutAddendum() *File {
	file := createTestACHFile()
	payment := file.Schedules[0].(*ACHSchedule).Payments[0].(*ACHPayment)
	payment.StandardEntryClassCode = "CTX"
	payment.Addenda = nil // CTX requires addendum
	return file
}

func createMixedPaymentSchedule() *File {
	file := createTestACHFile()
	// Add check payment to ACH schedule (not allowed)
	file.Schedules[0].(*ACHSchedule).Payments = append(
		file.Schedules[0].(*ACHSchedule).Payments,
		&CheckPayment{RecordCode: "12"},
	)
	return file
}

// Helper to validate file structure matches
func validateFileStructure(t *testing.T, expected, actual *File) {
	// Validate header (trim spaces for comparison since fixed-width format preserves padding)
	if strings.TrimSpace(expected.Header.InputSystem) != strings.TrimSpace(actual.Header.InputSystem) {
		t.Errorf("Header mismatch: expected %s, got %s",
			strings.TrimSpace(expected.Header.InputSystem), strings.TrimSpace(actual.Header.InputSystem))
	}

	// Validate schedule count
	if len(expected.Schedules) != len(actual.Schedules) {
		t.Fatalf("Schedule count mismatch: expected %d, got %d",
			len(expected.Schedules), len(actual.Schedules))
	}

	// Validate each schedule
	for i, expectedSchedule := range expected.Schedules {
		actualSchedule := actual.Schedules[i]

		// Check schedule type
		if getScheduleType(expectedSchedule) != getScheduleType(actualSchedule) {
			t.Errorf("Schedule %d type mismatch", i)
		}

		// Validate payment count
		expectedPayments := getPayments(expectedSchedule)
		actualPayments := getPayments(actualSchedule)

		if len(expectedPayments) != len(actualPayments) {
			t.Errorf("Schedule %d payment count mismatch: expected %d, got %d",
				i, len(expectedPayments), len(actualPayments))
		}
	}

	// Validate trailer
	if expected.Trailer.TotalCountPayments != actual.Trailer.TotalCountPayments {
		t.Errorf("Trailer payment count mismatch: expected %d, got %d",
			expected.Trailer.TotalCountPayments, actual.Trailer.TotalCountPayments)
	}
	if expected.Trailer.TotalAmountPayments != actual.Trailer.TotalAmountPayments {
		t.Errorf("Trailer amount mismatch: expected %d, got %d",
			expected.Trailer.TotalAmountPayments, actual.Trailer.TotalAmountPayments)
	}
}

func getScheduleType(s Schedule) string {
	switch s.(type) {
	case *ACHSchedule:
		return "ACH"
	case *CheckSchedule:
		return "CHECK"
	default:
		return "UNKNOWN"
	}
}

func getPayments(s Schedule) []Payment {
	switch schedule := s.(type) {
	case *ACHSchedule:
		return schedule.Payments
	case *CheckSchedule:
		return schedule.Payments
	default:
		return nil
	}
}

func padLeft(s string, length int, pad rune) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(pad), length-len(s)) + s
}
