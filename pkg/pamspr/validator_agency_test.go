package pamspr

import (
	"strings"
	"testing"
)

// TestValidateVAPayment tests Veterans Affairs payment validation
func TestValidateVAPayment(t *testing.T) {
	validator := NewValidator()

	// Test valid VA ACH payment
	validVARecon := "01" + // StationCode (2 chars)
		"02" + // FinCode (2 chars)
		"A" + // AppropCode (1 char)
		"B" + // AddressSeqCode (1 char)
		"CD" + // PolicyPreCode (2 chars)
		"EF" + // PolicyNum (2 chars)
		"123456789012" + // PayPeriodInfo (12 chars)
		strings.Repeat(" ", 78) // Pad to 100 chars

	achPayment := &ACHPayment{
		Reconcilement: validVARecon,
	}
	err := validator.validateVAPayment(achPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid VA ACH payment: %v", err)
	}

	// Test valid VA Check payment
	validVACheckRecon := "01" + // StationCode (2 chars)
		"02" + // FinCode (2 chars)
		"C" + // CourtesyCode (1 char)
		"A" + // AppropCode (1 char)
		"B" + // AddressSeqCode (1 char)
		"CD" + // PolicyPreCode (2 chars)
		"EF" + // PolicyNum (2 chars)
		"123456789012" + // PayPeriodInfo (12 chars)
		"GHI" + // NameCode (3 chars)
		strings.Repeat(" ", 74) // Pad to 100 chars (2+2+1+1+1+2+2+12+3 = 26, need 74 more)

	checkPayment := &CheckPayment{
		Reconcilement: validVACheckRecon,
	}
	err = validator.validateVAPayment(checkPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid VA check payment: %v", err)
	}

	// Test invalid reconcilement length
	shortRecon := &ACHPayment{
		Reconcilement: "TOO SHORT",
	}
	err = validator.validateVAPayment(shortRecon)
	if err == nil {
		t.Error("Expected error for short VA reconcilement")
	}
	if !strings.Contains(err.Error(), "100 characters") {
		t.Errorf("Expected length error, got: %v", err)
	}

	// Test missing station code
	noStationRecon := strings.Repeat(" ", 2) + // Empty StationCode
		"02" + // FinCode
		strings.Repeat(" ", 96) // Pad to 100 chars
	noStationPayment := &ACHPayment{
		Reconcilement: noStationRecon,
	}
	err = validator.validateVAPayment(noStationPayment)
	if err == nil {
		t.Error("Expected error for missing VA station code")
	}
	if !strings.Contains(err.Error(), "Reconcilement.StationCode is required") {
		t.Errorf("Expected station code error, got: %v", err)
	}

	// Test missing FIN code
	noFinRecon := "01" + // StationCode
		strings.Repeat(" ", 2) + // Empty FinCode
		strings.Repeat(" ", 96) // Pad to 100 chars
	noFinPayment := &ACHPayment{
		Reconcilement: noFinRecon,
	}
	err = validator.validateVAPayment(noFinPayment)
	if err == nil {
		t.Error("Expected error for missing VA FIN code")
	}
	if !strings.Contains(err.Error(), "Reconcilement.FinCode is required") {
		t.Errorf("Expected FIN code error, got: %v", err)
	}

	// Test oversized FIN code (the parser extracts fixed positions, so we need to test the validation logic)
	oversizedFinRecon := "01" + // StationCode
		"123" + // FinCode - parser will extract "12" but we want to test validation
		strings.Repeat(" ", 95) // Pad to 100 chars
	oversizedFinPayment := &ACHPayment{
		Reconcilement: oversizedFinRecon,
	}
	err = validator.validateVAPayment(oversizedFinPayment)
	// This should pass because parser extracts first 2 chars only
	if err != nil {
		t.Errorf("Unexpected error for valid extracted FIN code: %v", err)
	}
}

