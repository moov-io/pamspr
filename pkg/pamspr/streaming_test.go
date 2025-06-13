package pamspr

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// TestReader_ProcessPaymentsOnly tests payment-only processing
func TestReader_ProcessPaymentsOnly(t *testing.T) {
	testData := createTestFileContent()
	reader := NewReader(strings.NewReader(testData))

	var payments []Payment
	var scheduleIndices []int
	var paymentIndices []int

	err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
		payments = append(payments, payment)
		scheduleIndices = append(scheduleIndices, scheduleIndex)
		paymentIndices = append(paymentIndices, paymentIndex)
		return true
	})

	if err != nil {
		t.Fatalf("ProcessPaymentsOnly failed: %v", err)
	}

	// Verify we got the expected payments
	if len(payments) != 2 {
		t.Errorf("Expected 2 payments, got %d", len(payments))
	}

	// Verify payment details
	if achPayment, ok := payments[0].(*ACHPayment); ok {
		if achPayment.Amount != 150000 {
			t.Errorf("Expected first payment amount 150000, got %d", achPayment.Amount)
		}
		payeeName := strings.TrimSpace(achPayment.PayeeName)
		if payeeName != "ACME CORPORATION" {
			t.Errorf("Expected payee name 'ACME CORPORATION', got '%s'", payeeName)
		}
	} else {
		t.Error("First payment should be ACH payment")
	}

	// Verify statistics
	stats := reader.GetStats()
	if stats.PaymentsProcessed != 2 {
		t.Errorf("Expected 2 payments processed, got %d", stats.PaymentsProcessed)
	}
	if stats.LinesProcessed == 0 {
		t.Error("Expected lines to be processed")
	}
}

// TestReader_ProcessFile tests full file processing with callbacks
func TestReader_ProcessFile(t *testing.T) {
	testData := createTestFileContent()
	reader := NewReader(strings.NewReader(testData))

	var schedules []Schedule
	var payments []Payment
	var recordTypes []string

	scheduleCallback := func(schedule Schedule, scheduleIndex int) bool {
		schedules = append(schedules, schedule)
		return true
	}

	paymentCallback := func(payment Payment, scheduleIndex, paymentIndex int) bool {
		payments = append(payments, payment)
		return true
	}

	recordCallback := func(recordType string, lineNumber int, line string) {
		recordTypes = append(recordTypes, recordType)
	}

	err := reader.ProcessFile(scheduleCallback, paymentCallback, recordCallback)
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	// Verify schedules
	if len(schedules) != 1 {
		t.Errorf("Expected 1 schedule, got %d", len(schedules))
	}

	// Verify payments
	if len(payments) != 2 {
		t.Errorf("Expected 2 payments, got %d", len(payments))
	}

	// Verify record types were captured (includes schedule trailer now)
	expectedRecordTypes := []string{"01", "02", "02", "T "}
	if len(recordTypes) < len(expectedRecordTypes) {
		t.Errorf("Expected at least %d record types, got %d", len(expectedRecordTypes), len(recordTypes))
	}
}

// TestReader_ValidateFileStructureOnly tests structure-only validation
func TestReader_ValidateFileStructureOnly(t *testing.T) {
	testData := createTestFileContent()
	reader := NewReader(strings.NewReader(testData))

	err := reader.ValidateFileStructureOnly()
	if err != nil {
		t.Fatalf("ValidateFileStructureOnly failed: %v", err)
	}

	stats := reader.GetStats()
	if stats.LinesProcessed == 0 {
		t.Error("Expected lines to be processed")
	}
}

// TestReader_Configuration tests different reader configurations
func TestReader_Configuration(t *testing.T) {
	testData := createTestFileContent()

	// Test with custom configuration
	config := &ReaderConfig{
		BufferSize:         32 * 1024,
		EnableValidation:   false,
		CollectErrors:      true,
		MaxErrors:          10,
		SkipInvalidRecords: true,
	}

	reader := NewReaderWithConfig(strings.NewReader(testData), config)

	paymentCount := 0
	err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
		paymentCount++
		return true
	})

	if err != nil {
		t.Fatalf("ProcessPaymentsOnly with custom config failed: %v", err)
	}

	if paymentCount != 2 {
		t.Errorf("Expected 2 payments, got %d", paymentCount)
	}
}

