# PAM SPR File Format Documentation

This directory contains the extracted and organized documentation for the Payment Automation Manager (PAM) Standard Payment Request (SPR) file format, Version 5.0.2, Release 11.1.0, dated December 17, 2021.

The original document has been split into separate files for easier navigation and reference.

## Table of Contents

### Section 1: Input File Specification

- [1.1 General Instructions](01-general-instructions.md)
- [1.2 General Structure of File](02-general-structure.md)
- [1.3 File Structure Validations](03-file-structure-validations.md)
- [1.4 Hexadecimal Character Validation](04-hexadecimal-character-validation.md)
- [1.5 Validation for Balancing](05-validation-for-balancing.md)
- [1.6 Validation for Same Day ACH (SDA)](06-validation-same-day-ach.md)
- [1.7 Derived Data Elements](07-derived-data-elements.md)
- [1.8 Input Management (IM) Interface Information](08-input-management-interface.md)
- [1.9 Specification Notes](09-specification-notes.md)

### Section 2: File Record Specifications

- [2.1 File Header Record](record-01-file-header.md)
- [2.2 ACH Schedule Header Record](record-02-ach-schedule-header.md)
- [2.3 Check Schedule Header Record](record-03-check-schedule-header.md)
- [2.4 ACH Payment Data Record](record-04-ach-payment-data.md)
- [2.5 Check Payment Data Record](record-05-check-payment-data.md)
- [2.6 ACH Addendum Record](record-06-ach-addendum.md)
- [2.7 CARS TAS/BETC Record](record-07-cars-tasbetc.md)
- [2.8 Check Stub Record](record-08-check-stub.md)
- [2.9 Procurement Record (deleted)](record-09-procurement-deleted.md)
- [2.10 DNP Record](record-10-dnp.md)
- [2.11 Schedule Trailer Control Record](record-11-schedule-trailer.md)
- [2.12 File Trailer Control Record](record-12-file-trailer.md)

### Section 3: Appendices

- [3.1 Appendix A - ACH Transaction Codes](appendix-a-ach-transaction-codes.md)
- [3.2 Appendix B - Agency Specific Values](appendix-b-agency-specific-values.md)
- [3.3 Appendix C - Addressing Reference Information](appendix-c-addressing-reference.md)
- [3.4 Appendix D - Glossary of Terms](appendix-d-glossary.md)
- [3.5 Appendix E - PaymentTypeCode Values](appendix-e-paymenttypecode-values.md)

### Document History

- [Document History](document-history.md) - Complete version history and changes

## Quick Reference

### Key Information
- **File Format**: Fixed-width records of 850 characters
- **Current Version**: 5.0.2
- **Record Types**: H, 01, 02, 03, 04, 11, 12, 13, G, DD, T, E
- **Maximum SDA Amount**: $1,000,000 per payment

### File Structure Overview
```
File Header Record (H)
  ACH Schedule Header Record (01)
    ACH Payment Data Record (02)
    ACH Addendum Record (03 or 04)
    CARS TAS/BETC Record (G)
    DNP Record (DD)
  ACH Schedule Trailer Control Record (T)
  Check Schedule Header Record (11)
    Check Payment Data Record (12)
    Check Stub Record (13)
    CARS TAS/BETC Record (G)
    DNP Record (DD)
  Check Schedule Trailer Control Record (T)
File Trailer Control Record (E)
```

### Important Notes
- Each schedule can contain only one type of payment (ACH or Check)
- ACH Payment Data Records must be in Routing Number order within the schedule
- All records for one payment must be together in the file
- Records associated with Payment Data Records can be received in any order (Addenda, CARS, Stub)

## Original Source
The complete original specification is available in: `InputFileSpecifications-StandardPaymentRequest-5.0.2-v11.1.0-12-17-21.md`

## Related Files
- `.docx` version: `InputFileSpecifications-StandardPaymentRequest-5.0.2-v11.1.0-12-17-21.docx`
- `.pdf` version: `InputFileSpecifications-StandardPaymentRequest-5.0.2-v11.1.0-12-17-21.pdf`