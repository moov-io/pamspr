package pamspr

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ValidationError represents a validation error with context
type ValidationError struct {
	Field   string
	Value   string
	Rule    string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: field=%s, value=%s, rule=%s, message=%s",
		e.Field, e.Value, e.Rule, e.Message)
}

// NewValidationError creates a ValidationError with consistent formatting
func NewValidationError(field, value, rule, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   value,
		Rule:    rule,
		Message: message,
	}
}

// NewFieldRequiredError creates a ValidationError for missing required fields
func NewFieldRequiredError(field string) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   "",
		Rule:    "required",
		Message: fmt.Sprintf("%s is required", field),
	}
}

// NewFieldLengthError creates a ValidationError for incorrect field lengths
func NewFieldLengthError(field, value string, expected, actual int) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   value,
		Rule:    "length",
		Message: fmt.Sprintf("%s must be exactly %d characters, got %d", field, expected, actual),
	}
}

// NewFieldFormatError creates a ValidationError for invalid field formats
func NewFieldFormatError(field, value, expectedFormat string) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   value,
		Rule:    "format",
		Message: fmt.Sprintf("%s must %s", field, expectedFormat),
	}
}

// WrapValidationError wraps an existing error as a ValidationError with additional context
func WrapValidationError(field, value, rule string, err error) ValidationError {
	return ValidationError{
		Field:   field,
		Value:   value,
		Rule:    rule,
		Message: err.Error(),
	}
}

// Validator provides validation methods for PAM SPR records
type Validator struct {
	// Configuration
	AllowedPaymentTypes map[string]bool
	ValidALCs           map[string]bool
	CustomAgencyRuleID  string // Agency-specific rule ID (e.g., "SSA-A", "SSA-Daily")
}

// NewValidator creates a new validator with default configuration
func NewValidator() *Validator {
	return &Validator{
		AllowedPaymentTypes: map[string]bool{
			"Allotment":       true,
			"Annuity":         true,
			"ChildSupport":    true,
			"Daily Benefit":   true,
			"Education":       true,
			"Fee":             true,
			"Insurance":       true,
			"Miscellaneous":   true,
			"Monthly Benefit": true,
			"Refund":          true,
			"Salary":          true,
			"Thrift":          true,
			"Travel":          true,
			"Vendor":          true,
		},
		ValidALCs: make(map[string]bool), // Populated from agency configuration
	}
}

// File Structure Validations
// ValidateFileStructure is now in validator_structure.go

// Field-level validations
func (v *Validator) ValidateFileHeader(header *FileHeader) error {
	// Record code validation
	if header.RecordCode != "H " {
		return ValidationError{
			Field:   "RecordCode",
			Value:   header.RecordCode,
			Rule:    "exact_match",
			Message: "must be 'H '",
		}
	}

	// Version validation
	if header.StandardPaymentVersion != CurrentSPRVersion {
		return ValidationError{
			Field:   "StandardPaymentVersion",
			Value:   header.StandardPaymentVersion,
			Rule:    "version",
			Message: fmt.Sprintf("current published version must be '%s'", CurrentSPRVersion),
		}
	}

	// Same Day ACH flag validation
	if header.IsRequestedForSameDayACH != SDAFlagDisabled && header.IsRequestedForSameDayACH != SDAFlagEnabled && header.IsRequestedForSameDayACH != "" {
		return ValidationError{
			Field:   "IsRequestedForSameDayACH",
			Value:   header.IsRequestedForSameDayACH,
			Rule:    "valid_values",
			Message: fmt.Sprintf("must be '%s', '%s', or blank", SDAFlagDisabled, SDAFlagEnabled),
		}
	}

	return nil
}