// TestValidateSSAPayment tests Social Security Administration payment validation
func TestValidateSSAPayment(t *testing.T) {
	validator := NewValidator()

	// Test valid SSA payment
	validSSARecon := "1" + // ProgramServiceCenterCode (1 char)
		"AB" + // PaymentIDCode (2 chars)
		"C" + // TINIndicatorOffset (1 char)
		strings.Repeat(" ", 96) // Pad to 100 chars

	ssaPayment := &ACHPayment{
		Reconcilement: validSSARecon,
	}
	err := validator.validateSSAPayment(ssaPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid SSA payment: %v", err)
	}

	// Test SSA-A variant (no TIN Indicator Offset)
	validator.CustomAgencyRuleID = "SSA-A"
	validSSAARecon := "2" + // ProgramServiceCenterCode (1 char)
		"CD" + // PaymentIDCode (2 chars)
		strings.Repeat(" ", 97) // Pad to 100 chars (no TIN offset)

	ssaAPayment := &ACHPayment{
		Reconcilement: validSSAARecon,
	}
	err = validator.validateSSAPayment(ssaAPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid SSA-A payment: %v", err)
	}
	validator.CustomAgencyRuleID = "" // Reset

	// Test invalid reconcilement length
	shortSSARecon := &ACHPayment{
		Reconcilement: "TOO SHORT",
	}
	err = validator.validateSSAPayment(shortSSARecon)
	if err == nil {
		t.Error("Expected error for short SSA reconcilement")
	}
	if !strings.Contains(err.Error(), "100 characters") {
		t.Errorf("Expected length error, got: %v", err)
	}

	// Test missing Program Service Center Code
	noPSCRecon := strings.Repeat(" ", 1) + // Empty PSC Code
		"AB" + // PaymentIDCode
		"C" + // TINIndicatorOffset
		strings.Repeat(" ", 96) // Pad to 100 chars
	noPSCPayment := &ACHPayment{
		Reconcilement: noPSCRecon,
	}
	err = validator.validateSSAPayment(noPSCPayment)
	if err == nil {
		t.Error("Expected error for missing SSA program service center code")
	}
	if !strings.Contains(err.Error(), "Reconcilement.ProgramServiceCenterCode is required") {
		t.Errorf("Expected PSC error, got: %v", err)
	}

	// Test missing Payment ID Code
	noPaymentIDRecon := "1" + // PSC Code
		strings.Repeat(" ", 2) + // Empty PaymentIDCode
		"C" + // TINIndicatorOffset
		strings.Repeat(" ", 96) // Pad to 100 chars
	noPaymentIDPayment := &ACHPayment{
		Reconcilement: noPaymentIDRecon,
	}
	err = validator.validateSSAPayment(noPaymentIDPayment)
	if err == nil {
		t.Error("Expected error for missing SSA payment ID code")
	}
	if !strings.Contains(err.Error(), "Reconcilement.PaymentIDCode is required") {
		t.Errorf("Expected payment ID error, got: %v", err)
	}

	// Test oversized PSC Code (parser extracts position 0:1, so "12" becomes "1")
	oversizedPSCRecon := "12" + // PSC Code - parser will extract "1"
		"AB" + // PaymentIDCode
		"C" + // TINIndicatorOffset
		strings.Repeat(" ", 95) // Pad to 100 chars
	oversizedPSCPayment := &ACHPayment{
		Reconcilement: oversizedPSCRecon,
	}
	err = validator.validateSSAPayment(oversizedPSCPayment)
	// This should pass because parser extracts first 1 char only
	if err != nil {
		t.Errorf("Unexpected error for valid extracted PSC code: %v", err)
	}

	// Test oversized Payment ID Code (parser extracts position 1:3, so "ABC" becomes "AB")
	oversizedPaymentIDRecon := "1" + // PSC Code
		"ABC" + // PaymentIDCode - parser will extract "AB"
		"D" + // TINIndicatorOffset
		strings.Repeat(" ", 95) // Pad to 100 chars
	oversizedPaymentIDPayment := &ACHPayment{
		Reconcilement: oversizedPaymentIDRecon,
	}
	err = validator.validateSSAPayment(oversizedPaymentIDPayment)
	// This should pass because parser extracts first 2 chars only
	if err != nil {
		t.Errorf("Unexpected error for valid extracted payment ID: %v", err)
	}

	// Test oversized TIN Indicator Offset (parser extracts position 3:4, so "CD" becomes "C")
	oversizedTINRecon := "1" + // PSC Code
		"AB" + // PaymentIDCode
		"CD" + // TINIndicatorOffset - parser will extract "C"
		strings.Repeat(" ", 95) // Pad to 100 chars
	oversizedTINPayment := &ACHPayment{
		Reconcilement: oversizedTINRecon,
	}
	err = validator.validateSSAPayment(oversizedTINPayment)
	// This should pass because parser extracts first 1 char only
	if err != nil {
		t.Errorf("Unexpected error for valid extracted TIN offset: %v", err)
	}
}

