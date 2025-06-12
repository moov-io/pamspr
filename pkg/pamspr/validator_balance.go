package pamspr

import "fmt"

// ScheduleBalanceInfo holds balance calculation results for a schedule
type ScheduleBalanceInfo struct {
	Records  int64
	Payments int64
	Amount   int64
}

// FileBalanceInfo holds balance calculation results for an entire file
type FileBalanceInfo struct {
	TotalRecords  int64
	TotalPayments int64
	TotalAmount   int64
}

// calculateACHScheduleBalance calculates balance information for an ACH schedule
func (v *Validator) calculateACHScheduleBalance(schedule *ACHSchedule) ScheduleBalanceInfo {
	balance := ScheduleBalanceInfo{
		Records: 2, // Schedule header + trailer
	}

	for _, payment := range schedule.Payments {
		balance.Records++ // Payment record
		balance.Payments++

		if achPayment, ok := payment.(*ACHPayment); ok {
			balance.Amount += achPayment.Amount

			// Count associated records
			balance.Records += int64(len(achPayment.Addenda))
			balance.Records += int64(len(achPayment.CARSTASBETC))
			if achPayment.DNP != nil {
				balance.Records++
			}
		}
	}

	return balance
}

// calculateCheckScheduleBalance calculates balance information for a check schedule
func (v *Validator) calculateCheckScheduleBalance(schedule *CheckSchedule) ScheduleBalanceInfo {
	balance := ScheduleBalanceInfo{
		Records: 2, // Schedule header + trailer
	}

	for _, payment := range schedule.Payments {
		balance.Records++ // Payment record
		balance.Payments++

		if checkPayment, ok := payment.(*CheckPayment); ok {
			balance.Amount += checkPayment.Amount

			// Count associated records
			if checkPayment.Stub != nil {
				balance.Records++
			}
			balance.Records += int64(len(checkPayment.CARSTASBETC))
			if checkPayment.DNP != nil {
				balance.Records++
			}
		}
	}

	return balance
}

// validateScheduleTrailer validates that schedule trailer matches calculated values
func (v *Validator) validateScheduleTrailer(trailer *ScheduleTrailer, balance ScheduleBalanceInfo, scheduleType string) error {
	if trailer == nil {
		return NewFieldRequiredError(fmt.Sprintf("%sScheduleTrailer", scheduleType))
	}

	if trailer.ScheduleCount != balance.Payments {
		return ValidationError{
			Field:   "ScheduleTrailer.ScheduleCount",
			Value:   fmt.Sprintf("%d", trailer.ScheduleCount),
			Rule:    "balance",
			Message: fmt.Sprintf("expected %d payments, got %d", balance.Payments, trailer.ScheduleCount),
		}
	}

	if trailer.ScheduleAmount != balance.Amount {
		return ValidationError{
			Field:   "ScheduleTrailer.ScheduleAmount",
			Value:   fmt.Sprintf("%d", trailer.ScheduleAmount),
			Rule:    "balance",
			Message: fmt.Sprintf("expected amount %d, got %d", balance.Amount, trailer.ScheduleAmount),
		}
	}

	return nil
}

// validateFileTrailer validates that file trailer matches calculated values
func (v *Validator) validateFileTrailer(trailer *FileTrailer, balance FileBalanceInfo) error {
	if trailer.TotalCountRecords != balance.TotalRecords {
		return ValidationError{
			Field:   "FileTrailer.TotalCountRecords",
			Value:   fmt.Sprintf("%d", trailer.TotalCountRecords),
			Rule:    "balance",
			Message: fmt.Sprintf("expected %d records, got %d", balance.TotalRecords, trailer.TotalCountRecords),
		}
	}

	if trailer.TotalCountPayments != balance.TotalPayments {
		return ValidationError{
			Field:   "FileTrailer.TotalCountPayments",
			Value:   fmt.Sprintf("%d", trailer.TotalCountPayments),
			Rule:    "balance",
			Message: fmt.Sprintf("expected %d payments, got %d", balance.TotalPayments, trailer.TotalCountPayments),
		}
	}

	if trailer.TotalAmountPayments != balance.TotalAmount {
		return ValidationError{
			Field:   "FileTrailer.TotalAmountPayments",
			Value:   fmt.Sprintf("%d", trailer.TotalAmountPayments),
			Rule:    "balance",
			Message: fmt.Sprintf("expected amount %d, got %d", balance.TotalAmount, trailer.TotalAmountPayments),
		}
	}

	return nil
}

// ValidateBalancing validates that all totals in the file are balanced
// This replaces the original large ValidateBalancing function with a more focused approach
func (v *Validator) ValidateBalancing(file *File) error {
	// Calculate file totals
	fileBalance := FileBalanceInfo{
		TotalRecords: 2, // Header + Trailer
	}

	// Process each schedule
	for _, schedule := range file.Schedules {
		var scheduleBalance ScheduleBalanceInfo
		var err error

		switch s := schedule.(type) {
		case *ACHSchedule:
			scheduleBalance = v.calculateACHScheduleBalance(s)
			err = v.validateScheduleTrailer(s.Trailer, scheduleBalance, "ACH")

		case *CheckSchedule:
			scheduleBalance = v.calculateCheckScheduleBalance(s)
			err = v.validateScheduleTrailer(s.Trailer, scheduleBalance, "Check")
		}

		if err != nil {
			return err
		}

		// Add schedule totals to file totals
		fileBalance.TotalRecords += scheduleBalance.Records
		fileBalance.TotalPayments += scheduleBalance.Payments
		fileBalance.TotalAmount += scheduleBalance.Amount
	}

	// Validate file trailer
	return v.validateFileTrailer(file.Trailer, fileBalance)
}
