package pamspr

import (
	"strings"
	"testing"
)

// TestFieldPadding_PadLeft tests left padding functionality
func TestFieldPadding_PadLeft(t *testing.T) {
	fp := &FieldPadding{}

	tests := []struct {
		name     string
		input    string
		length   int
		padChar  rune
		expected string
	}{
		{"pad with zeros", "123", 5, '0', "00123"},
		{"pad with spaces", "ABC", 7, ' ', "    ABC"},
		{"exact length", "12345", 5, '0', "12345"},
		{"truncate when too long", "123456789", 5, '0', "12345"},
		{"empty string", "", 3, 'X', "XXX"},
		{"single character", "A", 4, '-', "---A"},
		{"pad with special char", "test", 8, '*', "****test"},
		{"special chars", "test", 10, '0', "000000test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.PadLeft(tt.input, tt.length, tt.padChar)
			if result != tt.expected {
				t.Errorf("PadLeft(%q, %d, %q) = %q, want %q", tt.input, tt.length, tt.padChar, result, tt.expected)
			}
			if len(result) != tt.length && len(tt.input) < tt.length {
				t.Errorf("PadLeft result length = %d, want %d", len(result), tt.length)
			}
		})
	}
}

// TestFieldPadding_PadRight tests right padding functionality
func TestFieldPadding_PadRight(t *testing.T) {
	fp := &FieldPadding{}

	tests := []struct {
		name     string
		input    string
		length   int
		padChar  rune
		expected string
	}{
		{"pad with spaces", "ABC", 7, ' ', "ABC    "},
		{"pad with zeros", "123", 5, '0', "12300"},
		{"exact length", "12345", 5, ' ', "12345"},
		{"truncate when too long", "123456789", 5, ' ', "12345"},
		{"empty string", "", 4, 'Y', "YYYY"},
		{"single character", "Z", 3, '+', "Z++"},
		{"pad with special char", "test", 8, '#', "test####"},
		{"long word", "word", 10, '0', "word000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.PadRight(tt.input, tt.length, tt.padChar)
			if result != tt.expected {
				t.Errorf("PadRight(%q, %d, %q) = %q, want %q", tt.input, tt.length, tt.padChar, result, tt.expected)
			}
			if len(result) != tt.length && len(tt.input) < tt.length {
				t.Errorf("PadRight result length = %d, want %d", len(result), tt.length)
			}
		})
	}
}

// TestFieldPadding_PadNumeric tests numeric padding functionality
func TestFieldPadding_PadNumeric(t *testing.T) {
	fp := &FieldPadding{}

	tests := []struct {
		name     string
		input    string
		length   int
		expected string
	}{
		{"simple number", "123", 5, "00123"},
		{"with spaces", "1 2 3", 5, "00123"},
		{"with letters", "ABC123DEF", 5, "00123"},
		{"with special chars", "1-2-3", 5, "00123"},
		{"mixed alphanumeric", "A1B2C3", 5, "00123"},
		{"empty string", "", 3, "000"},
		{"no digits", "ABC", 3, "000"},
		{"leading zeros", "000123", 8, "00000123"},
		{"phone number format", "(555) 123-4567", 10, "5551234567"},
		{"decimal number", "12.34", 5, "01234"},
		{"negative sign", "-123", 5, "00123"},
		{"currency format", "$123.45", 5, "12345"},
		{"scientific notation", "1.23E+4", 5, "01234"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.PadNumeric(tt.input, tt.length)
			if result != tt.expected {
				t.Errorf("PadNumeric(%q, %d) = %q, want %q", tt.input, tt.length, result, tt.expected)
			}
			if len(result) != tt.length {
				t.Errorf("PadNumeric result length = %d, want %d", len(result), tt.length)
			}
		})
	}
}

