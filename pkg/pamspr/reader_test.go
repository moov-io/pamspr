package pamspr

import (
	"strings"
	"testing"
)

func TestReaderBasicFile(t *testing.T) {
	// Create a minimal valid PAM SPR file
	fileContent := createTestFile(t, []string{
		"H TEST_SYSTEM                             5020" + strings.Repeat(" ", 804),
		"01TEST00TEST001     Salary                   PPD12345678 123456789 " + strings.Repeat(" ", 783),
		"02EMP001          0000100000S1JOHN DOE                           123 MAIN ST                        " +
			"                                   WASHINGTON                 DC20001     021000021        1234567890       " +
			"22         123456789JOHN DOE SECONDARY                PAY001              " + strings.Repeat(" ", 100) +
			"123456789110000000000                                                                           " +
			"                                                       " + strings.Repeat(" ", 317),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000100000" + strings.Repeat(" ", 812),
		"E 000000000000000005000000000000000001000000000000100000" + strings.Repeat(" ", 794),
	})

	reader := NewReader(strings.NewReader(fileContent))
	file, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Verify file structure
	if file.Header == nil {
		t.Fatal("File header is nil")
	}
	if strings.TrimSpace(file.Header.InputSystem) != "TEST_SYSTEM" {
		t.Errorf("Expected input system 'TEST_SYSTEM', got %s", file.Header.InputSystem)
	}
	if file.Header.StandardPaymentVersion != "502" {
		t.Errorf("Expected version '502', got %s", file.Header.StandardPaymentVersion)
	}

	// Verify schedules
	if len(file.Schedules) != 1 {
		t.Fatalf("Expected 1 schedule, got %d", len(file.Schedules))
	}

	achSchedule, ok := file.Schedules[0].(*ACHSchedule)
	if !ok {
		t.Fatal("Expected ACH schedule")
	}
	if strings.TrimSpace(achSchedule.Header.ScheduleNumber) != "00TEST001" {
		t.Errorf("Expected schedule number '00TEST001', got %s", achSchedule.Header.ScheduleNumber)
	}

	// Verify payments
	if len(achSchedule.Payments) != 1 {
		t.Fatalf("Expected 1 payment, got %d", len(achSchedule.Payments))
	}

	payment := achSchedule.Payments[0].(*ACHPayment)
	if payment.Amount != 100000 {
		t.Errorf("Expected amount 100000, got %d", payment.Amount)
	}
	if strings.TrimSpace(payment.PayeeName) != "JOHN DOE" {
		t.Errorf("Expected payee name 'JOHN DOE', got %s", payment.PayeeName)
	}

	// Verify trailer
	if file.Trailer == nil {
		t.Fatal("File trailer is nil")
	}
	if file.Trailer.TotalCountRecords != 5 {
		t.Errorf("Expected 5 records, got %d", file.Trailer.TotalCountRecords)
	}
}

func TestReaderWithAddenda(t *testing.T) {
	fileContent := createTestFile(t, []string{
		"H TEST_SYSTEM                             5020" + strings.Repeat(" ", 804),
		"01TEST00TEST001     Salary                   PPD12345678 123456789 " + strings.Repeat(" ", 783),
		"02EMP001          0000100000S1JOHN DOE                           123 MAIN ST                        " +
			"                                   WASHINGTON                 DC20001     021000021        1234567890       " +
			"22         123456789JOHN DOE SECONDARY                PAY001              " + strings.Repeat(" ", 100) +
			"123456789110000000000                                                                           " +
			"                                                       " + strings.Repeat(" ", 317),
		"03PAY001              INVOICE 12345" + strings.Repeat(" ", 815),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000100000" + strings.Repeat(" ", 812),
		"E 000000000000000006000000000000000001000000000000100000" + strings.Repeat(" ", 794),
	})

	reader := NewReader(strings.NewReader(fileContent))
	file, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	achSchedule := file.Schedules[0].(*ACHSchedule)
	payment := achSchedule.Payments[0].(*ACHPayment)

	if len(payment.Addenda) != 1 {
		t.Fatalf("Expected 1 addendum, got %d", len(payment.Addenda))
	}

	expectedAddenda := "INVOICE 12345" + strings.Repeat(" ", 67) // 80 chars total
	if payment.Addenda[0].AddendaInformation != expectedAddenda {
		t.Errorf("Unexpected addenda information: %s", payment.Addenda[0].AddendaInformation)
	}
}

