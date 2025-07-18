// Package pamspr provides functionality for reading, writing, and validating
// Payment Automation Manager (PAM) Standard Payment Request (SPR) files
//
// Design Philosophy:
// - Direct field access is preferred over getter/setter methods
// - Interfaces only provide essential polymorphic behavior
// - Use type assertions for type-specific fields: payment.(*ACHPayment).RoutingNumber
// - Minimal interface design with only required methods
package pamspr

import "fmt"

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
	GetTrailer() *ScheduleTrailer
	SetTrailer(*ScheduleTrailer)
	Validate() error
}

// Removed ACHScheduleAccessor and CheckScheduleAccessor interfaces
// Use direct type assertions to access type-specific fields

// Payment represents either an ACH or Check payment
type Payment interface {
	GetPaymentID() string
	GetAmount() int64 // Amount in cents
	SetAmount(int64)
	GetPayeeName() string
	GetRecordCode() string
	GetPaymentType() PaymentType
	GetReconcilement() string
	Validate() error
}

// Removed ACHPaymentAccessor and CheckPaymentAccessor interfaces
// Use direct type assertions to access type-specific fields

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

// GetTrailer returns the schedule trailer
func (s *ACHSchedule) GetTrailer() *ScheduleTrailer {
	return s.Trailer
}

// SetTrailer sets the schedule trailer
func (s *ACHSchedule) SetTrailer(trailer *ScheduleTrailer) {
	s.Trailer = trailer
}

// Access s.Header directly - no getter methods provided for type-specific fields