// TestValidateRRBPayment tests Railroad Retirement Board payment validation
func TestValidateRRBPayment(t *testing.T) {
	validator := NewValidator()

	// Test valid RRB payment with all required fields
	validRRBRecon := "AB" + // BeneficiarySymbol (2 chars)
		"C" + // PrefixCode (1 char)
		"D" + // PayeeCode (1 char)
		"E" + // ObjectCode (1 char)
		strings.Repeat(" ", 95) // Filler (95 chars)

	achPayment := &ACHPayment{
		Reconcilement: validRRBRecon,
	}
	err := validator.validateRRBPayment(achPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid RRB payment: %v", err)
	}

	// Test valid RRB check payment
	checkPayment := &CheckPayment{
		Reconcilement: validRRBRecon,
	}
	err = validator.validateRRBPayment(checkPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid RRB check payment: %v", err)
	}

	// Test invalid reconcilement length
	shortRecon := "ABCDE" // Only 5 characters
	shortPayment := &ACHPayment{
		Reconcilement: shortRecon,
	}
	err = validator.validateRRBPayment(shortPayment)
	if err == nil {
		t.Error("Expected error for short reconcilement field")
	}
	if !strings.Contains(err.Error(), "must be exactly 100 characters") {
		t.Errorf("Expected length error, got: %v", err)
	}

	// Test missing Beneficiary Symbol (empty first 2 positions)
	missingBeneficiaryRecon := "  " + // Empty BeneficiarySymbol (2 chars)
		"C" + // PrefixCode (1 char)
		"D" + // PayeeCode (1 char)
		"E" + // ObjectCode (1 char)
		strings.Repeat(" ", 95) // Filler (95 chars)
	missingBeneficiaryPayment := &ACHPayment{
		Reconcilement: missingBeneficiaryRecon,
	}
	err = validator.validateRRBPayment(missingBeneficiaryPayment)
	if err == nil {
		t.Error("Expected error for missing Beneficiary Symbol")
	}
	if !strings.Contains(err.Error(), "BeneficiarySymbol must be exactly 2") {
		t.Errorf("Expected BeneficiarySymbol error, got: %v", err)
	}

	// Test missing Prefix Code
	missingPrefixRecon := "AB" + // BeneficiarySymbol (2 chars)
		" " + // Empty PrefixCode (1 char)
		"D" + // PayeeCode (1 char)
		"E" + // ObjectCode (1 char)
		strings.Repeat(" ", 95) // Filler (95 chars)
	missingPrefixPayment := &ACHPayment{
		Reconcilement: missingPrefixRecon,
	}
	err = validator.validateRRBPayment(missingPrefixPayment)
	if err == nil {
		t.Error("Expected error for missing Prefix Code")
	}
	if !strings.Contains(err.Error(), "Prefix Code must be exactly 1") {
		t.Errorf("Expected Prefix Code error, got: %v", err)
	}

	// Test missing Payee Code
	missingPayeeRecon := "AB" + // BeneficiarySymbol (2 chars)
		"C" + // PrefixCode (1 char)
		" " + // Empty PayeeCode (1 char)
		"E" + // ObjectCode (1 char)
		strings.Repeat(" ", 95) // Filler (95 chars)
	missingPayeePayment := &ACHPayment{
		Reconcilement: missingPayeeRecon,
	}
	err = validator.validateRRBPayment(missingPayeePayment)
	if err == nil {
		t.Error("Expected error for missing Payee Code")
	}
	if !strings.Contains(err.Error(), "Payee Code must be exactly 1") {
		t.Errorf("Expected Payee Code error, got: %v", err)
	}

	// Test missing Object Code
	missingObjectRecon := "AB" + // BeneficiarySymbol (2 chars)
		"C" + // PrefixCode (1 char)
		"D" + // PayeeCode (1 char)
		" " + // Empty ObjectCode (1 char)
		strings.Repeat(" ", 95) // Filler (95 chars)
	missingObjectPayment := &ACHPayment{
		Reconcilement: missingObjectRecon,
	}
	err = validator.validateRRBPayment(missingObjectPayment)
	if err == nil {
		t.Error("Expected error for missing Object Code")
	}
	if !strings.Contains(err.Error(), "Object Code must be exactly 1") {
		t.Errorf("Expected Object Code error, got: %v", err)
	}
}