// ACH Payment Validations
func (v *Validator) ValidateACHPayment(payment *ACHPayment) error {
	// Amount validation
	if payment.Amount < 0 {
		return ValidationError{
			Field:   "Amount",
			Value:   fmt.Sprintf("%d", payment.Amount),
			Rule:    "positive",
			Message: "amount cannot be negative",
		}
	}

	// Payee name required
	if strings.TrimSpace(payment.PayeeName) == "" {
		return ValidationError{
			Field:   "PayeeName",
			Rule:    "required",
			Message: "payee name is required",
		}
	}

	// Routing number validation (9 digits, valid checksum)
	if err := v.validateRoutingNumber(payment.RoutingNumber); err != nil {
		return WrapValidationError("RoutingNumber", payment.RoutingNumber, "routing_number", err)
	}

	// Account number validation
	if strings.TrimSpace(payment.AccountNumber) == "" || strings.Trim(payment.AccountNumber, "0") == "" {
		return ValidationError{
			Field:   "AccountNumber",
			Value:   payment.AccountNumber,
			Rule:    "required",
			Message: "account number cannot be blank or all zeros",
		}
	}

	// ACH Transaction code validation
	if !ValidACHTransactionCodes[payment.ACH_TransactionCode] {
		return ValidationError{
			Field:   "ACH_TransactionCode",
			Value:   payment.ACH_TransactionCode,
			Rule:    "valid_values",
			Message: "invalid ACH transaction code",
		}
	}

	// TIN validation
	if payment.TIN != "" && !v.isValidTIN(payment.TIN) {
		return ValidationError{
			Field:   "TIN",
			Value:   payment.TIN,
			Rule:    "tin_format",
			Message: fmt.Sprintf("TIN must be %d numeric digits", TINLength),
		}
	}

	// Payment ID required and unique within schedule
	if strings.TrimSpace(payment.PaymentID) == "" {
		return ValidationError{
			Field:   "PaymentID",
			Rule:    "required",
			Message: "payment ID is required",
		}
	}

	// IAT-specific validations
	if payment.StandardEntryClassCode == "IAT" {
		if strings.TrimSpace(payment.PayeeAddressLine1) == "" {
			return ValidationError{
				Field:   "PayeeAddressLine1",
				Rule:    "required_for_iat",
				Message: "address line 1 is required for IAT payments",
			}
		}
		if strings.TrimSpace(payment.CityName) == "" {
			return ValidationError{
				Field:   "CityName",
				Rule:    "required_for_iat",
				Message: "city name is required for IAT payments",
			}
		}
		if strings.TrimSpace(payment.CountryCodeText) == "" || payment.CountryCodeText == CountryCodeEmpty {
			return ValidationError{
				Field:   "CountryCodeText",
				Rule:    "required_for_iat",
				Message: "valid country code is required for IAT payments",
			}
		}
	}

	return nil
}

// Check Payment Validations
func (v *Validator) ValidateCheckPayment(payment *CheckPayment) error {
	// Amount validation (no zero dollar checks)
	if payment.Amount <= 0 {
		return ValidationError{
			Field:   "Amount",
			Value:   fmt.Sprintf("%d", payment.Amount),
			Rule:    "positive_non_zero",
			Message: "check amount must be greater than zero",
		}
	}

	// Payee name required
	if strings.TrimSpace(payment.PayeeName) == "" {
		return ValidationError{
			Field:   "PayeeName",
			Rule:    "required",
			Message: "payee name is required",
		}
	}

	// Address validation for mailability
	if strings.TrimSpace(payment.PayeeAddressLine1) == "" {
		// Mark as suspect for manual review, not invalid
		// This would be handled in a separate suspect validation pass
	}

	// TIN validation
	if payment.TIN != "" && !v.isValidTIN(payment.TIN) {
		return ValidationError{
			Field:   "TIN",
			Value:   payment.TIN,
			Rule:    "tin_format",
			Message: fmt.Sprintf("TIN must be %d numeric digits", TINLength),
		}
	}

	// Payment ID required and unique within schedule
	if strings.TrimSpace(payment.PaymentID) == "" {
		return ValidationError{
			Field:   "PaymentID",
			Rule:    "required",
			Message: "payment ID is required",
		}
	}

	return nil
}

// Same Day ACH Validations
func (v *Validator) ValidateSameDayACH(file *File) error {
	if file.Header.IsRequestedForSameDayACH != SDAFlagEnabled {
		return nil // Not a Same Day ACH file
	}

	// Rule: Can only contain ACH schedules
	for _, schedule := range file.Schedules {
		if _, ok := schedule.(*ACHSchedule); !ok {
			return ValidationError{
				Rule:    "sda_ach_only",
				Message: "Same Day ACH files can only contain ACH schedules",
			}
		}
	}

	// Rule: Individual payments must be <= $1,000,000
	maxSDAAmount := int64(MaxSDAAmountCents)
	for _, schedule := range file.Schedules {
		if achSchedule, ok := schedule.(*ACHSchedule); ok {
			for _, payment := range achSchedule.Payments {
				if achPayment, ok := payment.(*ACHPayment); ok {
					if achPayment.Amount > maxSDAAmount {
						return ValidationError{
							Field:   "Amount",
							Value:   fmt.Sprintf("%d", achPayment.Amount),
							Rule:    "sda_max_amount",
							Message: fmt.Sprintf("payment amount exceeds Same Day ACH limit of $%d", maxSDAAmount/100),
						}
					}
				}
			}
		}
	}

	// Rule: SEC code cannot be IAT
	for _, schedule := range file.Schedules {
		if achSchedule, ok := schedule.(*ACHSchedule); ok {
			if achSchedule.Header.StandardEntryClassCode == "IAT" {
				return ValidationError{
					Field:   "StandardEntryClassCode",
					Value:   "IAT",
					Rule:    "sda_no_iat",
					Message: "IAT payments are not allowed for Same Day ACH",
				}
			}
		}
	}

	return nil
}

