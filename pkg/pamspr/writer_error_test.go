package pamspr

import (
	"bytes"
	"strings"
	"testing"
)

// TestWriter_ErrorHandling tests that the writer properly returns errors instead of panicking
func TestWriter_ErrorHandling(t *testing.T) {
	// Test individual format functions for error accumulation
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Clear any existing errors
	writer.errors = writer.errors[:0]

	// Test a field that's too long
	oversizedString := strings.Repeat("X", 100)
	result := writer.formatField(oversizedString, 10)

	// Should return truncated value
	if len(result) != 10 {
		t.Errorf("formatField should return truncated value, got length %d", len(result))
	}

	// Should have accumulated an error
	if len(writer.errors) == 0 {
		t.Errorf("formatField should accumulate error for oversized input")
	}

	// Create a simple file that will cause a truncation error during Write()
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              oversizedString, // Too long for 40 chars
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

	// Write should return an error due to field truncation
	err := writer.Write(file)
	if err == nil {
		t.Errorf("Writer.Write() should return error for oversized fields, but got nil")
	}

	// Verify the error mentions field formatting
	if !strings.Contains(err.Error(), "field formatting") {
		t.Errorf("Error should mention field formatting, got: %s", err.Error())
	}
}

// TestWriter_NoErrorsForValidData tests that valid data doesn't produce errors
func TestWriter_NoErrorsForValidData(t *testing.T) {
	// Test individual format functions with valid data
	var buf bytes.Buffer
	writer := NewWriter(&buf)

	// Clear any existing errors
	writer.errors = writer.errors[:0]

	// Test a field that's within limits
	validString := "TEST"
	result := writer.formatField(validString, 10)

	// Should return padded value
	if len(result) != 10 {
		t.Errorf("formatField should return padded value, got length %d", len(result))
	}

	// Should not have accumulated any errors
	if len(writer.errors) != 0 {
		t.Errorf("formatField should not accumulate errors for valid input, got %d errors", len(writer.errors))
	}

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
