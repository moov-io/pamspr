package pamspr

// Field Length Constants
const (
	// Core field lengths defined by SPR specification
	TINLength              = 9   // Tax Identification Number
	ReconcilementLength    = 100 // Reconcilement field for all payment types
	RoutingNumberLength    = 9   // ACH routing number
	AccountNumberMaxLength = 17  // Maximum account number length
	PaymentIDLength        = 20  // Payment identifier
	PayeeNameMaxLength     = 35  // Maximum payee name length
	AddressLineMaxLength   = 35  // Maximum address line length
	CityNameMaxLength      = 27  // Maximum city name length
	StateNameMaxLength     = 10  // Maximum state name length
	StateCodeLength        = 2   // State code (e.g., "CA", "NY")
	PostalCodeLength       = 5   // Base postal code length
	PostalCodeExtLength    = 5   // Postal code extension length
	CountryCodeLength      = 2   // Country code length
	TransactionCodeLength  = 2   // ACH transaction code length
	ScheduleNumberLength   = 14  // Schedule number field length
)

// Amount and Validation Constants
const (
	// Same Day ACH limits per Federal Reserve regulations
	MaxSDAAmountCents = 100000000 // $1,000,000 in cents (Same Day ACH limit)

	// Minimum valid amounts
	MinPaymentAmount = 1 // Minimum 1 cent for payments

	// Amount field lengths in SPR format
	AmountFieldLength      = 10 // Standard amount field length in characters
	LargeAmountFieldLength = 15 // Large amount field for file totals
	SmallAmountFieldLength = 8  // Small amount field for counts
)

// Record Structure Constants
const (
	// RecordLength = 850 is already defined in file.go
	RecordCodeLength = 2 // Length of record type code
	FillerLength     = 0 // Base filler length (varies by record type)
)

// Agency-Specific Field Lengths
const (
	// VA (Veterans Affairs) specific field lengths
	VAStationCodeLength   = 2  // VA station code
	VAFinCodeLength       = 2  // VA financial code
	VACourtesyCodeLength  = 1  // VA courtesy code
	VAAppropCodeLength    = 1  // VA appropriation code
	VAPolicyNumLength     = 2  // VA policy number
	VAPayPeriodInfoLength = 12 // VA pay period information
	VANameCodeLength      = 3  // VA name code

	// SSA (Social Security Administration) specific field lengths
	SSAProgramServiceCenterCodeLength = 1 // SSA program service center code
	SSAPaymentIDCodeLength            = 2 // SSA payment ID code
	SSATINIndicatorOffsetLength       = 1 // SSA TIN indicator offset

	// RRB (Railroad Retirement Board) specific field lengths
	RRBBeneficiarySymbolLength = 2 // RRB beneficiary symbol
	RRBPrefixCodeLength        = 1 // RRB prefix code
	RRBPayeeCodeLength         = 1 // RRB payee code
	RRBObjectCodeLength        = 1 // RRB object code

	// CCC (Commodity Credit Corporation) specific field lengths
	CCCTOPPaymentAgencyIDLength = 2 // CCC TOP payment agency ID
	CCCTOPAgencySiteIDLength    = 2 // CCC TOP agency site ID
)

// Valid ACH Transaction Codes
// Reference: NACHA Operating Rules
var ValidACHTransactionCodes = map[string]bool{
	"22": true, // Credit (deposit) to checking account
	"23": true, // Pre-note for credit to checking account
	"24": true, // Zero dollar amount to checking account
	"32": true, // Credit (deposit) to savings account
	"33": true, // Pre-note for credit to savings account
	"34": true, // Zero dollar amount to savings account
	"42": true, // Credit (deposit) to general ledger account
	"43": true, // Pre-note for credit to general ledger account
	"52": true, // Credit (deposit) to loan account
	"53": true, // Pre-note for credit to loan account
}

// Valid Payment Types per SPR specification
var ValidPaymentTypes = map[string]bool{
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
}

// Valid Standard Entry Class Codes for ACH
var ValidSECCodes = map[string]bool{
	"CCD": true, // Corporate Credit/Debit
	"PPD": true, // Prearranged Payment & Deposit
	"IAT": true, // International ACH Transaction
	"CTX": true, // Corporate Trade Exchange
}

// Valid TIN Indicator Values
var ValidTINIndicators = map[string]bool{
	"1": true, // SSN (Social Security Number)
	"2": true, // EIN (Employer Identification Number)
	"3": true, // ITIN (Individual Taxpayer Identification Number)
}

// Valid TOP Offset Indicators
var ValidTOPOffsetIndicators = map[string]bool{
	"0": true, // Not subject to offset
	"1": true, // Subject to offset
}

// Valid Check Payment Enclosure Codes
var ValidCheckEnclosureCodes = map[string]bool{
	"nameonly": true, // Name only
	"letter":   true, // Letter
	"stub":     true, // Stub
	"insert":   true, // Insert
	"":         true, // Blank (no enclosure)
}

// Same Day ACH Processing Time Windows (in hours)
const (
	SDASubmissionDeadlineHour = 14 // 2:00 PM ET submission deadline
	SDASettlementHour         = 17 // 5:00 PM ET settlement
)

// Validation Rule Constants
const (
	MaxSchedulesPerFile    = 999    // Maximum number of schedules per file
	MaxPaymentsPerSchedule = 999999 // Maximum payments per schedule

	// File size limits
	MaxFileSizeBytes  = 100 * 1024 * 1024               // 100MB maximum file size
	MaxRecordsPerFile = MaxFileSizeBytes / RecordLength // Approximately 123,294 records
)

// Error Message Templates
const (
	ErrMsgFieldRequired = "%s is required"
	ErrMsgFieldLength   = "%s must be exactly %d characters, got %d"
	ErrMsgFieldFormat   = "%s must %s"
	ErrMsgAmountLimit   = "amount %d exceeds maximum limit of %d cents"
	ErrMsgInvalidValue  = "%s contains invalid value: %s"
)

// Version Constants
const (
	CurrentSPRVersion = "502" // Current published SPR version per Treasury specification
)

// Same Day ACH Constants
const (
	SDAFlagEnabled  = "1"
	SDAFlagDisabled = "0"
)

// Routing Number Validation Constants
const (
	RoutingNumberModulus = 10   // Check digit modulus for routing number validation
	MinHexCharacter      = 0x40 // Minimum valid hex character value
)

// Payment Identification Constants
const (
	MaxPaymentIdentificationLines = 14 // Maximum payment ID lines for check stubs
)

// Country Code Constants
const (
	CountryCodeEmpty = "00" // Empty/invalid country code
)

// CTX Constants
const (
	CTXEDISegmentIdentifier = "ISA" // EDI segment identifier for CTX addenda
	CTXEDISegmentLength     = 3     // Length of EDI segment identifier
)