// Balancing Validations - moved to validator_balance.go for better organization
// This function is now implemented in smaller, focused functions

// Helper functions
func (v *Validator) validateRoutingNumber(rtn string) error {
	if len(rtn) != RoutingNumberLength {
		return fmt.Errorf("routing number must be %d digits", RoutingNumberLength)
	}

	// Check if all digits
	if _, err := strconv.Atoi(rtn); err != nil {
		return fmt.Errorf("routing number must contain only digits")
	}

	// Validate first two digits
	firstTwo, _ := strconv.Atoi(rtn[:2])
	validRanges := [][]int{{0, 12}, {21, 32}, {61, 72}, {80, 80}}
	valid := false
	for _, r := range validRanges {
		if firstTwo >= r[0] && firstTwo <= r[1] {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid routing number prefix")
	}

	// Check digit validation (modulus 10)
	sum := 0
	weights := []int{3, 7, 1, 3, 7, 1, 3, 7, 1}
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(rtn[i]))
		sum += digit * weights[i]
	}
	if sum%RoutingNumberModulus != 0 {
		return fmt.Errorf("invalid routing number check digit")
	}

	return nil
}

func (v *Validator) isValidTIN(tin string) bool {
	if len(tin) != TINLength {
		return false
	}
	_, err := strconv.Atoi(tin)
	return err == nil
}

func (v *Validator) validateACHPaymentOrder(schedule *ACHSchedule) error {
	if len(schedule.Payments) < 2 {
		return nil // No order to validate
	}

	var lastRTN string
	for i, payment := range schedule.Payments {
		if achPayment, ok := payment.(*ACHPayment); ok {
			if i > 0 && achPayment.RoutingNumber < lastRTN {
				return fmt.Errorf("ACH payments must be in routing number order")
			}
			lastRTN = achPayment.RoutingNumber
		}
	}
	return nil
}

// CTX Validation Rules
func (v *Validator) ValidateCTXAddendum(payment *ACHPayment) error {
	if payment.StandardEntryClassCode != "CTX" {
		return nil
	}

	// CTX must have at least one addendum
	if len(payment.Addenda) == 0 {
		return ValidationError{
			Field:   "Addenda",
			Rule:    "ctx_required",
			Message: "CTX payments must have at least one addendum record",
		}
	}

	// First addendum must start with ISA
	if len(payment.Addenda) > 0 && len(payment.Addenda[0].AddendaInformation) >= CTXEDISegmentLength {
		if payment.Addenda[0].AddendaInformation[:CTXEDISegmentLength] != CTXEDISegmentIdentifier {
			return ValidationError{
				Field:   "Addenda[0]",
				Rule:    "ctx_isa_required",
				Message: "first CTX addendum must start with ISA segment",
			}
		}
	}

	// Validate EDI structure
	for _, addendum := range payment.Addenda {
		if addendum.RecordCode != string(RecordTypeACHAddendumCTX) {
			return ValidationError{
				Field:   "Addenda.RecordCode",
				Value:   addendum.RecordCode,
				Rule:    "ctx_record_code",
				Message: fmt.Sprintf("CTX addenda must use record code '%s'", RecordTypeACHAddendumCTX),
			}
		}
	}

	return nil
}

// Agency-specific validations based on Custom Agency Rule ID
func (v *Validator) ValidateAgencySpecific(payment Payment, customAgencyRuleID string) error {
	switch customAgencyRuleID {
	case "IRS":
		return v.validateIRSPayment(payment)
	case "VA", "VACP":
		return v.validateVAPayment(payment)
	case "SSA", "SSA-Daily", "SSA-A":
		return v.validateSSAPayment(payment)
	case "RRB":
		return v.validateRRBPayment(payment)
	case "CCC":
		return v.validateCCCPayment(payment)
	}
	return nil
}

