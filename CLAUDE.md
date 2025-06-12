# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go library for working with Payment Automation Manager (PAM) Standard Payment Request (SPR) files used by the U.S. Treasury. It handles reading, writing, and validating payment files for federal agencies processing salaries, benefits, refunds, and vendor payments.

## Technical Specification

The authoritative source of truth for the PAM SPR file format is defined in the structured documentation in the `./docs/` folder. The specification has been organized into separate, easily-referenced files:

**Core Specification Files:**
- `./docs/README.md` - Complete index and navigation guide
- `./docs/01-general-instructions.md` through `./docs/09-specification-notes.md` - General requirements
- `./docs/record-01-file-header.md` through `./docs/record-12-file-trailer.md` - Record specifications  
- `./docs/appendix-a-ach-transaction-codes.md` through `./docs/appendix-e-paymenttypecode-values.md` - Reference data

**IMPORTANT**: When implementing any feature, parsing logic, validation rules, or data structures, ALWAYS refer to the relevant specification files in `./docs/` first. These files contain the complete technical details including:
- Field positions, lengths, and data types
- Validation rules and error codes
- Agency-specific requirements
- Record structure and ordering requirements

When in doubt about any implementation detail, consult the appropriate specification file before proceeding.

## Development Commands

```bash
# Build the CLI tool
go build -o pamspr cmd/pamspr/main.go

# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for a specific package
go test ./pkg/pamspr/

# Run a specific test
go test -run TestReaderParseFile ./pkg/pamspr/

# Install the CLI tool globally
go install ./cmd/pamspr

# Format code
go fmt ./...

# Run the example code
go run examples/examples.go
```

## Architecture

The codebase follows a hierarchical data model:
- **File** contains multiple Schedules
- **Schedule** contains multiple Payments  
- **Payment** can be either ACH or Check type

Key components:
- `pkg/pamspr/file.go`: Core data structures and record types
- `pkg/pamspr/reader.go`: Parses SPR files with lookahead capability
- `pkg/pamspr/writer.go`: Formats and writes valid SPR files
- `pkg/pamspr/validator.go`: Enforces business rules and balancing
- `cmd/pamspr/main.go`: CLI tool for validation and conversion

## Important Implementation Details

1. **Fixed Format**: SPR files use fixed-width records of 850 characters. All fields must be properly padded.

2. **Record Types**: Files contain hierarchical records (H, 01, 02, 03, 04, 11, 12, 13, G, DD, T, E) that must appear in specific order.

3. **Validation**: The validator enforces file structure, balancing (totals must match), and agency-specific rules. Always validate files after modifications.

4. **Agency Support**: Special handling exists for IRS, SSA, VA, RRB, and CCC agencies with specific validation rules.

5. **Same Day ACH**: SDA payments have additional validation requirements including time windows and amount limits.

## Testing Approach

- Unit tests exist for all major components (*_test.go files)
- Test data uses realistic payment scenarios
- Tests validate both successful parsing and error conditions
- Use table-driven tests for multiple scenarios