// TestValidateCCCPayment tests Commodity Credit Corporation payment validation
func TestValidateCCCPayment(t *testing.T) {
	validator := NewValidator()

	// Test valid CCC payment with TOP Payment Agency ID and Site ID
	validCCCRecon := "AB" + // TOPPaymentAgencyID (2 chars)
		"CD" + // TOPAgencySiteID (2 chars)
		strings.Repeat(" ", 96) // Filler (96 chars)

	achPayment := &ACHPayment{
		Reconcilement: validCCCRecon,
	}
	err := validator.validateCCCPayment(achPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid CCC payment: %v", err)
	}

	// Test valid CCC check payment
	checkPayment := &CheckPayment{
		Reconcilement: validCCCRecon,
	}
	err = validator.validateCCCPayment(checkPayment)
	if err != nil {
		t.Errorf("Unexpected error for valid CCC check payment: %v", err)
	}

	// Test CCC payment with empty TOP fields (should be valid)
	emptyCCCRecon := "  " + // Empty TOPPaymentAgencyID (2 chars)
		"  " + // Empty TOPAgencySiteID (2 chars)
		strings.Repeat(" ", 96) // Filler (96 chars)
	emptyPayment := &ACHPayment{
		Reconcilement: emptyCCCRecon,
	}
	err = validator.validateCCCPayment(emptyPayment)
	if err != nil {
		t.Errorf("Unexpected error for CCC payment with empty TOP fields: %v", err)
	}

	// Test invalid reconcilement length
	shortRecon := "ABCD" // Only 4 characters
	shortPayment := &ACHPayment{
		Reconcilement: shortRecon,
	}
	err = validator.validateCCCPayment(shortPayment)
	if err == nil {
		t.Error("Expected error for short reconcilement field")
	}
	if !strings.Contains(err.Error(), "must be exactly 100 characters") {
		t.Errorf("Expected length error, got: %v", err)
	}

	// Test invalid TOP Payment Agency ID with numeric characters
	invalidAgencyRecon := "A1" + // Invalid TOPPaymentAgencyID (contains numeric)
		"CD" + // TOPAgencySiteID (2 chars)
		strings.Repeat(" ", 96) // Filler (96 chars)
	invalidAgencyPayment := &ACHPayment{
		Reconcilement: invalidAgencyRecon,
	}
	err = validator.validateCCCPayment(invalidAgencyPayment)
	if err == nil {
		t.Error("Expected error for invalid TOP Payment Agency ID")
	}
	if !strings.Contains(err.Error(), "must contain only alphabetic characters") {
		t.Errorf("Expected alphabetic character error, got: %v", err)
	}

	// Test invalid TOP Agency Site ID with numeric characters
	invalidSiteRecon := "AB" + // TOPPaymentAgencyID (2 chars)
		"C2" + // Invalid TOPAgencySiteID (contains numeric)
		strings.Repeat(" ", 96) // Filler (96 chars)
	invalidSitePayment := &ACHPayment{
		Reconcilement: invalidSiteRecon,
	}
	err = validator.validateCCCPayment(invalidSitePayment)
	if err == nil {
		t.Error("Expected error for invalid TOP Agency Site ID")
	}
	if !strings.Contains(err.Error(), "must contain only alphabetic characters") {
		t.Errorf("Expected alphabetic character error, got: %v", err)
	}

	// Test single character TOP Payment Agency ID (should fail)
	singleCharAgencyRecon := "A" + // Single char TOPPaymentAgencyID (invalid length)
		" " + // Pad to make position 2
		"CD" + // TOPAgencySiteID (2 chars)
		strings.Repeat(" ", 96) // Filler (96 chars)
	singleCharPayment := &ACHPayment{
		Reconcilement: singleCharAgencyRecon,
	}
	err = validator.validateCCCPayment(singleCharPayment)
	if err == nil {
		t.Error("Expected error for single character TOP Payment Agency ID")
	}
	if !strings.Contains(err.Error(), "must be exactly 2 alphabetic characters") {
		t.Errorf("Expected length error for TOP Payment Agency ID, got: %v", err)
	}
}

