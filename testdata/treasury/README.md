# Treasury-Approved SPR Test Files

## Overview

This directory is reserved for **real** Treasury-approved PAM SPR test files obtained from the U.S. Treasury. These files complement the synthetic test files but provide validation against actual production data formats.

## Status

🔴 **NO TREASURY FILES OBTAINED YET**

We are currently seeking Treasury-approved test files. Please contact:
- **PAM.SAT@fiscal.treasury.gov** for test data access
- **FS.AgencyOutreach@fiscal.treasury.gov** for documentation
- **pam.help.desk@fiscal.treasury.gov** or 816-414-2340 for support

## Directory Structure

```
treasury/
├── valid/                    # Valid Treasury test files
│   ├── ach/                 # ACH payment files (PPD, CCD, CTX)
│   ├── check/               # Check payment files with stubs
│   └── mixed/               # Multi-schedule files (ACH + Check)
├── invalid/                 # Invalid files for error testing
└── agency/                  # Agency-specific files
    ├── IRS/                 # Internal Revenue Service files
    ├── VA/                  # Veterans Affairs files  
    ├── SSA/                 # Social Security Administration files
    ├── RRB/                 # Railroad Retirement Board files
    └── CCC/                 # Commodity Credit Corporation files
```

## Requested File Types

### Valid Files
- **ACH Files**: PPD, CCD, and CTX payment files with proper addenda
- **Check Files**: Check payments with stubs and special handling codes
- **Mixed Files**: Files containing both ACH and check schedules
- **Agency Files**: Files specific to each federal agency with their reconcilement patterns
- **Edge Cases**: Files with maximum field lengths, special characters, large amounts

### Invalid Files
- **Format Errors**: Malformed records, incorrect lengths, invalid characters
- **Validation Errors**: Unbalanced files, incorrect totals, missing required fields
- **Business Rule Violations**: Invalid agency codes, prohibited payment combinations

## Usage

Once Treasury files are obtained:

1. **Place files in appropriate directories** based on payment type and agency
2. **Create corresponding test cases** in the test suite
3. **Compare results** with synthetic file parsing to validate accuracy
4. **Document any differences** found between synthetic and real file formats
5. **Update parser** if real files reveal format variations not in specification

## Security Considerations

Treasury test files may contain:
- Redacted but realistic payment data
- Sensitive agency information patterns
- Production-like formatting requirements

Ensure proper handling according to any data use agreements with Treasury.

## Integration with Test Suite

```go
// Example test structure for Treasury files
func TestTreasuryFiles(t *testing.T) {
    treasuryFiles := []string{
        "testdata/treasury/valid/ach/treasury_ach_ppd.spr",
        "testdata/treasury/valid/check/treasury_check_stub.spr",
        "testdata/treasury/agency/IRS/treasury_irs_refund.spr",
    }
    
    for _, filename := range treasuryFiles {
        t.Run(filename, func(t *testing.T) {
            if _, err := os.Stat(filename); os.IsNotExist(err) {
                t.Skip("Treasury file not available yet")
            }
            
            file, err := pamspr.ReadFile(filename)
            if err != nil {
                t.Fatalf("Failed to read Treasury file: %v", err)
            }
            
            if err := file.Validate(); err != nil {
                t.Errorf("Treasury file validation failed: %v", err)
            }
        })
    }
}
```

## Contributing

When Treasury files become available:

1. **Document the source** and date received
2. **Add test cases** that exercise the files
3. **Note any parser changes** required for compatibility
4. **Update documentation** with findings
5. **Preserve file integrity** - do not modify Treasury-provided files

## Priority Request List

High priority Treasury test files needed:

1. **IRS Tax Refund Files** - ACH PPD with reconcilement data
2. **SSA Benefit Files** - Large volume benefit payments
3. **VA Disability Files** - Check payments with special handling
4. **Multi-Agency Files** - Files with payments to multiple agencies
5. **Error Cases** - Files that should fail validation for testing error handling

The goal is to have Treasury files that cover all code paths in our parser and validator.