func (v *Validator) validateIRSPayment(payment Payment) error {
	// IRS-specific reconcilement field validation
	switch p := payment.(type) {
	case *ACHPayment:
		if len(p.Reconcilement) != ReconcilementLength {
			return ValidationError{
				Field:   "Reconcilement",
				Rule:    "irs_format",
				Message: fmt.Sprintf("IRS reconcilement must be %d characters", ReconcilementLength),
			}
		}
	case *CheckPayment:
		if len(p.Reconcilement) != ReconcilementLength {
			return ValidationError{
				Field:   "Reconcilement",
				Rule:    "irs_format",
				Message: fmt.Sprintf("IRS reconcilement must be %d characters", ReconcilementLength),
			}
		}
	}
	return nil
}

func (v *Validator) validateVAPayment(payment Payment) error {
	recon := payment.GetReconcilement()

	// Validate reconcilement field length
	if len(recon) != ReconcilementLength {
		return ValidationError{
			Field:   "Reconcilement",
			Rule:    "va_recon_length",
			Message: fmt.Sprintf("VA reconcilement must be %d characters", ReconcilementLength),
		}
	}

	// Parse VA reconcilement fields
	parser := &AgencyReconcilementParser{}
	isACH := false
	if _, ok := payment.(*ACHPayment); ok {
		isACH = true
	}
	fields := parser.ParseVAReconcilement(recon, isACH)

	// Validate required Station Code (2 chars)
	stationCode := fields["StationCode"]
	if len(stationCode) == 0 {
		return NewFieldRequiredError("Reconcilement.StationCode")
	}
	if len(stationCode) > 2 {
		return ValidationError{
			Field:   "Reconcilement.StationCode",
			Rule:    "va_station_length",
			Message: "VA station code must be 2 characters or less",
		}
	}
	// TODO: Validate against approved VA station codes list when available
	// Contact Treasury for valid station code ranges

	// Validate required FIN Code (2 chars)
	finCode := fields["FinCode"]
	if len(finCode) == 0 {
		return NewFieldRequiredError("Reconcilement.FinCode")
	}
	if len(finCode) > 2 {
		return ValidationError{
			Field:   "Reconcilement.FinCode",
			Rule:    "va_fin_length",
			Message: "VA FIN code must be 2 characters or less",
		}
	}
	// TODO: Validate against approved VA FIN codes list when available
	// Contact Treasury for valid FIN code ranges

	// Check payment specific validations
	if _, ok := payment.(*CheckPayment); ok {
		// Validate courtesy code for checks
		courtesyCode := fields["CourtesyCode"]
		if len(courtesyCode) > 1 {
			return ValidationError{
				Field:   "Reconcilement.CourtesyCode",
				Rule:    "va_courtesy_length",
				Message: "VA courtesy code must be 1 character or less",
			}
		}
		// TODO: Validate against approved VA courtesy codes when available
		// Contact Treasury for valid courtesy code values

		// Validate appropriation code
		appropCode := fields["AppropCode"]
		if len(appropCode) > 1 {
			return ValidationError{
				Field:   "Reconcilement.AppropCode",
				Rule:    "va_approp_length",
				Message: "VA appropriation code must be 1 character or less",
			}
		}
		// TODO: Validate against VA appropriation codes when available

		// Validate policy number format
		policyNum := fields["PolicyNum"]
		if len(policyNum) > 2 {
			return ValidationError{
				Field:   "Reconcilement.PolicyNum",
				Rule:    "va_policy_length",
				Message: "VA policy number must be 2 characters or less",
			}
		}
		// TODO: Implement policy number format validation when specifications available

		// Validate name code for checks
		nameCode := fields["NameCode"]
		if len(nameCode) > 3 {
			return ValidationError{
				Field:   "Reconcilement.NameCode",
				Rule:    "va_name_code_length",
				Message: "VA name code must be 3 characters or less",
			}
		}
	}

	// Validate payment amount constraints
	// TODO: Implement VA-specific payment amount limits when available
	// TODO: Validate payment type restrictions for VA when available
	// TODO: Implement cross-field validation rules when business requirements available

	return nil
}

