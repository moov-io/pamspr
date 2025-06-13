package pamspr

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

// mockWriter simulates write errors
type mockWriter struct {
	err error
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	if m.err != nil {
		return 0, m.err
	}
	return len(p), nil
}

// TestNewWriter tests the NewWriter constructor
func TestNewWriter(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	if writer == nil {
		t.Fatal("NewWriter returned nil")
	}
	if writer.validator == nil {
		t.Error("writer has nil validator")
	}
}

// TestWriter_Write_CompleteACHFile tests writing a complete ACH file with all record types
func TestWriter_Write_CompleteACHFile(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TREASURY SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:              "01",
					AgencyACHText:           "TEST",
					ScheduleNumber:          "00000000000001",
					PaymentTypeCode:         "Vendor",
					StandardEntryClassCode:  "PPD",
					AgencyLocationCode:      "12345678",
					FederalEmployerIDNumber: "123456789 ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{
							RecordCode:              "02",
							AgencyAccountIdentifier: "ACC001          ",
							Amount:                  100000,
							PayeeName:               "TEST PAYEE",
							RoutingNumber:           "021000021",
							AccountNumber:           "1234567890",
							ACH_TransactionCode:     "22",
							PaymentID:               "PAY001              ",
							StandardEntryClassCode:  "PPD",
							Addenda: []*ACHAddendum{
								{
									RecordCode:         "03",
									PaymentID:          "PAY001              ",
									AddendaInformation: "TEST ADDENDUM",
								},
							},
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 100000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   6, // H + 01 + 02 + 03 + T + E
			TotalCountPayments:  1,
			TotalAmountPayments: 100000,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	// Verify correct number of lines
	expectedLines := 6 // H + 01 + 02 + 03 + T + E
	if len(lines) != expectedLines {
		t.Errorf("Expected %d lines, got %d", expectedLines, len(lines))
	}

	// Verify each line is exactly 850 characters
	for i, line := range lines {
		if len(line) != RecordLength {
			t.Errorf("Line %d has incorrect length: expected %d, got %d", i+1, RecordLength, len(line))
		}
	}

	// Verify line record codes
	expectedCodes := []string{"H ", "01", "02", "03", "T ", "E "}
	for i, expectedCode := range expectedCodes {
		if i < len(lines) && lines[i][:2] != expectedCode {
			t.Errorf("Line %d: expected record code %s, got %s", i+1, expectedCode, lines[i][:2])
		}
	}
}

// TestWriter_Write_CheckFile tests writing a check file
func TestWriter_Write_CheckFile(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "CHECK SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&CheckSchedule{
				Header: &CheckScheduleHeader{
					RecordCode:                "11",
					ScheduleNumber:            "00000000000001",
					PaymentTypeCode:           "Refund",
					AgencyLocationCode:        "87654321",
					CheckPaymentEnclosureCode: "stub",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&CheckPayment{
							RecordCode:              "12",
							AgencyAccountIdentifier: "CHK001          ",
							Amount:                  150000,
							PayeeName:               "CHECK PAYEE",
							PaymentID:               "CHK001              ",
							Stub: &CheckStub{
								RecordCode: "13",
								PaymentID:  "CHK001              ",
								PaymentIdentificationLines: [14]string{
									"INVOICE #12345",
									"DATE: 2024-01-01",
								},
							},
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 150000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   6, // H + 11 + 12 + 13 + T + E
			TotalCountPayments:  1,
			TotalAmountPayments: 150000,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	output := buf.String()
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")

	// Verify check schedule structure
	expectedCodes := []string{"H ", "11", "12", "13", "T ", "E "}
	for i, expectedCode := range expectedCodes {
		if i < len(lines) && lines[i][:2] != expectedCode {
			t.Errorf("Line %d: expected record code %s, got %s", i+1, expectedCode, lines[i][:2])
		}
	}
}

// TestWriter_WriteError tests writer error handling
func TestWriter_WriteError(t *testing.T) {
	mockWriter := &mockWriter{err: errors.New("write error")}
	writer := NewWriter(mockWriter)

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TEST",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   2,
			TotalCountPayments:  0,
			TotalAmountPayments: 0,
		},
	}

	err := writer.Write(file)
	if err == nil {
		t.Error("Expected write error, got nil")
	}
	if !strings.Contains(err.Error(), "write error") {
		t.Errorf("Expected 'write error' in error message, got: %v", err)
	}
}

// TestWriter_GetStats tests statistics tracking
func TestWriter_GetStats(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Initial stats should be zero
	recordCount, paymentCount, scheduleCount, totalAmount := writer.GetStats()
	if recordCount != 0 || paymentCount != 0 || scheduleCount != 0 || totalAmount != 0 {
		t.Errorf("Initial stats should be zero, got: records=%d, payments=%d, schedules=%d, amount=%d",
			recordCount, paymentCount, scheduleCount, totalAmount)
	}

	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TEST",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: []Schedule{
			&ACHSchedule{
				Header: &ACHScheduleHeader{
					RecordCode:              "01",
					AgencyACHText:           "TEST",
					ScheduleNumber:          "00000000000001",
					PaymentTypeCode:         "Vendor",
					StandardEntryClassCode:  "PPD",
					AgencyLocationCode:      "12345678",
					FederalEmployerIDNumber: "123456789 ",
				},
				BaseSchedule: BaseSchedule{
					Payments: []Payment{
						&ACHPayment{
							RecordCode:              "02",
							AgencyAccountIdentifier: "ACC001          ",
							Amount:                  100000,
							PayeeName:               "TEST PAYEE",
							RoutingNumber:           "021000021",
							AccountNumber:           "1234567890",
							ACH_TransactionCode:     "22",
							PaymentID:               "PAY001              ",
						},
					},
					Trailer: &ScheduleTrailer{
						RecordCode:     "T ",
						ScheduleCount:  1,
						ScheduleAmount: 100000,
					},
				},
			},
		},
		Trailer: &FileTrailer{
			RecordCode:          "E ",
			TotalCountRecords:   5, // H + 01 + 02 + T + E
			TotalCountPayments:  1,
			TotalAmountPayments: 100000,
		},
	}

	err := writer.Write(file)
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}

	// Check stats after writing
	recordCount, paymentCount, scheduleCount, totalAmount = writer.GetStats()
	if recordCount != 5 {
		t.Errorf("Expected record count 5, got %d", recordCount)
	}
	if paymentCount != 1 {
		t.Errorf("Expected payment count 1, got %d", paymentCount)
	}
	if scheduleCount != 1 {
		t.Errorf("Expected schedule count 1, got %d", scheduleCount)
	}
	if totalAmount != 100000 {
		t.Errorf("Expected total amount 100000, got %d", totalAmount)
	}
}
