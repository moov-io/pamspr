package pamspr

import (
	"testing"
)

func TestFileStructure(t *testing.T) {
	// Test creating a basic file structure
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TEST_SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: make([]Schedule, 0),
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   2,
			TotalCountPayments:  0,
			TotalAmountPayments: 0,
		},
	}

	if file.Header.RecordCode != "H " {
		t.Errorf("Expected header record code 'H ', got %s", file.Header.RecordCode)
	}
	if file.Trailer.RecordCode != "E " {
		t.Errorf("Expected trailer record code 'E ', got %s", file.Trailer.RecordCode)
	}
}

func TestACHScheduleInterface(t *testing.T) {
	schedule := &ACHSchedule{
		Header: &ACHScheduleHeader{
			RecordCode:             "01",
			ScheduleNumber:         "TEST001",
			PaymentTypeCode:        "Salary",
			StandardEntryClassCode: "PPD",
			AgencyLocationCode:     "12345678",
		},
		BaseSchedule: BaseSchedule{
			ScheduleNumber: "TEST001",
			PaymentType:    "Salary",
			ALC:            "12345678",
			Payments:       make([]Payment, 0),
		},
	}

	// Test interface methods
	if schedule.GetScheduleNumber() != "TEST001" {
		t.Errorf("Expected schedule number TEST001, got %s", schedule.GetScheduleNumber())
	}
	if schedule.GetPaymentType() != PaymentTypeACH {
		t.Errorf("Expected payment type ACH, got %v", schedule.GetPaymentType())
	}
	if len(schedule.GetPayments()) != 0 {
		t.Errorf("Expected 0 payments, got %d", len(schedule.GetPayments()))
	}
}

func TestCheckScheduleInterface(t *testing.T) {
	schedule := &CheckSchedule{
		Header: &CheckScheduleHeader{
			RecordCode:                "11",
			ScheduleNumber:            "CHK001",
			PaymentTypeCode:           "Vendor",
			AgencyLocationCode:        "87654321",
			CheckPaymentEnclosureCode: "stub",
		},
		BaseSchedule: BaseSchedule{
			ScheduleNumber: "CHK001",
			PaymentType:    "Vendor",
			ALC:            "87654321",
			Payments:       make([]Payment, 0),
		},
	}

	// Test interface methods
	if schedule.GetScheduleNumber() != "CHK001" {
		t.Errorf("Expected schedule number CHK001, got %s", schedule.GetScheduleNumber())
	}
	if schedule.GetPaymentType() != PaymentTypeCheck {
		t.Errorf("Expected payment type Check, got %v", schedule.GetPaymentType())
	}
}

func TestACHPaymentInterface(t *testing.T) {
	payment := &ACHPayment{
		RecordCode:    "02",
		PaymentID:     "PAY001",
		Amount:        150000, // $1,500.00
		PayeeName:     "JOHN DOE",
		RoutingNumber: "021000021",
		AccountNumber: "1234567890",
	}

	// Test interface methods
	if payment.GetPaymentID() != "PAY001" {
		t.Errorf("Expected payment ID PAY001, got %s", payment.GetPaymentID())
	}
	if payment.GetAmount() != 150000 {
		t.Errorf("Expected amount 150000, got %d", payment.GetAmount())
	}
	if payment.GetPayeeName() != "JOHN DOE" {
		t.Errorf("Expected payee name JOHN DOE, got %s", payment.GetPayeeName())
	}
}

func TestCheckPaymentInterface(t *testing.T) {
	payment := &CheckPayment{
		RecordCode: "12",
		PaymentID:  "CHK001",
		Amount:     100000, // $1,000.00
		PayeeName:  "ABC COMPANY",
	}

	// Test interface methods
	if payment.GetPaymentID() != "CHK001" {
		t.Errorf("Expected payment ID CHK001, got %s", payment.GetPaymentID())
	}
	if payment.GetAmount() != 100000 {
		t.Errorf("Expected amount 100000, got %d", payment.GetAmount())
	}
	if payment.GetPayeeName() != "ABC COMPANY" {
		t.Errorf("Expected payee name ABC COMPANY, got %s", payment.GetPayeeName())
	}
}

