package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/pamspr/pkg/pamspr"
)

// RecordBuilder helps build fixed-width records
type RecordBuilder struct {
	data      []byte
	recordLen int
}

// NewRecordBuilder creates a new record builder
func NewRecordBuilder() *RecordBuilder {
	return &RecordBuilder{
		data:      make([]byte, pamspr.RecordLength),
		recordLen: pamspr.RecordLength,
	}
}

// SetField sets a field value at the correct position with proper padding
func (rb *RecordBuilder) SetField(recordCode, fieldName, value string) error {
	fieldDefs := pamspr.GetFieldDefinitions(recordCode)
	if fieldDefs == nil {
		return fmt.Errorf("unknown record code: %s", recordCode)
	}

	fieldDef, ok := fieldDefs[fieldName]
	if !ok {
		return fmt.Errorf("unknown field %s for record code %s", fieldName, recordCode)
	}

	// Initialize with spaces
	for i := 0; i < rb.recordLen; i++ {
		if rb.data[i] == 0 {
			rb.data[i] = ' '
		}
	}

	// Truncate or pad the value to fit the field
	padded := padField(value, fieldDef.Length)

	// Copy the padded value to the correct position (convert to 0-based indexing)
	start := fieldDef.Start - 1
	copy(rb.data[start:start+fieldDef.Length], []byte(padded))

	return nil
}

// String returns the complete record as a string
func (rb *RecordBuilder) String() string {
	// Ensure the entire record is filled with spaces if not already set
	for i := 0; i < rb.recordLen; i++ {
		if rb.data[i] == 0 {
			rb.data[i] = ' '
		}
	}
	return string(rb.data)
}

// padField pads a field value to the specified length
func padField(value string, length int) string {
	if len(value) >= length {
		return value[:length]
	}

	// For numeric fields (all digits), pad with zeros on the left
	if isNumeric(value) {
		return fmt.Sprintf("%0*s", length, value)
	}

	// For text fields, pad with spaces on the right
	return fmt.Sprintf("%-*s", length, value)
}