// TestReader_ErrorHandling tests error handling and recovery
func TestReader_ErrorHandling(t *testing.T) {
	// Create test data with an invalid record code that will definitely be caught
	validData := createTestFileContent()
	lines := strings.Split(strings.TrimRight(validData, "\n"), "\n")

	// Replace a valid record with an invalid record code that will be caught
	// We'll add a completely invalid record after the first payment
	insertIndex := -1
	for i, line := range lines {
		if strings.HasPrefix(line, "02") {
			insertIndex = i + 1
			break
		}
	}

	if insertIndex > 0 {
		// Insert an invalid record
		invalidLine := "ZZ" + strings.Repeat("X", 848) // Invalid record code "ZZ"
		newLines := make([]string, 0, len(lines)+1)
		newLines = append(newLines, lines[:insertIndex]...)
		newLines = append(newLines, invalidLine)
		newLines = append(newLines, lines[insertIndex:]...)
		lines = newLines
	}

	testData := strings.Join(lines, "\n") + "\n"

	// Test with skip invalid records enabled - should continue but collect errors
	config := DefaultConfig()
	config.SkipInvalidRecords = true
	config.CollectErrors = true

	reader := NewReaderWithConfig(strings.NewReader(testData), config)

	paymentCount := 0
	err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
		paymentCount++
		return true
	})

	// Should not fail with skip enabled
	if err != nil {
		t.Fatalf("Expected processing to continue with skip enabled, got error: %v", err)
	}

	// Should have processed the valid payments
	if paymentCount == 0 {
		t.Error("Expected some payments to be processed")
	}

	// Test with skip invalid records disabled - should fail fast
	config2 := DefaultConfig()
	config2.SkipInvalidRecords = false
	reader2 := NewReaderWithConfig(strings.NewReader(testData), config2)

	err = reader2.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
		return true
	})

	// Should fail with skip disabled
	if err == nil {
		t.Error("Expected error with skip disabled")
	}
}

// TestStreamingWriter tests the streaming writer functionality
func TestStreamingWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Write file header
	header := &FileHeader{
		RecordCode:               "H ",
		InputSystem:              "TEST STREAMING SYSTEM",
		StandardPaymentVersion:   "502",
		IsRequestedForSameDayACH: "0",
	}

	err := writer.WriteFileHeader(header)
	if err != nil {
		t.Fatalf("WriteFileHeader failed: %v", err)
	}

	// Write schedule
	schedule := &ACHSchedule{
		Header: &ACHScheduleHeader{
			RecordCode:              "01",
			AgencyACHText:           "TEST",
			ScheduleNumber:          "00000000000001",
			PaymentTypeCode:         "Vendor",
			StandardEntryClassCode:  "CCD",
			AgencyLocationCode:      "12345678",
			FederalEmployerIDNumber: "1234567890",
		},
	}

	err = writer.WriteScheduleHeader(schedule)
	if err != nil {
		t.Fatalf("WriteScheduleHeader failed: %v", err)
	}

	// Write payment
	payment := &ACHPayment{
		RecordCode:                   "02",
		AgencyAccountIdentifier:      "ACC123456789012 ",
		Amount:                       100000,
		AgencyPaymentTypeCode:        "V",
		IsTOP_Offset:                 "0",
		PayeeName:                    "TEST PAYEE",
		PayeeAddressLine1:            "123 TEST STREET",
		CityName:                     "TEST CITY",
		StateName:                    "TEST STATE",
		StateCodeText:                "TS",
		PostalCode:                   "12345",
		RoutingNumber:                "021000021",
		AccountNumber:                "123456789",
		ACH_TransactionCode:          "22",
		PaymentID:                    "PAY000000001",
		TIN:                          "123456789",
		PaymentRecipientTINIndicator: "1",
		AdditionalPayeeTINIndicator:  "0",
		AmountEligibleForOffset:      "0000100000",
	}

	err = writer.WritePayment(payment)
	if err != nil {
		t.Fatalf("WritePayment failed: %v", err)
	}

	// Write schedule trailer
	scheduleTrailer := &ScheduleTrailer{
		RecordCode:     "T ",
		ScheduleCount:  1,
		ScheduleAmount: 100000,
	}

	err = writer.WriteScheduleTrailer(scheduleTrailer)
	if err != nil {
		t.Fatalf("WriteScheduleTrailer failed: %v", err)
	}

	// Write file trailer
	fileTrailer := &FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   5, // H + 01 + 02 + T + E
		TotalCountPayments:  1,
		TotalAmountPayments: 100000,
	}

	err = writer.WriteFileTrailer(fileTrailer)
	if err != nil {
		t.Fatalf("WriteFileTrailer failed: %v", err)
	}

	// Verify output
	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	if len(lines) != 5 {
		t.Errorf("Expected 5 lines, got %d", len(lines))
	}

	// Verify line lengths
	for i, line := range lines {
		if len(line) != RecordLength {
			t.Errorf("Line %d has incorrect length: expected %d, got %d", i+1, RecordLength, len(line))
		}
	}

	// Verify record codes
	expectedCodes := []string{"H ", "01", "02", "T ", "E "}
	for i, expectedCode := range expectedCodes {
		if !strings.HasPrefix(lines[i], expectedCode) {
			t.Errorf("Line %d: expected record code %q, got %q", i+1, expectedCode, lines[i][:2])
		}
	}

	// Check statistics
	recordCount, paymentCount, scheduleCount, totalAmount := writer.GetStats()
	if recordCount != 5 {
		t.Errorf("Expected 5 records written, got %d", recordCount)
	}
	if paymentCount != 1 {
		t.Errorf("Expected 1 payment written, got %d", paymentCount)
	}
	if scheduleCount != 1 {
		t.Errorf("Expected 1 schedule written, got %d", scheduleCount)
	}
	if totalAmount != 100000 {
		t.Errorf("Expected total amount 100000, got %d", totalAmount)
	}
}

