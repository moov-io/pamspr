package pamspr

import (
	"strings"
	"testing"
)

func TestValidationError(t *testing.T) {
	err := ValidationError{
		Field:   "Amount",
		Value:   "-100",
		Rule:    "positive",
		Message: "amount cannot be negative",
	}

	expected := "validation error: field=Amount, value=-100, rule=positive, message=amount cannot be negative"
	if err.Error() != expected {
		t.Errorf("Expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestValidateFileHeader(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		header    *FileHeader
		expectErr bool
		errField  string
	}{
		{
			name: "Valid header",
			header: &FileHeader{
				RecordCode:               "H ",
				InputSystem:              "TEST_SYSTEM",
				StandardPaymentVersion:   "502",
				IsRequestedForSameDayACH: "0",
			},
			expectErr: false,
		},
		{
			name: "Invalid record code",
			header: &FileHeader{
				RecordCode:             "XX",
				StandardPaymentVersion: "502",
			},
			expectErr: true,
			errField:  "RecordCode",
		},
		{
			name: "Invalid version",
			header: &FileHeader{
				RecordCode:             "H ",
				StandardPaymentVersion: "500",
			},
			expectErr: true,
			errField:  "StandardPaymentVersion",
		},
		{
			name: "Invalid SDA flag",
			header: &FileHeader{
				RecordCode:               "H ",
				StandardPaymentVersion:   "502",
				IsRequestedForSameDayACH: "X",
			},
			expectErr: true,
			errField:  "IsRequestedForSameDayACH",
		},
		{
			name: "Blank SDA flag is valid",
			header: &FileHeader{
				RecordCode:               "H ",
				StandardPaymentVersion:   "502",
				IsRequestedForSameDayACH: "",
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateFileHeader(tt.header)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tt.expectErr && err != nil {
				if valErr, ok := err.(ValidationError); ok {
					if valErr.Field != tt.errField {
						t.Errorf("Expected error on field %s, got %s", tt.errField, valErr.Field)
					}
				}
			}
		})
	}
}

func TestValidateACHPayment(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		payment   *ACHPayment
		expectErr bool
		errField  string
	}{
		{
			name: "Valid ACH payment",
			payment: &ACHPayment{
				Amount:                       100000,
				PayeeName:                    "JOHN DOE",
				RoutingNumber:                "021000021",
				AccountNumber:                "1234567890",
				ACHTransactionCode:           "22",
				PaymentID:                    "PAY001",
				TIN:                          "123456789",
				PaymentRecipientTINIndicator: "1",
			},
			expectErr: false,
		},
		{
			name: "Negative amount",
			payment: &ACHPayment{
				Amount:             -100,
				PayeeName:          "JOHN DOE",
				RoutingNumber:      "021000021",
				AccountNumber:      "1234567890",
				ACHTransactionCode: "22",
				PaymentID:          "PAY001",
			},
			expectErr: true,
			errField:  "Amount",
		},
		{
			name: "Missing payee name",
			payment: &ACHPayment{
				Amount:             100000,
				PayeeName:          "",
				RoutingNumber:      "021000021",
				AccountNumber:      "1234567890",
				ACHTransactionCode: "22",
				PaymentID:          "PAY001",
			},
			expectErr: true,
			errField:  "PayeeName",
		},
		{
			name: "Invalid routing number",
			payment: &ACHPayment{
				Amount:             100000,
				PayeeName:          "JOHN DOE",
				RoutingNumber:      "123456789", // Invalid checksum
				AccountNumber:      "1234567890",
				ACHTransactionCode: "22",
				PaymentID:          "PAY001",
			},
			expectErr: true,
			errField:  "RoutingNumber",
		},
		{
			name: "All zeros account number",
			payment: &ACHPayment{
				Amount:             100000,
				PayeeName:          "JOHN DOE",
				RoutingNumber:      "021000021",
				AccountNumber:      "00000000000000000",
				ACHTransactionCode: "22",
				PaymentID:          "PAY001",
			},
			expectErr: true,
			errField:  "AccountNumber",
		},
		{
			name: "Invalid transaction code",
			payment: &ACHPayment{
				Amount:             100000,
				PayeeName:          "JOHN DOE",
				RoutingNumber:      "021000021",
				AccountNumber:      "1234567890",
				ACHTransactionCode: "99",
				PaymentID:          "PAY001",
			},
			expectErr: true,
			errField:  "ACHTransactionCode",
		},
		{
			name: "Invalid TIN",
			payment: &ACHPayment{
				Amount:             100000,
				PayeeName:          "JOHN DOE",
				RoutingNumber:      "021000021",
				AccountNumber:      "1234567890",
				ACHTransactionCode: "22",
				PaymentID:          "PAY001",
				TIN:                "12345", // Too short
			},
			expectErr: true,
			errField:  "TIN",
		},
		{
			name: "Missing payment ID",
			payment: &ACHPayment{
				Amount:             100000,
				PayeeName:          "JOHN DOE",
				RoutingNumber:      "021000021",
				AccountNumber:      "1234567890",
				ACHTransactionCode: "22",
				PaymentID:          "",
			},
			expectErr: true,
			errField:  "PaymentID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateACHPayment(tt.payment)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tt.expectErr && err != nil {
				if valErr, ok := err.(ValidationError); ok {
					if valErr.Field != tt.errField {
						t.Errorf("Expected error on field %s, got %s", tt.errField, valErr.Field)
					}
				}
			}
		})
	}
}

