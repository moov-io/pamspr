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
	if !strings.Contains(err.Error(), "station code is required") {
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
	if !strings.Contains(err.Error(), "FIN code is required") {
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
	if !strings.Contains(err.Error(), "program service center code is required") {
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
	if !strings.Contains(err.Error(), "payment ID code is required") {
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