func TestReaderWithCARSTASBETC(t *testing.T) {
	fileContent := createTestFile(t, []string{
		"H TEST_SYSTEM                             5020" + strings.Repeat(" ", 804),
		"01TEST00TEST001     Salary                   PPD12345678 123456789 " + strings.Repeat(" ", 783),
		"02EMP001          0000100000S1JOHN DOE                           123 MAIN ST                        " +
			"                                   WASHINGTON                 DC20001     021000021        1234567890       " +
			"22         123456789JOHN DOE SECONDARY                PAY001              " + strings.Repeat(" ", 100) +
			"123456789110000000000                                                                           " +
			"                                                       " + strings.Repeat(" ", 317),
		"G PAY001              0101201220212021X1234001SALARY0100001000000" + strings.Repeat(" ", 785),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000100000" + strings.Repeat(" ", 812),
		"E 000000000000000006000000000000000001000000000000100000" + strings.Repeat(" ", 794),
	})

	reader := NewReader(strings.NewReader(fileContent))
	file, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	achSchedule := file.Schedules[0].(*ACHSchedule)
	payment := achSchedule.Payments[0].(*ACHPayment)

	if len(payment.CARSTASBETC) != 1 {
		t.Fatalf("Expected 1 CARS record, got %d", len(payment.CARSTASBETC))
	}

	cars := payment.CARSTASBETC[0]
	if cars.AgencyIdentifier != "012" {
		t.Errorf("Expected agency identifier '012', got %s", cars.AgencyIdentifier)
	}
	if cars.AccountClassificationAmount != 100000 {
		t.Errorf("Expected amount 100000, got %d", cars.AccountClassificationAmount)
	}
}

func TestReaderCheckSchedule(t *testing.T) {
	fileContent := createTestFile(t, []string{
		"H TEST_SYSTEM                             5020" + strings.Repeat(" ", 804),
		"1100CHK001       Vendor                   87654321" + strings.Repeat(" ", 9) + "stub      " + strings.Repeat(" ", 781),
		"12VEND001         0000150000V1ABC COMPANY LLC                    789 BUSINESS BLVD                  " +
			"SUITE 100                          " + strings.Repeat(" ", 35) + strings.Repeat(" ", 35) +
			"NEW YORK                   " + strings.Repeat(" ", 10) + "NY10001     000" + strings.Repeat(" ", 14) +
			strings.Repeat(" ", 40) + "   INVOICE #12345" + strings.Repeat(" ", 38) +
			"PO #67890" + strings.Repeat(" ", 46) + "123456789ABC SECONDARY NAME                CHK001              " +
			strings.Repeat(" ", 100) + strings.Repeat(" ", 50) + "123456789" + strings.Repeat(" ", 50) +
			"210000000000                                00" + strings.Repeat(" ", 111),
		"13CHK001              " + ("Line 1" + strings.Repeat(" ", 49)) + ("Line 2" + strings.Repeat(" ", 49)) +
			strings.Repeat(" ", 55) + strings.Repeat(" ", 55) + strings.Repeat(" ", 55) +
			strings.Repeat(" ", 55) + strings.Repeat(" ", 55) + strings.Repeat(" ", 55) +
			strings.Repeat(" ", 55) + strings.Repeat(" ", 55) + strings.Repeat(" ", 55) +
			strings.Repeat(" ", 55) + strings.Repeat(" ", 55) + strings.Repeat(" ", 55) +
			strings.Repeat(" ", 58),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000150000" + strings.Repeat(" ", 812),
		"E 000000000000000006000000000000000001000000000000150000" + strings.Repeat(" ", 794),
	})

	reader := NewReader(strings.NewReader(fileContent))
	file, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(file.Schedules) != 1 {
		t.Fatalf("Expected 1 schedule, got %d", len(file.Schedules))
	}

	checkSchedule, ok := file.Schedules[0].(*CheckSchedule)
	if !ok {
		t.Fatal("Expected check schedule")
	}

	if checkSchedule.Header.CheckPaymentEnclosureCode != "stub" {
		t.Errorf("Expected enclosure code 'stub', got %s", checkSchedule.Header.CheckPaymentEnclosureCode)
	}

	payment := checkSchedule.Payments[0].(*CheckPayment)
	if payment.Amount != 150000 {
		t.Errorf("Expected amount 150000, got %d", payment.Amount)
	}

	if payment.Stub == nil {
		t.Fatal("Expected stub record")
	}
	if payment.Stub.PaymentIdentificationLines[0] != "Line 1"+strings.Repeat(" ", 49) {
		t.Errorf("Unexpected stub line 1: %s", payment.Stub.PaymentIdentificationLines[0])
	}
}