// TestFieldPadding_TruncateOrPad tests truncation and padding
func TestFieldPadding_TruncateOrPad(t *testing.T) {
	fp := &FieldPadding{}

	tests := []struct {
		name     string
		input    string
		length   int
		padRight bool
		expected string
	}{
		{"pad right with spaces", "ABC", 7, true, "ABC    "},
		{"pad left with spaces", "ABC", 7, false, "    ABC"},
		{"truncate from right", "ABCDEFGH", 5, true, "ABCDE"},
		{"truncate from left", "ABCDEFGH", 5, false, "ABCDE"},
		{"exact length right", "ABCDE", 5, true, "ABCDE"},
		{"exact length left", "ABCDE", 5, false, "ABCDE"},
		{"empty string right", "", 3, true, "   "},
		{"empty string left", "", 3, false, "   "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fp.TruncateOrPad(tt.input, tt.length, tt.padRight)
			if result != tt.expected {
				t.Errorf("TruncateOrPad(%q, %d, %t) = %q, want %q", tt.input, tt.length, tt.padRight, result, tt.expected)
			}
			if len(result) != tt.length {
				t.Errorf("TruncateOrPad result length = %d, want %d", len(result), tt.length)
			}
		})
	}
}

// TestFormatUtils_FormatAmount tests amount formatting to dollar strings
func TestFormatUtils_FormatAmount(t *testing.T) {
	fu := &FormatUtils{}

	tests := []struct {
		name     string
		cents    int64
		expected string
	}{
		{"zero amount", 0, "0.00"},
		{"simple amount", 12345, "123.45"},
		{"single digit cents", 105, "1.05"},
		{"no cents", 10000, "100.00"},
		{"large amount", 999999999, "9999999.99"},
		{"negative amount", -12345, "-123.45"},
		{"single cent", 1, "0.01"},
		{"ten cents", 10, "0.10"},
		{"million dollars", 100000000, "1000000.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fu.FormatAmount(tt.cents)
			if result != tt.expected {
				t.Errorf("FormatAmount(%d) = %q, want %q", tt.cents, result, tt.expected)
			}
		})
	}
}

