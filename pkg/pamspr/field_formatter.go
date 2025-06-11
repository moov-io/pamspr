package pamspr

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// FieldFormatter handles automatic field formatting using field definitions and struct tags
type FieldFormatter struct {
	validator *Validator
}

// NewFieldFormatter creates a new field formatter
func NewFieldFormatter(validator *Validator) *FieldFormatter {
	return &FieldFormatter{
		validator: validator,
	}
}

// FormatterConfig defines formatting options for a field
type FormatterConfig struct {
	FieldName     string
	Length        int
	Required      bool
	FormatType    FieldFormatType
	PadChar       rune
	Justification FieldJustification
}

// FieldFormatType represents different field formatting types
type FieldFormatType string

const (
	FormatText      FieldFormatType = "text"    // Default text formatting
	FormatNumeric   FieldFormatType = "numeric" // Zero-padded numbers
	FormatAmount    FieldFormatType = "amount"  // Currency amounts in cents
	FormatFiller    FieldFormatType = "filler"  // Filler spaces
	FormatNoJustify FieldFormatType = "nojust"  // No justification
)

// FieldJustification represents field alignment
type FieldJustification string

const (
	JustifyLeft  FieldJustification = "left"
	JustifyRight FieldJustification = "right"
	JustifyNone  FieldJustification = "none"
)

// FormatRecord formats a struct into a fixed-width record using field definitions
func (f *FieldFormatter) FormatRecord(record interface{}, recordCode string) (string, error) {
	// Get field definitions for this record type
	fieldDefs := GetFieldDefinitions(recordCode)
	if fieldDefs == nil {
		return "", fmt.Errorf("no field definitions found for record code: %s", recordCode)
	}

	// Create output buffer
	output := make([]byte, RecordLength)
	for i := range output {
		output[i] = ' ' // Initialize with spaces
	}

	// Use reflection to iterate over struct fields
	recordValue := reflect.ValueOf(record)
	recordType := reflect.TypeOf(record)

	// Handle pointer types
	if recordValue.Kind() == reflect.Ptr {
		recordValue = recordValue.Elem()
		recordType = recordType.Elem()
	}

	// Process each struct field
	for i := 0; i < recordValue.NumField(); i++ {
		field := recordValue.Field(i)
		fieldType := recordType.Field(i)

		// Skip unexported fields
		if !field.CanInterface() {
			continue
		}

		// Get field name (use tag if present, otherwise struct field name)
		fieldName := fieldType.Name
		if tag := fieldType.Tag.Get("pamspr"); tag != "" {
			fieldName = tag
		}

		// Skip fields not in field definitions
		fieldDef, exists := fieldDefs[fieldName]
		if !exists {
			continue
		}

		// Format the field value
		formattedValue, err := f.formatFieldValue(field.Interface(), fieldDef, fieldType)
		if err != nil {
			return "", fmt.Errorf("formatting field %s: %w", fieldName, err)
		}

		// Place formatted value in correct position
		start := fieldDef.Start - 1 // Convert to 0-based
		end := fieldDef.End
		if start < 0 || end > len(output) {
			return "", fmt.Errorf("field %s position out of bounds: %d-%d", fieldName, fieldDef.Start, fieldDef.End)
		}

		copy(output[start:end], []byte(formattedValue))
	}

	return string(output), nil
}

// formatFieldValue formats a single field value based on its type and field definition
func (f *FieldFormatter) formatFieldValue(value interface{}, fieldDef FieldDefinition, fieldType reflect.StructField) (string, error) {
	// Get formatting configuration from struct tag or defaults
	config := f.getFormatterConfig(fieldDef, fieldType)

	// Convert value to string
	stringValue := f.valueToString(value)

	// Apply formatting based on type
	switch config.FormatType {
	case FormatNumeric:
		return f.formatNumeric(stringValue, config.Length), nil
	case FormatAmount:
		if intValue, ok := value.(int64); ok {
			return f.formatAmount(intValue, config.Length), nil
		}
		return f.formatNumeric(stringValue, config.Length), nil
	case FormatFiller:
		return strings.Repeat(" ", config.Length), nil
	case FormatNoJustify:
		return f.formatFieldNoJustify(stringValue, config.Length), nil
	case FormatText:
		if config.Justification == JustifyRight {
			return f.formatFieldRightJustified(stringValue, config.Length, config.PadChar), nil
		}
		return f.formatField(stringValue, config.Length), nil
	default:
		// This should never happen if all enum values are handled above
		return "", fmt.Errorf("unknown format type: %s", config.FormatType)
	}
}

// getFormatterConfig extracts formatting configuration from struct tags and field definition
func (f *FieldFormatter) getFormatterConfig(fieldDef FieldDefinition, fieldType reflect.StructField) FormatterConfig {
	config := FormatterConfig{
		FieldName:     fieldType.Name,
		Length:        fieldDef.Length,
		Required:      fieldDef.Required,
		FormatType:    FormatText,
		PadChar:       ' ',
		Justification: JustifyLeft,
	}

	// Parse struct tag for formatting options
	if tag := fieldType.Tag.Get("format"); tag != "" {
		parts := strings.Split(tag, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			switch {
			case part == "numeric":
				config.FormatType = FormatNumeric
				config.Justification = JustifyRight
				config.PadChar = '0'
			case part == "amount":
				config.FormatType = FormatAmount
				config.Justification = JustifyRight
				config.PadChar = '0'
			case part == "filler":
				config.FormatType = FormatFiller
			case part == "nojust":
				config.FormatType = FormatNoJustify
			case part == "right":
				config.Justification = JustifyRight
			case part == "left":
				config.Justification = JustifyLeft
			case strings.HasPrefix(part, "pad="):
				if len(part) > 4 {
					config.PadChar = rune(part[4])
				}
			}
		}
	}

	return config
}

// valueToString converts any value to a string
func (f *FieldFormatter) valueToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		if v {
			return "1"
		}
		return "0"
	default:
		return fmt.Sprintf("%v", value)
	}
}

// Helper formatting methods (reuse existing writer methods)
func (f *FieldFormatter) formatField(value string, length int) string {
	if len(value) > length {
		return value[:length]
	}
	return value + strings.Repeat(" ", length-len(value))
}

func (f *FieldFormatter) formatFieldRightJustified(value string, length int, padChar rune) string {
	value = strings.TrimSpace(value)
	if len(value) > length {
		return value[:length]
	}
	return strings.Repeat(string(padChar), length-len(value)) + value
}

func (f *FieldFormatter) formatFieldNoJustify(value string, length int) string {
	if len(value) > length {
		return value[:length]
	}
	if len(value) < length {
		return value + strings.Repeat(" ", length-len(value))
	}
	return value
}

func (f *FieldFormatter) formatNumeric(value string, length int) string {
	// Remove non-numeric characters
	numeric := ""
	for _, r := range value {
		if r >= '0' && r <= '9' {
			numeric += string(r)
		}
	}

	if len(numeric) > length {
		return numeric[:length]
	}
	return strings.Repeat("0", length-len(numeric)) + numeric
}

func (f *FieldFormatter) formatAmount(cents int64, length int) string {
	return f.formatNumeric(fmt.Sprintf("%d", cents), length)
}