func TestReaderMultipleSchedules(t *testing.T) {
	fileContent := createTestFile(t, []string{
		"H TEST_SYSTEM                             5020" + strings.Repeat(" ", 804),
		// First ACH schedule
		"01TEST00TEST001     Salary                   PPD12345678 123456789 " + strings.Repeat(" ", 783),
		"02EMP001          0000100000S1JOHN DOE                           123 MAIN ST                        " +
			"                                   WASHINGTON                 DC20001     021000021        1234567890       " +
			"22         123456789JOHN DOE SECONDARY                PAY001              " + strings.Repeat(" ", 100) +
			"123456789110000000000                                                                           " +
			"                                                       " + strings.Repeat(" ", 317),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000100000" + strings.Repeat(" ", 812),
		// Second ACH schedule
		"01TEST00TEST002     Vendor                   CCD12345678 123456789 " + strings.Repeat(" ", 783),
		"02VEND001         0000200000V1VENDOR COMPANY                     456 OAK AVE                        " +
			"                                   ARLINGTON                  VA22201     021000021        0987654321       " +
			"32         987654321                                   PAY002              " + strings.Repeat(" ", 100) +
			"987654321210000000000                                                                           " +
			"                                                       " + strings.Repeat(" ", 316),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000200000" + strings.Repeat(" ", 812),
		"E 000000000000000008000000000000000002000000000000300000" + strings.Repeat(" ", 794),
	})

	reader := NewReader(strings.NewReader(fileContent))
	file, err := reader.Read()
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if len(file.Schedules) != 2 {
		t.Fatalf("Expected 2 schedules, got %d", len(file.Schedules))
	}

	// Verify first schedule
	schedule1 := file.Schedules[0].(*ACHSchedule)
	if schedule1.Header.PaymentTypeCode != "Salary"+strings.Repeat(" ", 19) {
		t.Errorf("Expected payment type 'Salary', got %s", schedule1.Header.PaymentTypeCode)
	}

	// Verify second schedule
	schedule2 := file.Schedules[1].(*ACHSchedule)
	if schedule2.Header.PaymentTypeCode != "Vendor"+strings.Repeat(" ", 19) {
		t.Errorf("Expected payment type 'Vendor', got %s", schedule2.Header.PaymentTypeCode)
	}

	// Verify totals
	if file.Trailer.TotalCountPayments != 2 {
		t.Errorf("Expected 2 payments, got %d", file.Trailer.TotalCountPayments)
	}
	if file.Trailer.TotalAmountPayments != 300000 {
		t.Errorf("Expected total amount 300000, got %d", file.Trailer.TotalAmountPayments)
	}
}

func TestReaderInvalidRecordLength(t *testing.T) {
	fileContent := "H TEST_SYSTEM 502 0" // Too short

	reader := NewReader(strings.NewReader(fileContent))
	_, err := reader.Read()
	if err == nil {
		t.Error("Expected error for invalid record length")
	}
	if !strings.Contains(err.Error(), "invalid record length") {
		t.Errorf("Expected invalid record length error, got: %v", err)
	}
}

func TestReaderMissingFileTrailer(t *testing.T) {
	fileContent := createTestFile(t, []string{
		"H TEST_SYSTEM                             5020" + strings.Repeat(" ", 804),
		"01TEST00TEST001     Salary                   PPD12345678 123456789 " + strings.Repeat(" ", 783),
		"02EMP001          0000100000S1JOHN DOE                           123 MAIN ST                        " +
			"                                   WASHINGTON                 DC20001     021000021        1234567890       " +
			"22         123456789JOHN DOE SECONDARY                PAY001              " + strings.Repeat(" ", 100) +
			"123456789110000000000                                                                           " +
			"                                                       " + strings.Repeat(" ", 317),
		"T " + strings.Repeat(" ", 10) + "00000001" + strings.Repeat(" ", 3) + "000000000100000" + strings.Repeat(" ", 812),
		// Missing file trailer
	})

	reader := NewReader(strings.NewReader(fileContent))
	_, err := reader.Read()
	if err == nil {
		t.Error("Expected error for missing file trailer")
	}
}

func TestReaderPushBackLine(t *testing.T) {
	reader := &Reader{}

	// Test pushing back a line
	reader.pushBackLine("TEST LINE")
	if reader.nextLine == nil || *reader.nextLine != "TEST LINE" {
		t.Error("Failed to push back line")
	}

	// Test scanning after pushback
	line, ok := reader.scanLine()
	if !ok || line != "TEST LINE" {
		t.Error("Failed to read pushed back line")
	}
	if reader.nextLine != nil {
		t.Error("nextLine should be nil after reading")
	}
}

// Helper function to create test file content
func createTestFile(t *testing.T, lines []string) string {
	t.Helper()
	var result strings.Builder
	for i, line := range lines {
		if len(line) != RecordLength {
			t.Fatalf("Line %d has invalid length: expected %d, got %d", i+1, RecordLength, len(line))
		}
		result.WriteString(line)
		if i < len(lines)-1 {
			result.WriteString("\n")
		}
	}
	return result.String()
}