// Validate validates the ACH schedule
func (s *ACHSchedule) Validate() error {
	if s.Header == nil {
		return fmt.Errorf("ACH schedule header is required")
	}
	if len(s.Payments) == 0 {
		return fmt.Errorf("ACH schedule must have at least one payment")
	}
	for i, payment := range s.Payments {
		if err := payment.Validate(); err != nil {
			return fmt.Errorf("payment %d validation failed: %w", i, err)
		}
	}
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

// GetTrailer returns the schedule trailer
func (s *CheckSchedule) GetTrailer() *ScheduleTrailer {
	return s.Trailer
}

// SetTrailer sets the schedule trailer
func (s *CheckSchedule) SetTrailer(trailer *ScheduleTrailer) {
	s.Trailer = trailer
}

// Access s.Header directly - no getter methods provided for type-specific fields

// Validate validates the check schedule
func (s *CheckSchedule) Validate() error {
	if s.Header == nil {
		return fmt.Errorf("check schedule header is required")
	}
	if len(s.Payments) == 0 {
		return fmt.Errorf("check schedule must have at least one payment")
	}
	for i, payment := range s.Payments {
		if err := payment.Validate(); err != nil {
			return fmt.Errorf("payment %d validation failed: %w", i, err)
		}
	}
	return nil
}

// FileHeader represents the file header record
type FileHeader struct {
	RecordCode               string `pamspr:"RecordCode"`
	InputSystem              string `pamspr:"InputSystem"`
	StandardPaymentVersion   string `pamspr:"StandardPaymentVersion"`
	IsRequestedForSameDayACH string `pamspr:"IsRequestedForSameDayACH"`
	Filler                   string `pamspr:"Filler"`
}

// FileTrailer represents the file trailer record
type FileTrailer struct {
	RecordCode          string `pamspr:"RecordCode"`
	TotalCountRecords   int64  `pamspr:"TotalCountRecords" format:"numeric"`
	TotalCountPayments  int64  `pamspr:"TotalCountPayments" format:"numeric"`
	TotalAmountPayments int64  `pamspr:"TotalAmountPayments" format:"numeric"`
}

// ACHScheduleHeader represents ACH schedule header record
type ACHScheduleHeader struct {
	RecordCode              string `pamspr:"RecordCode"`
	AgencyACHText           string `pamspr:"AgencyACHText"`
	ScheduleNumber          string `pamspr:"ScheduleNumber" format:"numeric"`
	PaymentTypeCode         string `pamspr:"PaymentTypeCode"`
	StandardEntryClassCode  string `pamspr:"StandardEntryClassCode"`
	AgencyLocationCode      string `pamspr:"AgencyLocationCode" format:"numeric"`
	Filler1                 string `pamspr:"Filler1"`
	FederalEmployerIDNumber string `pamspr:"FederalEmployerIDNumber"`
	Filler2                 string `pamspr:"Filler2"`
}

// CheckScheduleHeader represents check schedule header record
type CheckScheduleHeader struct {
	RecordCode                string // "11"
	ScheduleNumber            string // 14 chars
	PaymentTypeCode           string // 25 chars
	AgencyLocationCode        string // 8 digits
	Filler1                   string // 9 chars
	CheckPaymentEnclosureCode string // 10 chars: "nameonly", "letter", "stub", "insert", or blank
	Filler2                   string // 782 chars
}

// ACHPayment represents an ACH payment data record
type ACHPayment struct {
	RecordCode                   string // "02"
	AgencyAccountIdentifier      string // 16 chars
	Amount                       int64  // 10 digits, amount in cents
	AgencyPaymentTypeCode        string // 1 char
	IsTOP_Offset                 string // "0" or "1"
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
	ACH_TransactionCode          string // 2 digits
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
	Filler                       string // 284 chars

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

// SetAmount sets the amount in cents
func (p *ACHPayment) SetAmount(amount int64) {
	p.Amount = amount
}

// GetPayeeName returns the payee name
func (p *ACHPayment) GetPayeeName() string {
	return p.PayeeName
}

// GetRecordCode returns the record code
func (p *ACHPayment) GetRecordCode() string {
	return p.RecordCode
}

// GetPaymentType returns the payment type
func (p *ACHPayment) GetPaymentType() PaymentType {
	return PaymentTypeACH
}

// GetReconcilement returns the reconcilement field
func (p *ACHPayment) GetReconcilement() string {
	return p.Reconcilement
}

// Use direct field access for type-specific fields:
// - p.RoutingNumber, p.AccountNumber, p.StandardEntryClassCode
// - p.Addenda, p.CARSTASBETC, p.DNP
//
// Helper methods for common operations:

// AddAddendum appends an addendum to the payment
func (p *ACHPayment) AddAddendum(addendum *ACHAddendum) {
	p.Addenda = append(p.Addenda, addendum)
}

// AddCARSTASBETC appends a CARS/TAS/BETC record to the payment
func (p *ACHPayment) AddCARSTASBETC(car *CARSTASBETC) {
	p.CARSTASBETC = append(p.CARSTASBETC, car)
}

// Validate validates the ACH payment
func (p *ACHPayment) Validate() error {
	if p.Amount <= 0 {
		return fmt.Errorf("ACH payment amount must be positive, got %d", p.Amount)
	}
	if p.PayeeName == "" {
		return fmt.Errorf("ACH payment must have a payee name")
	}
	if p.PaymentID == "" {
		return fmt.Errorf("ACH payment must have a payment ID")
	}
	return nil
}

// CheckPayment represents a check payment data record
type CheckPayment struct {
	RecordCode                   string // "12"
	AgencyAccountIdentifier      string // 16 chars
	Amount                       int64  // 10 digits, amount in cents
	AgencyPaymentTypeCode        string // 1 char
	IsTOP_Offset                 string // "0" or "1"
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
	Filler1                      string // 14 chars
	CountryName                  string // 40 chars
	ConsularCode                 string // 3 chars (Geo Code)
	CheckLegendText1             string // 55 chars
	CheckLegendText2             string // 55 chars
	PayeeIdentifier_Secondary    string // 9 chars
	PartyName_Secondary          string // 35 chars
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
	Filler2                      string // 87 chars

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

// SetAmount sets the amount in cents
func (p *CheckPayment) SetAmount(amount int64) {
	p.Amount = amount
}

// GetPayeeName returns the payee name
func (p *CheckPayment) GetPayeeName() string {
	return p.PayeeName
}

// GetRecordCode returns the record code
func (p *CheckPayment) GetRecordCode() string {
	return p.RecordCode
}

// GetPaymentType returns the payment type
func (p *CheckPayment) GetPaymentType() PaymentType {
	return PaymentTypeCheck
}

// GetReconcilement returns the reconcilement field
func (p *CheckPayment) GetReconcilement() string {
	return p.Reconcilement
}

// Use direct field access for type-specific fields:
// - p.Stub, p.CARSTASBETC, p.DNP
//
// Helper methods for common operations:

// AddCARSTASBETC appends a CARS/TAS/BETC record to the payment
func (p *CheckPayment) AddCARSTASBETC(car *CARSTASBETC) {
	p.CARSTASBETC = append(p.CARSTASBETC, car)
}

// Validate validates the check payment
func (p *CheckPayment) Validate() error {
	if p.Amount <= 0 {
		return fmt.Errorf("check payment amount must be positive, got %d", p.Amount)
	}
	if p.PayeeName == "" {
		return fmt.Errorf("check payment must have a payee name")
	}
	if p.PaymentID == "" {
		return fmt.Errorf("check payment must have a payment ID")
	}
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

// Type assertion helpers for cleaner code
// Direct type assertions are preferred: payment.(*ACHPayment) or schedule.(*ACHSchedule)
// These helpers provide convenience but direct assertions are more idiomatic

// AsACHPayment safely converts a Payment to *ACHPayment
func AsACHPayment(payment Payment) (*ACHPayment, bool) {
	if achPayment, ok := payment.(*ACHPayment); ok {
		return achPayment, true
	}
	return nil, false
}

// AsCheckPayment safely converts a Payment to *CheckPayment
func AsCheckPayment(payment Payment) (*CheckPayment, bool) {
	if checkPayment, ok := payment.(*CheckPayment); ok {
		return checkPayment, true
	}
	return nil, false
}

// AsACHSchedule safely converts a Schedule to *ACHSchedule
func AsACHSchedule(schedule Schedule) (*ACHSchedule, bool) {
	if achSchedule, ok := schedule.(*ACHSchedule); ok {
		return achSchedule, true
	}
	return nil, false
}

// AsCheckSchedule safely converts a Schedule to *CheckSchedule
func AsCheckSchedule(schedule Schedule) (*CheckSchedule, bool) {
	if checkSchedule, ok := schedule.(*CheckSchedule); ok {
		return checkSchedule, true
	}
	return nil, false
}

// StandardEntryClassCode represents ACH SEC codes
type StandardEntryClassCode string

const (
	SECCodeCCD StandardEntryClassCode = "CCD" // Corporate Credit/Debit
	SECCodePPD StandardEntryClassCode = "PPD" // Prearranged Payment & Deposit
	SECCodeIAT StandardEntryClassCode = "IAT" // International ACH Transaction
	SECCodeCTX StandardEntryClassCode = "CTX" // Corporate Trade Exchange
)
