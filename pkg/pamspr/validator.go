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
func (v *Validator) ValidateFileStructure(file *File) error {
	// Rule: File must have header and trailer
	if file.Header == nil {
		return ValidationError{Field: "FileHeader", Rule: "required", Message: "file header is required"}
	}
	if file.Trailer == nil {
		return ValidationError{Field: "FileTrailer", Rule: "required", Message: "file trailer is required"}
	}

	// Rule: A schedule can contain only one type of payment
	for i, schedule := range file.Schedules {
		switch s := schedule.(type) {
		case *ACHSchedule:
			for _, payment := range s.Payments {
				if _, ok := payment.(*ACHPayment); !ok {
					return ValidationError{
						Field:   fmt.Sprintf("Schedule[%d]", i),
						Rule:    "payment_type_consistency",
						Message: "ACH schedule cannot contain non-ACH payments",
					}
				}
			}
		case *CheckSchedule:
			for _, payment := range s.Payments {
				if _, ok := payment.(*CheckPayment); !ok {
					return ValidationError{
						Field:   fmt.Sprintf("Schedule[%d]", i),
						Rule:    "payment_type_consistency",
						Message: "Check schedule cannot contain non-Check payments",
					}
				}
			}
		}
	}

	// Rule: ACH payments must be in routing number order within schedule
	for i, schedule := range file.Schedules {
		if achSchedule, ok := schedule.(*ACHSchedule); ok {
			if err := v.validateACHPaymentOrder(achSchedule); err != nil {
				return ValidationError{
					Field:   fmt.Sprintf("Schedule[%d]", i),
					Rule:    "routing_number_order",
					Message: err.Error(),
				}
			}
		}
	}

	// Rule: Same Day ACH validation
	if file.Header != nil && file.Header.IsRequestedForSameDayACH == "1" {
		if err := v.ValidateSameDayACH(file); err != nil {
			return err
		}
	}

	// Rule: CTX payment validation
	for _, schedule := range file.Schedules {
		if achSchedule, ok := schedule.(*ACHSchedule); ok {
			for _, payment := range achSchedule.Payments {
				if achPayment, ok := payment.(*ACHPayment); ok {
					if err := v.ValidateCTXAddendum(achPayment); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

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
	if header.StandardPaymentVersion != "502" {
		return ValidationError{
			Field:   "StandardPaymentVersion",
			Value:   header.StandardPaymentVersion,
			Rule:    "version",
			Message: "current published version must be '502'",
		}
	}

	// Same Day ACH flag validation
	if header.IsRequestedForSameDayACH != "0" && header.IsRequestedForSameDayACH != "1" && header.IsRequestedForSameDayACH != "" {
		return ValidationError{
			Field:   "IsRequestedForSameDayACH",
			Value:   header.IsRequestedForSameDayACH,
			Rule:    "valid_values",
			Message: "must be '0', '1', or blank",
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
		return ValidationError{
			Field:   "RoutingNumber",
			Value:   payment.RoutingNumber,
			Rule:    "routing_number",
			Message: err.Error(),
		}
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
	validTransactionCodes := map[string]bool{
		"22": true, "23": true, "24": true,
		"32": true, "33": true, "34": true,
		"42": true, "43": true,
		"52": true, "53": true,
	}
	if !validTransactionCodes[payment.ACHTransactionCode] {
		return ValidationError{
			Field:   "ACHTransactionCode",
			Value:   payment.ACHTransactionCode,
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
			Message: "TIN must be 9 numeric digits",
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
		if strings.TrimSpace(payment.CountryCodeText) == "" || payment.CountryCodeText == "00" {
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
			Message: "TIN must be 9 numeric digits",
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
	if file.Header.IsRequestedForSameDayACH != "1" {
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
	maxSDAAmount := int64(100000000) // $1,000,000 in cents
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

// Balancing Validations
func (v *Validator) ValidateBalancing(file *File) error {
	// Calculate actual totals
	totalRecords := int64(2) // Header + Trailer
	totalPayments := int64(0)
	totalAmount := int64(0)

	for _, schedule := range file.Schedules {
		scheduleRecords := int64(2) // Schedule header + trailer
		schedulePayments := int64(0)
		scheduleAmount := int64(0)

		switch s := schedule.(type) {
		case *ACHSchedule:
			for _, payment := range s.Payments {
				scheduleRecords++ // Payment record
				schedulePayments++
				if achPayment, ok := payment.(*ACHPayment); ok {
					scheduleAmount += achPayment.Amount

					// Count associated records
					scheduleRecords += int64(len(achPayment.Addenda))
					scheduleRecords += int64(len(achPayment.CARSTASBETC))
					if achPayment.DNP != nil {
						scheduleRecords++
					}
				}
			}

			// Validate schedule trailer
			if s.Trailer != nil {
				if s.Trailer.ScheduleCount != schedulePayments {
					return ValidationError{
						Field:   "ScheduleTrailer.ScheduleCount",
						Value:   fmt.Sprintf("%d", s.Trailer.ScheduleCount),
						Rule:    "balance",
						Message: fmt.Sprintf("expected %d payments, got %d", schedulePayments, s.Trailer.ScheduleCount),
					}
				}
				if s.Trailer.ScheduleAmount != scheduleAmount {
					return ValidationError{
						Field:   "ScheduleTrailer.ScheduleAmount",
						Value:   fmt.Sprintf("%d", s.Trailer.ScheduleAmount),
						Rule:    "balance",
						Message: fmt.Sprintf("expected amount %d, got %d", scheduleAmount, s.Trailer.ScheduleAmount),
					}
				}
			}
		case *CheckSchedule:
			// Similar logic for check schedules
			for _, payment := range s.Payments {
				scheduleRecords++ // Payment record
				schedulePayments++
				if checkPayment, ok := payment.(*CheckPayment); ok {
					scheduleAmount += checkPayment.Amount

					// Count associated records
					if checkPayment.Stub != nil {
						scheduleRecords++
					}
					scheduleRecords += int64(len(checkPayment.CARSTASBETC))
					if checkPayment.DNP != nil {
						scheduleRecords++
					}
				}
			}

			// Validate schedule trailer
			if s.Trailer != nil {
				if s.Trailer.ScheduleCount != schedulePayments {
					return ValidationError{
						Field:   "ScheduleTrailer.ScheduleCount",
						Value:   fmt.Sprintf("%d", s.Trailer.ScheduleCount),
						Rule:    "balance",
						Message: fmt.Sprintf("expected %d payments, got %d", schedulePayments, s.Trailer.ScheduleCount),
					}
				}
				if s.Trailer.ScheduleAmount != scheduleAmount {
					return ValidationError{
						Field:   "ScheduleTrailer.ScheduleAmount",
						Value:   fmt.Sprintf("%d", s.Trailer.ScheduleAmount),
						Rule:    "balance",
						Message: fmt.Sprintf("expected amount %d, got %d", scheduleAmount, s.Trailer.ScheduleAmount),
					}
				}
			}
		}

		totalRecords += scheduleRecords
		totalPayments += schedulePayments
		totalAmount += scheduleAmount
	}

	// Validate file trailer
	if file.Trailer.TotalCountRecords != totalRecords {
		return ValidationError{
			Field:   "FileTrailer.TotalCountRecords",
			Value:   fmt.Sprintf("%d", file.Trailer.TotalCountRecords),
			Rule:    "balance",
			Message: fmt.Sprintf("expected %d records, got %d", totalRecords, file.Trailer.TotalCountRecords),
		}
	}

	if file.Trailer.TotalCountPayments != totalPayments {
		return ValidationError{
			Field:   "FileTrailer.TotalCountPayments",
			Value:   fmt.Sprintf("%d", file.Trailer.TotalCountPayments),
			Rule:    "balance",
			Message: fmt.Sprintf("expected %d payments, got %d", totalPayments, file.Trailer.TotalCountPayments),
		}
	}

	if file.Trailer.TotalAmountPayments != totalAmount {
		return ValidationError{
			Field:   "FileTrailer.TotalAmountPayments",
			Value:   fmt.Sprintf("%d", file.Trailer.TotalAmountPayments),
			Rule:    "balance",
			Message: fmt.Sprintf("expected amount %d, got %d", totalAmount, file.Trailer.TotalAmountPayments),
		}
	}

	return nil
}

// Helper functions
func (v *Validator) validateRoutingNumber(rtn string) error {
	if len(rtn) != 9 {
		return fmt.Errorf("routing number must be 9 digits")
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
	if sum%10 != 0 {
		return fmt.Errorf("invalid routing number check digit")
	}

	return nil
}

func (v *Validator) isValidTIN(tin string) bool {
	if len(tin) != 9 {
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
	if len(payment.Addenda) > 0 && len(payment.Addenda[0].AddendaInformation) >= 3 {
		if payment.Addenda[0].AddendaInformation[:3] != "ISA" {
			return ValidationError{
				Field:   "Addenda[0]",
				Rule:    "ctx_isa_required",
				Message: "first CTX addendum must start with ISA segment",
			}
		}
	}

	// Validate EDI structure
	for _, addendum := range payment.Addenda {
		if addendum.RecordCode != "04" {
			return ValidationError{
				Field:   "Addenda.RecordCode",
				Value:   addendum.RecordCode,
				Rule:    "ctx_record_code",
				Message: "CTX addenda must use record code '04'",
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
		if len(p.Reconcilement) != 100 {
			return ValidationError{
				Field:   "Reconcilement",
				Rule:    "irs_format",
				Message: "IRS reconcilement must be 100 characters",
			}
		}
	case *CheckPayment:
		if len(p.Reconcilement) != 100 {
			return ValidationError{
				Field:   "Reconcilement",
				Rule:    "irs_format",
				Message: "IRS reconcilement must be 100 characters",
			}
		}
	}
	return nil
}

func (v *Validator) validateVAPayment(payment Payment) error {
	recon := payment.GetReconcilement()

	// Validate reconcilement field length
	if len(recon) != 100 {
		return ValidationError{
			Field:   "Reconcilement",
			Rule:    "va_recon_length",
			Message: "VA reconcilement must be 100 characters",
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
		return ValidationError{
			Field:   "Reconcilement.StationCode",
			Rule:    "va_station_required",
			Message: "VA station code is required",
		}
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
		return ValidationError{
			Field:   "Reconcilement.FinCode",
			Rule:    "va_fin_required",
			Message: "VA FIN code is required",
		}
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
	if len(recon) != 100 {
		return ValidationError{
			Field:   "Reconcilement",
			Rule:    "ssa_recon_length",
			Message: "SSA reconcilement must be 100 characters",
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
		return ValidationError{
			Field:   "Reconcilement.ProgramServiceCenterCode",
			Rule:    "ssa_psc_required",
			Message: "SSA program service center code is required",
		}
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
		return ValidationError{
			Field:   "Reconcilement.PaymentIDCode",
			Rule:    "ssa_payment_id_required",
			Message: "SSA payment ID code is required",
		}
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
	// RRB-specific validations are not currently implemented
	// Future enhancements may add validation for RRB payment requirements
	_ = payment
	return nil
}

func (v *Validator) validateCCCPayment(payment Payment) error {
	// CCC-specific validations are not currently implemented
	// Future enhancements may add validation for CCC payment requirements
	_ = payment
	return nil
}

// Hexadecimal character validation
func (v *Validator) ValidateHexCharacters(data string) error {
	for i, r := range data {
		if r < 0x40 { // Invalid hex characters are 0x00-0x3F
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
