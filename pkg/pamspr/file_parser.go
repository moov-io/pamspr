package pamspr

import (
	"fmt"
	"strconv"
	"strings"
)

// FileParser handles parsing of file header and trailer records
type FileParser struct {
	validator *Validator
}

// NewFileParser creates a new file parser
func NewFileParser(validator *Validator) *FileParser {
	return &FileParser{
		validator: validator,
	}
}

// ParseFileHeader parses a file header record ("H ")
func (p *FileParser) ParseFileHeader(line string) (*FileHeader, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("H ")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for file header")
	}

	header := &FileHeader{
		RecordCode:               extractField(line, fields["RecordCode"]),
		InputSystem:              extractField(line, fields["InputSystem"]),
		StandardPaymentVersion:   extractField(line, fields["StandardPaymentVersion"]),
		IsRequestedForSameDayACH: extractField(line, fields["IsRequestedForSameDayACH"]),
	}

	// Validate header
	if p.validator != nil {
		if err := p.validator.ValidateFileHeader(header); err != nil {
			return nil, err
		}
	}

	return header, nil
}

// ParseFileTrailer parses a file trailer record ("E ")
func (p *FileParser) ParseFileTrailer(line string) (*FileTrailer, error) {
	if len(line) != RecordLength {
		return nil, fmt.Errorf("invalid record length: expected %d, got %d", RecordLength, len(line))
	}

	fields := GetFieldDefinitions("E ")
	if fields == nil {
		return nil, fmt.Errorf("no field definitions for file trailer")
	}

	trailer := &FileTrailer{
		RecordCode:          extractField(line, fields["RecordCode"]),
		TotalCountRecords:   parseAmount(extractField(line, fields["TotalCountRecords"])),
		TotalCountPayments:  parseAmount(extractField(line, fields["TotalCountPayments"])),
		TotalAmountPayments: parseAmount(extractField(line, fields["TotalAmountPayments"])),
	}

	return trailer, nil
}

// Helper functions
func extractField(line string, field FieldDefinition) string {
	if field.Start > len(line) || field.End > len(line) {
		return ""
	}
	return line[field.Start-1 : field.End]
}

func extractFieldTrimmed(line string, field FieldDefinition) string {
	return strings.TrimSpace(extractField(line, field))
}

func parseAmount(s string) int64 {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}
	amount, _ := strconv.ParseInt(s, 10, 64)
	return amount
}
