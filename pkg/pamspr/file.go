// Package pamspr provides functionality for reading, writing, and validating
// Payment Automation Manager (PAM) Standard Payment Request (SPR) files
package pamspr

// RecordType represents the type of record in the file
type RecordType string

const (
	RecordTypeFileHeader          RecordType = "H "
	RecordTypeACHScheduleHeader   RecordType = "01"
	RecordTypeACHPayment          RecordType = "02"
	RecordTypeACHAddendum         RecordType = "03"
	RecordTypeACHAddendumCTX      RecordType = "04"
	RecordTypeCheckScheduleHeader RecordType = "11"
	RecordTypeCheckPayment        RecordType = "12"
	RecordTypeCheckStub           RecordType = "13"
	RecordTypeCARSTASBETC         RecordType = "G "
	RecordTypeDNP                 RecordType = "DD"
	RecordTypeScheduleTrailer     RecordType = "T "
	RecordTypeFileTrailer         RecordType = "E "
)

// RecordLength is the fixed length of each record
const RecordLength = 850

// File represents a complete PAM SPR file
type File struct {
	Header    *FileHeader
	Schedules []Schedule
	Trailer   *FileTrailer
}

// Schedule represents either an ACH or Check schedule
type Schedule interface {
	GetScheduleNumber() string
	GetPaymentType() PaymentType
	GetPayments() []Payment
	Validate() error
}

// Payment represents either an ACH or Check payment
type Payment interface {
	GetPaymentID() string
	GetAmount() int64 // Amount in cents
	GetPayeeName() string
	Validate() error
}

// BaseSchedule contains common schedule fields
type BaseSchedule struct {
	ScheduleNumber string
	PaymentType    string
	ALC            string // Agency Location Code
	Payments       []Payment
	Trailer        *ScheduleTrailer
}

// ACHSchedule represents an ACH payment schedule
type ACHSchedule struct {
	BaseSchedule
	Header *ACHScheduleHeader
}

// GetScheduleNumber returns the schedule number
func (s *ACHSchedule) GetScheduleNumber() string {
	return s.ScheduleNumber
}

// GetPaymentType returns the payment type
func (s *ACHSchedule) GetPaymentType() PaymentType {
	return PaymentTypeACH
}

// GetPayments returns the payments in the schedule
func (s *ACHSchedule) GetPayments() []Payment {
	return s.Payments
}

// Validate validates the ACH schedule
func (s *ACHSchedule) Validate() error {
	// TODO: Implement validation
	return nil
}

// CheckSchedule represents a Check payment schedule
type CheckSchedule struct {
	BaseSchedule
	Header *CheckScheduleHeader
}

// GetScheduleNumber returns the schedule number
func (s *CheckSchedule) GetScheduleNumber() string {
	return s.ScheduleNumber
}

// GetPaymentType returns the payment type
func (s *CheckSchedule) GetPaymentType() PaymentType {
	return PaymentTypeCheck
}

// GetPayments returns the payments in the schedule
func (s *CheckSchedule) GetPayments() []Payment {
	return s.Payments
}

// Validate validates the check schedule
func (s *CheckSchedule) Validate() error {
	// TODO: Implement validation
	return nil
}

// FileHeader represents the file header record
type FileHeader struct {
	RecordCode               string // "H "
	InputSystem              string // 40 chars
	StandardPaymentVersion   string // "502"
	IsRequestedForSameDayACH string // "0" or "1"
}

// FileTrailer represents the file trailer record
type FileTrailer struct {
	RecordCode          string // "E "
	TotalCountRecords   int64  // 18 digits
	TotalCountPayments  int64  // 18 digits
	TotalAmountPayments int64  // 18 digits, amount in cents
}

// ACHScheduleHeader represents ACH schedule header record
type ACHScheduleHeader struct {
	RecordCode              string // "01"
	AgencyACHText           string // 4 chars
	ScheduleNumber          string // 14 chars
	PaymentTypeCode         string // 25 chars
	StandardEntryClassCode  string // 3 chars: CCD, PPD, IAT, CTX
	AgencyLocationCode      string // 8 digits
	FederalEmployerIDNumber string // 10 chars
}

// CheckScheduleHeader represents check schedule header record
type CheckScheduleHeader struct {
	RecordCode                string // "11"
	ScheduleNumber            string // 14 chars
	PaymentTypeCode           string // 25 chars
	AgencyLocationCode        string // 8 digits
	CheckPaymentEnclosureCode string // 10 chars: "nameonly", "letter", "stub", "insert", or blank
}

// ACHPayment represents an ACH payment data record
type ACHPayment struct {
	RecordCode                   string // "02"
	AgencyAccountIdentifier      string // 16 chars
	Amount                       int64  // 10 digits, amount in cents
	AgencyPaymentTypeCode        string // 1 char
	IsTOPOffset                  string // "0" or "1"
	PayeeName                    string // 35 chars
	PayeeAddressLine1            string // 35 chars
	PayeeAddressLine2            string // 35 chars
	CityName                     string // 27 chars
	StateName                    string // 10 chars
	StateCodeText                string // 2 chars
	PostalCode                   string // 5 chars
	PostalCodeExtension          string // 5 chars
	CountryCodeText              string // 2 chars
	RoutingNumber                string // 9 digits
	AccountNumber                string // 17 chars
	ACHTransactionCode           string // 2 digits
	PayeeIdentifierAdditional    string // 9 chars (Secondary TIN)
	PayeeNameAdditional          string // 35 chars (Secondary Name)
	PaymentID                    string // 20 chars
	Reconcilement                string // 100 chars
	TIN                          string // 9 chars
	PaymentRecipientTINIndicator string // 1 char: "1"=SSN, "2"=EIN, "3"=ITIN
	AdditionalPayeeTINIndicator  string // 1 char
	AmountEligibleForOffset      string // 10 digits
	PayeeAddressLine3            string // 35 chars
	PayeeAddressLine4            string // 35 chars
	CountryName                  string // 40 chars
	ConsularCode                 string // 3 chars (Geo Code)
	SubPaymentTypeCode           string // 32 chars
	PayerMechanism               string // 20 chars
	PaymentDescriptionCode       string // 2 chars

	// Associated records
	Addenda     []*ACHAddendum
	CARSTASBETC []*CARSTASBETC
	DNP         *DNPRecord

	// StandardEntryClassCode is set from the schedule header
	StandardEntryClassCode string
}