// TestFormatUtils_ParseAmount tests parsing dollar amounts to cents
func TestFormatUtils_ParseAmount(t *testing.T) {
	fu := &FormatUtils{}

	tests := []struct {
		name     string
		input    string
		expected int64
		wantErr  bool
	}{
		{"simple amount", "123.45", 12345, false},
		{"no decimal", "123", 12300, false},
		{"with dollar sign", "$123.45", 12345, false},
		{"with commas", "1,234.56", 123456, false},
		{"single decimal digit", "123.5", 12350, false},
		{"three decimal digits", "123.456", 12345, false},
		{"zero amount", "0.00", 0, false},
		{"zero no decimal", "0", 0, false},
		{"just decimal", ".50", 50, false},
		{"multiple decimals", "12.34.56", 1234, false}, // Only first decimal counts
		{"letters and numbers", "ABC123.45DEF", 12345, false},
		{"only letters", "ABCDEF", 0, false},
		{"empty string", "", 0, false},
		{"spaces", " 123.45 ", 12345, false},
		{"negative sign", "-123.45", 12345, false}, // Negative sign is ignored
		{"parentheses", "(123.45)", 12345, false},
		{"currency with cents", "$1,234.56", 123456, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := fu.ParseAmount(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAmount(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if result != tt.expected {
				t.Errorf("ParseAmount(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// TestFormatUtils_FormatTIN tests TIN formatting with dashes
func TestFormatUtils_FormatTIN(t *testing.T) {
	fu := &FormatUtils{}

	tests := []struct {
		name     string
		tin      string
		tinType  string
		expected string
	}{
		{"SSN format", "123456789", "1", "123-45-6789"},
		{"EIN format", "123456789", "2", "12-3456789"},
		{"ITIN format", "123456789", "3", "123456789"},
		{"unknown type", "123456789", "9", "123456789"},
		{"with existing dashes SSN", "123-45-6789", "1", "123-45-6789"},
		{"with spaces", "123 456 789", "1", "123-45-6789"},
		{"with letters", "12A34B6789", "1", "12A34B6789"},
		{"too short", "12345", "1", "12345"},
		{"too long", "1234567890", "1", "1234567890"},
		{"empty string", "", "1", ""},
		{"all letters", "ABCDEFGHI", "1", "ABCDEFGHI"},
		{"mixed with special chars", "1(2)3-4[5]6{7}8|9", "1", "123-45-6789"},
		{"EIN with letters", "AB123456C", "2", "AB123456C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fu.FormatTIN(tt.tin, tt.tinType)
			if result != tt.expected {
				t.Errorf("FormatTIN(%q, %q) = %q, want %q", tt.tin, tt.tinType, result, tt.expected)
			}
		})
	}
}

// TestFormatUtils_CleanAddress tests address cleaning functionality
func TestFormatUtils_CleanAddress(t *testing.T) {
	fu := &FormatUtils{}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"replace quotes", `123 "Main" St`, "123 'Main' St"},
		{"replace angle brackets", "123 <Main> St", "123 (Main) St"},
		{"replace ampersand", "Smith & Jones", "Smith + Jones"},
		{"remove non-printable", "123\x00\x01\x02 Main St", "123 Main St"},
		{"multiple spaces", "123    Main     St", "123 Main St"},
		{"leading/trailing spaces", "  123 Main St  ", "123 Main St"},
		{"mixed problematic chars", `123 "Main" & <Company>`, "123 'Main' + (Company)"},
		{"unicode characters", "123 Straße München", "123 Stra e M nchen"},
		{"control characters", "123\tMain\nSt\r", "123 Main St"},
		{"empty string", "", ""},
		{"only spaces", "   ", ""},
		{"only problematic chars", `"<>&`, "'()+"},
		{"normal address", "123 Main Street", "123 Main Street"},
		{"apartment number", "123 Main St Apt #5", "123 Main St Apt #5"},
		{"special chars in name", "O'Brien & Associates", "O'Brien + Associates"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := fu.CleanAddress(tt.input)
			if result != tt.expected {
				t.Errorf("CleanAddress(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestNewFileBuilder tests FileBuilder creation
func TestNewFileBuilder(t *testing.T) {
	fb := NewFileBuilder()
	if fb == nil {
		t.Fatal("NewFileBuilder returned nil")
	}
	if fb.file == nil {
		t.Error("FileBuilder file is nil")
	}
	if fb.file.Schedules == nil {
		t.Error("FileBuilder schedules is nil")
	}
	if len(fb.file.Schedules) != 0 {
		t.Errorf("FileBuilder schedules length = %d, want 0", len(fb.file.Schedules))
	}
	if fb.errors == nil {
		t.Error("FileBuilder errors is nil")
	}
	if len(fb.errors) != 0 {
		t.Errorf("FileBuilder errors length = %d, want 0", len(fb.errors))
	}
}

// TestFileBuilder_WithHeader tests header setting
func TestFileBuilder_WithHeader(t *testing.T) {
	tests := []struct {
		name        string
		inputSystem string
		version     string
		sameDayACH  bool
		expectedSDA string
	}{
		{"with same day ACH", "TEST SYSTEM", "502", true, "1"},
		{"without same day ACH", "TEST SYSTEM", "502", false, "0"},
		{"empty input system", "", "502", false, "0"},
		{"different version", "PROD SYSTEM", "501", true, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := NewFileBuilder()
			result := fb.WithHeader(tt.inputSystem, tt.version, tt.sameDayACH)

			// Should return same builder for chaining
			if result != fb {
				t.Error("WithHeader should return same FileBuilder for chaining")
			}

			if fb.file.Header == nil {
				t.Fatal("Header was not set")
			}

			if fb.file.Header.RecordCode != "H " {
				t.Errorf("Header RecordCode = %q, want %q", fb.file.Header.RecordCode, "H ")
			}
			if fb.file.Header.InputSystem != tt.inputSystem {
				t.Errorf("Header InputSystem = %q, want %q", fb.file.Header.InputSystem, tt.inputSystem)
			}
			if fb.file.Header.StandardPaymentVersion != tt.version {
				t.Errorf("Header StandardPaymentVersion = %q, want %q", fb.file.Header.StandardPaymentVersion, tt.version)
			}
			if fb.file.Header.IsRequestedForSameDayACH != tt.expectedSDA {
				t.Errorf("Header IsRequestedForSameDayACH = %q, want %q", fb.file.Header.IsRequestedForSameDayACH, tt.expectedSDA)
			}
		})
	}
}

// TestFileBuilder_StartACHSchedule tests ACH schedule creation
func TestFileBuilder_StartACHSchedule(t *testing.T) {
	fb := NewFileBuilder()

	result := fb.StartACHSchedule("000000000001", "Vendor", "12345678", "CCD")

	// Should return same builder for chaining
	if result != fb {
		t.Error("StartACHSchedule should return same FileBuilder for chaining")
	}

	if fb.currentSchedule == nil {
		t.Fatal("Current schedule was not set")
	}

	achSchedule, ok := fb.currentSchedule.(*ACHSchedule)
	if !ok {
		t.Fatal("Current schedule is not an ACH schedule")
	}

	if achSchedule.Header.RecordCode != "01" {
		t.Errorf("ACH schedule RecordCode = %q, want %q", achSchedule.Header.RecordCode, "01")
	}
	if achSchedule.Header.ScheduleNumber != "000000000001" {
		t.Errorf("ACH schedule ScheduleNumber = %q, want %q", achSchedule.Header.ScheduleNumber, "000000000001")
	}
	if achSchedule.Header.PaymentTypeCode != "Vendor" {
		t.Errorf("ACH schedule PaymentTypeCode = %q, want %q", achSchedule.Header.PaymentTypeCode, "Vendor")
	}
	if achSchedule.Header.AgencyLocationCode != "12345678" {
		t.Errorf("ACH schedule AgencyLocationCode = %q, want %q", achSchedule.Header.AgencyLocationCode, "12345678")
	}
	if achSchedule.Header.StandardEntryClassCode != "CCD" {
		t.Errorf("ACH schedule StandardEntryClassCode = %q, want %q", achSchedule.Header.StandardEntryClassCode, "CCD")
	}
}

// TestFileBuilder_StartCheckSchedule tests Check schedule creation
func TestFileBuilder_StartCheckSchedule(t *testing.T) {
	fb := NewFileBuilder()

	result := fb.StartCheckSchedule("000000000002", "Refund", "87654321", "stub")

	// Should return same builder for chaining
	if result != fb {
		t.Error("StartCheckSchedule should return same FileBuilder for chaining")
	}

	if fb.currentSchedule == nil {
		t.Fatal("Current schedule was not set")
	}

	checkSchedule, ok := fb.currentSchedule.(*CheckSchedule)
	if !ok {
		t.Fatal("Current schedule is not a Check schedule")
	}

	if checkSchedule.Header.RecordCode != "11" {
		t.Errorf("Check schedule RecordCode = %q, want %q", checkSchedule.Header.RecordCode, "11")
	}
	if checkSchedule.Header.ScheduleNumber != "000000000002" {
		t.Errorf("Check schedule ScheduleNumber = %q, want %q", checkSchedule.Header.ScheduleNumber, "000000000002")
	}
	if checkSchedule.Header.PaymentTypeCode != "Refund" {
		t.Errorf("Check schedule PaymentTypeCode = %q, want %q", checkSchedule.Header.PaymentTypeCode, "Refund")
	}
	if checkSchedule.Header.AgencyLocationCode != "87654321" {
		t.Errorf("Check schedule AgencyLocationCode = %q, want %q", checkSchedule.Header.AgencyLocationCode, "87654321")
	}
	if checkSchedule.Header.CheckPaymentEnclosureCode != "stub" {
		t.Errorf("Check schedule CheckPaymentEnclosureCode = %q, want %q", checkSchedule.Header.CheckPaymentEnclosureCode, "stub")
	}
}

// TestFileBuilder_AddACHPayment tests adding ACH payments
func TestFileBuilder_AddACHPayment(t *testing.T) {
	fb := NewFileBuilder()

	// Test adding payment without schedule (should error)
	payment := &ACHPayment{
		AgencyAccountIdentifier: "TEST123",
		Amount:                  15000,
		PayeeName:               "John Doe",
		PaymentID:               "PAY001",
	}

	fb.AddACHPayment(payment)
	if len(fb.errors) == 0 {
		t.Error("Expected error when adding payment without schedule")
	}

	// Test adding payment to ACH schedule
	fb = NewFileBuilder()
	fb.StartACHSchedule("000000000001", "Vendor", "12345678", "CCD")
	fb.AddACHPayment(payment)

	if len(fb.errors) != 0 {
		t.Errorf("Unexpected errors: %v", fb.errors)
	}

	achSchedule := fb.currentSchedule.(*ACHSchedule)
	if len(achSchedule.Payments) != 1 {
		t.Errorf("Expected 1 payment, got %d", len(achSchedule.Payments))
	}

	addedPayment := achSchedule.Payments[0].(*ACHPayment)
	if addedPayment.RecordCode != "02" {
		t.Errorf("Payment RecordCode = %q, want %q", addedPayment.RecordCode, "02")
	}
	if addedPayment.PaymentID != "PAY001" {
		t.Errorf("Payment PaymentID = %q, want %q", addedPayment.PaymentID, "PAY001")
	}

	// Test adding ACH payment to Check schedule (should error)
	fb.StartCheckSchedule("000000000002", "Refund", "87654321", "stub")
	fb.AddACHPayment(payment)

	if len(fb.errors) == 0 {
		t.Error("Expected error when adding ACH payment to Check schedule")
	}
}

// TestFileBuilder_AddCheckPayment tests adding Check payments
func TestFileBuilder_AddCheckPayment(t *testing.T) {
	fb := NewFileBuilder()

	// Test adding payment without schedule (should error)
	payment := &CheckPayment{
		AgencyAccountIdentifier: "CHK123",
		Amount:                  25000,
		PayeeName:               "Jane Smith",
		PaymentID:               "CHK001",
	}

	fb.AddCheckPayment(payment)
	if len(fb.errors) == 0 {
		t.Error("Expected error when adding payment without schedule")
	}

	// Test adding payment to Check schedule
	fb = NewFileBuilder()
	fb.StartCheckSchedule("000000000002", "Refund", "87654321", "stub")
	fb.AddCheckPayment(payment)

	if len(fb.errors) != 0 {
		t.Errorf("Unexpected errors: %v", fb.errors)
	}

	checkSchedule := fb.currentSchedule.(*CheckSchedule)
	if len(checkSchedule.Payments) != 1 {
		t.Errorf("Expected 1 payment, got %d", len(checkSchedule.Payments))
	}

	addedPayment := checkSchedule.Payments[0].(*CheckPayment)
	if addedPayment.RecordCode != "12" {
		t.Errorf("Payment RecordCode = %q, want %q", addedPayment.RecordCode, "12")
	}
	if addedPayment.PaymentID != "CHK001" {
		t.Errorf("Payment PaymentID = %q, want %q", addedPayment.PaymentID, "CHK001")
	}

	// Test adding Check payment to ACH schedule (should error)
	fb.StartACHSchedule("000000000001", "Vendor", "12345678", "CCD")
	fb.AddCheckPayment(payment)

	if len(fb.errors) == 0 {
		t.Error("Expected error when adding Check payment to ACH schedule")
	}
}

// TestFileBuilder_Build tests building complete files
func TestFileBuilder_Build(t *testing.T) {
	// Test build with errors
	fb := NewFileBuilder()
	fb.AddACHPayment(&ACHPayment{}) // This will add an error

	file, err := fb.Build()
	if err == nil {
		t.Error("Expected error when building with errors")
	}
	if file != nil {
		t.Error("Expected nil file when build fails")
	}

	// Test successful build
	fb = NewFileBuilder()
	fb.WithHeader("TEST SYSTEM", "502", false)
	fb.StartACHSchedule("000000000001", "Vendor", "12345678", "CCD")
	fb.AddACHPayment(&ACHPayment{
		AgencyAccountIdentifier: "TEST123",
		Amount:                  15000,
		PayeeName:               "John Doe",
		PaymentID:               "PAY001",
	})

	file, err = fb.Build()
	if err != nil {
		t.Fatalf("Unexpected error building file: %v", err)
	}
	if file == nil {
		t.Fatal("Expected non-nil file")
	}

	// Verify file structure
	if file.Header == nil {
		t.Error("File header is nil")
	}
	if len(file.Schedules) != 1 {
		t.Errorf("Expected 1 schedule, got %d", len(file.Schedules))
	}
	if file.Trailer == nil {
		t.Error("File trailer is nil")
	}

	// Verify schedule trailer was calculated
	achSchedule := file.Schedules[0].(*ACHSchedule)
	if achSchedule.Trailer == nil {
		t.Error("Schedule trailer is nil")
	}
	if achSchedule.Trailer.ScheduleCount != 1 {
		t.Errorf("Schedule count = %d, want 1", achSchedule.Trailer.ScheduleCount)
	}
	if achSchedule.Trailer.ScheduleAmount != 15000 {
		t.Errorf("Schedule amount = %d, want 15000", achSchedule.Trailer.ScheduleAmount)
	}

	// Verify file trailer was calculated
	if file.Trailer.TotalCountPayments != 1 {
		t.Errorf("Total payments = %d, want 1", file.Trailer.TotalCountPayments)
	}
	if file.Trailer.TotalAmountPayments != 15000 {
		t.Errorf("Total amount = %d, want 15000", file.Trailer.TotalAmountPayments)
	}
}

// TestFileBuilder_MultipleSchedules tests building files with multiple schedules
func TestFileBuilder_MultipleSchedules(t *testing.T) {
	fb := NewFileBuilder()
	fb.WithHeader("TEST SYSTEM", "502", false)

	// Add ACH schedule
	fb.StartACHSchedule("000000000001", "Vendor", "12345678", "CCD")
	fb.AddACHPayment(&ACHPayment{
		AgencyAccountIdentifier: "TEST123",
		Amount:                  15000,
		PayeeName:               "John Doe",
		PaymentID:               "PAY001",
	})

	// Add Check schedule
	fb.StartCheckSchedule("000000000002", "Refund", "87654321", "stub")
	fb.AddCheckPayment(&CheckPayment{
		AgencyAccountIdentifier: "CHK123",
		Amount:                  25000,
		PayeeName:               "Jane Smith",
		PaymentID:               "CHK001",
	})

	file, err := fb.Build()
	if err != nil {
		t.Fatalf("Unexpected error building file: %v", err)
	}

	if len(file.Schedules) != 2 {
		t.Errorf("Expected 2 schedules, got %d", len(file.Schedules))
	}

	// Verify total amounts
	if file.Trailer.TotalCountPayments != 2 {
		t.Errorf("Total payments = %d, want 2", file.Trailer.TotalCountPayments)
	}
	if file.Trailer.TotalAmountPayments != 40000 {
		t.Errorf("Total amount = %d, want 40000", file.Trailer.TotalAmountPayments)
	}
}

// TestAgencyReconcilementParser_ParseIRSReconcilement tests IRS reconcilement parsing
func TestAgencyReconcilementParser_ParseIRSReconcilement(t *testing.T) {
	arp := &AgencyReconcilementParser{}

	// Test with invalid length
	result := arp.ParseIRSReconcilement("short", "BONDS")
	if len(result) != 0 {
		t.Error("Expected empty result for invalid length")
	}

	// Test BONDS format (100 character string)
	bondsRecon := "2023" + "12" + "01" + "02" + "03" + "A" + "TEST" + "123" + "X" + "Y" + "Z" + strings.Repeat("A", 33) + "B" + strings.Repeat("C", 33) + strings.Repeat("0", 10)

	result = arp.ParseIRSReconcilement(bondsRecon, "BONDS")

	expectedFields := []string{
		"TaxPeriodYear", "TaxPeriodMonth", "MFTCode", "ServiceCenterCode",
		"DistrictOfficeCode", "FileTINCode", "NameControl", "PlanReportNumber",
		"SplitRefundCode", "InjuredSpouseCode", "DebtBypassIndicator",
		"BondName1", "BondREGCode", "BondName2",
	}

	for _, field := range expectedFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Expected field %q not found in BONDS result", field)
		}
	}

	// Test standard format
	standardRecon := "2023" + "12" + "01" + "02" + "03" + "A" + "TEST" + "123" + "X" + "Y" + "Z" + strings.Repeat("1", 14) + strings.Repeat("2", 10) + strings.Repeat("0", 53)

	result = arp.ParseIRSReconcilement(standardRecon, "STANDARD")

	standardFields := []string{
		"TaxPeriodYear", "TaxPeriodMonth", "MFTCode", "ServiceCenterCode",
		"DistrictOfficeCode", "FileTINCode", "NameControl", "PlanReportNumber",
		"SplitRefundCode", "InjuredSpouseCode", "DebtBypassIndicator",
		"DocumentLocatorNumber", "CheckDetailEnclosureCode",
	}

	for _, field := range standardFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Expected field %q not found in standard result", field)
		}
	}

	if result["TaxPeriodYear"] != "2023" {
		t.Errorf("TaxPeriodYear = %q, want %q", result["TaxPeriodYear"], "2023")
	}
	if result["TaxPeriodMonth"] != "12" {
		t.Errorf("TaxPeriodMonth = %q, want %q", result["TaxPeriodMonth"], "12")
	}
}

// TestAgencyReconcilementParser_ParseVAReconcilement tests VA reconcilement parsing
func TestAgencyReconcilementParser_ParseVAReconcilement(t *testing.T) {
	arp := &AgencyReconcilementParser{}

	// Test with invalid length
	result := arp.ParseVAReconcilement("short", true)
	if len(result) != 0 {
		t.Error("Expected empty result for invalid length")
	}

	// Test ACH format (100 character string)
	achRecon := "12" + "34" + "A" + "B" + "56" + "78" + strings.Repeat("9", 12) + strings.Repeat("0", 78)

	result = arp.ParseVAReconcilement(achRecon, true)

	achFields := []string{
		"StationCode", "FinCode", "AppropCode", "AddressSeqCode",
		"PolicyPreCode", "PolicyNum", "PayPeriodInfo",
	}

	for _, field := range achFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Expected field %q not found in ACH result", field)
		}
	}

	if result["StationCode"] != "12" {
		t.Errorf("StationCode = %q, want %q", result["StationCode"], "12")
	}
	if result["FinCode"] != "34" {
		t.Errorf("FinCode = %q, want %q", result["FinCode"], "34")
	}

	// Test Check format
	checkRecon := "12" + "34" + "C" + "A" + "B" + "56" + "78" + strings.Repeat("9", 12) + "NAM" + strings.Repeat("0", 74)

	result = arp.ParseVAReconcilement(checkRecon, false)

	checkFields := []string{
		"StationCode", "FinCode", "CourtesyCode", "AppropCode",
		"AddressSeqCode", "PolicyPreCode", "PolicyNum", "PayPeriodInfo", "NameCode",
	}

	for _, field := range checkFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Expected field %q not found in Check result", field)
		}
	}

	if result["CourtesyCode"] != "C" {
		t.Errorf("CourtesyCode = %q, want %q", result["CourtesyCode"], "C")
	}
	if result["NameCode"] != "NAM" {
		t.Errorf("NameCode = %q, want %q", result["NameCode"], "NAM")
	}
}

