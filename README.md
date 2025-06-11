# PAM SPR - Payment Automation Manager Standard Payment Request

[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/pamspr/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/pamspr)](https://goreportcard.com/report/github.com/moov-io/pamspr)

A Go library for reading, writing, and validating Payment Automation Manager (PAM) Standard Payment Request (SPR) files used by the U.S. Treasury Bureau of the Fiscal Service for federal agency payment processing.

## Overview

PAM SPR files are used by federal agencies to process various types of payments including:

- **Salaries and Benefits**: Federal employee payroll, retirement, and benefit payments
- **Vendor Payments**: Payments to contractors and service providers  
- **Tax Refunds**: Individual and business tax refund processing
- **Social Security**: SSA benefit payments and Medicare reimbursements
- **Other Federal Payments**: Education loans, agricultural subsidies, and more

The library supports both **ACH (Automated Clearing House)** and **Check** payment methods, with comprehensive validation for Same Day ACH requirements and agency-specific rules.

## Features

### Core Functionality
- ✅ **Read PAM SPR Files**: Parse fixed-width 850-character records
- ✅ **Write PAM SPR Files**: Generate compliant payment files with proper formatting  
- ✅ **Validate Files**: Comprehensive validation including balancing and business rules
- ✅ **ACH Support**: CCD, PPD, IAT, and CTX Standard Entry Class codes
- ✅ **Check Support**: Paper check generation with customizable stubs
- ✅ **Same Day ACH**: Validation for expedited ACH processing requirements

### Advanced Features
- ✅ **Agency-Specific Rules**: Built-in support for IRS, SSA, VA, RRB, and CCC agencies
- ✅ **File Builder**: Fluent API for programmatic file construction
- ✅ **Utility Functions**: Field formatting, amount conversion, address cleaning
- ✅ **CLI Tool**: Command-line interface for file operations
- ✅ **Comprehensive Testing**: Extensive test coverage with realistic scenarios

## Installation

```bash
go get github.com/moov-io/pamspr
```

Or install the CLI tool:

```bash
go install github.com/moov-io/pamspr/cmd/pamspr@latest
```

## Quick Start

### Reading a PAM SPR File

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
    // Open the file
    file, err := os.Open("payments.spr")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Parse the file
    reader := pamspr.NewReader(file)
    pamFile, err := reader.Read()
    if err != nil {
        log.Fatal(err)
    }

    // Display file information
    fmt.Printf("Input System: %s\n", pamFile.Header.InputSystem)
    fmt.Printf("Total Payments: %d\n", pamFile.Trailer.TotalCountPayments)
    fmt.Printf("Total Amount: $%.2f\n", float64(pamFile.Trailer.TotalAmountPayments)/100)
}
```

### Creating an ACH Payment File

```go
package main

import (
    "log"
    "os"
    
    "github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
    // Build file using fluent API
    builder := pamspr.NewFileBuilder()
    
    file, err := builder.
        WithHeader("MY_AGENCY_SYSTEM", "502", false).
        StartACHSchedule("SCH001", "Salary", "12345678", "PPD").
        AddACHPayment(&pamspr.ACHPayment{
            AgencyAccountIdentifier:      "EMP001",
            Amount:                       250000, // $2,500.00
            AgencyPaymentTypeCode:        "S",
            IsTOPOffset:                  "1",
            PayeeName:                    "JOHN DOE",
            PayeeAddressLine1:            "123 MAIN ST",
            CityName:                     "WASHINGTON",
            StateCodeText:                "DC",
            PostalCode:                   "20001",
            RoutingNumber:                "021000021",
            AccountNumber:                "1234567890",
            ACHTransactionCode:           "22", // Checking Credit
            PaymentID:                    "PAY001",
            TIN:                          "123456789",
            PaymentRecipientTINIndicator: "1", // SSN
        }).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }

    // Write to file
    output, err := os.Create("ach_payments.spr")
    if err != nil {
        log.Fatal(err)
    }
    defer output.Close()

    writer := pamspr.NewWriter(output)
    if err := writer.Write(file); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("ACH payment file created successfully!")
}
```

### Creating a Check Payment File

```go
package main