func TestRecordTypes(t *testing.T) {
	tests := []struct {
		name     string
		record   RecordType
		expected string
	}{
		{"FileHeader", RecordTypeFileHeader, "H "},
		{"ACHScheduleHeader", RecordTypeACHScheduleHeader, "01"},
		{"ACHPayment", RecordTypeACHPayment, "02"},
		{"ACHAddendum", RecordTypeACHAddendum, "03"},
		{"ACHAddendumCTX", RecordTypeACHAddendumCTX, "04"},
		{"CheckScheduleHeader", RecordTypeCheckScheduleHeader, "11"},
		{"CheckPayment", RecordTypeCheckPayment, "12"},
		{"CheckStub", RecordTypeCheckStub, "13"},
		{"CARSTASBETC", RecordTypeCARSTASBETC, "G "},
		{"DNP", RecordTypeDNP, "DD"},
		{"ScheduleTrailer", RecordTypeScheduleTrailer, "T "},
		{"FileTrailer", RecordTypeFileTrailer, "E "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.record) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(tt.record))
			}
		})
	}
}

func TestStandardEntryClassCodes(t *testing.T) {
	tests := []struct {
		name     string
		code     StandardEntryClassCode
		expected string
	}{
		{"CCD", SECCodeCCD, "CCD"},
		{"PPD", SECCodePPD, "PPD"},
		{"IAT", SECCodeIAT, "IAT"},
		{"CTX", SECCodeCTX, "CTX"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if string(tt.code) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(tt.code))
			}
		})
	}
}

func TestACHAddendum(t *testing.T) {
	// Test standard addendum (03)
	addendum := &ACHAddendum{
		RecordCode:         "03",
		PaymentID:          "PAY001",
		AddendaInformation: "INVOICE 12345",
	}

	if addendum.RecordCode != "03" {
		t.Errorf("Expected record code 03, got %s", addendum.RecordCode)
	}
	if len(addendum.AddendaInformation) > 80 {
		t.Errorf("Addenda information too long for record code 03")
	}

	// Test CTX addendum (04)
	ctxAddendum := &ACHAddendum{
		RecordCode:         "04",
		PaymentID:          "PAY002",
		AddendaInformation: "ISA*00*          *00*          *ZZ*123456789      *ZZ*987654321      *210101*1200*U*00401*000000001*0*P*>",
	}

	if ctxAddendum.RecordCode != "04" {
		t.Errorf("Expected record code 04, got %s", ctxAddendum.RecordCode)
	}
}

func TestCARSTASBETC(t *testing.T) {
	cars := &CARSTASBETC{
		RecordCode:                  "G ",
		PaymentID:                   "PAY001",
		AgencyIdentifier:            "012",
		MainAccountCode:             "1234",
		SubAccountCode:              "001",
		BusinessEventTypeCode:       "SALARY01",
		AccountClassificationAmount: 150000,
		IsCredit:                    "0",
	}

	if cars.RecordCode != "G " {
		t.Errorf("Expected record code 'G ', got %s", cars.RecordCode)
	}
	if cars.AccountClassificationAmount != 150000 {
		t.Errorf("Expected amount 150000, got %d", cars.AccountClassificationAmount)
	}
}

func TestCheckStub(t *testing.T) {
	stub := &CheckStub{
		RecordCode: "13",
		PaymentID:  "CHK001",
		PaymentIdentificationLines: [14]string{
			"Invoice: 12345",
			"Date: 01/01/2024",
			"Amount: $1,000.00",
			"", "", "", "", "", "", "", "", "", "", "",
		},
	}

	if stub.RecordCode != "13" {
		t.Errorf("Expected record code 13, got %s", stub.RecordCode)
	}
	if stub.PaymentIdentificationLines[0] != "Invoice: 12345" {
		t.Errorf("Expected first line to be 'Invoice: 12345', got %s", stub.PaymentIdentificationLines[0])
	}
}