// TestAgencyReconcilementParser_ParseSSAReconcilement tests SSA reconcilement parsing
func TestAgencyReconcilementParser_ParseSSAReconcilement(t *testing.T) {
	arp := &AgencyReconcilementParser{}

	// Test with invalid length
	result := arp.ParseSSAReconcilement("short", "SSA")
	if len(result) != 0 {
		t.Error("Expected empty result for invalid length")
	}

	// Test SSA format with TIN Indicator Offset (100 character string)
	ssaRecon := "A" + "BC" + "D" + strings.Repeat("0", 96)

	result = arp.ParseSSAReconcilement(ssaRecon, "SSA")

	expectedFields := []string{
		"ProgramServiceCenterCode", "PaymentIDCode", "TINIndicatorOffset",
	}

	for _, field := range expectedFields {
		if _, exists := result[field]; !exists {
			t.Errorf("Expected field %q not found in SSA result", field)
		}
	}

	if result["ProgramServiceCenterCode"] != "A" {
		t.Errorf("ProgramServiceCenterCode = %q, want %q", result["ProgramServiceCenterCode"], "A")
	}
	if result["PaymentIDCode"] != "BC" {
		t.Errorf("PaymentIDCode = %q, want %q", result["PaymentIDCode"], "BC")
	}
	if result["TINIndicatorOffset"] != "D" {
		t.Errorf("TINIndicatorOffset = %q, want %q", result["TINIndicatorOffset"], "D")
	}

	// Test SSA-A format without TIN Indicator Offset
	result = arp.ParseSSAReconcilement(ssaRecon, "SSA-A")

	if _, exists := result["TINIndicatorOffset"]; exists {
		t.Error("TINIndicatorOffset should not exist for SSA-A format")
	}

	if result["ProgramServiceCenterCode"] != "A" {
		t.Errorf("ProgramServiceCenterCode = %q, want %q", result["ProgramServiceCenterCode"], "A")
	}
	if result["PaymentIDCode"] != "BC" {
		t.Errorf("PaymentIDCode = %q, want %q", result["PaymentIDCode"], "BC")
	}
}

// BenchmarkFieldPadding_PadLeft benchmarks left padding performance
func BenchmarkFieldPadding_PadLeft(b *testing.B) {
	fp := &FieldPadding{}
	for i := 0; i < b.N; i++ {
		fp.PadLeft("test", 10, '0')
	}
}

// BenchmarkFieldPadding_PadNumeric benchmarks numeric padding performance
func BenchmarkFieldPadding_PadNumeric(b *testing.B) {
	fp := &FieldPadding{}
	for i := 0; i < b.N; i++ {
		fp.PadNumeric("ABC123DEF456", 10)
	}
}

// BenchmarkFormatUtils_ParseAmount benchmarks amount parsing performance
func BenchmarkFormatUtils_ParseAmount(b *testing.B) {
	fu := &FormatUtils{}
	for i := 0; i < b.N; i++ {
		fu.ParseAmount("$1,234.56")
	}
}

// BenchmarkFormatUtils_CleanAddress benchmarks address cleaning performance
func BenchmarkFormatUtils_CleanAddress(b *testing.B) {
	fu := &FormatUtils{}
	address := `123 "Main" Street & <Company> Name`
	for i := 0; i < b.N; i++ {
		fu.CleanAddress(address)
	}
}
