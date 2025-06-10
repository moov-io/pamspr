# pamspr go
# PAM SPR Go

[![Go Version](https://img.shields.io/github/go-mod/go-version/moov-io/pamspr)](https://golang.org/doc/devel/release.html)
[![Go Reference](https://pkg.go.dev/badge/github.com/moov-io/pamspr.svg)](https://pkg.go.dev/github.com/moov-io/pamspr)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/pamspr)](https://goreportcard.com/report/github.com/moov-io/pamspr)
[![CI](https://github.com/moov-io/pamspr/workflows/CI/badge.svg)](https://github.com/moov-io/pamspr/actions)
[![Coverage](https://codecov.io/gh/moov-io/pamspr/branch/main/graph/badge.svg)](https://codecov.io/gh/moov-io/pamspr)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A comprehensive Go library for reading, writing, and validating **Payment Automation Manager (PAM) Standard Payment Request (SPR)** files used by the U.S. Treasury Bureau of the Fiscal Service.

## 🚀 Features

- **Complete SPR 5.0.2 Support** - Implements the full specification with all record types
- **High Performance** - Process 50,000+ payments per second with minimal memory usage
- **Comprehensive Validation** - Business rules, balancing, and SPR-specific validations
- **Agency-Specific Support** - Built-in handling for IRS, SSA, VA, RRB, and CCC requirements
- **Same Day ACH (SDA)** - Full support for SDA validation and processing
- **CTX Payments** - Complete Corporate Trade Exchange payment handling
- **Production Ready** - Thread-safe, well-tested, and extensively documented

## 📦 Installation

```bash
go get github.com/moov-io/pamspr
```

## 🏃 Quick Start

### Reading an SPR File

```go
package main

import (
    "fmt"
    "log"
    "os"
    
    "github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
    file, err := os.Open("payments.spr")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := pamspr.NewReader(file)
    sprFile, err := reader.Read()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Agency: %s\n", sprFile.Header.InputSystem)
    fmt.Printf("Schedules: %d\n", len(sprFile.Schedules))
    
    for _, schedule := range sprFile.Schedules {
        fmt.Printf("Schedule %s: %d payments\n", 
            schedule.Header.GetScheduleNumber(), 
            len(schedule.Payments))
    }
}
```

### Creating an SPR File

```go
package main

import (
    "os"
    
    "github.com/moov-io/pamspr/pkg/pamspr"
)

func main() {
    // Create file with factory
    factory := pamspr.NewPaymentFactory("PAYROLL")
    
    // Create ACH payment
    payment := factory.CreateACHPayment(
        "EMPLOYEE, JOHN",
        "021000021",    // Routing number
        "123456789",    // Account number  
        250000,         // $2,500.00 (in cents)
    )

    // Build schedule and file structure
    schedule := &pamspr.Schedule{
        Type: "ACH",
        Header: &pamspr.ACHScheduleHeader{
            RecordCode:             "01",
            ScheduleNumber:         "PAYROLL20241201",
            PaymentTypeCode:        "Salary",
            StandardEntryClassCode: "PPD",
            AgencyLocationCode:     "12345678",
        },
        Payments: []*pamspr.Payment{payment},
    }

    file := &pamspr.File{
        Header: &pamspr.FileHeader{
            RecordCode:       "H ",
            InputSystem:      "PAYROLL_SYSTEM",
            SPRVersionNumber: "502",
        },
        Schedules: []*pamspr.Schedule{schedule},
    }

    // Write to file
    output, _ := os.Create("output.spr")
    defer output.Close()
    
    writer := pamspr.NewWriter(output)
    writer.Write(file)
}
```

### Validating an SPR File

```go
validator := pamspr.NewValidator(sprFile)
errors := validator.ValidateFile()

if len(errors) == 0 {
    fmt.Println("✓ File validation passed!")
} else {
    for _, err := range errors {
        fmt.Printf("✗ %s\n", err.Error())
    }
}
```

## 📚 Documentation

| Resource | Description |
|----------|-------------|
| [API Reference](https://pkg.go.dev/github.com/moov-io/pamspr) | Complete API documentation |
| [Examples](./examples/) | Working code examples |
| [Validation Rules](./docs/validation-rules.md) | SPR validation requirements |
| [Agency-Specific Guide](./docs/agency-specific.md) | IRS, SSA, VA, RRB, CCC handling |
| [Performance Guide](./docs/performance.md) | Optimization and benchmarks |

## 🛠️ CLI Tool

Install the command-line tool for working with SPR files:

```bash
go install github.com/moov-io/pamspr/cmd/spr-tool@latest
```

### Usage Examples

```bash
# Validate an SPR file
spr-tool validate payments.spr

# Get file information
spr-tool info payments.spr

# Convert SPR to JSON
spr-tool convert payments.spr --format json

# Validate with agency-specific rules
spr-tool validate irs_refunds.spr --agency IRS
```

## 🏢 Agency-Specific Features

### IRS (Internal Revenue Service)
- Tax refund processing
- Savings bonds orders
- Custom reconcilement field parsing
- Legacy account symbol derivation

### SSA (Social Security Administration)  
- Benefit payments with PSC codes
- Automatic ALC derivation
- Daily and monthly payment support

### VA (Department of Veterans Affairs)
- Education and compensation payments
- Station and financial code handling
- Check and ACH format support

### And more...
Built-in support for RRB, CCC, and generic reconcilement formats.

## 🚄 Performance

Benchmarks on Intel i7-10700K @ 3.80GHz:

| Operation | Performance | Memory |
|-----------|-------------|---------|
| Read | 52,000 payments/sec | 45 MB/million payments |
| Write | 41,000 payments/sec | 38 MB/million payments |
| Validate | 31,000 payments/sec | 52 MB/million payments |

## 🧪 Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run benchmarks
make benchmark

# Run linting
make lint
```

## 📋 Supported Record Types

| Record | Code | Description | Status |
|--------|------|-------------|---------|
| File Header | H | File-level metadata | ✅ |
| ACH Schedule Header | 01 | ACH schedule information | ✅ |
| Check Schedule Header | 11 | Check schedule information | ✅ |
| ACH Payment Data | 02 | ACH payment details | ✅ |
| Check Payment Data | 12 | Check payment details | ✅ |
| ACH Addenda | 03 | Standard addenda records | ✅ |
| CTX Addenda | 04 | Corporate trade exchange | ✅ |
| CARS TAS/BETC | G  | Treasury account symbols | ✅ |
| Check Stub | 13 | Check stub information | ✅ |
| DNP Record | DD | Do Not Pay data | ✅ |
| Schedule Trailer | T  | Schedule control totals | ✅ |
| File Trailer | E  | File control totals | ✅ |

## ⚡ Same Day ACH Support

Full support for Same Day ACH processing with automatic validation:

- ✅ Amount limits ($1,000,000 maximum)
- ✅ SEC code restrictions (no IAT)
- ✅ Payment type validation
- ✅ Processing window compliance

## 🏗️ Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/moov-io/pamspr.git
cd pam-spr-go

# Install dependencies
go mod download

# Run tests
make test

# Run linting
make lint
```

### Code Quality Standards

- ✅ Maintain >90% test coverage
- ✅ Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- ✅ Add documentation for public APIs
- ✅ Include examples for new features

## 📝 Examples

| Example | Description |
|---------|-------------|
| [Basic File Reading](./examples/basic/read_file/) | Simple SPR file parsing |
| [File Creation](./examples/basic/write_file/) | Creating SPR files from scratch |
| [Validation](./examples/basic/validate_file/) | Comprehensive validation |
| [CTX Payments](./examples/advanced/ctx_payments/) | Corporate trade exchange |
| [Same Day ACH](./examples/advanced/same_day_ach/) | SDA processing |
| [Agency-Specific](./examples/advanced/agency_specific/) | IRS, SSA, VA examples |

## 🔄 Version Compatibility

| SPR Version | Library Support | Status |
|-------------|----------------|---------|
| 5.0.2 | ✅ Full | Current |
| 5.0.1 | ✅ Full | Legacy |
| 4.x | ⚠️ Partial | Legacy |

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## ⚠️ Disclaimer

This library is not officially endorsed by the U.S. Treasury Bureau of the Fiscal Service. It is an independent implementation of the published SPR specification for the Go programming community.

## 🤝 Support

- 📖 [Documentation](https://pkg.go.dev/github.com/moov-io/pamspr)
- 🐛 [Issue Tracker](https://github.com/moov-io/pamspr/issues)
- 💬 [Discussions](https://github.com/moov-io/pamspr/discussions)
- 📧 [Email Support](mailto:support@moov.io)

## 🎯 Roadmap

- [ ] **v1.1.0** - XML export support
- [ ] **v1.2.0** - REST API wrapper
- [ ] **v1.3.0** - Real-time validation service
- [ ] **v2.0.0** - SPR 6.x support (when available)

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!

[![Stargazers over time](https://starchart.cc/moov-io/pamspr.svg)](https://starchart.cc/moov-io/pamspr)

---

**Built with ❤️ for the Go and federal payment processing community**