// GetPaymentID returns the payment ID
func (p *ACHPayment) GetPaymentID() string {
	return p.PaymentID
}

// GetAmount returns the amount in cents
func (p *ACHPayment) GetAmount() int64 {
	return p.Amount
}

// GetPayeeName returns the payee name
func (p *ACHPayment) GetPayeeName() string {
	return p.PayeeName
}

// Validate validates the ACH payment
func (p *ACHPayment) Validate() error {
	// TODO: Implement validation
	return nil
}

// CheckPayment represents a check payment data record
type CheckPayment struct {
	RecordCode                   string // "12"
	AgencyAccountIdentifier      string // 16 chars
	Amount                       int64  // 10 digits, amount in cents
	AgencyPaymentTypeCode        string // 1 char
	IsTOPOffset                  string // "0" or "1"
	PayeeName                    string // 35 chars
	PayeeAddressLine1            string // 35 chars
	PayeeAddressLine2            string // 35 chars
	PayeeAddressLine3            string // 35 chars
	PayeeAddressLine4            string // 35 chars
	CityName                     string // 27 chars
	StateName                    string // 10 chars
	StateCodeText                string // 2 chars
	PostalCode                   string // 5 chars
	PostalCodeExtension          string // 5 chars
	PostNetBarcodeDeliveryPoint  string // 3 chars
	CountryName                  string // 40 chars
	ConsularCode                 string // 3 chars (Geo Code)
	CheckLegendText1             string // 55 chars
	CheckLegendText2             string // 55 chars
	PayeeIdentifierSecondary     string // 9 chars
	PartyNameSecondary           string // 35 chars
	PaymentID                    string // 20 chars
	Reconcilement                string // 100 chars
	SpecialHandling              string // 50 chars
	TIN                          string // 9 chars
	USPSIntelligentMailBarcode   string // 50 chars
	PaymentRecipientTINIndicator string // 1 char
	SecondaryPayeeTINIndicator   string // 1 char
	AmountEligibleForOffset      string // 10 digits
	SubPaymentTypeCode           string // 32 chars
	PayerMechanism               string // 20 chars
	PaymentDescriptionCode       string // 2 chars

	// Associated records
	Stub        *CheckStub
	CARSTASBETC []*CARSTASBETC
	DNP         *DNPRecord
}

// GetPaymentID returns the payment ID
func (p *CheckPayment) GetPaymentID() string {
	return p.PaymentID
}

// GetAmount returns the amount in cents
func (p *CheckPayment) GetAmount() int64 {
	return p.Amount
}

// GetPayeeName returns the payee name
func (p *CheckPayment) GetPayeeName() string {
	return p.PayeeName
}

// Validate validates the check payment
func (p *CheckPayment) Validate() error {
	// TODO: Implement validation
	return nil
}

// ACHAddendum represents an ACH addendum record
type ACHAddendum struct {
	RecordCode         string // "03" or "04"
	PaymentID          string // 20 chars
	AddendaInformation string // 80 chars for "03", 800 chars for "04" (CTX)
}

// CARSTASBETC represents a CARS TAS/BETC record
type CARSTASBETC struct {
	RecordCode                    string // "G "
	PaymentID                     string // 20 chars
	SubLevelPrefixCode            string // 2 chars
	AllocationTransferAgencyID    string // 3 chars
	AgencyIdentifier              string // 3 chars
	BeginningPeriodOfAvailability string // 4 chars
	EndingPeriodOfAvailability    string // 4 chars
	AvailabilityTypeCode          string // 1 char
	MainAccountCode               string // 4 chars
	SubAccountCode                string // 3 chars
	BusinessEventTypeCode         string // 8 chars
	AccountClassificationAmount   int64  // 10 digits
	IsCredit                      string // 1 char: "0" or "1"
}

// CheckStub represents a check stub record
type CheckStub struct {
	RecordCode                 string     // "13"
	PaymentID                  string     // 20 chars
	PaymentIdentificationLines [14]string // 14 lines of 55 chars each
}

// DNPRecord represents a DNP record
type DNPRecord struct {
	RecordCode string // "DD"
	PaymentID  string // 20 chars
	DNPDetail  string // 766 chars
}

// ScheduleTrailer represents a schedule trailer record
type ScheduleTrailer struct {
	RecordCode     string // "T "
	ScheduleCount  int64  // 8 digits
	ScheduleAmount int64  // 15 digits, amount in cents
}

// PaymentType represents the type of payment
type PaymentType int

const (
	PaymentTypeUnknown PaymentType = iota
	PaymentTypeACH
	PaymentTypeCheck
)

// StandardEntryClassCode represents ACH SEC codes
type StandardEntryClassCode string

const (
	SECCodeCCD StandardEntryClassCode = "CCD" // Corporate Credit/Debit
	SECCodePPD StandardEntryClassCode = "PPD" // Prearranged Payment & Deposit
	SECCodeIAT StandardEntryClassCode = "IAT" // International ACH Transaction
	SECCodeCTX StandardEntryClassCode = "CTX" // Corporate Trade Exchange
)
