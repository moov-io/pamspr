package pamspr

import (
	"fmt"
	"strings"
	"testing"
)

func TestSecureExtractField(t *testing.T) {
	tests := []struct {
		name        string
		line        string
		field       FieldDefinition
		fieldName   string
		config      *SecurityConfig
		expected    string
		expectError bool
		errorType   string
	}{
		{
			name:      "valid field extraction",
			line:      "ABCDEFGHIJ",
			field:     FieldDefinition{Start: 1, End: 5},
			fieldName: "TestField",
			config:    DefaultSecurityConfig(),
			expected:  "ABCDE",
		},
		{
			name:        "start position too small",
			line:        "ABCDEFGHIJ",
			field:       FieldDefinition{Start: 0, End: 5},
			fieldName:   "TestField",
			config:      DefaultSecurityConfig(),
			expectError: true,
			errorType:   "*pamspr.FieldExtractionError",
		},
		{
			name:        "end before start",
			line:        "ABCDEFGHIJ",
			field:       FieldDefinition{Start: 5, End: 3},
			fieldName:   "TestField",
			config:      DefaultSecurityConfig(),
			expectError: true,
			errorType:   "*pamspr.FieldExtractionError",
		},
		{
			name:        "start position exceeds line length",
			line:        "ABCDE",
			field:       FieldDefinition{Start: 10, End: 15},
			fieldName:   "TestField",
			config:      DefaultSecurityConfig(),
			expectError: true,
			errorType:   "*pamspr.FieldExtractionError",
		},
		{
			name:        "end position exceeds line length",
			line:        "ABCDE",
			field:       FieldDefinition{Start: 1, End: 10},
			fieldName:   "TestField",
			config:      DefaultSecurityConfig(),
			expectError: true,
			errorType:   "*pamspr.FieldExtractionError",
		},
		{
			name:      "legacy mode allows out of bounds",
			line:      "ABCDE",
			field:     FieldDefinition{Start: 1, End: 10},
			fieldName: "TestField",
			config:    &SecurityConfig{EnableBoundsCheck: false},
			expected:  "", // Would panic in real scenario, but test framework may handle
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SecureExtractField(tt.line, tt.field, tt.fieldName, tt.config)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}

				// Check error type if specified
				if tt.errorType != "" {
					errorTypeName := fmt.Sprintf("%T", err)
					if !strings.Contains(errorTypeName, "FieldExtractionError") {
						t.Errorf("expected error type containing FieldExtractionError, got %s", errorTypeName)
					}
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
					return
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func TestSecureFormatField(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		length      int
		fieldName   string
		config      *SecurityConfig
		expected    string
		expectError bool
		errorType   string
	}{
		{
			name:      "exact length",
			value:     "ABCDE",
			length:    5,
			fieldName: "TestField",
			config:    DefaultSecurityConfig(),
			expected:  "ABCDE",
		},
		{
			name:      "pad with spaces",
			value:     "ABC",
			length:    5,
			fieldName: "TestField",
			config:    DefaultSecurityConfig(),
			expected:  "ABC  ",
		},
		{
			name:        "truncation error policy",
			value:       "ABCDEFGH",
			length:      5,
			fieldName:   "TestField",
			config:      DefaultSecurityConfig(),
			expectError: true,
			errorType:   "*pamspr.FieldTruncationError",
		},
		{
			name:      "truncation allow policy",
			value:     "ABCDEFGH",
			length:    5,
			fieldName: "TestField",
			config:    LegacySecurityConfig(),
			expected:  "ABCDE",
		},
		{
			name:      "truncation warn policy",
			value:     "ABCDEFGH",
			length:    5,
			fieldName: "TestField",
			config:    &SecurityConfig{TruncationPolicy: TruncationPolicyWarn},
			expected:  "ABCDE",
			// Note: In real implementation, this would also log a warning
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SecureFormatField(tt.value, tt.length, tt.fieldName, tt.config)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}

				// Check error type if specified
				if tt.errorType != "" {
					errorTypeName := fmt.Sprintf("%T", err)
					if !strings.Contains(errorTypeName, "FieldTruncationError") {
						t.Errorf("expected error type containing FieldTruncationError, got %s", errorTypeName)
					}
				}
			} else {
				if err != nil && tt.config.TruncationPolicy != TruncationPolicyWarn {
					t.Errorf("unexpected error: %v", err)
					return
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func TestValidateFieldDefinitions(t *testing.T) {
	tests := []struct {
		name        string
		recordType  string
		fields      map[string]FieldDefinition
		maxLength   int
		expectError bool
		errorMsg    string
	}{
		{
			name:       "valid fields",
			recordType: "01",
			fields: map[string]FieldDefinition{
				"Field1": {Start: 1, End: 5},
				"Field2": {Start: 6, End: 10},
			},
			maxLength: 20,
		},
		{
			name:       "overlapping fields",
			recordType: "01",
			fields: map[string]FieldDefinition{
				"Field1": {Start: 1, End: 5},
				"Field2": {Start: 3, End: 8},
			},
			maxLength:   20,
			expectError: true,
			errorMsg:    "overlaps",
		},
		{
			name:       "field starts at position 0",
			recordType: "01",
			fields: map[string]FieldDefinition{
				"Field1": {Start: 0, End: 5},
			},
			maxLength:   20,
			expectError: true,
			errorMsg:    "invalid",
		},
		{
			name:       "field end exceeds record length",
			recordType: "01",
			fields: map[string]FieldDefinition{
				"Field1": {Start: 1, End: 25},
			},
			maxLength:   20,
			expectError: true,
			errorMsg:    "exceeds",
		},
		{
			name:       "start greater than end",
			recordType: "01",
			fields: map[string]FieldDefinition{
				"Field1": {Start: 10, End: 5},
			},
			maxLength:   20,
			expectError: true,
			errorMsg:    "greater than",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateFieldDefinitions(tt.recordType, tt.fields, tt.maxLength)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
					return
				}
				if tt.errorMsg != "" && !strings.Contains(err.Error(), tt.errorMsg) {
					t.Errorf("expected error containing %q, got %q", tt.errorMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

func TestFieldExtractionError(t *testing.T) {
	err := &FieldExtractionError{
		FieldName: "TestField",
		Line:      "ABCDEFGHIJ",
		Start:     1,
		End:       15,
		Reason:    "end position exceeds line length",
	}

	expected := "field extraction error for TestField: end position exceeds line length (position 1-15, line length 10)"
	if err.Error() != expected {
		t.Errorf("expected error message %q, got %q", expected, err.Error())
	}
}

func TestFieldTruncationError(t *testing.T) {
	err := &FieldTruncationError{
		FieldName:      "TestField",
		OriginalValue:  "ABCDEFGH",
		TruncatedValue: "ABCDE",
		MaxLength:      5,
	}

	expected := "field TestField truncated from 8 to 5 characters: \"ABCDEFGH\" -> \"ABCDE\""
	if err.Error() != expected {
		t.Errorf("expected error message %q, got %q", expected, err.Error())
	}
}

func TestSecurityConfigurations(t *testing.T) {
	t.Run("default config is secure", func(t *testing.T) {
		config := DefaultSecurityConfig()
		if config.TruncationPolicy != TruncationPolicyError {
			t.Error("default config should use error policy for truncation")
		}
		if !config.EnableBoundsCheck {
			t.Error("default config should enable bounds checking")
		}
	})

	t.Run("legacy config preserves old behavior", func(t *testing.T) {
		config := LegacySecurityConfig()
		if config.TruncationPolicy != TruncationPolicyAllow {
			t.Error("legacy config should allow truncation")
		}
		if !config.EnableBoundsCheck {
			t.Error("legacy config should still enable bounds checking for safety")
		}
	})
}