func (v *Validator) validateSSAPayment(payment Payment) error {
	recon := payment.GetReconcilement()

	// Validate reconcilement field length
	if len(recon) != ReconcilementLength {
		return ValidationError{
			Field:   "Reconcilement",
			Rule:    "ssa_recon_length",
			Message: fmt.Sprintf("SSA reconcilement must be %d characters", ReconcilementLength),
		}
	}

	// Get the agency rule ID to determine SSA variant
	ruleID := v.CustomAgencyRuleID
	if ruleID == "" {
		ruleID = "SSA" // Default to standard SSA
	}

	// Parse SSA reconcilement fields
	parser := &AgencyReconcilementParser{}
	fields := parser.ParseSSAReconcilement(recon, ruleID)

	// Validate Program Service Center Code (1 char, required)
	pscCode := fields["ProgramServiceCenterCode"]
	if len(pscCode) == 0 {
		return NewFieldRequiredError("Reconcilement.ProgramServiceCenterCode")
	}
	if len(pscCode) > 1 {
		return ValidationError{
			Field:   "Reconcilement.ProgramServiceCenterCode",
			Rule:    "ssa_psc_length",
			Message: "SSA program service center code must be 1 character",
		}
	}
	// TODO: Validate against approved SSA program service center codes when available
	// Contact Treasury for valid PSC code values (typically 0-9, A-Z)

	// Validate Payment ID Code (2 chars, required)
	paymentIDCode := fields["PaymentIDCode"]
	if len(paymentIDCode) == 0 {
		return NewFieldRequiredError("Reconcilement.PaymentIDCode")
	}
	if len(paymentIDCode) > 2 {
		return ValidationError{
			Field:   "Reconcilement.PaymentIDCode",
			Rule:    "ssa_payment_id_length",
			Message: "SSA payment ID code must be 2 characters or less",
		}
	}
	// TODO: Validate against approved SSA payment ID codes when available
	// Contact Treasury for valid payment ID code values

	// TIN Indicator Offset validation (not applicable for SSA-A)
	if ruleID != "SSA-A" {
		tinOffset := fields["TINIndicatorOffset"]
		if len(tinOffset) > 1 {
			return ValidationError{
				Field:   "Reconcilement.TINIndicatorOffset",
				Rule:    "ssa_tin_offset_length",
				Message: "SSA TIN indicator offset must be 1 character or less",
			}
		}
		// TODO: Validate TIN offset values when business rules available
		// Typically should be numeric or specific codes
	}

	// Validate payment type constraints for SSA
	// TODO: Implement SSA-specific payment type restrictions when available
	// TODO: Validate benefit type codes when specifications available
	// TODO: Implement SSA payment amount limits when available
	// TODO: Cross-validate PSC and Payment ID combinations when rules available

	return nil
}

func (v *Validator) validateRRBPayment(payment Payment) error {
	// RRB-specific reconcilement field validation
	recon := payment.GetReconcilement()
	if len(recon) != ReconcilementLength {
		return NewFieldLengthError("Reconcilement", recon, 100, len(recon))
	}

	// Parse RRB reconcilement fields using parser
	parser := &AgencyReconcilementParser{}
	fields := parser.ParseRRBReconcilement(recon)

	// Validate required fields are present and properly formatted
	if beneficiarySymbol := fields["BeneficiarySymbol"]; len(beneficiarySymbol) != 2 {
		return NewFieldFormatError("BeneficiarySymbol", beneficiarySymbol, "be exactly 2 alphanumeric characters")
	}

	if prefixCode := fields["PrefixCode"]; len(prefixCode) != 1 {
		return ValidationError{
			Field:   "PrefixCode",
			Value:   prefixCode,
			Rule:    "rrb_prefix_code",
			Message: "RRB Prefix Code must be exactly 1 alphanumeric character",
		}
	}

	if payeeCode := fields["PayeeCode"]; len(payeeCode) != 1 {
		return ValidationError{
			Field:   "PayeeCode",
			Value:   payeeCode,
			Rule:    "rrb_payee_code",
			Message: "RRB Payee Code must be exactly 1 alphanumeric character",
		}
	}

	if objectCode := fields["ObjectCode"]; len(objectCode) != 1 {
		return ValidationError{
			Field:   "ObjectCode",
			Value:   objectCode,
			Rule:    "rrb_object_code",
			Message: "RRB Object Code must be exactly 1 alphanumeric character",
		}
	}

	// TODO: Add business rule validation when agency provides valid code ranges:
	// - Valid Beneficiary Symbol values for different RRB programs
	// - Valid Prefix Code values for railroad retirement vs survivor benefits
	// - Valid Payee Code values for different recipient types
	// - Valid Object Code values for PACER integration

	return nil
}

