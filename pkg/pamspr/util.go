package pamspr

import (
	"fmt"
	"strings"
	"unicode"
)

// PadLeft is deprecated - use SecurePadLeft instead
// Kept only for tests that explicitly test padding behavior
func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s[:length]
	}
	return strings.Repeat(string(padChar), length-len(s)) + s
}

// PadRight is deprecated - use SecurePadRight instead
// Kept only for tests that explicitly test padding behavior
func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s[:length]
	}
	return s + strings.Repeat(string(padChar), length-len(s))
}

// PadNumeric is deprecated - use SecurePadNumeric instead
// Kept only for tests that explicitly test padding behavior
func PadNumeric(s string, length int) string {
	// Remove non-numeric characters using strings.Builder for efficiency
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity
	for _, r := range s {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	return PadLeft(builder.String(), length, '0')
}

// TruncateOrPad is deprecated - use SecureTruncateOrPad instead
// Kept only for tests that explicitly test padding behavior
func TruncateOrPad(s string, length int, padRight bool) string {
	if padRight {
		return PadRight(s, length, ' ')
	}
	return PadLeft(s, length, ' ')
}

// FormatCents converts cents to a formatted dollar string
func FormatCents(cents int64) string {
	dollars := float64(cents) / 100.0
	return fmt.Sprintf("%.2f", dollars)
}

// ParseAmount converts a dollar amount string to cents
func ParseAmount(amount string) (int64, error) {
	// Remove non-numeric characters except decimal point using strings.Builder
	var builder strings.Builder
	builder.Grow(len(amount)) // Pre-allocate capacity
	hasDecimal := false
	for _, r := range amount {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		} else if r == '.' && !hasDecimal {
			hasDecimal = true
			builder.WriteRune(r)
		}
	}
	cleaned := builder.String()

	// Parse as float and convert to cents
	var cents int64
	if hasDecimal {
		var dollars, fractional int
		parts := strings.Split(cleaned, ".")
		fmt.Sscanf(parts[0], "%d", &dollars)
		if len(parts) > 1 && len(parts[1]) > 0 {
			// Ensure we handle fractional cents correctly
			fracStr := parts[1]
			if len(fracStr) == 1 {
				fracStr += "0" // .5 becomes .50
			} else if len(fracStr) > 2 {
				fracStr = fracStr[:2] // Truncate to 2 decimal places
			}
			fmt.Sscanf(fracStr, "%d", &fractional)
		}
		cents = int64(dollars*100 + fractional)
	} else {
		// No decimal, assume whole dollars
		var dollars int
		fmt.Sscanf(cleaned, "%d", &dollars)
		cents = int64(dollars * 100)
	}

	return cents, nil
}

