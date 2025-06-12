package pamspr

import (
	"fmt"
	"strings"
)

// TruncationPolicy defines how to handle field values that are too long
type TruncationPolicy int

const (
	TruncationPolicyError TruncationPolicy = iota // Return error on truncation
	TruncationPolicyWarn                          // Log warning and truncate
	TruncationPolicyAllow                         // Silent truncation (current behavior)
)

// SecurityConfig holds configuration for field security and validation
type SecurityConfig struct {
	TruncationPolicy  TruncationPolicy
	EnableBoundsCheck bool
	EnableLengthCheck bool
}

// DefaultSecurityConfig returns a secure configuration
func DefaultSecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		TruncationPolicy:  TruncationPolicyError, // Secure default
		EnableBoundsCheck: true,
		EnableLengthCheck: true,
	}
}

// LegacySecurityConfig returns a configuration that matches current behavior
func LegacySecurityConfig() *SecurityConfig {
	return &SecurityConfig{
		TruncationPolicy:  TruncationPolicyAllow, // Legacy behavior
		EnableBoundsCheck: true,
		EnableLengthCheck: false,
	}
}

// FieldExtractionError represents an error during field extraction
type FieldExtractionError struct {
	FieldName string
	Line      string
	Start     int
	End       int
	Reason    string
}

func (e *FieldExtractionError) Error() string {
	return fmt.Sprintf("field extraction error for %s: %s (position %d-%d, line length %d)",
		e.FieldName, e.Reason, e.Start, e.End, len(e.Line))
}

// FieldTruncationError represents a field truncation event
type FieldTruncationError struct {
	FieldName      string
	OriginalValue  string
	TruncatedValue string
	MaxLength      int
}

func (e *FieldTruncationError) Error() string {
	return fmt.Sprintf("field %s truncated from %d to %d characters: %q -> %q",
		e.FieldName, len(e.OriginalValue), len(e.TruncatedValue), e.OriginalValue, e.TruncatedValue)
}

// SecureExtractField extracts a field with comprehensive bounds checking and error reporting
func SecureExtractField(line string, field FieldDefinition, fieldName string, config *SecurityConfig) (string, error) {
	if config == nil {
		config = DefaultSecurityConfig()
	}

	// Validate field definition
	if field.Start < 1 {
		return "", &FieldExtractionError{
			FieldName: fieldName,
			Line:      line,
			Start:     field.Start,
			End:       field.End,
			Reason:    "field start position must be >= 1",
		}
	}

	if field.End < field.Start {
		return "", &FieldExtractionError{
			FieldName: fieldName,
			Line:      line,
			Start:     field.Start,
			End:       field.End,
			Reason:    "field end position must be >= start position",
		}
	}

	// Always check bounds to prevent panics, but only return errors if enabled
	boundsError := ""
	if field.Start > len(line) {
		boundsError = fmt.Sprintf("start position %d exceeds line length %d", field.Start, len(line))
	} else if field.End > len(line) {
		boundsError = fmt.Sprintf("end position %d exceeds line length %d", field.End, len(line))
	}

	if boundsError != "" {
		if config.EnableBoundsCheck {
			return "", &FieldExtractionError{
				FieldName: fieldName,
				Line:      line,
				Start:     field.Start,
				End:       field.End,
				Reason:    boundsError,
			}
		} else {
			// Legacy behavior: return empty string for out of bounds
			return "", nil
		}
	}

	// Safe extraction
	return line[field.Start-1 : field.End], nil
}

// SecureExtractFieldTrimmed extracts and trims a field with bounds checking
func SecureExtractFieldTrimmed(line string, field FieldDefinition, fieldName string, config *SecurityConfig) (string, error) {
	value, err := SecureExtractField(line, field, fieldName, config)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value), nil
}

// SecureFormatField formats a field value with truncation policy enforcement
func SecureFormatField(value string, length int, fieldName string, config *SecurityConfig) (string, error) {
	if config == nil {
		config = DefaultSecurityConfig()
	}

	if len(value) > length {
		switch config.TruncationPolicy {
		case TruncationPolicyError:
			return "", &FieldTruncationError{
				FieldName:      fieldName,
				OriginalValue:  value,
				TruncatedValue: value[:length],
				MaxLength:      length,
			}
		case TruncationPolicyWarn:
			// TODO: Add logging framework integration
			// For now, return error with warning context
			return value[:length], &FieldTruncationError{
				FieldName:      fieldName,
				OriginalValue:  value,
				TruncatedValue: value[:length],
				MaxLength:      length,
			}
		case TruncationPolicyAllow:
			return value[:length], nil
		}
	}

	// Pad with spaces if needed
	if len(value) < length {
		return value + strings.Repeat(" ", length-len(value)), nil
	}

	return value, nil
}

// ValidateFieldDefinitions checks field definitions for overlaps and out-of-bounds
func ValidateFieldDefinitions(recordType string, fields map[string]FieldDefinition, maxLength int) error {
	positions := make([]bool, maxLength+1) // 1-indexed

	for fieldName, field := range fields {
		// Validate field boundaries
		if field.Start < 1 {
			return fmt.Errorf("field %s start position %d is invalid (must be >= 1)", fieldName, field.Start)
		}
		if field.End > maxLength {
			return fmt.Errorf("field %s end position %d exceeds record length %d", fieldName, field.End, maxLength)
		}
		if field.Start > field.End {
			return fmt.Errorf("field %s start position %d is greater than end position %d", fieldName, field.Start, field.End)
		}

		// Check for overlaps
		for pos := field.Start; pos <= field.End; pos++ {
			if positions[pos] {
				return fmt.Errorf("field %s position %d overlaps with another field", fieldName, pos)
			}
			positions[pos] = true
		}
	}

	return nil
}
