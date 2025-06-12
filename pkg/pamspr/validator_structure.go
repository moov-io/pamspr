package pamspr

import "fmt"

// ValidateFileStructure validates the overall structure and rules of a PAM SPR file
func (v *Validator) ValidateFileStructure(file *File) error {
	// Validate required components
	if err := v.validateRequiredComponents(file); err != nil {
		return err
	}

	// Validate payment type consistency within schedules
	if err := v.validateSchedulePaymentConsistency(file); err != nil {
		return err
	}

	// Validate ACH-specific rules
	if err := v.validateACHRules(file); err != nil {
		return err
	}

	// Validate Same Day ACH if applicable
	if err := v.validateSameDayACHIfApplicable(file); err != nil {
		return err
	}

	return nil
}

// validateRequiredComponents ensures file has required header and trailer
func (v *Validator) validateRequiredComponents(file *File) error {
	if file.Header == nil {
		return NewFieldRequiredError("FileHeader")
	}
	if file.Trailer == nil {
		return NewFieldRequiredError("FileTrailer")
	}
	return nil
}

// validateSchedulePaymentConsistency ensures each schedule contains only one payment type
func (v *Validator) validateSchedulePaymentConsistency(file *File) error {
	for i, schedule := range file.Schedules {
		if err := v.validateSingleScheduleConsistency(schedule, i); err != nil {
			return err
		}
	}
	return nil
}

// validateSingleScheduleConsistency validates payment type consistency for a single schedule
func (v *Validator) validateSingleScheduleConsistency(schedule Schedule, index int) error {
	switch s := schedule.(type) {
	case *ACHSchedule:
		return v.validateACHScheduleConsistency(s, index)
	case *CheckSchedule:
		return v.validateCheckScheduleConsistency(s, index)
	default:
		return ValidationError{
			Field:   fmt.Sprintf("Schedule[%d]", index),
			Rule:    "schedule_type",
			Message: "unknown schedule type",
		}
	}
}

// validateACHScheduleConsistency ensures ACH schedule contains only ACH payments
func (v *Validator) validateACHScheduleConsistency(schedule *ACHSchedule, index int) error {
	for _, payment := range schedule.Payments {
		if _, ok := payment.(*ACHPayment); !ok {
			return ValidationError{
				Field:   fmt.Sprintf("Schedule[%d]", index),
				Rule:    "payment_type_consistency",
				Message: "ACH schedule cannot contain non-ACH payments",
			}
		}
	}
	return nil
}

// validateCheckScheduleConsistency ensures Check schedule contains only Check payments
func (v *Validator) validateCheckScheduleConsistency(schedule *CheckSchedule, index int) error {
	for _, payment := range schedule.Payments {
		if _, ok := payment.(*CheckPayment); !ok {
			return ValidationError{
				Field:   fmt.Sprintf("Schedule[%d]", index),
				Rule:    "payment_type_consistency",
				Message: "Check schedule cannot contain non-Check payments",
			}
		}
	}
	return nil
}

// validateACHRules validates ACH-specific rules including payment order and CTX addenda
func (v *Validator) validateACHRules(file *File) error {
	for i, schedule := range file.Schedules {
		if achSchedule, ok := schedule.(*ACHSchedule); ok {
			// Validate payment order
			if err := v.validateACHSchedulePaymentOrder(achSchedule, i); err != nil {
				return err
			}

			// Validate CTX addenda for all payments
			if err := v.validateACHScheduleCTXPayments(achSchedule); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateACHSchedulePaymentOrder validates ACH payments are in routing number order
func (v *Validator) validateACHSchedulePaymentOrder(schedule *ACHSchedule, index int) error {
	if err := v.validateACHPaymentOrder(schedule); err != nil {
		return ValidationError{
			Field:   fmt.Sprintf("Schedule[%d]", index),
			Rule:    "routing_number_order",
			Message: err.Error(),
		}
	}
	return nil
}

// validateACHScheduleCTXPayments validates CTX addenda for all payments in a schedule
func (v *Validator) validateACHScheduleCTXPayments(schedule *ACHSchedule) error {
	for _, payment := range schedule.Payments {
		if achPayment, ok := payment.(*ACHPayment); ok {
			if err := v.ValidateCTXAddendum(achPayment); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateSameDayACHIfApplicable validates Same Day ACH rules if the flag is set
func (v *Validator) validateSameDayACHIfApplicable(file *File) error {
	if file.Header != nil && file.Header.IsRequestedForSameDayACH == SDAFlagEnabled {
		return v.ValidateSameDayACH(file)
	}
	return nil
}