// TestParseRRBReconcilement tests RRB reconcilement parsing
func TestParseRRBReconcilement(t *testing.T) {
	parser := &AgencyReconcilementParser{}

	// Test valid RRB reconcilement
	validRecon := "AB" + // BeneficiarySymbol
		"C" + // PrefixCode
		"D" + // PayeeCode
		"E" + // ObjectCode
		strings.Repeat(" ", 95) // Filler

	fields := parser.ParseRRBReconcilement(validRecon)

	expectedFields := map[string]string{
		"BeneficiarySymbol": "AB",
		"PrefixCode":        "C",
		"PayeeCode":         "D",
		"ObjectCode":        "E",
	}

	for key, expected := range expectedFields {
		if actual, exists := fields[key]; !exists {
			t.Errorf("Missing field %s", key)
		} else if actual != expected {
			t.Errorf("Field %s: expected %s, got %s", key, expected, actual)
		}
	}

	// Test invalid length reconcilement
	shortRecon := "ABCDE"
	shortFields := parser.ParseRRBReconcilement(shortRecon)
	if len(shortFields) != 0 {
		t.Errorf("Expected empty result for invalid length, got: %v", shortFields)
	}
}

// TestParseCCCReconcilement tests CCC reconcilement parsing
func TestParseCCCReconcilement(t *testing.T) {
	parser := &AgencyReconcilementParser{}

	// Test valid CCC reconcilement
	validRecon := "AB" + // TOPPaymentAgencyID
		"CD" + // TOPAgencySiteID
		strings.Repeat(" ", 96) // Filler

	fields := parser.ParseCCCReconcilement(validRecon)

	expectedFields := map[string]string{
		"TOPPaymentAgencyID": "AB",
		"TOPAgencySiteID":    "CD",
	}

	for key, expected := range expectedFields {
		if actual, exists := fields[key]; !exists {
			t.Errorf("Missing field %s", key)
		} else if actual != expected {
			t.Errorf("Field %s: expected %s, got %s", key, expected, actual)
		}
	}

	// Test empty TOP fields (should parse but be empty)
	emptyRecon := "  " + // Empty TOPPaymentAgencyID
		"  " + // Empty TOPAgencySiteID
		strings.Repeat(" ", 96) // Filler

	emptyFields := parser.ParseCCCReconcilement(emptyRecon)
	if emptyFields["TOPPaymentAgencyID"] != "" {
		t.Errorf("Expected empty TOP Payment Agency ID, got: '%s'", emptyFields["TOPPaymentAgencyID"])
	}
	if emptyFields["TOPAgencySiteID"] != "" {
		t.Errorf("Expected empty TOP Agency Site ID, got: '%s'", emptyFields["TOPAgencySiteID"])
	}

	// Test invalid length reconcilement
	shortRecon := "ABCD"
	shortFields := parser.ParseCCCReconcilement(shortRecon)
	if len(shortFields) != 0 {
		t.Errorf("Expected empty result for invalid length, got: %v", shortFields)
	}
}

