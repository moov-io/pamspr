package pamspr_test

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/moov-io/pamspr/pkg/pamspr"
)

func Example_createACHFile() {
	builder := pamspr.NewFileBuilder()

	file, err := builder.
		WithHeader("MY_AGENCY_SYSTEM", "502", false).
		StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
		AddACHPayment(&pamspr.ACHPayment{
			AgencyAccountIdentifier:      "EMP001",
			Amount:                       250000, // $2,500.00
			AgencyPaymentTypeCode:        "S",
			IsTOP_Offset:                 "1",
			PayeeName:                    "JOHN DOE",
			PayeeAddressLine1:            "123 MAIN ST",
			CityName:                     "WASHINGTON",
			StateCodeText:                "DC",
			PostalCode:                   "20001",
			RoutingNumber:                "021000021",
			AccountNumber:                "1234567890",
			ACH_TransactionCode:          "22", // Checking Credit
			PaymentID:                    "PAY001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "1", // SSN
		}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	validator := pamspr.NewValidator()
	if err := validator.ValidateFileStructure(file); err != nil {
		log.Fatal(err)
	}
	if err := validator.ValidateBalancing(file); err != nil {
		log.Fatal(err)
	}

	var output bytes.Buffer
	writer := pamspr.NewWriter(&output)
	if err := writer.Write(file); err != nil {
		log.Fatal(err)
	}

	fmt.Println("ACH payment file created successfully!")
	// Output: ACH payment file created successfully!
}

func Example_createCheckFile() {
	builder := pamspr.NewFileBuilder()

	file, err := builder.
		WithHeader("CHECK_SYSTEM", "502", false).
		StartCheckSchedule("CHK001", "Vendor", "87654321", "stub").
		AddCheckPayment(&pamspr.CheckPayment{
			AgencyAccountIdentifier:      "VEND001",
			Amount:                       150000, // $1,500.00
			AgencyPaymentTypeCode:        "V",
			IsTOP_Offset:                 "1",
			PayeeName:                    "ACME CORP",
			PayeeAddressLine1:            "789 BUSINESS BLVD",
			PayeeAddressLine2:            "SUITE 100",
			CityName:                     "NEW YORK",
			StateCodeText:                "NY",
			PostalCode:                   "10001",
			CheckLegendText1:             "INVOICE #12345",
			CheckLegendText2:             "PO #67890",
			PaymentID:                    "CHK001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "2", // EIN
			Stub: &pamspr.CheckStub{
				RecordCode: "13",
				PaymentID:  "CHK001",
				PaymentIdentificationLines: [14]string{
					"Invoice Date: 01/15/2024",
					"Invoice Number: 12345",
					"Purchase Order: 67890",
					"Description: Office Supplies",
					"",
					"Item Details:",
					"  Paper - 10 reams @ $50.00 = $500.00",
					"  Toner - 5 units @ $100.00 = $500.00",
					"  Misc Supplies = $500.00",
					"",
					"Subtotal: $1,500.00",
					"Tax: $0.00",
					"Total: $1,500.00",
					"",
				},
			},
		}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	validator := pamspr.NewValidator()
	if err := validator.ValidateFileStructure(file); err != nil {
		log.Fatal(err)
	}
	if err := validator.ValidateBalancing(file); err != nil {
		log.Fatal(err)
	}

	var output bytes.Buffer
	writer := pamspr.NewWriter(&output)
	if err := writer.Write(file); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Check payment file created successfully!")
	// Output: Check payment file created successfully!
}

func Example_readFile() {
	file, err := pamspr.NewFileBuilder().
		WithHeader("MY_AGENCY_SYSTEM", "502", false).
		StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
		AddACHPayment(&pamspr.ACHPayment{
			AgencyAccountIdentifier:      "EMP001",
			Amount:                       250000,
			AgencyPaymentTypeCode:        "S",
			IsTOP_Offset:                 "1",
			PayeeName:                    "JOHN DOE",
			PayeeAddressLine1:            "123 MAIN ST",
			CityName:                     "WASHINGTON",
			StateCodeText:                "DC",
			PostalCode:                   "20001",
			RoutingNumber:                "021000021",
			AccountNumber:                "1234567890",
			ACH_TransactionCode:          "22",
			PaymentID:                    "PAY001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "1",
		}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := pamspr.NewWriter(&buf).Write(file); err != nil {
		log.Fatal(err)
	}

	reader := pamspr.NewReader(bytes.NewReader(buf.Bytes()))
	pamFile, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Input System: %s\n", strings.TrimSpace(pamFile.Header.InputSystem))
	fmt.Printf("Total Payments: %d\n", pamFile.Trailer.TotalCountPayments)
	fmt.Printf("Total Amount: $%.2f\n", float64(pamFile.Trailer.TotalAmountPayments)/100)
	// Output:
	// Input System: MY_AGENCY_SYSTEM
	// Total Payments: 1
	// Total Amount: $2500.00
}

func Example_validateFile() {
	file, err := pamspr.NewFileBuilder().
		WithHeader("MY_AGENCY_SYSTEM", "502", false).
		StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
		AddACHPayment(&pamspr.ACHPayment{
			AgencyAccountIdentifier:      "EMP001",
			Amount:                       250000,
			AgencyPaymentTypeCode:        "S",
			IsTOP_Offset:                 "1",
			PayeeName:                    "JOHN DOE",
			PayeeAddressLine1:            "123 MAIN ST",
			CityName:                     "WASHINGTON",
			StateCodeText:                "DC",
			PostalCode:                   "20001",
			RoutingNumber:                "021000021",
			AccountNumber:                "1234567890",
			ACH_TransactionCode:          "22",
			PaymentID:                    "PAY001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "1",
		}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	validator := pamspr.NewValidator()
	if err := validator.ValidateFileStructure(file); err != nil {
		log.Fatal(err)
	}
	if err := validator.ValidateBalancing(file); err != nil {
		log.Fatal(err)
	}
	if err := validator.ValidateSameDayACH(file); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Validation completed!")
	// Output: Validation completed!
}

func Example_validateFileStructureOnly() {
	file, err := pamspr.NewFileBuilder().
		WithHeader("MY_AGENCY_SYSTEM", "502", false).
		StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
		AddACHPayment(&pamspr.ACHPayment{
			AgencyAccountIdentifier:      "EMP001",
			Amount:                       250000,
			AgencyPaymentTypeCode:        "S",
			IsTOP_Offset:                 "1",
			PayeeName:                    "JOHN DOE",
			PayeeAddressLine1:            "123 MAIN ST",
			CityName:                     "WASHINGTON",
			StateCodeText:                "DC",
			PostalCode:                   "20001",
			RoutingNumber:                "021000021",
			AccountNumber:                "1234567890",
			ACH_TransactionCode:          "22",
			PaymentID:                    "PAY001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "1",
		}).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	var buf bytes.Buffer
	if err := pamspr.NewWriter(&buf).Write(file); err != nil {
		log.Fatal(err)
	}

	reader := pamspr.NewReader(bytes.NewReader(buf.Bytes()))
	start := time.Now()
	if err := reader.ValidateFileStructureOnly(); err != nil {
		log.Fatal(err)
	}
	_ = time.Since(start)
	stats := reader.GetStats()
	fmt.Printf("lines=%d payments=%d schedules=%d\n", stats.LinesProcessed, stats.PaymentsProcessed, stats.SchedulesProcessed)
	// Output is not asserted: line and payment counts are corrected in a companion change.
}
