package pamspr

import (
	"bytes"
	"strings"
	"testing"
)

func sampleACHPayment(id string, amount int64) *ACHPayment {
	return &ACHPayment{
		AgencyAccountIdentifier:      "EMP001",
		Amount:                       amount,
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
		PaymentID:                    id,
		TIN:                          "123456789",
		PaymentRecipientTINIndicator: "1",
	}
}

func sampleCheckPayment(id string, amount int64) *CheckPayment {
	return &CheckPayment{
		AgencyAccountIdentifier:      "VEND001",
		Amount:                       amount,
		AgencyPaymentTypeCode:        "V",
		IsTOP_Offset:                 "1",
		PayeeName:                    "ACME CORP",
		PayeeAddressLine1:            "789 BUSINESS BLVD",
		CityName:                     "NEW YORK",
		StateCodeText:                "NY",
		PostalCode:                   "10001",
		PaymentID:                    id,
		TIN:                          "123456789",
		PaymentRecipientTINIndicator: "2",
	}
}

func writeBuiltFile(t *testing.T, file *File) []byte {
	t.Helper()
	var buf bytes.Buffer
	if err := NewWriter(&buf).Write(file); err != nil {
		t.Fatalf("Write() error = %v", err)
	}
	return buf.Bytes()
}

func TestProcessFileGetScheduleNumber(t *testing.T) {
	file, err := NewFileBuilder().
		WithHeader("TEST_SYSTEM", "502", false).
		StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
		AddACHPayment(sampleACHPayment("PAY001", 150000)).
		StartCheckSchedule("CHK001", "Vendor", "87654321", "stub").
		AddCheckPayment(sampleCheckPayment("CHK001", 100000)).
		Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}

	data := writeBuiltFile(t, file)

	parsed, err := NewReader(bytes.NewReader(data)).Read()
	if err != nil {
		t.Fatalf("Read() error = %v", err)
	}
	if len(parsed.Schedules) != 2 {
		t.Fatalf("Read() schedules = %d, want 2", len(parsed.Schedules))
	}

	var got []string
	err = NewReader(bytes.NewReader(data)).ProcessFile(
		func(schedule Schedule, _ int) bool {
			got = append(got, schedule.GetScheduleNumber())
			return true
		},
		func(Payment, int, int) bool { return true },
		nil,
	)
	if err != nil {
		t.Fatalf("ProcessFile() error = %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("ProcessFile() schedules = %d, want 2", len(got))
	}
	for i, schedule := range parsed.Schedules {
		want := schedule.GetScheduleNumber()
		if want == "" {
			t.Fatalf("Read() schedule %d has empty GetScheduleNumber()", i)
		}
		if got[i] != want {
			t.Errorf("ProcessFile() schedule %d GetScheduleNumber() = %q, want %q", i, got[i], want)
		}
	}
}

func TestProcessFileFileTrailerCallbackOnce(t *testing.T) {
	data := []byte(createTestFileContent())

	var trailerCount int
	var trailerLines []int
	var recordCodes []string
	err := NewReader(bytes.NewReader(data)).ProcessFile(
		func(Schedule, int) bool { return true },
		func(Payment, int, int) bool { return true },
		func(recordType string, lineNumber int, _ string) {
			recordCodes = append(recordCodes, recordType)
			if recordType == "E " {
				trailerCount++
				trailerLines = append(trailerLines, lineNumber)
			}
		},
	)
	if err != nil {
		t.Fatalf("ProcessFile() error = %v", err)
	}

	if trailerCount != 1 {
		t.Fatalf("file trailer callback count = %d, want 1 (codes=%v)", trailerCount, recordCodes)
	}
	if len(trailerLines) != 1 {
		t.Fatalf("file trailer line numbers = %v, want one entry", trailerLines)
	}

	codeCounts := map[string]int{}
	for _, code := range recordCodes {
		codeCounts[code]++
	}
	if codeCounts["H "] != 1 {
		t.Errorf("header callback count = %d, want 1", codeCounts["H "])
	}
	if codeCounts["01"] != 1 {
		t.Errorf("ACH schedule header callback count = %d, want 1", codeCounts["01"])
	}
}

func TestValidateFileStructureOnlyStats(t *testing.T) {
	data := []byte(createTestFileContent())
	lineCount := int64(len(strings.Split(strings.TrimRight(string(data), "\n"), "\n")))

	reader := NewReader(bytes.NewReader(data))
	if err := reader.ValidateFileStructureOnly(); err != nil {
		t.Fatalf("ValidateFileStructureOnly() error = %v", err)
	}

	stats := reader.GetStats()
	if stats.LinesProcessed != lineCount {
		t.Errorf("LinesProcessed = %d, want %d", stats.LinesProcessed, lineCount)
	}
	if stats.PaymentsProcessed != 2 {
		t.Errorf("PaymentsProcessed = %d, want 2", stats.PaymentsProcessed)
	}
	if stats.SchedulesProcessed != 1 {
		t.Errorf("SchedulesProcessed = %d, want 1", stats.SchedulesProcessed)
	}
}

func TestValidateFileStructureOnlyStatsFiveRecordFile(t *testing.T) {
	file, err := NewFileBuilder().
		WithHeader("TEST_SYSTEM", "502", false).
		StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
		AddACHPayment(sampleACHPayment("PAY001", 150000)).
		Build()
	if err != nil {
		t.Fatalf("Build() error = %v", err)
	}
	data := writeBuiltFile(t, file)

	lineCount := int64(len(strings.Split(strings.TrimRight(string(data), "\n"), "\n")))
	if lineCount != 5 {
		t.Fatalf("fixture line count = %d, want 5 (H, 01, 02, T, E)", lineCount)
	}

	reader := NewReader(bytes.NewReader(data))
	if err := reader.ValidateFileStructureOnly(); err != nil {
		t.Fatalf("ValidateFileStructureOnly() error = %v", err)
	}

	stats := reader.GetStats()
	if stats.LinesProcessed != 5 {
		t.Errorf("LinesProcessed = %d, want 5", stats.LinesProcessed)
	}
	if stats.PaymentsProcessed != 1 {
		t.Errorf("PaymentsProcessed = %d, want 1", stats.PaymentsProcessed)
	}
	if stats.SchedulesProcessed != 1 {
		t.Errorf("SchedulesProcessed = %d, want 1", stats.SchedulesProcessed)
	}
}