// TestValidateAgencySpecificIntegration tests the main ValidateAgencySpecific function
func TestValidateAgencySpecificIntegration(t *testing.T) {
	validator := NewValidator()

	// Test that the new implementations are called correctly
	validVARecon := "01" + // StationCode
		"02" + // FinCode
		"A" + // AppropCode
		"B" + // AddressSeqCode
		"CD" + // PolicyPreCode
		"EF" + // PolicyNum
		"123456789012" + // PayPeriodInfo
		strings.Repeat(" ", 78) // Pad to 100 chars

	vaPayment := &ACHPayment{
		Reconcilement: validVARecon,
	}

	// Test VA validation through public interface
	err := validator.ValidateAgencySpecific(vaPayment, "VA")
	if err != nil {
		t.Errorf("Unexpected error for valid VA payment through public interface: %v", err)
	}

	// Test RRB validation through public interface
	validRRBRecon := "AB" + // BeneficiarySymbol
		"C" + // PrefixCode
		"D" + // PayeeCode
		"E" + // ObjectCode
		strings.Repeat(" ", 95) // Filler

	rrBPayment := &ACHPayment{
		Reconcilement: validRRBRecon,
	}

	err = validator.ValidateAgencySpecific(rrBPayment, "RRB")
	if err != nil {
		t.Errorf("Unexpected error for valid RRB payment through public interface: %v", err)
	}

	// Test CCC validation through public interface
	validCCCRecon := "AB" + // TOPPaymentAgencyID
		"CD" + // TOPAgencySiteID
		strings.Repeat(" ", 96) // Filler

	cccPayment := &ACHPayment{
		Reconcilement: validCCCRecon,
	}

	err = validator.ValidateAgencySpecific(cccPayment, "CCC")
	if err != nil {
		t.Errorf("Unexpected error for valid CCC payment through public interface: %v", err)
	}

	// Test SSA validation through public interface
	validSSARecon := "1" + // PSC Code
		"AB" + // PaymentIDCode
		"C" + // TINIndicatorOffset
		strings.Repeat(" ", 96) // Pad to 100 chars

	ssaPayment := &ACHPayment{
		Reconcilement: validSSARecon,
	}

	err = validator.ValidateAgencySpecific(ssaPayment, "SSA")
	if err != nil {
		t.Errorf("Unexpected error for valid SSA payment through public interface: %v", err)
	}

	// Test invalid VA payment through public interface
	invalidVAPayment := &ACHPayment{
		Reconcilement: "TOO SHORT",
	}
	err = validator.ValidateAgencySpecific(invalidVAPayment, "VA")
	if err == nil {
		t.Error("Expected error for invalid VA payment through public interface")
	}
}

// TODO: Add tests for RRB validation when implemented
// TODO: Add tests for CCC validation when implemented
// TODO: Add tests for enhanced IRS validation when implemented
// TODO: Add integration tests with synthetic agency test files
// TODO: Add tests for cross-field validation rules when business requirements available
// TODO: Add tests for payment type restrictions when agency specifications available
// TODO: Add tests for amount limits when agency requirements available