// FormatTIN formats a TIN with dashes (for display only)
func FormatTIN(tin string, tinType string) string {
	// Remove non-numeric characters using strings.Builder
	var builder strings.Builder
	builder.Grow(9) // TINs are 9 digits max
	for _, r := range tin {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	cleaned := builder.String()

	if len(cleaned) != 9 {
		return tin // Return as-is if not 9 digits
	}

	switch tinType {
	case "1": // SSN
		return fmt.Sprintf("%s-%s-%s", cleaned[:3], cleaned[3:5], cleaned[5:])
	case "2": // EIN
		return fmt.Sprintf("%s-%s", cleaned[:2], cleaned[2:])
	default:
		return cleaned
	}
}

// CleanAddress removes special characters that might cause issues
func CleanAddress(addr string) string {
	// Replace problematic characters
	replacements := map[rune]rune{
		'"': '\'',
		'<': '(',
		'>': ')',
		'&': '+',
	}

	// Use strings.Builder for efficient string building
	var builder strings.Builder
	builder.Grow(len(addr)) // Pre-allocate capacity
	for _, r := range addr {
		if replacement, ok := replacements[r]; ok {
			builder.WriteRune(replacement)
		} else if r >= 0x20 && r <= 0x7E { // Printable ASCII
			builder.WriteRune(r)
		} else {
			builder.WriteRune(' ') // Replace non-printable with space
		}
	}
	result := builder.String()

	// Trim and collapse multiple spaces
	result = strings.TrimSpace(result)
	for strings.Contains(result, "  ") {
		result = strings.ReplaceAll(result, "  ", " ")
	}

	return result
}

// FileBuilder provides a fluent interface for building PAM SPR files
type FileBuilder struct {
	file            *File
	currentSchedule Schedule
	errors          []error
}

// NewFileBuilder creates a new file builder
func NewFileBuilder() *FileBuilder {
	return &FileBuilder{
		file: &File{
			Schedules: make([]Schedule, 0),
		},
		errors: make([]error, 0),
	}
}

// WithHeader sets the file header
func (fb *FileBuilder) WithHeader(inputSystem, version string, sameDayACH bool) *FileBuilder {
	fb.file.Header = &FileHeader{
		RecordCode:             "H ",
		InputSystem:            inputSystem,
		StandardPaymentVersion: version,
	}

	if sameDayACH {
		fb.file.Header.IsRequestedForSameDayACH = "1"
	} else {
		fb.file.Header.IsRequestedForSameDayACH = "0"
	}

	return fb
}

// StartACHSchedule starts building an ACH schedule
func (fb *FileBuilder) StartACHSchedule(scheduleNum, paymentType, alc, secCode string) *FileBuilder {
	// Save current schedule if exists
	if fb.currentSchedule != nil {
		fb.file.Schedules = append(fb.file.Schedules, fb.currentSchedule)
	}

	schedule := &ACHSchedule{
		Header: &ACHScheduleHeader{
			RecordCode:             "01",
			ScheduleNumber:         scheduleNum,
			PaymentTypeCode:        paymentType,
			AgencyLocationCode:     alc,
			StandardEntryClassCode: secCode,
		},
		BaseSchedule: BaseSchedule{
			ScheduleNumber: scheduleNum,
			PaymentType:    paymentType,
			ALC:            alc,
			Payments:       make([]Payment, 0),
		},
	}

	fb.currentSchedule = schedule
	return fb
}

// StartCheckSchedule starts building a check schedule
func (fb *FileBuilder) StartCheckSchedule(scheduleNum, paymentType, alc, enclosureCode string) *FileBuilder {
	// Save current schedule if exists
	if fb.currentSchedule != nil {
		fb.file.Schedules = append(fb.file.Schedules, fb.currentSchedule)
	}

	schedule := &CheckSchedule{
		Header: &CheckScheduleHeader{
			RecordCode:                "11",
			ScheduleNumber:            scheduleNum,
			PaymentTypeCode:           paymentType,
			AgencyLocationCode:        alc,
			CheckPaymentEnclosureCode: enclosureCode,
		},
		BaseSchedule: BaseSchedule{
			ScheduleNumber: scheduleNum,
			PaymentType:    paymentType,
			ALC:            alc,
			Payments:       make([]Payment, 0),
		},
	}

	fb.currentSchedule = schedule
	return fb
}

// AddACHPayment adds an ACH payment to the current schedule
func (fb *FileBuilder) AddACHPayment(payment *ACHPayment) *FileBuilder {
	if fb.currentSchedule == nil {
		fb.errors = append(fb.errors, fmt.Errorf("no schedule started"))
		return fb
	}

	if achSchedule, ok := fb.currentSchedule.(*ACHSchedule); ok {
		payment.RecordCode = "02"
		achSchedule.Payments = append(achSchedule.Payments, payment)
	} else {
		fb.errors = append(fb.errors, fmt.Errorf("current schedule is not an ACH schedule"))
	}

	return fb
}

// AddCheckPayment adds a check payment to the current schedule
func (fb *FileBuilder) AddCheckPayment(payment *CheckPayment) *FileBuilder {
	if fb.currentSchedule == nil {
		fb.errors = append(fb.errors, fmt.Errorf("no schedule started"))
		return fb
	}

	if checkSchedule, ok := fb.currentSchedule.(*CheckSchedule); ok {
		payment.RecordCode = "12"
		checkSchedule.Payments = append(checkSchedule.Payments, payment)
	} else {
		fb.errors = append(fb.errors, fmt.Errorf("current schedule is not a check schedule"))
	}

	return fb
}

// Build finalizes the file and calculates totals
func (fb *FileBuilder) Build() (*File, error) {
	if len(fb.errors) > 0 {
		return nil, fmt.Errorf("build errors: %v", fb.errors)
	}

	// Add last schedule
	if fb.currentSchedule != nil {
		fb.file.Schedules = append(fb.file.Schedules, fb.currentSchedule)
	}

	// Calculate schedule trailers
	for _, schedule := range fb.file.Schedules {
		switch s := schedule.(type) {
		case *ACHSchedule:
			fb.calculateScheduleTrailer(&s.BaseSchedule, s.Payments)
			s.Trailer = s.BaseSchedule.Trailer
		case *CheckSchedule:
			fb.calculateScheduleTrailer(&s.BaseSchedule, s.Payments)
			s.Trailer = s.BaseSchedule.Trailer
		}
	}

	// Calculate file trailer
	fb.calculateFileTrailer()

	return fb.file, nil
}

func (fb *FileBuilder) calculateScheduleTrailer(base *BaseSchedule, payments []Payment) {
	count := int64(len(payments))
	amount := int64(0)

	for _, payment := range payments {
		amount += payment.GetAmount()
	}

	base.Trailer = &ScheduleTrailer{
		RecordCode:     "T ",
		ScheduleCount:  count,
		ScheduleAmount: amount,
	}
}

func (fb *FileBuilder) calculateFileTrailer() {
	totalRecords := int64(2) // Header + Trailer
	totalPayments := int64(0)
	totalAmount := int64(0)

	for _, schedule := range fb.file.Schedules {
		totalRecords += 2 // Schedule header + trailer

		switch s := schedule.(type) {
		case *ACHSchedule:
			for _, payment := range s.Payments {
				totalRecords++
				totalPayments++
				if achPay, ok := payment.(*ACHPayment); ok {
					totalAmount += achPay.Amount
					totalRecords += int64(len(achPay.Addenda))
					totalRecords += int64(len(achPay.CARSTASBETC))
					if achPay.DNP != nil {
						totalRecords++
					}
				}
			}
		case *CheckSchedule:
			for _, payment := range s.Payments {
				totalRecords++
				totalPayments++
				if checkPay, ok := payment.(*CheckPayment); ok {
					totalAmount += checkPay.Amount
					if checkPay.Stub != nil {
						totalRecords++
					}
					totalRecords += int64(len(checkPay.CARSTASBETC))
					if checkPay.DNP != nil {
						totalRecords++
					}
				}
			}
		}
	}

	fb.file.Trailer = &FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   totalRecords,
		TotalCountPayments:  totalPayments,
		TotalAmountPayments: totalAmount,
	}
}

