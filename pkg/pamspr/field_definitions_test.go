package pamspr

import (
	"strings"
	"testing"
)

func TestFieldDefinitions(t *testing.T) {
	// Test that all field positions are valid
	if err := ValidateFieldPositions(); err != nil {
		t.Fatalf("Field position validation failed: %v", err)
	}
}

func TestGetFieldDefinitions(t *testing.T) {
	tests := []struct {
		recordCode    string
		shouldExist   bool
		expectedField string // A field that should exist in the definition
	}{
		{"H ", true, "InputSystem"},
		{"01", true, "AgencyACHText"},
		{"02", true, "PayeeName"},
		{"03", true, "AddendaInformation"},
		{"04", true, "AddendaInformation"},
		{"11", true, "CheckPaymentEnclosureCode"},
		{"12", true, "PayeeName"},
		{"13", true, "Line1"},
		{"G ", true, "AccountClassificationAmount"},
		{"DD", true, "DNPDetail"},
		{"T ", true, "ScheduleCount"},
		{"E ", true, "TotalCountRecords"},
		{"XX", false, ""},
	}

	for _, tt := range tests {
		t.Run("Record_"+tt.recordCode, func(t *testing.T) {
			fields := GetFieldDefinitions(tt.recordCode)
			
			if tt.shouldExist {
				if fields == nil {
					t.Errorf("Expected field definitions for record %s, got nil", tt.recordCode)
					return
				}
				
				if tt.expectedField != "" {
					if _, exists := fields[tt.expectedField]; !exists {
						t.Errorf("Expected field %s to exist in record %s", tt.expectedField, tt.recordCode)
					}
				}
			} else {
				if fields != nil {
					t.Errorf("Expected no field definitions for record %s, got %v", tt.recordCode, fields)
				}
			}
		})
	}
}

func TestFieldDefinitionConsistency(t *testing.T) {
	// Test specific field position calculations
	tests := []struct {
		recordCode string
		fieldName  string
		wantStart  int
		wantEnd    int
		wantLength int
	}{
		{"H ", "RecordCode", 1, 2, 2},
		{"H ", "InputSystem", 3, 42, 40},
		{"02", "RecordCode", 1, 2, 2},
		{"02", "Amount", 19, 28, 10},
		{"02", "PayeeName", 31, 65, 35},
		{"02", "RoutingNumber", 187, 195, 9},
		{"12", "RecordCode", 1, 2, 2},
		{"12", "Amount", 19, 28, 10},
		{"12", "PayeeName", 31, 65, 35},
		{"13", "Line1", 23, 77, 55},
		{"13", "Line2", 78, 132, 55},
		{"T ", "ScheduleCount", 13, 20, 8},
		{"E ", "TotalCountRecords", 3, 20, 18},
	}

	for _, tt := range tests {
		t.Run(tt.recordCode+"_"+tt.fieldName, func(t *testing.T) {
			fields := GetFieldDefinitions(tt.recordCode)
			if fields == nil {
				t.Fatalf("No field definitions for record %s", tt.recordCode)
			}

			field, exists := fields[tt.fieldName]
			if !exists {
				t.Fatalf("Field %s not found in record %s", tt.fieldName, tt.recordCode)
			}

			if field.Start != tt.wantStart {
				t.Errorf("Field %s start: want %d, got %d", tt.fieldName, tt.wantStart, field.Start)
			}
			if field.End != tt.wantEnd {
				t.Errorf("Field %s end: want %d, got %d", tt.fieldName, tt.wantEnd, field.End)
			}
			if field.Length != tt.wantLength {
				t.Errorf("Field %s length: want %d, got %d", tt.fieldName, tt.wantLength, field.Length)
			}
		})
	}
}

func TestCheckStubFieldPositions(t *testing.T) {
	// Test that check stub lines are positioned correctly
	fields := GetFieldDefinitions("13")
	if fields == nil {
		t.Fatal("No field definitions for check stub record")
	}

	// Verify all 14 lines are present and correctly positioned
	expectedLines := []struct {
		name      string
		wantStart int
		wantEnd   int
	}{
		{"Line1", 23, 77},
		{"Line2", 78, 132},
		{"Line3", 133, 187},
		{"Line4", 188, 242},
		{"Line5", 243, 297},
		{"Line6", 298, 352},
		{"Line7", 353, 407},
		{"Line8", 408, 462},
		{"Line9", 463, 517},
		{"Line10", 518, 572},
		{"Line11", 573, 627},
		{"Line12", 628, 682},
		{"Line13", 683, 737},
		{"Line14", 738, 792},
	}

	for _, line := range expectedLines {
		field, exists := fields[line.name]
		if !exists {
			t.Errorf("Check stub line %s not found", line.name)
			continue
		}

		if field.Start != line.wantStart {
			t.Errorf("Line %s start: want %d, got %d", line.name, line.wantStart, field.Start)
		}
		if field.End != line.wantEnd {
			t.Errorf("Line %s end: want %d, got %d", line.name, line.wantEnd, field.End)
		}
		if field.Length != 55 {
			t.Errorf("Line %s length: want 55, got %d", line.name, field.Length)
		}
	}
}

func TestFieldPositionsNoOverlap(t *testing.T) {
	// Test that no non-filler fields overlap within each record type
	recordTypes := []string{"H ", "01", "02", "03", "04", "11", "12", "13", "G ", "DD", "T ", "E "}

	for _, recordCode := range recordTypes {
		t.Run("Record_"+recordCode, func(t *testing.T) {
			fields := GetFieldDefinitions(recordCode)
			if fields == nil {
				return // Skip if no definitions
			}

			// Create a map of positions to field names
			positionMap := make(map[int]string)

			for fieldName, field := range fields {
				// Skip filler fields in overlap checking
				if strings.Contains(fieldName, "Filler") {
					continue
				}

				for pos := field.Start; pos <= field.End; pos++ {
					if existingField, exists := positionMap[pos]; exists {
						t.Errorf("Position %d in record %s is used by both %s and %s", 
							pos, recordCode, existingField, fieldName)
					}
					positionMap[pos] = fieldName
				}
			}
		})
	}
}

func TestNewFieldDef(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		length   int
		required bool
		wantEnd  int
	}{
		{"single char", 1, 1, true, 1},
		{"two chars", 1, 2, true, 2},
		{"ten chars", 10, 10, false, 19},
		{"large field", 100, 100, false, 199},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := NewFieldDef(tt.start, tt.length, tt.required)
			
			if field.Start != tt.start {
				t.Errorf("Start: want %d, got %d", tt.start, field.Start)
			}
			if field.End != tt.wantEnd {
				t.Errorf("End: want %d, got %d", tt.wantEnd, field.End)
			}
			if field.Length != tt.length {
				t.Errorf("Length: want %d, got %d", tt.length, field.Length)
			}
			if field.Required != tt.required {
				t.Errorf("Required: want %v, got %v", tt.required, field.Required)
			}
		})
	}
}