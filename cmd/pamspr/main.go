package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
	var (
		validate = flag.Bool("validate", false, "Validate a PAM SPR file")
		info     = flag.Bool("info", false, "Display file information")
		convert  = flag.Bool("convert", false, "Convert file to JSON")
		create   = flag.String("create", "", "Create a sample file (ach or check)")
		input    = flag.String("input", "", "Input file path")
		output   = flag.String("output", "", "Output file path")
	)

	flag.Parse()

	switch {
	case *validate:
		if *input == "" {
			log.Fatal("Input file required for validation")
		}
		validateFile(*input)

	case *info:
		if *input == "" {
			log.Fatal("Input file required")
		}
		displayFileInfo(*input)

	case *convert:
		if *input == "" || *output == "" {
			log.Fatal("Both input and output files required for conversion")
		}
		convertToJSON(*input, *output)

	case *create != "":
		if *output == "" {
			log.Fatal("Output file required")
		}
		createSampleFile(*create, *output)

	default:
		flag.Usage()
	}
}

func validateFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := pamspr.NewReader(file)
	pamFile, err := reader.Read()
	if err != nil {
		log.Fatalf("Validation failed: %v", err)
	}

	// Additional validation
	validator := pamspr.NewValidator()

	// Validate file structure
	if err := validator.ValidateFileStructure(pamFile); err != nil {
		log.Fatalf("File structure validation failed: %v", err)
	}

	// Validate balancing
	if err := validator.ValidateBalancing(pamFile); err != nil {
		log.Fatalf("Balancing validation failed: %v", err)
	}

	// Validate Same Day ACH if applicable
	if err := validator.ValidateSameDayACH(pamFile); err != nil {
		log.Fatalf("Same Day ACH validation failed: %v", err)
	}

	fmt.Println("✓ File validation passed")
	fmt.Printf("  Schedules: %d\n", len(pamFile.Schedules))
	fmt.Printf("  Total Payments: %d\n", pamFile.Trailer.TotalCountPayments)
	fmt.Printf("  Total Amount: $%.2f\n", float64(pamFile.Trailer.TotalAmountPayments)/100)
}

func displayFileInfo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	reader := pamspr.NewReader(file)
	pamFile, err := reader.Read()
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println("PAM SPR File Information")
	fmt.Println("========================")
	fmt.Printf("Input System: %s\n", pamFile.Header.InputSystem)
	fmt.Printf("Version: %s\n", pamFile.Header.StandardPaymentVersion)
	fmt.Printf("Same Day ACH: %s\n", pamFile.Header.IsRequestedForSameDayACH)
	fmt.Printf("\nSchedules: %d\n", len(pamFile.Schedules))

	for i, schedule := range pamFile.Schedules {
		fmt.Printf("\nSchedule %d:\n", i+1)
		fmt.Printf("  Number: %s\n", schedule.GetScheduleNumber())

		switch s := schedule.(type) {
		case *pamspr.ACHSchedule:
			fmt.Printf("  Type: ACH\n")
			fmt.Printf("  SEC Code: %s\n", s.Header.StandardEntryClassCode)
			fmt.Printf("  ALC: %s\n", s.Header.AgencyLocationCode)
			fmt.Printf("  Payment Type: %s\n", s.Header.PaymentTypeCode)
		case *pamspr.CheckSchedule:
			fmt.Printf("  Type: Check\n")
			fmt.Printf("  Enclosure: %s\n", s.Header.CheckPaymentEnclosureCode)
			fmt.Printf("  ALC: %s\n", s.Header.AgencyLocationCode)
			fmt.Printf("  Payment Type: %s\n", s.Header.PaymentTypeCode)
		}

		payments := schedule.GetPayments()
		fmt.Printf("  Payments: %d\n", len(payments))

		// Show first 3 payments
		max := 3
		if len(payments) < max {
			max = len(payments)
		}

		for j := 0; j < max; j++ {
			payment := payments[j]
			fmt.Printf("    Payment %d: ID=%s, Amount=$%.2f, Payee=%s\n",
				j+1,
				payment.GetPaymentID(),
				float64(payment.GetAmount())/100,
				payment.GetPayeeName())
		}

		if len(payments) > 3 {
			fmt.Printf("    ... and %d more payments\n", len(payments)-3)
		}
	}

	fmt.Printf("\nFile Totals:\n")
	fmt.Printf("  Records: %d\n", pamFile.Trailer.TotalCountRecords)
	fmt.Printf("  Payments: %d\n", pamFile.Trailer.TotalCountPayments)
	fmt.Printf("  Amount: $%.2f\n", float64(pamFile.Trailer.TotalAmountPayments)/100)
}