// AgencyReconcilementParser provides parsing for agency-specific reconcilement fields
type AgencyReconcilementParser struct{}

// ParseIRSReconcilement parses IRS-specific reconcilement data
func (arp *AgencyReconcilementParser) ParseIRSReconcilement(recon string, depositorAccount string) map[string]string {
	result := make(map[string]string)

	if len(recon) != 100 {
		return result
	}

	if depositorAccount == "BONDS" {
		// IRS Savings Bonds format
		result["TaxPeriodYear"] = strings.TrimSpace(recon[0:4])
		result["TaxPeriodMonth"] = strings.TrimSpace(recon[4:6])
		result["MFTCode"] = strings.TrimSpace(recon[6:8])
		result["ServiceCenterCode"] = strings.TrimSpace(recon[8:10])
		result["DistrictOfficeCode"] = strings.TrimSpace(recon[10:12])
		result["FileTINCode"] = strings.TrimSpace(recon[12:13])
		result["NameControl"] = strings.TrimSpace(recon[13:17])
		result["PlanReportNumber"] = strings.TrimSpace(recon[17:20])
		result["SplitRefundCode"] = strings.TrimSpace(recon[20:21])
		result["InjuredSpouseCode"] = strings.TrimSpace(recon[21:22])
		result["DebtBypassIndicator"] = strings.TrimSpace(recon[22:23])
		result["BondName1"] = strings.TrimSpace(recon[23:56])
		result["BondREGCode"] = strings.TrimSpace(recon[56:57])
		result["BondName2"] = strings.TrimSpace(recon[57:90])
	} else {
		// Standard IRS format
		result["TaxPeriodYear"] = strings.TrimSpace(recon[0:4])
		result["TaxPeriodMonth"] = strings.TrimSpace(recon[4:6])
		result["MFTCode"] = strings.TrimSpace(recon[6:8])
		result["ServiceCenterCode"] = strings.TrimSpace(recon[8:10])
		result["DistrictOfficeCode"] = strings.TrimSpace(recon[10:12])
		result["FileTINCode"] = strings.TrimSpace(recon[12:13])
		result["NameControl"] = strings.TrimSpace(recon[13:17])
		result["PlanReportNumber"] = strings.TrimSpace(recon[17:20])
		result["SplitRefundCode"] = strings.TrimSpace(recon[20:21])
		result["InjuredSpouseCode"] = strings.TrimSpace(recon[21:22])
		result["DebtBypassIndicator"] = strings.TrimSpace(recon[22:23])
		result["DocumentLocatorNumber"] = strings.TrimSpace(recon[23:37])
		result["CheckDetailEnclosureCode"] = strings.TrimSpace(recon[37:47])
	}

	return result
}