func TestDNPRecord(t *testing.T) {
	dnp := &DNPRecord{
		RecordCode: "DD",
		PaymentID:  "PAY001",
		DNPDetail:  "AGENCY SPECIFIC DATA",
	}

	if dnp.RecordCode != "DD" {
		t.Errorf("Expected record code DD, got %s", dnp.RecordCode)
	}
	if dnp.DNPDetail != "AGENCY SPECIFIC DATA" {
		t.Errorf("Expected DNP detail 'AGENCY SPECIFIC DATA', got %s", dnp.DNPDetail)
	}
}

func TestScheduleTrailer(t *testing.T) {
	trailer := &ScheduleTrailer{
		RecordCode:     "T ",
		ScheduleCount:  5,
		ScheduleAmount: 750000, // $7,500.00
	}

	if trailer.RecordCode != "T " {
		t.Errorf("Expected record code 'T ', got %s", trailer.RecordCode)
	}
	if trailer.ScheduleCount != 5 {
		t.Errorf("Expected count 5, got %d", trailer.ScheduleCount)
	}
	if trailer.ScheduleAmount != 750000 {
		t.Errorf("Expected amount 750000, got %d", trailer.ScheduleAmount)
	}
}

func TestFileWithSchedules(t *testing.T) {
	// Create a complete file with ACH and Check schedules
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TEST_SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:             "01",
					ScheduleNumber:         "ACH001",
					PaymentTypeCode:        "Salary",
					StandardEntryClassCode: "PPD",
				},
				BaseSchedule: BaseSchedule{
					ScheduleNumber: "ACH001",
					PaymentType:    "Salary",
					Payments: []Payment{
						&ACHPayment{
							PaymentID: "PAY001",
							Amount:    100000,
							PayeeName: "TEST PAYEE",
						},
					},
				},
			},
			&CheckSchedule{
				Header: &CheckScheduleHeader{
					RecordCode:      "11",
					ScheduleNumber:  "CHK001",
					PaymentTypeCode: "Vendor",
				},
				BaseSchedule: BaseSchedule{
					ScheduleNumber: "CHK001",
					PaymentType:    "Vendor",
					Payments: []Payment{
						&CheckPayment{
							PaymentID: "CHK001",
							Amount:    200000,
							PayeeName: "TEST VENDOR",
						},
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   8,
			TotalCountPayments:  2,
			TotalAmountPayments: 300000,
		},
	}

	if len(file.Schedules) != 2 {
		t.Errorf("Expected 2 schedules, got %d", len(file.Schedules))
	}

	// Verify ACH schedule
	achSchedule, ok := file.Schedules[0].(*ACHSchedule)
	if !ok {
		t.Error("First schedule should be ACH schedule")
	}
	if len(achSchedule.GetPayments()) != 1 {
		t.Errorf("Expected 1 ACH payment, got %d", len(achSchedule.GetPayments()))
	}

	// Verify Check schedule
	checkSchedule, ok := file.Schedules[1].(*CheckSchedule)
	if !ok {
		t.Error("Second schedule should be Check schedule")
	}
	if len(checkSchedule.GetPayments()) != 1 {
		t.Errorf("Expected 1 Check payment, got %d", len(checkSchedule.GetPayments()))
	}

	// Verify totals
	if file.Trailer.TotalCountPayments != 2 {
		t.Errorf("Expected 2 total payments, got %d", file.Trailer.TotalCountPayments)
	}
	if file.Trailer.TotalAmountPayments != 300000 {
		t.Errorf("Expected total amount 300000, got %d", file.Trailer.TotalAmountPayments)
	}
}