// TestWriter_Configuration tests streaming writer configuration
func TestWriter_Configuration(t *testing.T) {
	var buf bytes.Buffer

	config := &WriterConfig{
		BufferSize:         16 * 1024,
		EnableValidation:   true,
		FlushInterval:      1, // Flush after every record
		ChecksumValidation: false,
	}

	writer := NewWriterWithConfig(&buf, config)

	// Write minimal file
	header := &FileHeader{
		RecordCode:               "H ",
		InputSystem:              "CONFIG TEST",
		StandardPaymentVersion:   "502",
		IsRequestedForSameDayACH: "0",
	}

	err := writer.WriteFileHeader(header)
	if err != nil {
		t.Fatalf("WriteFileHeader with config failed: %v", err)
	}

	trailer := &FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   2,
		TotalCountPayments:  0,
		TotalAmountPayments: 0,
	}

	err = writer.WriteFileTrailer(trailer)
	if err != nil {
		t.Fatalf("WriteFileTrailer with config failed: %v", err)
	}

	output := buf.String()
	if len(output) == 0 {
		t.Error("Expected output to be written")
	}
}

// TestStreamingRoundTrip tests reading and writing the same file
func TestStreamingRoundTrip(t *testing.T) {
	originalData := createTestFileContent()

	// Parse with streaming reader
	reader := NewReader(strings.NewReader(originalData))

	var schedules []Schedule
	var allPayments []Payment

	// Collect all data
	err := reader.ProcessFile(
		func(schedule Schedule, scheduleIndex int) bool {
			schedules = append(schedules, schedule)
			return true
		},
		func(payment Payment, scheduleIndex, paymentIndex int) bool {
			allPayments = append(allPayments, payment)
			return true
		},
		nil,
	)
	if err != nil {
		t.Fatalf("ProcessFile failed: %v", err)
	}

	// Create minimal header and trailer for test
	header := &FileHeader{
		RecordCode:               "H ",
		InputSystem:              "TEST SYSTEM",
		StandardPaymentVersion:   "502",
		IsRequestedForSameDayACH: "0",
	}

	trailer := &FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   5, // Approximate
		TotalCountPayments:  int64(len(allPayments)),
		TotalAmountPayments: 0, // Will be calculated
	}

	// Calculate total amount
	for _, payment := range allPayments {
		trailer.TotalAmountPayments += payment.GetAmount()
	}

	// Write with streaming writer
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	err = writer.WriteFileHeader(header)
	if err != nil {
		t.Fatalf("WriteFileHeader failed: %v", err)
	}

	// Write first schedule and its payments
	if len(schedules) > 0 {
		err = writer.WriteScheduleHeader(schedules[0])
		if err != nil {
			t.Fatalf("WriteScheduleHeader failed: %v", err)
		}

		for _, payment := range allPayments {
			err = writer.WritePayment(payment)
			if err != nil {
				t.Fatalf("WritePayment failed: %v", err)
			}
		}

		scheduleTrailer := &ScheduleTrailer{
			RecordCode:     "T ",
			ScheduleCount:  int64(len(allPayments)),
			ScheduleAmount: trailer.TotalAmountPayments,
		}

		err = writer.WriteScheduleTrailer(scheduleTrailer)
		if err != nil {
			t.Fatalf("WriteScheduleTrailer failed: %v", err)
		}
	}

	err = writer.WriteFileTrailer(trailer)
	if err != nil {
		t.Fatalf("WriteFileTrailer failed: %v", err)
	}

	// Verify output is valid
	output := buf.String()
	if len(output) == 0 {
		t.Error("No output generated")
	}

	// Parse the generated output to verify it's valid
	reader2 := NewReader(strings.NewReader(output))
	err = reader2.ValidateFileStructureOnly()
	if err != nil {
		t.Fatalf("Round-trip validation failed: %v", err)
	}
}

