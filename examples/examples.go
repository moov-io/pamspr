package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moov-io/pamspr/pkg/pamspr"
)

// Example 1: Create an ACH payment file
func CreateACHFile() {
	// Create file structure
	file := &pamspr.File{
		Header: &pamspr.FileHeader{
			RecordCode:               "H ",
			InputSystem:              "AGENCY_SYSTEM_001",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: make([]pamspr.Schedule, 0),
	}

	// Create ACH schedule
	achSchedule := &pamspr.ACHSchedule{
		Header: &pamspr.ACHScheduleHeader{
			RecordCode:              "01",
			AgencyACHText:           "AGCY",
			ScheduleNumber:          "SCH001",
			PaymentTypeCode:         "Salary",
			StandardEntryClassCode:  "PPD",
			AgencyLocationCode:      "12345678",
			FederalEmployerIDNumber: "123456789",
		},
		BaseSchedule: pamspr.BaseSchedule{
			ScheduleNumber: "SCH001",
			PaymentType:    "Salary",
			ALC:            "12345678",
			Payments:       make([]pamspr.Payment, 0),
		},
	}

	// Add ACH payments
	payment1 := &pamspr.ACHPayment{
		RecordCode:                   "02",
		AgencyAccountIdentifier:      "EMP001",
		Amount:                       250000, // $2,500.00
		AgencyPaymentTypeCode:        "S",
		IsTOPOffset:                  "1",
		PayeeName:                    "JOHN DOE",
		PayeeAddressLine1:            "123 MAIN ST",
		PayeeAddressLine2:            "APT 4B",
		CityName:                     "WASHINGTON",
		StateCodeText:                "DC",
		PostalCode:                   "20001",
		RoutingNumber:                "021000021",
		AccountNumber:                "1234567890",
		ACHTransactionCode:           "22", // Checking credit
		PaymentID:                    "PAY001",
		TIN:                          "123456789",
		PaymentRecipientTINIndicator: "1", // SSN
	}

	payment2 := &pamspr.ACHPayment{
		RecordCode:                   "02",
		AgencyAccountIdentifier:      "EMP002",
		Amount:                       350000, // $3,500.00
		AgencyPaymentTypeCode:        "S",
		IsTOPOffset:                  "1",
		PayeeName:                    "JANE SMITH",
		PayeeAddressLine1:            "456 OAK AVE",
		CityName:                     "ARLINGTON",
		StateCodeText:                "VA",
		PostalCode:                   "22201",
		RoutingNumber:                "021000021",
		AccountNumber:                "0987654321",
		ACHTransactionCode:           "32", // Savings credit
		PaymentID:                    "PAY002",
		TIN:                          "987654321",
		PaymentRecipientTINIndicator: "1", // SSN
	}

	// Add CARS TAS/BETC records
	payment1.CARSTASBETC = append(payment1.CARSTASBETC, &pamspr.CARSTASBETC{
		RecordCode:                  "G ",
		PaymentID:                   "PAY001",
		AgencyIdentifier:            "012",
		MainAccountCode:             "1234",
		SubAccountCode:              "001",
		BusinessEventTypeCode:       "SALARY01",
		AccountClassificationAmount: 250000,
		IsCredit:                    "0",
	})

	achSchedule.Payments = append(achSchedule.Payments, payment1, payment2)

	// Calculate schedule totals
	scheduleCount := int64(len(achSchedule.Payments))
	scheduleAmount := int64(0)
	for _, p := range achSchedule.Payments {
		if achPay, ok := p.(*pamspr.ACHPayment); ok {
			scheduleAmount += achPay.Amount
		}
	}

	achSchedule.Trailer = &pamspr.ScheduleTrailer{
		RecordCode:     "T ",
		ScheduleCount:  scheduleCount,
		ScheduleAmount: scheduleAmount,
	}

	file.Schedules = append(file.Schedules, achSchedule)

	// Calculate file totals
	totalRecords := int64(2) // Header + Trailer
	totalPayments := int64(0)
	totalAmount := int64(0)

	for _, schedule := range file.Schedules {
		if achSch, ok := schedule.(*pamspr.ACHSchedule); ok {
			totalRecords += 2 // Schedule header + trailer
			for _, payment := range achSch.Payments {
				totalRecords++ // Payment record
				totalPayments++
				if achPay, ok := payment.(*pamspr.ACHPayment); ok {
					totalAmount += achPay.Amount
					totalRecords += int64(len(achPay.CARSTASBETC))
				}
			}
		}
	}

	file.Trailer = &pamspr.FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   totalRecords,
		TotalCountPayments:  totalPayments,
		TotalAmountPayments: totalAmount,
	}

	// Write to file
	outputFile, err := os.Create("ach_payments.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := pamspr.NewWriter(outputFile)
	if err := writer.Write(file); err != nil {
		log.Fatal(err)
	}

	fmt.Println("ACH payment file created successfully")
}

// Example 2: Parse and validate an existing file
func ParseAndValidateFile(filename string) {
	// Open file
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	// Parse file
	reader := pamspr.NewReader(inputFile)
	file, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Display file information
	fmt.Printf("File Header:\n")
	fmt.Printf("  Input System: %s\n", file.Header.InputSystem)
	fmt.Printf("  Version: %s\n", file.Header.StandardPaymentVersion)
	fmt.Printf("  Same Day ACH: %s\n", file.Header.IsRequestedForSameDayACH)

	fmt.Printf("\nSchedules: %d\n", len(file.Schedules))

	for i, schedule := range file.Schedules {
		fmt.Printf("\nSchedule %d:\n", i+1)
		fmt.Printf("  Schedule Number: %s\n", schedule.GetScheduleNumber())
		fmt.Printf("  Payment Type: %d\n", schedule.GetPaymentType())
		fmt.Printf("  Payments: %d\n", len(schedule.GetPayments()))

		// Show payment details
		for j, payment := range schedule.GetPayments() {
			fmt.Printf("\n  Payment %d:\n", j+1)
			fmt.Printf("    Payment ID: %s\n", payment.GetPaymentID())
			fmt.Printf("    Amount: $%.2f\n", float64(payment.GetAmount())/100)
			fmt.Printf("    Payee: %s\n", payment.GetPayeeName())
		}
	}

	fmt.Printf("\nFile Totals:\n")
	fmt.Printf("  Total Records: %d\n", file.Trailer.TotalCountRecords)
	fmt.Printf("  Total Payments: %d\n", file.Trailer.TotalCountPayments)
	fmt.Printf("  Total Amount: $%.2f\n", float64(file.Trailer.TotalAmountPayments)/100)
}

// Example 3: Create a check payment file with stubs
func CreateCheckFile() {
	file := &pamspr.File{
		Header: &pamspr.FileHeader{
			RecordCode:               "H ",
			InputSystem:              "AGENCY_CHECK_SYS",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: make([]pamspr.Schedule, 0),
	}

	// Create check schedule
	checkSchedule := &pamspr.CheckSchedule{
		Header: &pamspr.CheckScheduleHeader{
			RecordCode:                "11",
			ScheduleNumber:            "CHK001",
			PaymentTypeCode:           "Vendor",
			AgencyLocationCode:        "87654321",
			CheckPaymentEnclosureCode: "stub",
		},
		BaseSchedule: pamspr.BaseSchedule{
			ScheduleNumber: "CHK001",
			PaymentType:    "Vendor",
			ALC:            "87654321",
			Payments:       make([]pamspr.Payment, 0),
		},
	}

	// Add check payment
	payment := &pamspr.CheckPayment{
		RecordCode:                   "12",
		AgencyAccountIdentifier:      "VEND001",
		Amount:                       150000, // $1,500.00
		AgencyPaymentTypeCode:        "V",
		IsTOPOffset:                  "1",
		PayeeName:                    "ABC COMPANY LLC",
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
	}

	// Add check stub
	payment.Stub = &pamspr.CheckStub{
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
	}

	checkSchedule.Payments = append(checkSchedule.Payments, payment)

	// Set trailer
	checkSchedule.Trailer = &pamspr.ScheduleTrailer{
		RecordCode:     "T ",
		ScheduleCount:  1,
		ScheduleAmount: 150000,
	}

	file.Schedules = append(file.Schedules, checkSchedule)

	// Set file trailer
	file.Trailer = &pamspr.FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   6, // Header + Schedule Header + Payment + Stub + Schedule Trailer + File Trailer
		TotalCountPayments:  1,
		TotalAmountPayments: 150000,
	}

	// Write file
	outputFile, err := os.Create("check_payments.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := pamspr.NewWriter(outputFile)
	if err := writer.Write(file); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Check payment file created successfully")
}

// Example 4: Validate Same Day ACH requirements
func ValidateSameDayACH() {
	validator := pamspr.NewValidator()

	file := &pamspr.File{
		Header: &pamspr.FileHeader{
			RecordCode:               "H ",
			InputSystem:              "SDA_SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "1", // Same Day ACH requested
		},
	}

	// Validate the file structure first
	err := validator.ValidateFileStructure(file)
	if err != nil {
		fmt.Printf("File validation failed: %v\n", err)
	} else {
		fmt.Println("File validation passed")
	}

	// This will pass validation
	validPayment := &pamspr.ACHPayment{
		Amount:             50000, // $500.00 - under SDA limit
		ACHTransactionCode: "22",
	}

	// This will fail validation - amount too high
	invalidPayment := &pamspr.ACHPayment{
		Amount:             200000000, // $2,000,000 - over SDA limit
		ACHTransactionCode: "22",
	}

	// Test valid payment
	err = validator.ValidateACHPayment(validPayment)
	if err != nil {
		fmt.Printf("Valid payment failed: %v\n", err)
	} else {
		fmt.Println("Valid payment passed validation")
	}

	// Test invalid payment
	err = validator.ValidateACHPayment(invalidPayment)
	if err != nil {
		fmt.Printf("Invalid payment failed as expected: %v\n", err)
	}
}

func main() {
	fmt.Println("PAM SPR Library Examples")
	fmt.Println("========================")

	// Example 1: Create ACH file
	fmt.Println("\n1. Creating ACH payment file...")
	CreateACHFile()

	// Example 2: Parse the created file
	fmt.Println("\n2. Parsing ACH payment file...")
	ParseAndValidateFile("ach_payments.txt")

	// Example 3: Create check file
	fmt.Println("\n3. Creating check payment file...")
	CreateCheckFile()

	// Example 4: Validate Same Day ACH
	fmt.Println("\n4. Testing Same Day ACH validation...")
	ValidateSameDayACH()
}