// ParseVAReconcilement parses VA-specific reconcilement data
func (arp *AgencyReconcilementParser) ParseVAReconcilement(recon string, isACH bool) map[string]string {
	result := make(map[string]string)

	if len(recon) != 100 {
		return result
	}

	result["StationCode"] = strings.TrimSpace(recon[0:2])
	result["FinCode"] = strings.TrimSpace(recon[2:4])

	if isACH {
		// ACH format (no courtesy code)
		result["AppropCode"] = strings.TrimSpace(recon[4:5])
		result["AddressSeqCode"] = strings.TrimSpace(recon[5:6])
		result["PolicyPreCode"] = strings.TrimSpace(recon[6:8])
		result["PolicyNum"] = strings.TrimSpace(recon[8:10])
		result["PayPeriodInfo"] = strings.TrimSpace(recon[10:22])
	} else {
		// Check format
		result["CourtesyCode"] = strings.TrimSpace(recon[4:5])
		result["AppropCode"] = strings.TrimSpace(recon[5:6])
		result["AddressSeqCode"] = strings.TrimSpace(recon[6:7])
		result["PolicyPreCode"] = strings.TrimSpace(recon[7:9])
		result["PolicyNum"] = strings.TrimSpace(recon[9:11])
		result["PayPeriodInfo"] = strings.TrimSpace(recon[11:23])
		result["NameCode"] = strings.TrimSpace(recon[23:26])
	}

	return result
}

// ParseSSAReconcilement parses SSA-specific reconcilement data
func (arp *AgencyReconcilementParser) ParseSSAReconcilement(recon string, ruleID string) map[string]string {
	result := make(map[string]string)

	if len(recon) != 100 {
		return result
	}

	result["ProgramServiceCenterCode"] = strings.TrimSpace(recon[0:1])
	result["PaymentIDCode"] = strings.TrimSpace(recon[1:3])

	if ruleID != "SSA-A" {
		// SSA and SSA-Daily have TIN Indicator Offset
		result["TINIndicatorOffset"] = strings.TrimSpace(recon[3:4])
	}

	return result
}

// ParseRRBReconcilement parses RRB-specific reconcilement data
// RRB reconcilement structure (100 characters):
// - Beneficiary Symbol (positions 1-2): 2-character alphanumeric field
// - Prefix Code (position 3): 1-character alphanumeric field
// - Payee Code (position 4): 1-character alphanumeric field
// - Object Code (position 5): 1-character alphanumeric field (mapped to PACER)
// - Filler (positions 6-100): 95 characters of filler
func (arp *AgencyReconcilementParser) ParseRRBReconcilement(recon string) map[string]string {
	result := make(map[string]string)

	if len(recon) != 100 {
		return result
	}

	result["BeneficiarySymbol"] = strings.TrimSpace(recon[0:2])
	result["PrefixCode"] = strings.TrimSpace(recon[2:3])
	result["PayeeCode"] = strings.TrimSpace(recon[3:4])
	result["ObjectCode"] = strings.TrimSpace(recon[4:5])
	// Filler at positions 5-99 (95 characters) - not included in result

	return result
}

// ParseCCCReconcilement parses CCC-specific reconcilement data
// CCC reconcilement structure (100 characters):
// - TOP Payment Agency ID (positions 1-2): 2-character alphabetic field
// - TOP Agency Site ID (positions 3-4): 2-character alphabetic field
// - Filler (positions 5-100): 96 characters of filler
// Business rule: If one payment within a schedule contains these values,
// the entire schedule will be sent to TOP with these values.
func (arp *AgencyReconcilementParser) ParseCCCReconcilement(recon string) map[string]string {
	result := make(map[string]string)

	if len(recon) != 100 {
		return result
	}

	result["TOPPaymentAgencyID"] = strings.TrimSpace(recon[0:2])
	result["TOPAgencySiteID"] = strings.TrimSpace(recon[2:4])
	// Filler at positions 4-99 (96 characters) - not included in result

	return result
}