func TestValidateIATPayment(t *testing.T) {
	validator := NewValidator()

	payment := &ACHPayment{
		Amount:                 100000,
		PayeeName:              "JOHN DOE",
		RoutingNumber:          "021000021",
		AccountNumber:          "1234567890",
		ACHTransactionCode:     "22",
		PaymentID:              "PAY001",
		StandardEntryClassCode: "IAT",
		// Missing required IAT fields
	}

	err := validator.ValidateACHPayment(payment)
	if err == nil {
		t.Error("Expected error for IAT payment missing required fields")
	}

	// Add required fields
	payment.PayeeAddressLine1 = "123 MAIN ST"
	payment.CityName = "LONDON"
	payment.CountryCodeText = "GB"

	err = validator.ValidateACHPayment(payment)
	if err != nil {
		t.Errorf("Unexpected error for valid IAT payment: %v", err)
	}
}

func TestValidateCheckPayment(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		payment   *CheckPayment
		expectErr bool
		errField  string
	}{
		{
			name: "Valid check payment",
			payment: &CheckPayment{
				Amount:    100000,
				PayeeName: "ABC COMPANY",
				PaymentID: "CHK001",
				TIN:       "123456789",
			},
			expectErr: false,
		},
		{
			name: "Zero amount check",
			payment: &CheckPayment{
				Amount:    0,
				PayeeName: "ABC COMPANY",
				PaymentID: "CHK001",
			},
			expectErr: true,
			errField:  "Amount",
		},
		{
			name: "Negative amount check",
			payment: &CheckPayment{
				Amount:    -100,
				PayeeName: "ABC COMPANY",
				PaymentID: "CHK001",
			},
			expectErr: true,
			errField:  "Amount",
		},
		{
			name: "Missing payee name",
			payment: &CheckPayment{
				Amount:    100000,
				PayeeName: "",
				PaymentID: "CHK001",
			},
			expectErr: true,
			errField:  "PayeeName",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCheckPayment(tt.payment)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tt.expectErr && err != nil {
				if valErr, ok := err.(ValidationError); ok {
					if valErr.Field != tt.errField {
						t.Errorf("Expected error on field %s, got %s", tt.errField, valErr.Field)
					}
				}
			}
		})
	}
}