func convertToJSON(inputFile, outputFile string) {
	// This would convert the PAM SPR file to JSON format
	// Implementation would use encoding/json package
	fmt.Printf("Converting %s to JSON format...\n", inputFile)
	fmt.Printf("Output will be saved to %s\n", outputFile)
	// TODO: Implement JSON conversion
}

func createSampleFile(fileType, outputFile string) {
	var file *pamspr.File
	var err error

	switch fileType {
	case "ach":
		file = createSampleACHFile()
	case "check":
		file = createSampleCheckFile()
	default:
		log.Fatalf("Unknown file type: %s (use 'ach' or 'check')", fileType)
	}

	// Write file
	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer output.Close()

	writer := pamspr.NewWriter(output)
	if err := writer.Write(file); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("Sample %s file created: %s\n", fileType, outputFile)
}

func createSampleACHFile() *pamspr.File {
	builder := pamspr.NewFileBuilder()

	file, err := builder.
		WithHeader("SAMPLE_AGENCY", "502", false).
		StartACHSchedule("TEST001", "Salary", "12345678", "PPD").
		AddACHPayment(&pamspr.ACHPayment{
			AgencyAccountIdentifier:      "EMP001",
			Amount:                       150000, // $1,500.00
			AgencyPaymentTypeCode:        "S",
			IsTOPOffset:                  "1",
			PayeeName:                    "JOHN SAMPLE",
			PayeeAddressLine1:            "123 TEST ST",
			CityName:                     "WASHINGTON",
			StateCodeText:                "DC",
			PostalCode:                   "20001",
			RoutingNumber:                "021000021",
			AccountNumber:                "1234567890",
			ACHTransactionCode:           "22",
			PaymentID:                    "SAMP001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "1",
		}).
		AddACHPayment(&pamspr.ACHPayment{
			AgencyAccountIdentifier:      "EMP002",
			Amount:                       200000, // $2,000.00
			AgencyPaymentTypeCode:        "S",
			IsTOPOffset:                  "1",
			PayeeName:                    "JANE EXAMPLE",
			PayeeAddressLine1:            "456 DEMO AVE",
			CityName:                     "ARLINGTON",
			StateCodeText:                "VA",
			PostalCode:                   "22201",
			RoutingNumber:                "021000021",
			AccountNumber:                "0987654321",
			ACHTransactionCode:           "32",
			PaymentID:                    "SAMP002",
			TIN:                          "987654321",
			PaymentRecipientTINIndicator: "1",
		}).
		Build()

	if err != nil {
		log.Fatalf("Error building file: %v", err)
	}

	return file
}

func createSampleCheckFile() *pamspr.File {
	builder := pamspr.NewFileBuilder()

	file, err := builder.
		WithHeader("SAMPLE_AGENCY", "502", false).
		StartCheckSchedule("CHK001", "Vendor", "87654321", "stub").
		AddCheckPayment(&pamspr.CheckPayment{
			AgencyAccountIdentifier:      "VEND001",
			Amount:                       100000, // $1,000.00
			AgencyPaymentTypeCode:        "V",
			IsTOPOffset:                  "1",
			PayeeName:                    "SAMPLE COMPANY INC",
			PayeeAddressLine1:            "789 BUSINESS BLVD",
			PayeeAddressLine2:            "SUITE 200",
			CityName:                     "NEW YORK",
			StateCodeText:                "NY",
			PostalCode:                   "10001",
			CheckLegendText1:             "INVOICE #SAMPLE123",
			CheckLegendText2:             "PO #DEMO456",
			PaymentID:                    "CHK001",
			TIN:                          "123456789",
			PaymentRecipientTINIndicator: "2",
			Stub: &pamspr.CheckStub{
				RecordCode: "13",
				PaymentID:  "CHK001",
				PaymentIdentificationLines: [14]string{
					"Sample Check Stub",
					"==================",
					"Invoice: SAMPLE123",
					"PO Number: DEMO456",
					"",
					"Description: Sample Services",
					"Amount: $1,000.00",
					"",
					"This is a sample check stub",
					"for demonstration purposes",
					"",
					"Thank you for your business",
					"",
					"",
				},
			},
		}).
		Build()

	if err != nil {
		log.Fatalf("Error building file: %v", err)
	}

	return file
}