func (v *Validator) validateCCCPayment(payment Payment) error {
	// CCC-specific reconcilement field validation
	recon := payment.GetReconcilement()
	if len(recon) != ReconcilementLength {
		return ValidationError{
			Field:   "Reconcilement",
			Rule:    "ccc_format",
			Message: fmt.Sprintf("CCC reconcilement field must be exactly %d characters, got %d", ReconcilementLength, len(recon)),
		}
	}

	// Parse CCC reconcilement fields using parser
	parser := &AgencyReconcilementParser{}
	fields := parser.ParseCCCReconcilement(recon)

	// Validate TOP Payment Agency ID (positions 1-2)
	if topPaymentAgencyID := fields["TOPPaymentAgencyID"]; len(topPaymentAgencyID) > 0 {
		if len(topPaymentAgencyID) != 2 {
			return ValidationError{
				Field:   "TOPPaymentAgencyID",
				Value:   topPaymentAgencyID,
				Rule:    "ccc_top_payment_agency_id",
				Message: "CCC TOP Payment Agency ID must be exactly 2 alphabetic characters when present",
			}
		}
		// Validate alphabetic characters only
		for _, r := range topPaymentAgencyID {
			if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
				return ValidationError{
					Field:   "TOPPaymentAgencyID",
					Value:   topPaymentAgencyID,
					Rule:    "ccc_top_payment_agency_id_format",
					Message: "CCC TOP Payment Agency ID must contain only alphabetic characters",
				}
			}
		}
	}

	// Validate TOP Agency Site ID (positions 3-4)
	if topAgencySiteID := fields["TOPAgencySiteID"]; len(topAgencySiteID) > 0 {
		if len(topAgencySiteID) != 2 {
			return ValidationError{
				Field:   "TOPAgencySiteID",
				Value:   topAgencySiteID,
				Rule:    "ccc_top_agency_site_id",
				Message: "CCC TOP Agency Site ID must be exactly 2 alphabetic characters when present",
			}
		}
		// Validate alphabetic characters only
		for _, r := range topAgencySiteID {
			if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')) {
				return ValidationError{
					Field:   "TOPAgencySiteID",
					Value:   topAgencySiteID,
					Rule:    "ccc_top_agency_site_id_format",
					Message: "CCC TOP Agency Site ID must contain only alphabetic characters",
				}
			}
		}
	}

	// TODO: Implement schedule-level validation rule:
	// "If one payment within a schedule contains TOP Agency/Site ID values,
	//  the entire schedule will be sent to TOP with these values."
	// This requires access to the parent schedule context.

	// TODO: Add business rule validation when agency provides requirements:
	// - Valid TOP Payment Agency ID codes for different CCC programs
	// - Valid TOP Agency Site ID codes for different agricultural regions
	// - Payment amount limits for commodity support payments
	// - Cross-validation with payment type codes

	return nil
}

// Hexadecimal character validation
func (v *Validator) ValidateHexCharacters(data string) error {
	for i, r := range data {
		if r < MinHexCharacter { // Invalid hex characters are below 0x40
			return ValidationError{
				Field:   "data",
				Rule:    "hex_validation",
				Message: fmt.Sprintf("invalid hex character at position %d: 0x%02X", i, r),
			}
		}
	}
	return nil
}

// Schedule number validation
func (v *Validator) ValidateScheduleNumber(scheduleNumber string) error {
	// Valid characters: A-Z, 0-9, dash (-)
	validPattern := regexp.MustCompile(`^[A-Z0-9\-]+$`)
	cleaned := strings.ToUpper(strings.TrimSpace(scheduleNumber))

	if cleaned == "" {
		return ValidationError{
			Field:   "ScheduleNumber",
			Rule:    "required",
			Message: "schedule number cannot be blank",
		}
	}

	if !validPattern.MatchString(cleaned) {
		return ValidationError{
			Field:   "ScheduleNumber",
			Value:   scheduleNumber,
			Rule:    "valid_characters",
			Message: "schedule number can only contain A-Z, 0-9, and dash (-)",
		}
	}

	return nil
}
