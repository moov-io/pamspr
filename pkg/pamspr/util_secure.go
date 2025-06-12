package pamspr

import (
	"fmt"
	"strings"
	"unicode"
)

// SecurePadLeft pads a string on the left with the specified character
// Returns error if truncation would occur
func SecurePadLeft(s string, length int, padChar rune, fieldName string) (string, error) {
	if len(s) > length {
		return "", &FieldTruncationError{
			FieldName:      fieldName,
			OriginalValue:  s,
			TruncatedValue: s[:length],
			MaxLength:      length,
		}
	}
	if len(s) == length {
		return s, nil
	}
	return strings.Repeat(string(padChar), length-len(s)) + s, nil
}

// SecurePadRight pads a string on the right with the specified character
// Returns error if truncation would occur
func SecurePadRight(s string, length int, padChar rune, fieldName string) (string, error) {
	if len(s) > length {
		return "", &FieldTruncationError{
			FieldName:      fieldName,
			OriginalValue:  s,
			TruncatedValue: s[:length],
			MaxLength:      length,
		}
	}
	if len(s) == length {
		return s, nil
	}
	return s + strings.Repeat(string(padChar), length-len(s)), nil
}

// SecurePadNumeric pads a numeric string with zeros on the left
// Returns error if the result would exceed the specified length
func SecurePadNumeric(s string, length int, fieldName string) (string, error) {
	// Remove non-numeric characters using strings.Builder for efficiency
	var builder strings.Builder
	builder.Grow(len(s)) // Pre-allocate capacity
	for _, r := range s {
		if unicode.IsDigit(r) {
			builder.WriteRune(r)
		}
	}
	cleaned := builder.String()

	if len(cleaned) > length {
		return "", &FieldTruncationError{
			FieldName:      fieldName,
			OriginalValue:  s,
			TruncatedValue: cleaned[:length],
			MaxLength:      length,
		}
	}

	return SecurePadLeft(cleaned, length, '0', fieldName)
}

// SecureTruncateOrPad truncates or pads a string to the exact length
// Returns error if truncation occurs
func SecureTruncateOrPad(s string, length int, padRight bool, fieldName string) (string, error) {
	if padRight {
		return SecurePadRight(s, length, ' ', fieldName)
	}
	return SecurePadLeft(s, length, ' ', fieldName)
}

// ValidateFieldLength validates that a field meets length requirements
// Returns error with details if validation fails
func ValidateFieldLength(value string, minLength, maxLength int, fieldName string) error {
	length := len(value)

	if length < minLength {
		return fmt.Errorf("%s must be at least %d characters, got %d", fieldName, minLength, length)
	}

	if length > maxLength {
		return &FieldTruncationError{
			FieldName:      fieldName,
			OriginalValue:  value,
			TruncatedValue: value[:maxLength],
			MaxLength:      maxLength,
		}
	}

	return nil
}

// FormatFieldWithValidation formats a field value with length validation
// This is the secure replacement for direct padding functions
func FormatFieldWithValidation(value string, length int, padChar rune, padRight bool, fieldName string) (string, error) {
	// First validate the field won't be truncated
	if len(value) > length {
		return "", &FieldTruncationError{
			FieldName:      fieldName,
			OriginalValue:  value,
			TruncatedValue: value[:length],
			MaxLength:      length,
		}
	}

	// Apply padding as needed
	if padRight {
		return SecurePadRight(value, length, padChar, fieldName)
	}
	return SecurePadLeft(value, length, padChar, fieldName)
}