import (
    "log"
    "os"
    
    "github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
    builder := pamspr.NewFileBuilder()
    
    file, err := builder.
        WithHeader("CHECK_SYSTEM", "502", false).
        StartCheckSchedule("CHK001", "Vendor", "87654321", "stub").
        AddCheckPayment(&pamspr.CheckPayment{
            AgencyAccountIdentifier:      "VEND001",
            Amount:                       150000, // $1,500.00
            AgencyPaymentTypeCode:        "V",
            IsTOPOffset:                  "1",
            PayeeName:                    "ACME CORP",
            PayeeAddressLine1:            "789 BUSINESS BLVD",
            PayeeAddressLine2:            "SUITE 100",
            CityName:                     "NEW YORK",
            StateCodeText:                "NY",
            PostalCode:                   "10001",
            CheckLegendText1:             "INVOICE #12345",
            CheckLegendText2:             "PO #67890",
            PaymentID:                    "CHK001",
            TIN:                          "123456789",
            PaymentRecipientTINIndicator: "2", // EIN
            Stub: &pamspr.CheckStub{
                RecordCode: "13",
                PaymentID:  "CHK001",
                PaymentIdentificationLines: [14]string{
                    "Invoice Date: 01/15/2024",
                    "Invoice Number: 12345",
                    "Purchase Order: 67890",
                    "Description: Office Supplies",
                    "",
                    "Item Details:",
                    "  Paper - 10 reams @ $50.00 = $500.00",
                    "  Toner - 5 units @ $100.00 = $500.00",
                    "  Misc Supplies = $500.00",
                    "",
                    "Subtotal: $1,500.00",
                    "Tax: $0.00",
                    "Total: $1,500.00",
                    "",
                },
            },
        }).
        Build()
    
    if err != nil {
        log.Fatal(err)
    }

    // Write to file
    output, err := os.Create("check_payments.spr")
    if err != nil {
        log.Fatal(err)
    }
    defer output.Close()

    writer := pamspr.NewWriter(output)
    if err := writer.Write(file); err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Check payment file created successfully!")
}
```

## Validation

The library includes comprehensive validation capabilities:

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
    // Read file
    file, err := os.Open("payments.spr")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := pamspr.NewReader(file)
    pamFile, err := reader.Read()
    if err != nil {
        log.Fatal(err)
    }

    // Create validator
    validator := pamspr.NewValidator()

    // Validate file structure
    if err := validator.ValidateFileStructure(pamFile); err != nil {
        log.Printf("Structure validation failed: %v", err)
    }

    // Validate balancing (totals match)
    if err := validator.ValidateBalancing(pamFile); err != nil {
        log.Printf("Balancing validation failed: %v", err)
    }

    // Validate Same Day ACH requirements
    if err := validator.ValidateSameDayACH(pamFile); err != nil {
        log.Printf("Same Day ACH validation failed: %v", err)
    }

    // Validate individual payments
    for _, schedule := range pamFile.Schedules {
        for _, payment := range schedule.GetPayments() {
            if achPayment, ok := payment.(*pamspr.ACHPayment); ok {
                if err := validator.ValidateACHPayment(achPayment); err != nil {
                    log.Printf("ACH payment validation failed: %v", err)
                }
            }
            if checkPayment, ok := payment.(*pamspr.CheckPayment); ok {
                if err := validator.ValidateCheckPayment(checkPayment); err != nil {
                    log.Printf("Check payment validation failed: %v", err)
                }
            }
        }
    }

    fmt.Println("Validation completed!")
}
```

## CLI Usage

The included command-line tool provides easy file operations:

### Validate a PAM SPR File
```bash
pamspr -validate -input payments.spr
```

### Display File Information
```bash
pamspr -info -input payments.spr
```

### Create Sample Files
```bash
# Create sample ACH file
pamspr -create ach -output sample_ach.spr

# Create sample Check file
pamspr -create check -output sample_check.spr
```

### Convert to JSON (Future Feature)
```bash
pamspr -convert -input payments.spr -output payments.json
```

## File Format Specifications

PAM SPR files use a hierarchical structure with fixed-width records:

### File Structure
```
File Header (H)
├── Schedule Header (01=ACH, 11=Check)
│   ├── Payment Record (02=ACH, 12=Check)
│   │   ├── Addendum Records (03, 04) [Optional]
│   │   ├── CARS TAS/BETC Records (G) [Optional]
│   │   ├── Check Stub (13) [Check Only]
│   │   └── DNP Record (DD) [Optional]
│   └── Schedule Trailer (T)
└── File Trailer (E)
```

### Record Types
- **H**: File Header - Contains system information and processing flags
- **01**: ACH Schedule Header - ACH-specific schedule information
- **11**: Check Schedule Header - Check-specific schedule information  
- **02**: ACH Payment Record - Individual ACH payment details
- **12**: Check Payment Record - Individual check payment details
- **03**: ACH Addendum - Additional payment information (80 chars)
- **04**: CTX Addendum - Corporate Trade Exchange data (800 chars)
- **13**: Check Stub - Check stub information with 14 lines of detail
- **G**: CARS TAS/BETC - Government accounting classification
- **DD**: DNP Record - Do Not Pay verification information
- **T**: Schedule Trailer - Schedule totals and counts
- **E**: File Trailer - File totals and counts