// TestReader_EarlyTermination tests stopping processing early
func TestReader_EarlyTermination(t *testing.T) {
	testData := createTestFileContent()
	reader := NewReader(strings.NewReader(testData))

	paymentCount := 0
	err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
		paymentCount++
		// Stop after first payment
		return paymentCount < 1
	})

	if err != nil {
		t.Fatalf("ProcessPaymentsOnly failed: %v", err)
	}

	if paymentCount != 1 {
		t.Errorf("Expected processing to stop after 1 payment, but processed %d", paymentCount)
	}
}

// Helper function to create test file content using existing writer
func createTestFileContent() string {
	// Create test data using existing structures to ensure proper formatting
	file := &File{
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
					AgencyACHText:           "AGCY",
					ScheduleNumber:          "00000000000001",
					PaymentTypeCode:         "Vendor",
					StandardEntryClassCode:  "CCD",
					AgencyLocationCode:      "12345678",
					FederalEmployerIDNumber: "1234567890",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{
							RecordCode:                   "02",
							AgencyAccountIdentifier:      "ACC123456789012 ",
							Amount:                       150000,
							AgencyPaymentTypeCode:        "V",
							IsTOP_Offset:                 "0",
							PayeeName:                    "ACME CORPORATION",
							PayeeAddressLine1:            "123 BUSINESS STREET",
							PayeeAddressLine2:            "SUITE 100",
							CityName:                     "NEW YORK",
							StateName:                    "NEW YORK",
							StateCodeText:                "NY",
							PostalCode:                   "10101",
							PostalCodeExtension:          "1234",
							CountryCodeText:              "US",
							RoutingNumber:                "021000021",
							AccountNumber:                "123456789012345",
							ACH_TransactionCode:          "22",
							PaymentID:                    "PAY000000001",
							TIN:                          "123456789",
							PaymentRecipientTINIndicator: "1",
							AdditionalPayeeTINIndicator:  "0",
							AmountEligibleForOffset:      "0000150000",
							CountryName:                  "UNITED STATES",
							SubPaymentTypeCode:           "VENDOR PAYMENT",
							PayerMechanism:               "ACH CREDIT",
						},
						&ACHPayment{
							RecordCode:                   "02",
							AgencyAccountIdentifier:      "ACC987654321098 ",
							Amount:                       75000,
							AgencyPaymentTypeCode:        "S",
							IsTOP_Offset:                 "1",
							PayeeName:                    "JOHN DOE",
							PayeeAddressLine1:            "456 MAIN STREET",
							CityName:                     "LOS ANGELES",
							StateName:                    "CALIFORNIA",
							StateCodeText:                "CA",
							PostalCode:                   "90001",
							CountryCodeText:              "US",
							RoutingNumber:                "122000247",
							AccountNumber:                "987654321",
							ACH_TransactionCode:          "27",
							PaymentID:                    "PAY000000002",
							TIN:                          "987654321",
							PaymentRecipientTINIndicator: "1",
							AdditionalPayeeTINIndicator:  "0",
							AmountEligibleForOffset:      "0000075000",
							CountryName:                  "UNITED STATES",
							SubPaymentTypeCode:           "SALARY PAYMENT",
							PayerMechanism:               "ACH DEBIT",
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  2,
						ScheduleAmount: 225000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   5,
			TotalCountPayments:  2,
			TotalAmountPayments: 225000,
		},
	}

	// Use existing writer to generate properly formatted content
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	err := writer.Write(file)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate test content: %v", err))
	}

	return buf.String()
}