func TestValidateRoutingNumber(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		rtn       string
		expectErr bool
	}{
		{"Valid routing number", "021000021", false},
		{"Valid routing number 2", "011401533", false},
		{"Too short", "12345678", true},
		{"Too long", "1234567890", true},
		{"Non-numeric", "12345678X", true},
		{"Invalid prefix", "991234567", true},
		{"Invalid checksum", "123456789", true},
		{"Valid prefix 00", "001234563", false}, // Government
		{"Valid prefix 12", "121234562", false},
		{"Valid prefix 21", "211234565", false}, // Thrift
		{"Valid prefix 32", "321234564", false},
		{"Valid prefix 61", "611234561", false}, // Electronic
		{"Valid prefix 72", "721234560", false},
		{"Valid prefix 80", "801234569", false}, // Traveler's checks
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.validateRoutingNumber(tt.rtn)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestValidateSameDayACH(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		file      *File
		expectErr bool
	}{
		{
			name: "Not SDA file",
			file: &File{
				Header: &FileHeader{
					IsRequestedForSameDayACH: "0",
				},
			},
			expectErr: false,
		},
		{
			name: "Valid SDA file",
			file: &File{
				Header: &FileHeader{
					IsRequestedForSameDayACH: "1",
				},
				Schedules: []Schedule{
					&ACHSchedule{
						Header: &ACHScheduleHeader{
							StandardEntryClassCode: "PPD",
						},
						BaseSchedule: BaseSchedule{
							Payments: []Payment{
								&ACHPayment{Amount: 50000}, // $500
							},
						},
					},
				},
			},
			expectErr: false,
		},
		{
			name: "SDA with check schedule",
			file: &File{
				Header: &FileHeader{
					IsRequestedForSameDayACH: "1",
				},
				Schedules: []Schedule{
					&CheckSchedule{},
				},
			},
			expectErr: true,
		},
		{
			name: "SDA with amount over limit",
			file: &File{
				Header: &FileHeader{
					IsRequestedForSameDayACH: "1",
				},
				Schedules: []Schedule{
					&ACHSchedule{
						Header: &ACHScheduleHeader{
							StandardEntryClassCode: "PPD",
						},
						BaseSchedule: BaseSchedule{
							Payments: []Payment{
								&ACHPayment{Amount: 200000000}, // $2,000,000
							},
						},
					},
				},
			},
			expectErr: true,
		},
		{
			name: "SDA with IAT",
			file: &File{
				Header: &FileHeader{
					IsRequestedForSameDayACH: "1",
				},
				Schedules: []Schedule{
					&ACHSchedule{
						Header: &ACHScheduleHeader{
							StandardEntryClassCode: "IAT",
						},
						BaseSchedule: BaseSchedule{
							Payments: []Payment{},
						},
					},
				},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateSameDayACH(tt.file)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestValidateBalancing(t *testing.T) {
	validator := NewValidator()

	// Create a properly balanced file
	file := &File{
		Header: &FileHeader{},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{Amount: 100000},
						&ACHPayment{Amount: 200000},
					},
					Trailer: &ScheduleTrailer{
						ScheduleCount:  2,
						ScheduleAmount: 300000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			TotalCountRecords:   6, // Header + Schedule Header + 2 Payments + Schedule Trailer + Trailer
			TotalCountPayments:  2,
			TotalAmountPayments: 300000,
		},
	}

	err := validator.ValidateBalancing(file)
	if err != nil {
		t.Errorf("Unexpected error for balanced file: %v", err)
	}

	// Test unbalanced payment count
	file.Trailer.TotalCountPayments = 3
	err = validator.ValidateBalancing(file)
	if err == nil {
		t.Error("Expected error for unbalanced payment count")
	}

	// Test unbalanced amount
	file.Trailer.TotalCountPayments = 2
	file.Trailer.TotalAmountPayments = 400000
	err = validator.ValidateBalancing(file)
	if err == nil {
		t.Error("Expected error for unbalanced amount")
	}

	// Test unbalanced record count
	file.Trailer.TotalAmountPayments = 300000
	file.Trailer.TotalCountRecords = 10
	err = validator.ValidateBalancing(file)
	if err == nil {
		t.Error("Expected error for unbalanced record count")
	}
}

func TestValidateCTXAddendum(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		payment   *ACHPayment
		expectErr bool
	}{
		{
			name: "Non-CTX payment",
			payment: &ACHPayment{
				StandardEntryClassCode: "PPD",
			},
			expectErr: false,
		},
		{
			name: "CTX without addendum",
			payment: &ACHPayment{
				StandardEntryClassCode: "CTX",
				Addenda:                []*ACHAddendum{},
			},
			expectErr: true,
		},
		{
			name: "CTX with valid addendum",
			payment: &ACHPayment{
				StandardEntryClassCode: "CTX",
				Addenda: []*ACHAddendum{
					{
						RecordCode:         "04",
						AddendaInformation: "ISA*00*          *00*          ",
					},
				},
			},
			expectErr: false,
		},
		{
			name: "CTX with wrong record code",
			payment: &ACHPayment{
				StandardEntryClassCode: "CTX",
				Addenda: []*ACHAddendum{
					{
						RecordCode:         "03",
						AddendaInformation: "ISA*00*          *00*          ",
					},
				},
			},
			expectErr: true,
		},
		{
			name: "CTX without ISA",
			payment: &ACHPayment{
				StandardEntryClassCode: "CTX",
				Addenda: []*ACHAddendum{
					{
						RecordCode:         "04",
						AddendaInformation: "BPR*00*          *00*          ",
					},
				},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateCTXAddendum(tt.payment)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestValidateHexCharacters(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		data      string
		expectErr bool
	}{
		{"Valid ASCII", "HELLO WORLD 123", false},
		{"Valid special chars", "TEST@#$%&*()", false},
		{"Invalid hex 0x00", "TEST\x00DATA", true},
		{"Invalid hex 0x1F", "TEST\x1FDATA", true},
		{"Invalid hex 0x3F", "TEST\x3FDATA", true},
		{"Valid hex 0x40", "TEST@DATA", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateHexCharacters(tt.data)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestValidateScheduleNumber(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		schedNum  string
		expectErr bool
	}{
		{"Valid alphanumeric", "TEST001", false},
		{"Valid with dash", "TEST-001", false},
		{"All uppercase", "SCHEDULE1", false},
		{"Lowercase converted", "test001", false}, // Should be converted to uppercase
		{"With spaces", "  TEST001  ", false},     // Should be trimmed
		{"Empty string", "", true},
		{"Invalid chars", "TEST@001", true},
		{"Invalid chars 2", "TEST.001", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateScheduleNumber(tt.schedNum)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestValidateAgencySpecific(t *testing.T) {
	validator := NewValidator()

	// Test IRS validation
	irsPayment := &ACHPayment{
		Reconcilement: strings.Repeat("X", 100), // Correct length
	}
	err := validator.ValidateAgencySpecific(irsPayment, "IRS")
	if err != nil {
		t.Errorf("Unexpected error for valid IRS payment: %v", err)
	}

	// Test IRS with wrong length
	irsPayment.Reconcilement = "TOO SHORT"
	err = validator.ValidateAgencySpecific(irsPayment, "IRS")
	if err == nil {
		t.Error("Expected error for invalid IRS reconcilement length")
	}

	// Test other agencies (placeholder implementations)
	err = validator.ValidateAgencySpecific(&ACHPayment{}, "VA")
	if err != nil {
		t.Errorf("Unexpected error for VA: %v", err)
	}

	err = validator.ValidateAgencySpecific(&ACHPayment{}, "SSA")
	if err != nil {
		t.Errorf("Unexpected error for SSA: %v", err)
	}

	err = validator.ValidateAgencySpecific(&ACHPayment{}, "UNKNOWN")
	if err != nil {
		t.Errorf("Unexpected error for unknown agency: %v", err)
	}
}

func TestValidateFileStructure(t *testing.T) {
	validator := NewValidator()

	tests := []struct {
		name      string
		file      *File
		expectErr bool
	}{
		{
			name: "Valid file structure",
			file: &File{
				Header: &FileHeader{},
				Schedules: []Schedule{
					&ACHSchedule{
						BaseSchedule: BaseSchedule{
							Payments: []Payment{
								&ACHPayment{RoutingNumber: "021000021"},
								&ACHPayment{RoutingNumber: "021000021"},
							},
						},
					},
				},
				Trailer: &FileTrailer{},
			},
			expectErr: false,
		},
		{
			name: "Missing header",
			file: &File{
				Schedules: []Schedule{},
				Trailer:   &FileTrailer{},
			},
			expectErr: true,
		},
		{
			name: "Missing trailer",
			file: &File{
				Header:    &FileHeader{},
				Schedules: []Schedule{},
			},
			expectErr: true,
		},
		{
			name: "Mixed payment types in ACH schedule",
			file: &File{
				Header: &FileHeader{},
				Schedules: []Schedule{
					&ACHSchedule{
						BaseSchedule: BaseSchedule{
							Payments: []Payment{
								&ACHPayment{},
								&CheckPayment{}, // Wrong type
							},
						},
					},
				},
				Trailer: &FileTrailer{},
			},
			expectErr: true,
		},
		{
			name: "ACH payments out of routing number order",
			file: &File{
				Header: &FileHeader{},
				Schedules: []Schedule{
					&ACHSchedule{
						BaseSchedule: BaseSchedule{
							Payments: []Payment{
								&ACHPayment{RoutingNumber: "121000248"},
								&ACHPayment{RoutingNumber: "021000021"}, // Lower RTN after higher
							},
						},
					},
				},
				Trailer: &FileTrailer{},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateFileStructure(tt.file)
			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