### Key Constraints
- **Fixed Length**: All records are exactly 850 characters
- **Balancing Required**: File and schedule totals must match payment sums
- **Sequential Processing**: Records must appear in hierarchical order
- **Amount Format**: All amounts stored as cents (integers)

## Agency-Specific Features

The library includes specialized support for federal agencies:

### IRS (Internal Revenue Service)
- Tax refund processing with reconcilement fields
- Bond payment handling for Savings Bonds
- Document locator number tracking

### SSA (Social Security Administration)  
- Benefit payment processing
- Program service center codes
- TIN indicator offset handling

### VA (Department of Veterans Affairs)
- Veteran benefit payments
- Station code and financing code support
- Policy number tracking

### RRB (Railroad Retirement Board)
- Railroad retirement benefit processing
- Specialized reconcilement parsing

### CCC (Commodity Credit Corporation)
- Agricultural subsidy payments
- Commodity-specific payment handling

## Same Day ACH Support

The library fully supports Same Day ACH processing with validation for:

- **Amount Limits**: $1,000,000 per transaction limit enforcement
- **Time Windows**: Processing deadline validation
- **SEC Code Restrictions**: Supported Standard Entry Class codes
- **Settlement Requirements**: Same-day settlement validation

## Utility Functions

The library provides helpful utilities for common operations:

### Field Formatting
```go
fp := &pamspr.FieldPadding{}

// Pad left with zeros
padded := fp.PadLeft("123", 5, '0') // "00123"

// Pad right with spaces  
padded = fp.PadRight("ABC", 7, ' ') // "ABC    "

// Extract and pad numeric values
numeric := fp.PadNumeric("ABC123DEF", 5) // "00123"
```

### Amount Handling
```go
fu := &pamspr.FormatUtils{}

// Convert cents to dollar string
dollars := fu.FormatAmount(12345) // "123.45"

// Parse dollar string to cents
cents, err := fu.ParseAmount("$1,234.56") // 123456, nil
```

### Address Cleaning
```go
fu := &pamspr.FormatUtils{}

// Clean problematic characters
clean := fu.CleanAddress(`123 "Main" & <Company>`) // `123 'Main' + (Company)`
```

### TIN Formatting
```go
fu := &pamspr.FormatUtils{}

// Format SSN
ssn := fu.FormatTIN("123456789", "1") // "123-45-6789"

// Format EIN  
ein := fu.FormatTIN("123456789", "2") // "12-3456789"
```

## Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./pkg/pamspr/

# Run specific test
go test -run TestReaderParseFile ./pkg/pamspr/

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Examples

See the `examples/` directory for complete working examples:

```bash
# Run example code
go run examples/examples.go
```

The examples demonstrate:
1. Creating ACH payment files with addenda and CARS records
2. Creating check payment files with stubs
3. Parsing and validating existing files  
4. Same Day ACH validation scenarios

## Contributing

We welcome contributions! Please see our [contributing guidelines](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/moov-io/pamspr.git
cd pamspr

# Run tests
go test ./...

# Build CLI tool
go build -o pamspr cmd/pamspr/main.go

# Format code
go fmt ./...

# Run linter
golangci-lint run
```

## Contributing & Future Work

We welcome contributions to improve the PAM SPR library! Here are areas where help is particularly needed:

### 🚀 **Planned Enhancements**

**Agency-Specific Payment Validation** - Currently, the library includes validation stubs for major federal agencies (VA, SSA, RRB, CCC) but these return "not implemented". We need subject matter experts to help implement the specific validation rules for each agency.

For detailed implementation plans and requirements, see [REFACTORING.md - Phase 5](REFACTORING.md#phase-5-department-specific-payment-validation).

### 📋 **How to Contribute**
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/agency-validation`)
3. Implement changes with comprehensive tests
4. Run all tests and ensure 100% pass rate
5. Submit a pull request with detailed description

**Note**: Agency-specific validation will require coordination with federal agency SMEs to ensure compliance with current requirements.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Support

- **Documentation**: See inline code documentation and examples
- **Issues**: Report bugs or request features via [GitHub Issues](https://github.com/moov-io/pamspr/issues)
- **Community**: Join the [Moov Community](https://community.moov.io) for discussions

## Related Projects

- [moov-io/ach](https://github.com/moov-io/ach) - ACH file processing library
- [moov-io/fed](https://github.com/moov-io/fed) - Federal Reserve routing number lookup
- [moov-io/paygate](https://github.com/moov-io/paygate) - Payment gateway service

---

**Disclaimer**: This library is not affiliated with the U.S. Treasury or any federal agency. It is an open-source implementation of the PAM SPR file format for integration purposes.