// isNumeric checks if a string contains only digits
func isNumeric(s string) bool {
	if s == "" {
		return true // Empty string is considered numeric (will be padded with zeros)
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// generateFileHeader creates a properly formatted file header
func generateFileHeader() string {
	rb := NewRecordBuilder()

	rb.SetField("H ", "RecordCode", "H ")
	rb.SetField("H ", "InputSystem", "SYNTHETIC_TEST_FILE_2025011001")
	rb.SetField("H ", "StandardPaymentVersion", "502")
	rb.SetField("H ", "IsRequestedForSameDayACH", "0")

	return rb.String()
}

// generateACHScheduleHeader creates a properly formatted ACH schedule header
func generateACHScheduleHeader(scheduleNum string) string {
	rb := NewRecordBuilder()

	rb.SetField("01", "RecordCode", "01")
	rb.SetField("01", "AgencyACHText", " ACH")
	rb.SetField("01", "ScheduleNumber", scheduleNum)
	rb.SetField("01", "PaymentTypeCode", "")
	rb.SetField("01", "StandardEntryClassCode", "PPD")
	rb.SetField("01", "AgencyLocationCode", "12345678")
	rb.SetField("01", "FederalEmployerIDNumber", "1234567890")

	return rb.String()
}

// generateACHPayment creates a properly formatted ACH payment record
func generateACHPayment(accountID, amount, paymentID string) string {
	rb := NewRecordBuilder()

	rb.SetField("02", "RecordCode", "02")
	rb.SetField("02", "AgencyAccountIdentifier", accountID)
	rb.SetField("02", "Amount", amount)
	rb.SetField("02", "AgencyPaymentTypeCode", "1")
	rb.SetField("02", "IsTOPOffset", " ")
	rb.SetField("02", "PayeeName", "JOHN DOE")
	rb.SetField("02", "PayeeAddressLine1", "123 MAIN STREET")
	rb.SetField("02", "PayeeAddressLine2", "APT 4B")
	rb.SetField("02", "CityName", "ANYTOWN")
	rb.SetField("02", "StateName", "CALIFORNIA")
	rb.SetField("02", "StateCodeText", "CA")
	rb.SetField("02", "PostalCode", "12345")
	rb.SetField("02", "PostalCodeExtension", "")
	rb.SetField("02", "CountryCodeText", "US")
	rb.SetField("02", "RoutingNumber", "123456789")
	rb.SetField("02", "AccountNumber", "12345678901234567")
	rb.SetField("02", "ACHTransactionCode", "22")
	rb.SetField("02", "PayeeIdentifierAdditional", "")
	rb.SetField("02", "PayeeNameAdditional", "JANE DOE")
	rb.SetField("02", "PaymentID", paymentID)
	rb.SetField("02", "Reconcilement", "SYNTHETIC TEST RECONCILEMENT DATA FOR ACH PAYMENT")
	rb.SetField("02", "TIN", "123456789")
	rb.SetField("02", "PaymentRecipientTINIndicator", "1")
	rb.SetField("02", "AdditionalPayeeTINIndicator", " ")
	rb.SetField("02", "AmountEligibleForOffset", amount)

	return rb.String()
}

// generateCheckScheduleHeader creates a properly formatted check schedule header
func generateCheckScheduleHeader(scheduleNum string) string {
	rb := NewRecordBuilder()

	rb.SetField("11", "RecordCode", "11")
	rb.SetField("11", "ScheduleNumber", scheduleNum)
	rb.SetField("11", "PaymentTypeCode", "")
	rb.SetField("11", "AgencyLocationCode", "12345678")
	rb.SetField("11", "CheckPaymentEnclosureCode", "")

	return rb.String()
}

// generateCheckPayment creates a properly formatted check payment record
func generateCheckPayment(accountID, amount, paymentID string) string {
	rb := NewRecordBuilder()

	rb.SetField("12", "RecordCode", "12")
	rb.SetField("12", "AgencyAccountIdentifier", accountID)
	rb.SetField("12", "Amount", amount)
	rb.SetField("12", "AgencyPaymentTypeCode", "1")
	rb.SetField("12", "IsTOPOffset", " ")
	rb.SetField("12", "PayeeName", "JANE SMITH")
	rb.SetField("12", "PayeeAddressLine1", "456 OAK AVENUE")
	rb.SetField("12", "PayeeAddressLine2", "")
	rb.SetField("12", "PayeeAddressLine3", "")
	rb.SetField("12", "PayeeAddressLine4", "")
	rb.SetField("12", "CityName", "ANOTHER CITY")
	rb.SetField("12", "StateName", "TEXAS")
	rb.SetField("12", "StateCodeText", "TX")
	rb.SetField("12", "PostalCode", "75001")
	rb.SetField("12", "PostalCodeExtension", "")
	rb.SetField("12", "PostNetBarcodeDeliveryPoint", "")
	rb.SetField("12", "CountryName", "UNITED STATES")
	rb.SetField("12", "ConsularCode", "")
	rb.SetField("12", "CheckLegendText1", "PAY TO THE ORDER OF")
	rb.SetField("12", "CheckLegendText2", "DOLLARS")
	rb.SetField("12", "PayeeIdentifierSecondary", "")
	rb.SetField("12", "PartyNameSecondary", "")
	rb.SetField("12", "PaymentID", paymentID)
	rb.SetField("12", "Reconcilement", "SYNTHETIC TEST RECONCILEMENT DATA FOR CHECK PAYMENT")
	rb.SetField("12", "SpecialHandling", "")
	rb.SetField("12", "TIN", "987654321")
	rb.SetField("12", "USPSIntelligentMailBarcode", "")
	rb.SetField("12", "PaymentRecipientTINIndicator", "1")
	rb.SetField("12", "SecondaryPayeeTINIndicator", " ")
	rb.SetField("12", "AmountEligibleForOffset", amount)

	return rb.String()
}

// generateScheduleTrailer creates a properly formatted schedule trailer
func generateScheduleTrailer(scheduleCount, scheduleAmount string) string {
	rb := NewRecordBuilder()

	rb.SetField("T ", "RecordCode", "T ")
	rb.SetField("T ", "ScheduleCount", scheduleCount)
	rb.SetField("T ", "ScheduleAmount", scheduleAmount)

	return rb.String()
}

// generateFileTrailer creates a properly formatted file trailer
func generateFileTrailer(totalRecords, totalPayments, totalAmount string) string {
	rb := NewRecordBuilder()

	rb.SetField("E ", "RecordCode", "E ")
	rb.SetField("E ", "TotalCountRecords", totalRecords)
	rb.SetField("E ", "TotalCountPayments", totalPayments)
	rb.SetField("E ", "TotalAmountPayments", totalAmount)

	return rb.String()
}

// writeFile writes content to a file
func writeFile(filename, content string) error {
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(filename, []byte(content), 0644)
}

func main() {
	baseDir := "./testdata/synthetic"

	// Generate simple ACH file
	fmt.Println("Generating simple ACH file...")
	achContent := []string{
		generateFileHeader(),
		generateACHScheduleHeader("00000000000001"),
		generateACHPayment("SYNTH00000000001", "0000001000", "PAYMENT000000000001"),
		generateScheduleTrailer("00000001", "000000000001000"),
		generateFileTrailer("000000000000000005", "000000000000000001", "000000000000001000"),
	}

	err := writeFile(filepath.Join(baseDir, "valid/synthetic_ach_simple.spr"), strings.Join(achContent, "\n")+"\n")
	if err != nil {
		fmt.Printf("Error writing ACH file: %v\n", err)
		return
	}

	// Generate simple check file
	fmt.Println("Generating simple check file...")
	checkContent := []string{
		generateFileHeader(),
		generateCheckScheduleHeader("00000000000002"),
		generateCheckPayment("SYNTH00000000002", "0000002500", "PAYMENT000000000002"),
		generateScheduleTrailer("00000001", "000000000002500"),
		generateFileTrailer("000000000000000005", "000000000000000001", "000000000000002500"),
	}

	err = writeFile(filepath.Join(baseDir, "valid/synthetic_check_simple.spr"), strings.Join(checkContent, "\n")+"\n")
	if err != nil {
		fmt.Printf("Error writing check file: %v\n", err)
		return
	}

	// Generate multi-schedule file
	fmt.Println("Generating multi-schedule file...")
	multiContent := []string{
		generateFileHeader(),
		// ACH Schedule
		generateACHScheduleHeader("00000000000001"),
		generateACHPayment("SYNTH00000000001", "0000001000", "PAYMENT000000000001"),
		generateACHPayment("SYNTH00000000002", "0000001500", "PAYMENT000000000002"),
		generateScheduleTrailer("00000002", "000000000002500"),
		// Check Schedule
		generateCheckScheduleHeader("00000000000002"),
		generateCheckPayment("SYNTH00000000003", "0000002000", "PAYMENT000000000003"),
		generateScheduleTrailer("00000001", "000000000002000"),
		// File trailer
		generateFileTrailer("000000000000000009", "000000000000000003", "000000000000004500"),
	}

	err = writeFile(filepath.Join(baseDir, "valid/synthetic_multi_schedule.spr"), strings.Join(multiContent, "\n")+"\n")
	if err != nil {
		fmt.Printf("Error writing multi-schedule file: %v\n", err)
		return
	}

	// Generate agency-specific files
	agencyFiles := map[string]struct {
		dir           string
		filename      string
		payeeName     string
		reconcilement string
	}{
		"IRS": {
			dir:           "agency/IRS",
			filename:      "synthetic_irs_refund.spr",
			payeeName:     "TAXPAYER REFUND RECIPIENT",
			reconcilement: "IRS TAX REFUND PAYMENT 2024",
		},
		"SSA": {
			dir:           "agency/SSA",
			filename:      "synthetic_ssa_benefit.spr",
			payeeName:     "SOCIAL SECURITY BENEFICIARY",
			reconcilement: "SSA RETIREMENT BENEFIT PAYMENT",
		},
		"VA": {
			dir:           "agency/VA",
			filename:      "synthetic_va_benefit.spr",
			payeeName:     "VETERAN BENEFIT RECIPIENT",
			reconcilement: "VA DISABILITY COMPENSATION PAYMENT",
		},
		"RRB": {
			dir:           "agency/RRB",
			filename:      "synthetic_rrb_annuity.spr",
			payeeName:     "RAILROAD RETIREMENT ANNUITANT",
			reconcilement: "RRB RETIREMENT ANNUITY PAYMENT",
		},
		"CCC": {
			dir:           "agency/CCC",
			filename:      "synthetic_ccc_payment.spr",
			payeeName:     "COMMODITY PROGRAM PARTICIPANT",
			reconcilement: "CCC COMMODITY SUPPORT PAYMENT",
		},
	}

	for agency, info := range agencyFiles {
		fmt.Printf("Generating %s agency file...\n", agency)

		// Create custom ACH payment for the agency
		rb := NewRecordBuilder()
		rb.SetField("02", "RecordCode", "02")
		rb.SetField("02", "AgencyAccountIdentifier", fmt.Sprintf("SYNTH%s0000001", agency))
		rb.SetField("02", "Amount", "0000003000") // $30.00
		rb.SetField("02", "AgencyPaymentTypeCode", "1")
		rb.SetField("02", "IsTOPOffset", " ")
		rb.SetField("02", "PayeeName", info.payeeName)
		rb.SetField("02", "PayeeAddressLine1", "789 GOVERNMENT PLAZA")
		rb.SetField("02", "PayeeAddressLine2", "")
		rb.SetField("02", "CityName", "WASHINGTON")
		rb.SetField("02", "StateName", "DISTRICT OF COLUMBIA")
		rb.SetField("02", "StateCodeText", "DC")
		rb.SetField("02", "PostalCode", "20001")
		rb.SetField("02", "PostalCodeExtension", "")
		rb.SetField("02", "CountryCodeText", "US")
		rb.SetField("02", "RoutingNumber", "123456789")
		rb.SetField("02", "AccountNumber", "98765432109876543")
		rb.SetField("02", "ACHTransactionCode", "22")
		rb.SetField("02", "PayeeIdentifierAdditional", "")
		rb.SetField("02", "PayeeNameAdditional", "")
		rb.SetField("02", "PaymentID", fmt.Sprintf("PAYMENT%s00000001", agency))
		rb.SetField("02", "Reconcilement", info.reconcilement)
		rb.SetField("02", "TIN", "111223333")
		rb.SetField("02", "PaymentRecipientTINIndicator", "1")
		rb.SetField("02", "AdditionalPayeeTINIndicator", " ")
		rb.SetField("02", "AmountEligibleForOffset", "0000003000")

		agencyContent := []string{
			generateFileHeader(),
			generateACHScheduleHeader("00000000000001"),
			rb.String(),
			generateScheduleTrailer("00000001", "000000000003000"),
			generateFileTrailer("000000000000000005", "000000000000000001", "000000000000003000"),
		}

		err = writeFile(filepath.Join(baseDir, info.dir, info.filename), strings.Join(agencyContent, "\n")+"\n")
		if err != nil {
			fmt.Printf("Error writing %s file: %v\n", agency, err)
			return
		}
	}

	fmt.Println("All synthetic SPR files generated successfully!")
	fmt.Println("Each record is exactly 850 characters long.")

	// Verify one file to show the format
	fmt.Println("\nVerifying format of generated files...")

	// Read and check the ACH file
	content, err := os.ReadFile(filepath.Join(baseDir, "valid/synthetic_ach_simple.spr"))
	if err != nil {
		fmt.Printf("Error reading verification file: %v\n", err)
		return
	}

	lines := strings.Split(strings.TrimRight(string(content), "\n"), "\n")
	fmt.Printf("Generated %d records:\n", len(lines))

	for i, line := range lines {
		length := len(line)
		fmt.Printf("Record %d: %d characters (expected: 850) - %s\n", i+1, length,
			func() string {
				if length == 850 {
					return "✓ CORRECT"
				}
				return "✗ INCORRECT"
			}())

		if length != 850 {
			fmt.Printf("  Content: %q\n", line)
		}
	}
}
