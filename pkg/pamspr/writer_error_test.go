package pamspr

import (
	"bytes"
	"strings"
	"testing"
)

// TestWriter_ErrorHandling tests that the writer properly returns errors
func TestWriter_ErrorHandling(t *testing.T) {
	// Create a simple file with invalid data
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              strings.Repeat("X", 50), // Too long for 40 chars
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

	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Write should succeed (streaming writer truncates automatically)
	err := writer.Write(file)
	if err != nil {
		t.Errorf("Writer.Write() should not fail on data truncation: %v", err)
	}

	// Verify output was generated
	output := buf.String()
	if len(output) == 0 {
		t.Errorf("Writer should produce output")
	}
}

// TestWriter_NoErrorsForValidData tests that valid data doesn't produce errors
func TestWriter_NoErrorsForValidData(t *testing.T) {
	// Create a simple valid file
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "TESTSYS", // Within 40 char limit
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

	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Write should succeed without errors
	err := writer.Write(file)
	if err != nil {
		t.Errorf("Writer.Write() should not return error for valid data, got: %s", err.Error())
	}

	// Verify output was written
	output := buf.String()
	if len(output) == 0 {
		t.Errorf("Writer should produce output for valid file")
	}
}
