# Synthetic PAM SPR Test Data

## Overview

This directory contains **synthetic** (artificially generated) PAM SPR test files created for testing purposes. These files are labeled as synthetic to distinguish them from actual test files that may be obtained from the U.S. Treasury.

## Important Notice

These files are **NOT** real payment files and should **NEVER** be used for actual payment processing. They are created solely for:
- Unit testing the PAM SPR reader/writer
- Validating parser functionality
- Testing agency-specific validation rules
- Development and debugging

## Directory Structure

```
synthetic/
├── valid/                          # Valid synthetic test files
│   ├── synthetic_ach_simple.spr    # Simple ACH payment with one transaction
│   ├── synthetic_check_simple.spr  # Simple check payment with one transaction
│   └── synthetic_multi_schedule.spr # Multiple schedules (ACH + Check)
├── invalid/                        # Invalid synthetic test files (future)
└── agency/                         # Agency-specific synthetic files
    ├── IRS/
    │   └── synthetic_irs_refund.spr     # IRS tax refund example
    ├── VA/
    │   └── synthetic_va_benefit.spr     # VA disability compensation
    ├── SSA/
    │   └── synthetic_ssa_benefit.spr    # SSA retirement benefit
    ├── RRB/
    │   └── synthetic_rrb_annuity.spr    # RRB railroad retirement
    └── CCC/
        └── synthetic_ccc_payment.spr    # CCC agricultural payment
```

## File Naming Convention

All synthetic files follow the naming pattern: `synthetic_[type]_[description].spr`

This clearly identifies them as synthetic test data.

## File Contents

### Valid Test Files

1. **synthetic_ach_simple.spr**
   - Single ACH schedule with one PPD payment
   - Tests basic ACH functionality
   - Amount: $100.00

2. **synthetic_check_simple.spr**
   - Single check schedule with one payment
   - Tests basic check functionality
   - Amount: $25,000.00

3. **synthetic_multi_schedule.spr**
   - Two ACH payments (CCD): $50.00 and $75.00
   - One check payment with stub: $100,000.00
   - Tests multiple schedule handling

### Agency-Specific Files

1. **IRS/synthetic_irs_refund.spr**
   - Tax refund payment via ACH
   - Includes IRS-specific reconcilement data
   - Tests IRS validation rules

2. **VA/synthetic_va_benefit.spr**
   - Veterans disability compensation via check
   - Includes station codes and FIN codes
   - Tests VA-specific validation

3. **SSA/synthetic_ssa_benefit.spr**
   - Social Security retirement benefit via ACH
   - Includes program service center codes
   - Tests SSA validation rules

4. **RRB/synthetic_rrb_annuity.spr**
   - Railroad retirement annuity via ACH
   - Includes employee ID and tier information
   - Tests RRB validation rules

5. **CCC/synthetic_ccc_payment.spr**
   - Agricultural program payment via check
   - Includes commodity and farm details
   - Tests CCC validation rules

## Obtaining Real Test Files

To obtain actual Treasury-approved test files:

1. **For Test Data**: Contact PAM.SAT@fiscal.treasury.gov
2. **For Documentation**: Contact FS.AgencyOutreach@fiscal.treasury.gov
3. **For Support**: Contact pam.help.desk@fiscal.treasury.gov or call 816-414-2340

### Requested Real Test Files

We are seeking real Treasury-approved SPR test files to enhance our test coverage. Please place any obtained files in:

```
testdata/treasury/
├── valid/                    # Valid Treasury test files
│   ├── ach/                 # ACH payment files
│   ├── check/               # Check payment files
│   └── mixed/               # Multi-schedule files
├── invalid/                 # Invalid files for error testing
└── agency/                  # Agency-specific files
    ├── IRS/
    ├── VA/
    ├── SSA/
    ├── RRB/
    └── CCC/
```

**Specifically requesting**:
- Real ACH payment files with various SEC codes (PPD, CCD, CTX)
- Real check payment files with stubs and special handling
- Multi-schedule files combining ACH and check payments
- Agency-specific files for each federal agency
- Invalid files that should fail validation (malformed records, wrong balancing, etc.)
- Files with edge cases (maximum field lengths, special characters, etc.)

These real files will help validate our parser against actual Treasury data formats and ensure compatibility with production systems.

## Usage in Tests

```go
// Example test using synthetic data
func TestSyntheticACHFile(t *testing.T) {
    file, err := pamspr.ReadFile("testdata/synthetic/valid/synthetic_ach_simple.spr")
    if err != nil {
        t.Fatalf("Failed to read synthetic file: %v", err)
    }
    
    // Validate the synthetic file
    if err := file.Validate(); err != nil {
        t.Errorf("Synthetic file validation failed: %v", err)
    }
}
```

## Contributing

When adding new synthetic test files:
1. Always prefix the filename with `synthetic_`
2. Document the test scenario in this README
3. Include realistic but fake data
4. Never use real payment information
5. Follow the SPR specification exactly

## Data Privacy

All data in synthetic files is completely fictional:
- Names are common placeholders (John Doe, Mary Smith, etc.)
- Addresses are generic (123 Main Street, etc.)
- Account numbers are sequential patterns
- TINs/SSNs are repeated digits (never real)
- All amounts are round numbers for easy verification