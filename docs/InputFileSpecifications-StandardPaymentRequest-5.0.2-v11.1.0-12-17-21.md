Payment Automation Manager (PAM)

Input File Specifications -- Standard Payment Request

Version 5.0.2

Release 11.1.0

December 17, 2021

# Table of Contents

Table of Contents ii

Document History iv

1 Input File Specification -- Standard Payment Request 1

1.1 General Instructions 1

1.2 General Structure of File 1

1.3 File Structure Validations 3

1.4 Hexadecimal Character Validation 4

1.5 Validation for Balancing 4

1.6 Validation for Same Day ACH (SDA) 6

1.7 Derived Data Elements 6

1.8 Input Management (IM) Interface Information 6

1.9 Specification Notes 7

2 File Record Specifications 8

2.1 File Header Record 8

2.2 ACH Schedule Header Record 10

2.3 Check Schedule Header Record 14

2.4 ACH Payment Data Record 19

2.5 Check Payment Data Record 34

2.6 ACH Addendum Record 50

*2.6.1* CTX Validation Rules 52

*2.6.2* ACH Addendum Record for CTX payments 53

2.7 CARS TAS/BETC Record 55

2.8 Check Stub Record 60

2.9 Procurement Record (deleted) 62

2.10 DNP Record 63

2.11 Schedule Trailer Control Record 64

2.12 File Trailer Control Record 66

3 Appendices 68

3.1 Appendix A - ACH Transaction Codes 68

3.2 Appendix B -- Agency Specific Values 69

*3.2.1* For Custom Agency Rule ID = "IRS" and Depositor Account Number
NOT equal "BONDS" 69

*3.2.2* For Custom Agency Rule ID = "IRS" and Depositor Account Number =
"BONDS" 72

*3.2.3* For Custom Agency Rule ID = "VA" or "VACP" and Check payments 74

*3.2.4* For Custom Agency Rule ID = "VA" or "VACP" and ACH payments 75

*3.2.5* For Custom Agency Rule ID = "SSA" and "SSA-Daily" 77

*3.2.6* For Custom Agency Rule ID = "SSA-A" 78

*3.2.7* For Custom Agency Rule ID = "SSA", "SSA-A", and "SSA-Daily" 78

*3.2.8* For Custom Agency Rule ID = "RRB" 79

*3.2.9* For Custom Agency Rule ID = "CCC" 81

*3.2.10* Generic Reconcilement Field 82

3.3 Appendix C - Addressing Reference Information 83

3.4 Appendix D - Glossary of Terms 86

3.5 Appendix E - PaymentTypeCode Values 90

4 Document History Continued 91

**Error! Hyperlink reference not valid.** **Error! Hyperlink reference
not valid.** **Error! Hyperlink reference not valid.** **Error!
Hyperlink reference not valid.** **Error! Hyperlink reference not
valid.** **Error! Hyperlink reference not valid.** **Error! Hyperlink
reference not valid.** **Error! Hyperlink reference not valid.**
**Error! Hyperlink reference not valid.** **Error! Hyperlink reference
not valid.** **Error! Hyperlink reference not valid.** **Error!
Hyperlink reference not valid.** **Error! Hyperlink reference not
valid.** **Error! Hyperlink reference not valid.** **Error! Hyperlink
reference not valid.** **Error! Hyperlink reference not valid.**
**Error! Hyperlink reference not valid.** **Error! Hyperlink reference
not valid.** **Error! Hyperlink reference not valid.** **Error!
Hyperlink reference not valid.** **Error! Hyperlink reference not
valid.** **Error! Hyperlink reference not valid.** **Error! Hyperlink
reference not valid.** **Error! Hyperlink reference not valid.**
**Error! Hyperlink reference not valid.** **Error! Hyperlink reference
not valid.** **Error! Hyperlink reference not valid.** **Error!
Hyperlink reference not valid.** **Error! Hyperlink reference not
valid.** **Error! Hyperlink reference not valid.** **Error! Hyperlink
reference not valid.** **Error! Hyperlink reference not valid.**
**Error! Hyperlink reference not valid.** **Error! Hyperlink reference
not valid.** **Error! Hyperlink reference not valid.** **Error!
Hyperlink reference not valid.** **Error! Hyperlink reference not
valid.** **Error! Hyperlink reference not valid.** **Error! Hyperlink
reference not valid.** **Error! Hyperlink reference not valid.**
**Error! Hyperlink reference not valid.** **Error! Hyperlink reference
not valid.** **Error! Hyperlink reference not valid.** **Error!
Hyperlink reference not valid.**

# Document History

+---------------+-------------+--------------------------------------------------------------+------------+
| **Version     | **Author**  | **Summary**                                                  | **Date of  |
| Number**      |             |                                                              | Version**  |
+:=============:+=============+==============================================================+============+
| 11.1.0        | Linda       | CR 16925                                                     | 12-17-2021 |
|               | Calder      |                                                              |            |
|               |             | 2.4 ACH Payment Data Record                                  |            |
|               |             |                                                              |            |
|               |             | 1.  Updated the Payee Identifier Additional field to match   |            |
|               |             |     the agency version. Removed secondary and change to      |            |
|               |             |     additional.                                              |            |
|               |             |                                                              |            |
|               |             | 2.  Updated the PayeeName Additional field to match the      |            |
|               |             |     agency version. Removed secondary and changed to         |            |
|               |             |     additional.                                              |            |
|               |             |                                                              |            |
|               |             | 3.  Updated the Additional Payee TIN Indicator field to      |            |
|               |             |     match the agency version. Removed secondary and changed  |            |
|               |             |     to additional.                                           |            |
|               |             |                                                              |            |
|               |             | Appendix D: Glossary of Terms                                |            |
|               |             |                                                              |            |
|               |             | > 4\. Updated the field name from "Party" to "Payee".        |            |
|               |             | >                                                            |            |
|               |             | > 5\. Updated the field name from "Party Name Secondary" to  |            |
|               |             | > "Payee Name Additional" and updated the Definition section |            |
|               |             | > from "secondary" to "additional".                          |            |
|               |             |                                                              |            |
|               |             | Little DR -- Appendix D: Glossary of Terms                   |            |
|               |             |                                                              |            |
|               |             | > 6\. Updated the Standard Entry Class Code field by adding  |            |
|               |             | > information for CTX.                                       |            |
|               |             | >                                                            |            |
|               |             | > 7\. Removed from the Availability Type Code "and merged    |            |
|               |             | > surplus accounts (M)" to match the agency version.         |            |
|               |             | >                                                            |            |
|               |             | > 8\. Removed from the Reconcilement field "your servicing   |            |
|               |             | > RFC for additional information" and updated with "Treasury |            |
|               |             | > for additional information" to match the agency version.   |            |
|               |             | >                                                            |            |
|               |             | > 9\. Removed references to "Procurement" as agencies are    |            |
|               |             | > not using the data.                                        |            |
|               |             |                                                              |            |
|               |             | CR PAM-17645                                                 |            |
|               |             |                                                              |            |
|               |             | 2.9 Procurement Record                                       |            |
|               |             |                                                              |            |
|               |             | > 10\. Deleted the record. Agencies are not using this data. |            |
|               |             |                                                              |            |
|               |             | Little DR -- Removed reference to "Procurement" as agencies  |            |
|               |             | are not using the data.                                      |            |
|               |             |                                                              |            |
|               |             | 1.2 General Structure of File                                |            |
|               |             |                                                              |            |
|               |             | > 11\. Removed references to "Procurement".                  |            |
|               |             |                                                              |            |
|               |             | 1.3 File Structure Validations                               |            |
|               |             |                                                              |            |
|               |             | > 12\. Removed reference to "Procurement".                   |            |
|               |             |                                                              |            |
|               |             | Little DR -- Glossary of Terms Country Name.                 |            |
|               |             |                                                              |            |
|               |             | > 13\. Removed "Used for Check Payments".                    |            |
|               |             |                                                              |            |
|               |             | PAM-20855 -- PAM Changes to Handle Increased Dollar Amount   |            |
|               |             | Level                                                        |            |
|               |             |                                                              |            |
|               |             | > 14\. Section 1.6 Validation for Same Day ACH (SDA) --      |            |
|               |             | > updated maximum payment amount to \$1,000,000.             |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 10.2.0        | Debbie      | PAM-19724 Documentation Updates for adding PICOE as a new    | 2021-07-01 |
|               | Jones       | SRF subscriber.                                              |            |
|               |             |                                                              |            |
|               |             | Change 1: Added "SRF" designation in Downstream Mapping      |            |
|               |             | column of [File Record                                       |            |
|               |             | Specification](#ach-schedule-header-record) for all fields   |            |
|               |             | passed to the SRF version 2.2.1.                             |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 9.2.0         | Linda       | CR 3274                                                      | 2020-01-06 |
|               | Calder      |                                                              |            |
|               |             | 1.  [Section 2.5 Check Payment Data                          |            |
|               |             |     Record](#check-payment-data-record):                     |            |
|               |             |                                                              |            |
|               |             | > Updated the Amount field to identify zero dollar checks as |            |
|               |             | > invalid payments.                                          |            |
|               |             |                                                              |            |
|               |             | 2.  [Section 1.6 Validation for Same Day ACH                 |            |
|               |             |     (SDA):](#validation-for-same-day-ach-sda)                |            |
|               |             |                                                              |            |
|               |             | > Updated the individual payment amounts maximum from        |            |
|               |             | > \$25,000 to \$100,000.                                     |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 9.0.1         | Cecelia     | CR 3297                                                      | 2019-6-17  |
|               | Walsh       |                                                              |            |
|               |             | 1.  [File Header](#file-header-record)                       |            |
|               |             |                                                              |            |
|               |             | <!-- -->                                                     |            |
|               |             |                                                              |            |
|               |             | 1.  Is Requested Same Day ACH: Added notes to support PAM    |            |
|               |             |     behavior when Agencies provide a "blank" value. Also     |            |
|               |             |     included reference to SDA Validations for Balancing in   |            |
|               |             |     Error Codes.                                             |            |
|               |             |                                                              |            |
|               |             | CR 2939 -- Document Only Update                              |            |
|               |             |                                                              |            |
|               |             | 2.  [Section 1.5 Validation for                              |            |
|               |             |     Balancing](#validation-for-balancing): Modified Error    |            |
|               |             |     Code for zero dollar CTX Prenote Transaction Codes to    |            |
|               |             |     match production.                                        |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 9.0.0         | Joshua      | CR-3101                                                      | 2018-12-15 |
|               | Stout       |                                                              |            |
|               |             | [2.1 File Header](#file-header-record)                       |            |
|               |             |                                                              |            |
|               |             | 1.  Added 'IsRequestedForSameDayACH' field to allow          |            |
|               |             |     customers to request Same Day ACH processing.            |            |
|               |             |                                                              |            |
|               |             | [1.6 Validation for Same Day ACH                             |            |
|               |             | (SDA)](#validation-for-same-day-ach-sda)                     |            |
|               |             |                                                              |            |
|               |             | 2.  Added validation and error reasons applicable when the   |            |
|               |             |     SPR payments have been requested for SDA.                |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| []{#v830      | Cecelia     | CR 2524                                                      | 2018-04-02 |
| .anchor}8.3.0 | Walsh/Linda |                                                              |            |
|               | Calder      | 1.  [Section 2.1 File Header Record:](#file-header-record)   |            |
|               |             |     Updated SPR version to "502".                            |            |
|               |             |                                                              |            |
|               |             | 2.  [Throughout the SPR Records](#file-header-record):       |            |
|               |             |     Numbering scheme changes to support Agency Outreach      |            |
|               |             |     training.                                                |            |
|               |             |                                                              |            |
|               |             | CR 2862                                                      |            |
|               |             |                                                              |            |
|               |             | 3.  [Section 2.4 ACH Payment Data                            |            |
|               |             |     Record:](#ach-payment-data-record)                       |            |
|               |             |                                                              |            |
|               |             | - Added Payment Description Code for TOP salary offsetting   |            |
|               |             |   purposes.                                                  |            |
|               |             |                                                              |            |
|               |             | - Reduced Filler to account for added payment description    |            |
|               |             |   code.                                                      |            |
|               |             |                                                              |            |
|               |             | 4.  [Section 2.5 Check Payment Data                          |            |
|               |             |     Record](#check-payment-data-record):                     |            |
|               |             |                                                              |            |
|               |             | - Added Payment Description Code for TOP salary offsetting   |            |
|               |             |   purposes.                                                  |            |
|               |             |                                                              |            |
|               |             | - Reduced Filler to account for added payment description    |            |
|               |             |   code.                                                      |            |
|               |             |                                                              |            |
|               |             | 5.  [Section 3.4 Appendix D -- Glossary of                   |            |
|               |             |     Terms](#appendix-d---glossary-of-terms): Added           |            |
|               |             |     description for new Payment Description code.            |            |
|               |             |                                                              |            |
|               |             | Document Only Change                                         |            |
|               |             |                                                              |            |
|               |             | 6.  [Section 3.2.7.2. Payment Legacy Account Symbol          |            |
|               |             |     Derivation                                               |            |
|               |             |     Rules:](#payment-legacy-account-symbol-derivation-rules) |            |
|               |             |                                                              |            |
|               |             | - Modified verbiage from "Record the derived value in the    |            |
|               |             |   Check Payment Detail's Legacy Symbol attribute" to "Map    |            |
|               |             |   the derived value to the Legacy Account Symbol attribute   |            |
|               |             |   for ACH and Check".                                        |            |
|               |             |                                                              |            |
|               |             | CR 1938                                                      |            |
|               |             |                                                              |            |
|               |             | 7.  [Section 2.4 ACH Payment Data Record (Various            |            |
|               |             |     Fields):](#ach-payment-data-record)                      |            |
|               |             |                                                              |            |
|               |             | - Added \"DoD-DCAS" to "Downstream Mapping" Column.          |            |
|               |             |                                                              |            |
|               |             | 8.  [Section 2.5 Check Payment Data Record (Various          |            |
|               |             |     Fields):](#check-payment-data-record)                    |            |
|               |             |                                                              |            |
|               |             | - Added "DoD-DCAS" to "Downstream Mapping" Column.           |            |
|               |             |                                                              |            |
|               |             | 9.  [Section 2.10 DNP Record (DNP Detail                     |            |
|               |             |     Field):](#dnp-record)                                    |            |
|               |             |                                                              |            |
|               |             | - Added "DoD-DCAS" to "Downstream Mapping" Column.           |            |
|               |             |                                                              |            |
|               |             | 10. [Section 2.9 Procurement Record (Various                 |            |
|               |             |     Fields):](#procurement-record-deleted)                   |            |
|               |             |                                                              |            |
|               |             | - Added "DoD-DCAS" to "Downstream Mapping" Column.           |            |
|               |             |                                                              |            |
|               |             | 11.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.2 ACH Schedule Header Record (Various           |            |
|               |             |   Fields):](#ach-schedule-header-record)Added "DoD-DCAS" to  |            |
|               |             |   "Downstream Mapping" Column.                               |            |
|               |             |                                                              |            |
|               |             | 12.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.3 Check Schedule Header (Various                |            |
|               |             |   Fields):](#check-schedule-header-record)Added "DoD-DCAS"   |            |
|               |             |   to "Downstream Mapping" Column.                            |            |
|               |             |                                                              |            |
|               |             | 13. [Section 2.9 Procurement Record (Various                 |            |
|               |             |     Fields):](#procurement-record-deleted)                   |            |
|               |             |                                                              |            |
|               |             | - Added "DNP" to "Downstream Mapping" Column.                |            |
|               |             |                                                              |            |
|               |             | 14. [Section 2.6 Addendum Record (Various                    |            |
|               |             |     Fields):](#ach-addendum-record)                          |            |
|               |             |                                                              |            |
|               |             | - Added "DNP" to "Downstream Mapping" Column.                |            |
|               |             |                                                              |            |
|               |             | 15.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.9 Procurement Records (Various                  |            |
|               |             |   Fields):](#procurement-record-deleted)Removed "IPP" from   |            |
|               |             |   "Downstream Mapping" Column.                               |            |
|               |             |                                                              |            |
|               |             | 16.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.7 CARS TAS/BETC Record (Various                 |            |
|               |             |   Fields):](#cars-tasbetc-record)Removed "IPP" from          |            |
|               |             |   "Downstream Mapping" Column.                               |            |
|               |             |                                                              |            |
|               |             | - Added "DoD-DCAS" to "Downstream Mapping"                   |            |
|               |             |                                                              |            |
|               |             | CR 2345                                                      |            |
|               |             |                                                              |            |
|               |             | 17. [Section 2.4 ACH Payment Data                            |            |
|               |             |     Record](#ach-payment-data-record):                       |            |
|               |             |                                                              |            |
|               |             | - Replaced verbiage "If blank default to 1. with "Any value  |            |
|               |             |   other than 0 will be defaulted to a 1."                    |            |
|               |             |                                                              |            |
|               |             | 18. [Section 2.5 Check Payment Data                          |            |
|               |             |     Record:](#check-payment-data-record)                     |            |
|               |             |                                                              |            |
|               |             | - Replaced verbiage "If blank default to 1" with "Any value  |            |
|               |             |   other than 0 will be defaulted to a 1."                    |            |
|               |             |                                                              |            |
|               |             | Document Only Change                                         |            |
|               |             |                                                              |            |
|               |             | 19. [Section 2.4 ACH Payment Data                            |            |
|               |             |     Record:](#ach-payment-data-record)                       |            |
|               |             |                                                              |            |
|               |             | - Updated reference to Fiscal Service Country Codes with new |            |
|               |             |   email.                                                     |            |
|               |             |                                                              |            |
|               |             | - Updated to reference Fiscal Service email with new domain  |            |
|               |             |   to reference Geo Codes.                                    |            |
|               |             |                                                              |            |
|               |             | 20.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.5 Check Payment Data                            |            |
|               |             |   Record:](#check-payment-data-record) Updated to reference  |            |
|               |             |   Fiscal Service email with new domain to reference Geo      |            |
|               |             |   Codes.                                                     |            |
|               |             |                                                              |            |
|               |             | 21. [Section 2.2 ACH Schedule Header                         |            |
|               |             |     Record:](#ach-schedule-header-record)                    |            |
|               |             |                                                              |            |
|               |             | - Added note to the Schedule Number field "Any lower-case    |            |
|               |             |   values received on input are stored as upper-case".        |            |
|               |             |                                                              |            |
|               |             | 22. [Section 2.3 Check Schedule Header                       |            |
|               |             |     Record:](#check-schedule-header-record)                  |            |
|               |             |                                                              |            |
|               |             | - Added note to the Schedule Number field "Any lower-case    |            |
|               |             |   values received on input are stored as upper-case".        |            |
|               |             |                                                              |            |
|               |             | CR 2925                                                      |            |
|               |             |                                                              |            |
|               |             | 23.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.4 ACH Payment Data                              |            |
|               |             |   Record](#ach-payment-data-record)Updated the Payment       |            |
|               |             |   Recipient TIN indicator Field Value to include "3"-ITIN as |            |
|               |             |   a selection.                                               |            |
|               |             |                                                              |            |
|               |             | - Updated the Payment Recipient TIN indicator Validation     |            |
|               |             |   Rules to include the number "3" which represents the ITIN  |            |
|               |             |   value.                                                     |            |
|               |             |                                                              |            |
|               |             | - Updated the Secondary Payee TIN Indicator Field Value to   |            |
|               |             |   include "3"-ITIN as a selection.                           |            |
|               |             |                                                              |            |
|               |             | - Updated the Secondary Payee TIN Indicator Validation Rules |            |
|               |             |   to include the number "3" which represents the ITIN value. |            |
|               |             |                                                              |            |
|               |             | 24.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.5 Check Payment Data                            |            |
|               |             |   Record](#check-payment-data-record)Updated the Payment     |            |
|               |             |   Recipient TIN indicator Field Value to include "3"-ITIN as |            |
|               |             |   a selection.                                               |            |
|               |             |                                                              |            |
|               |             | - Updated the Payment Recipient TIN indicator Validation     |            |
|               |             |   Rules to include the number "3" which represents the ITIN  |            |
|               |             |   value.                                                     |            |
|               |             |                                                              |            |
|               |             | - Updated the Secondary Payee TIN Indicator Field Value to   |            |
|               |             |   include "3"-ITIN as a selection.                           |            |
|               |             |                                                              |            |
|               |             | - Updated the Secondary Payee TIN Indicator Validation Rules |            |
|               |             |   to include the number "3" which represents the ITIN value. |            |
|               |             |                                                              |            |
|               |             | 25. [Section 3.4 Appendix                                    |            |
|               |             |     D-Glossary](#appendix-d---glossary-of-terms)             |            |
|               |             |                                                              |            |
|               |             | - Updated to reflect the addition of the ITIN.               |            |
|               |             |                                                              |            |
|               |             | Document Only Change                                         |            |
|               |             |                                                              |            |
|               |             | 26.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.4 ACH Payment Data                              |            |
|               |             |   Record](#ach-payment-data-record)Added note to include     |            |
|               |             |   that the Agency and Treasury agreed to the value for the   |            |
|               |             |   Subpayment Type Code.                                      |            |
|               |             |                                                              |            |
|               |             | 27.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.5 Check Payment Data                            |            |
|               |             |   Record](#check-payment-data-record)Added note to include   |            |
|               |             |   that the Agency and Treasury agreed to the value for the   |            |
|               |             |   Subpayment Type Code.                                      |            |
|               |             |                                                              |            |
|               |             | 28. [Section 2.4 ACH Payment Data                            |            |
|               |             |     Record](#ach-payment-data-record)                        |            |
|               |             |                                                              |            |
|               |             | - Updated note to clarify description of Secondary Payee TIN |            |
|               |             |   Indicator.                                                 |            |
|               |             |                                                              |            |
|               |             | 29.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.5 Check Payment Data                            |            |
|               |             |   Record](#check-payment-data-record)Updated note to clarify |            |
|               |             |   description of Secondary Payee TIN Indicator.              |            |
|               |             |                                                              |            |
|               |             | 30. [Appendix A: ACH Transaction                             |            |
|               |             |     Codes](#appendix-a---ach-transaction-codes)              |            |
|               |             |                                                              |            |
|               |             | - Updated Transaction Code 24 to include "with remittance"   |            |
|               |             |                                                              |            |
|               |             | - Updated Transaction Code 34 to include "with remittance'.  |            |
|               |             |                                                              |            |
|               |             | 31.                                                          |            |
|               |             |                                                              |            |
|               |             | - [Section 2.6.2 ACH Addendum Record for CTX                 |            |
|               |             |   Payments](#ach-addendum-record-for-ctx-payments)Added      |            |
|               |             |   wording to contact Agency Outreach for supplemental        |            |
|               |             |   information.                                               |            |
|               |             |                                                              |            |
|               |             | 32. [Appendix D: Glossary of                                 |            |
|               |             |     Terms](#appendix-d---glossary-of-terms)                  |            |
|               |             |                                                              |            |
|               |             | - Updated the PaymentTypeCode to reference Appendix E        |            |
|               |             |   instead of Appendix A.                                     |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 7.4.2         | Cecelia     | CR 1913, AMUQ 3/1/2016                                       | 2016-03-23 |
|               | Walsh       |                                                              |            |
|               |             | [Section 2.4](#ach-payment-data-record) and [Section         |            |
|               |             | 2.5](#check-payment-data-record): Modified Payer Mechanism   |            |
|               |             | value, "Book Entry" to be "BookEntry" to align with other    |            |
|               |             | similar values.                                              |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 7.4.1         | Cecelia     | CR 1721 :                                                    | 2016-01-26 |
|               | Walsh       |                                                              |            |
|               |             | **Section 1.5 Validation for Balancing:**                    |            |
|               |             |                                                              |            |
|               |             | Added rule for allowed Transaction Codes when dollar amount  |            |
|               |             | is greater than zero and SEC = CTX.                          |            |
|               |             |                                                              |            |
|               |             | Doc Only Changes:                                            |            |
|               |             |                                                              |            |
|               |             | **Section 2.2 ACH Schedule Header Record:**                  |            |
|               |             |                                                              |            |
|               |             | - Schedule Number: Added "all" in the event Schedule number  |            |
|               |             |   is less than 14 characters                                 |            |
|               |             |                                                              |            |
|               |             | - Payment Type Code: Added "all" clarification when payment  |            |
|               |             |   type is less than 25 characters                            |            |
|               |             |                                                              |            |
|               |             | **Section 2.3 Check Schedule Header Record:**                |            |
|               |             |                                                              |            |
|               |             | - Schedule Number: Added "all" in the event Schedule number  |            |
|               |             |   is less than 14 characters                                 |            |
|               |             |                                                              |            |
|               |             | - Payment Type Code: Added "all" clarification when payment  |            |
|               |             |   type is less than 25 characters                            |            |
|               |             |                                                              |            |
|               |             | **Section 2.4 ACH Payment Data Record:**                     |            |
|               |             |                                                              |            |
|               |             | - Amount: Updated Validation rules to reference Section 1.5  |            |
|               |             |   Validation for Balancing for all amount cross validations. |            |
|               |             |                                                              |            |
|               |             | - Payer Mechanism: Updated blank is an allowed value for     |            |
|               |             |   optional Payer Mechanism field.                            |            |
|               |             |                                                              |            |
|               |             | **Section 2.5 Check Payment Data Record:**                   |            |
|               |             |                                                              |            |
|               |             | Payer Mechanism: Updated blank is an allowed value for       |            |
|               |             | optional Payer Mechanism field.                              |            |
+---------------+-------------+--------------------------------------------------------------+------------+
| 7.4.0         | Ashley      | For CR 235: **CR 235 for release 7.4 SPR Enhancements:**     | 2015-10-19 |
|               | Shirk       |                                                              |            |
|               |             | **[Section 2.4 ACH Payment Data Record:]{.underline}**       |            |
|               |             |                                                              |            |
|               |             | [Renamed fields to align with SRF field naming]{.underline}  |            |
|               |             |                                                              |            |
|               |             | [Added Sub Payment Type Code, Payee Address Lines 3&4,       |            |
|               |             | Country Name, Consular Code, and Payer                       |            |
|               |             | Mechanism]{.underline}                                       |            |
|               |             |                                                              |            |
|               |             | **[Section 2.5 Check Payment Data Record:]{.underline}**     |            |
|               |             |                                                              |            |
|               |             | [Renamed fields to align with SRF field naming]{.underline}  |            |
|               |             |                                                              |            |
|               |             | [Added Sub Payment Type Code and Payer                       |            |
|               |             | Mechanism]{.underline}                                       |            |
|               |             |                                                              |            |
|               |             | **[Appendix D-Glossary of Terms:]{.underline}**              |            |
|               |             |                                                              |            |
|               |             | [Added new field definitions]{.underline}                    |            |
+---------------+-------------+--------------------------------------------------------------+------------+

# Input File Specification -- Standard Payment Request 

This document describes the standard format, data elements, validations
and agency specific information that agencies must use to submit payment
requests to the Treasury Bureau of Fiscal Service Payment Automation
Manager (PAM).

The Standard PAM Payment Request Format (SPR) is intended for agencies
to submit one file with one or more schedules. It is not intended to
allow for more than one file per schedule (i.e., one file can have
multiple schedules but one schedule cannot span across multiple files.)
In addition, one SPS summary certification should be submitted for each
[schedule]{.underline} in the file. One certification cannot cover
multiple schedules.

When a SPR is received by Treasury, its schedules are matched to a
certification from SPS based on the schedule number, ALC, amount and
item count. When preparing your SPR, be mindful of these matching
criteria.

Each field in the format has validation rules and indicates whether a
violation results in a file, schedule or payment rejection. It should be
noted that, until further notice, any schedule rejection will result in
a file rejection. In the future, Treasury will pursue offering a
configuration by agency to allow schedule or file resubmission.

## General Instructions

The development version of the Standard Payment Request may differ from
the version published to agencies. The current published version is
noted in the validation rules for the Standard Payment Request Version
field.

Agencies should:

- Ensure content of ACH files is in alignment with NACHA rules.

- Follow IAT rules for potential international payments.

- Comply with OFAC (Office of Foreign Assets Control) policies.

- Ensure payee data is correct so the transaction is payable (routing
  numbers, account numbers, mailing address, etc.).

- Comply with CARS reporting rules for TAS/BETC.

## General Structure of File

Each Record on the File shall be a length of 850 positions. The Records
must indicate an appropriate record code ("RC") and be structured as
indicated below. Records associated to the Check Payment Data Record and
the ACH Payment Data Record can be received in any order (i.e. Addenda,
CARS, Stub). All records for one payment must be together in the file.

- **File Header Record** RC=H (There can only be one file header record)

<!-- -->

- **ACH** Schedule Header Record RC=01(one - many Schedule Header
  Records in a file)

<!-- -->

- ACH Payment Data Record RC=02 (one - many ACH Payment Data Records in
  a schedule)

<!-- -->

- ACH Addendum Record RC=03

<!-- -->

- none -- one for each CCD or PPD Payment Data Record;

- none -- two for each IAT Payment Data Record

Or

- ACH Addendum Record RC=04

<!-- -->

- one-999 for each CTX Payment Data Record

<!-- -->

- CARS TAS/BETC Record RC=G (none-many for each ACH Payment Data Record)

- 

- DNP Record RC=DD (none- one for each ACH Payment Data Record)

<!-- -->

- **ACH** Schedule Trailer Control Record RC=T (one for each Schedule
  Header Record)

- **Check** Schedule Header Record RC=11 (one - many Schedule Header
  Records in a file)

<!-- -->

- Check Payment Data Record RC=12 (one - many Check Payment Data Records
  in a schedule)

<!-- -->

- Check Stub Record RC=13 (none unless CheckEnclosureCode = stub, then
  one for each Check Payment Data Record)

- CARS TAS/BETC Record RC=G (none - many for each Check Payment Data
  Record)

- 

- DNP Record RC=DD (none- one for each Check Payment Data Record)

<!-- -->

- **Check** Schedule Trailer Control Record RC=T (one for each Schedule
  Header Record)

- **File Trailer Control Record** RC=E (There can only be one File
  Trailer Control Record)

## File Structure Validations

The file structure must follow these rules:

  ------------------------------------------------------------------------
  **Rule**                      **Result**              **Error Code**
  ----------------------------- ----------------------- ------------------
  A schedule can contain only   If not, reject the file Error Reason Group
  one type of payment                                   1 Message 6

  If any schedule on a file is  File rejection          Error Reason Group
  rejected, the entire file                             1 Message 6
  will be rejected.                                     

  ACH Payment Data Records must If not, reject the file Error Reason Group
  be received in Routing Number                         1 Message 7
  order within the schedule.                            

  Check Payment Data Records                            
  can be received in any order.                         

  Records associated to the                             
  Check Payment Data and ACH                            
  Payment Data Records can be                           
  received in any order (i.e.                           
  Addenda, CARS, Stub).                                 

  All records for one payment   If not, reject the file Error Reason Group
  must have the same                                    1 Message 6
  information in the payment                            
  identifier fields.                                    
  ------------------------------------------------------------------------

Field types are as follows:

- A=alphabetic; A-Z, a-z, blanks, and special characters as noted below.

- N=numeric; 0-9.

- All dollar amount fields are numeric.

- AN= alphanumeric; A-Z, a-z, 0-9, blanks, and special characters as
  noted below.

![](media/image1.wmf){width="6.079861111111111in"
height="1.9291666666666667in"}

**Table 1**: Allowed Characters

## Hexadecimal Character Validation

  -----------------------------------------------------------------------
  **Rule**                **Result**              **Error Code**
  ----------------------- ----------------------- -----------------------
  Any data elements being If not, reject the file Error Reason Group 1
  validated and/or stored                         Message 5
  in PAM must be valid                            
  HEX characters.                                 

  -----------------------------------------------------------------------

- File formats must meet the following validation:

<!-- -->

- For alphanumeric and alphabetic fields, valid special characters are
  defined as all Hexadecimal characters with values greater than HEX
  "3F" (as shown in Table 1).

- Invalid Hexadecimal characters are defined as HEX values "00" through
  "3F".

## Validation for Balancing

+----------------------------------+--------------------+---------------+
| **Rule**                         | **Result**         | **Error       |
|                                  |                    | Code**        |
+==================================+====================+===============+
| If the ACH_Transaction Code      | If not, reject the | Error Reason  |
| indicates "Prenote" all of the   | file.              | Group 4       |
| payments in the schedule must be |                    | Message 5     |
| zero dollar amounts.             |                    |               |
+----------------------------------+--------------------+---------------+
| If zero dollar payment amounts   | If not, reject the | Error Reason  |
| are received in a schedule, the  | file               | Group 4       |
| ACH Transaction Code must        |                    | Message 4     |
| indicate "Prenote" or the SEC    |                    |               |
| code = CTX.                      |                    |               |
+----------------------------------+--------------------+---------------+
| If greater than zero dollar      | If not, reject the | Error Reason  |
| payment amounts are received in  | file.              | Group 4       |
| the schedule, the ACH            |                    |               |
| Transaction Code must ***not***  |                    | Message 3     |
| indicate "Zero Dollar Credit"    |                    |               |
| when SEC code = CTX.             |                    |               |
+----------------------------------+--------------------+---------------+
| The                              | If not, reject the | Error Reason  |
| ScheduleTrailer.ScheduleCount    | schedule           | Group 3       |
| must equal the accumulated       |                    | Message 6 for |
| number of Payments in the        |                    | ACH           |
| schedule.                        |                    | schedules.    |
|                                  |                    | Error Reason  |
|                                  |                    | Group 3       |
|                                  |                    | Message 4 for |
|                                  |                    | check         |
|                                  |                    | schedules.    |
+----------------------------------+--------------------+---------------+
| The                              | If not, reject the | Error Reason  |
| ScheduleTrailer.ScheduleAmount   | schedule           | Group 3       |
| must equal the accumulated       |                    | Message 5 for |
| amount of Payments in the        |                    | ACH           |
| schedule.                        |                    | schedules.    |
|                                  |                    | Error Reason  |
|                                  |                    | Group 3       |
|                                  |                    | Message 3 for |
|                                  |                    | check         |
|                                  |                    | schedules.    |
+----------------------------------+--------------------+---------------+
| The                              | If not, reject the | Error Reason  |
| FileTrailer.TotalCount_Records   | file               | Group 3       |
| must equal the accumulated       |                    | Message 2     |
| number of records on the file    |                    |               |
| (this includes header and        |                    |               |
| trailer records).                |                    |               |
+----------------------------------+--------------------+---------------+
| The                              | If not, reject the | Error Reason  |
| FileTrailer.TotalCount_Payments  | file               | Group 3       |
| must equal the accumulated       |                    | Message 2     |
| number of payments on the file.  |                    |               |
+----------------------------------+--------------------+---------------+
| The                              | If not, reject the | Error Reason  |
| FileTrailer.TotalAmount_Payments | file               | Group 3       |
| must equal the accumulated       |                    | Message 1     |
| amount of payments on the file.  |                    |               |
+----------------------------------+--------------------+---------------+

##  Validation for Same Day ACH (SDA) 

+----------------------------+----------------------+-----------------+
| **Rule**                   | **Result**           | **Error Code**  |
+============================+======================+=================+
| The SPR can only contain   | If not, and the      | Error Reason    |
| schedules with method of   | IsRequestedForSDA    | Group 4         |
| payment ACH                | value is "1", reject |                 |
|                            | the file.            | Message 7       |
+----------------------------+----------------------+-----------------+
| All individual payment     | If not, and the      | Error Reason    |
| amounts must be less than  | IsRequestedForSDA    | Group 4         |
| or equal to the MAX SDA    | value is "1", reject |                 |
| Amount (\$1,000,000)       | the file.            | Message 8       |
+----------------------------+----------------------+-----------------+
| Payment Type values must   | If not, and the      | Error Reason    |
| be allowed for SDA.        | IsRequestedForSDA    | Group 4         |
|                            | value is "1", reject |                 |
| Restricted Payment Types:  | the file.            | Message 9       |
|                            |                      |                 |
| None                       |                      |                 |
+----------------------------+----------------------+-----------------+
| SEC code must be a value   | If not, and the      | Error Reason    |
| other than IAT             | IsRequestedForSDA    | Group 4         |
|                            | value is "1", reject |                 |
|                            | the file.            | Message 10      |
+----------------------------+----------------------+-----------------+

## Derived Data Elements

  -------------------------------------------------------------------------
  **Data        **Derived From**   **Associated     **Associated Values**
  Element**                        With**           
  ------------- ------------------ ---------------- -----------------------
  Method of     Type of data       Schedule         ACH, Check
  Payment       records received                    

  Object Line 3 Enclosure Code     Check Payment    If Enclosure Code =
                                   Data Record      "stub" then Object Line
                                                    3 = "Per Enclosed
                                                    Mailing Notice"
  -------------------------------------------------------------------------

## Input Management (IM) Interface Information

Note: This section is subject to change in support of transitioning away
from IM.

+----------------------+-------------------------+---------------------+
| **Path**             | **Value**               | **Notes**           |
+======================+=========================+=====================+
| Original Filename    | FROXK.Agency.SPR.Unique | Dataset name        |
|                      |                         | including           |
|                      | Unique=Specified by the | delimiters is up to |
|                      | agency to make the      | 44 characters long. |
|                      | dataset unique for the  | Each node can       |
|                      | day.                    | contain up to 8     |
|                      |                         | characters.         |
|                      |                         |                     |
|                      |                         | If the same dataset |
|                      |                         | name is used within |
|                      |                         | a day, the previous |
|                      |                         | dataset will be     |
|                      |                         | overwritten.        |
+----------------------+-------------------------+---------------------+
| ControlNumber        | Cnnnnnn                 | Assigned by IM      |
+----------------------+-------------------------+---------------------+

## Specification Notes

- Field #s are purely for referential and discussion purposes; they are
  not part of the file's data.

- A "filler" or "blank" field is not validated.

- If PAM performs validations on a field, the validation rules and
  resulting error messages are provided in the "Validation rules" and
  "Error code" columns.

- Verbiage for the Error Messages is contained in the *SPD 102 PRF
  Validation Messages* document.

- If there is no validation rule ("n/a"), the field can contain any
  value including blanks

- Numeric fields that are not validated will be stored as zeros if
  blanks are received.

- Fields with dollar amounts do not have decimal points; the last two
  digits are the cents value. For example 1234567890 = \$12,345,678.90.

- Unless otherwise specified in the validation rules column, numeric
  fields are right justified, with zero (0) pad on the left; alpha and
  alphanumeric fields are left justified, with blank pad on the right.

- When left/right justification rules are listed in the validation rules
  column, PAM shall correct the field justification if needed. Unless
  otherwise specified, currency values are not formatted as zoned
  decimal.

# File Record Specifications

## File Header Record

+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **File Header Record**                                                                                                                                                                                                                             |
+----------------+--------------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------------+----------------+----------------+----------------+
| ***\#***       | **Field Name**           | **Type**       | **Field        | **Length**     | **Start**      | **End          | **Validation   | **Error Code**                                | **Stored       | **Notes**      | **Downstream   |
|                |                          |                | Value**        |                |                | Position**     | rules**        |                                               | Name**         |                | Mapping**      |
|                |                          |                |                |                | **Position**   |                |                |                                               |                |                |                |
+================+==========================+================+================+================+================+================+================+===============================================+================+================+================+
| H.01           | Record Code              | AN             | "H~~b~~"       | 2              | 1              | 2              | If invalid,    | Invalid: Reason Group 1 message 6. Missing or | n/a            | File Header    |                |
|                |                          |                |                |                |                |                | missing or out | out of order: Reason Group 1 message 4.       |                | Record         |                |
|                |                          |                |                |                |                |                | of order,      |                                               |                | Identifier =   |                |
|                |                          |                |                |                |                |                | reject the     |                                               |                | "H" followed   |                |
|                |                          |                |                |                |                |                | file           |                                               |                | by a space     |                |
+----------------+--------------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------------+----------------+----------------+----------------+
| H.02           | InputSystem              | AN             |                | 40             | 3              | 42             | This should    | Reason Group 1 message 6                      | Originating    | Used to        | DNP            |
|                |                          |                |                |                |                |                | match the      |                                               | Agency         | identify the   |                |
|                |                          |                |                |                |                |                | Agency         |                                               | Identifier     | sending entity | PPS            |
|                |                          |                |                |                |                |                | Identifier in  |                                               |                | (agency or     |                |
|                |                          |                |                |                |                |                | the PAM Agency |                                               |                | automated      |                |
|                |                          |                |                |                |                |                | Profile. If    |                                               |                | system)        |                |
|                |                          |                |                |                |                |                | not, reject    |                                               |                |                |                |
|                |                          |                |                |                |                |                | the file.      |                                               |                |                |                |
+----------------+--------------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------------+----------------+----------------+----------------+
| H.03           | Standard Payment Request | AN             |                | 3              | 43             | 45             | If invalid or  | Reason Group 1 message 6                      | n/a            | Identifies the |                |
|                | Version Number           |                |                |                |                |                | missing,       |                                               |                | version of the |                |
|                |                          |                |                |                |                |                | reject the     |                                               |                | Payment        |                |
|                |                          |                |                |                |                |                | file           |                                               |                | Request format |                |
|                |                          |                |                |                |                |                |                |                                               |                | used to        |                |
|                |                          |                |                |                |                |                | **Current      |                                               |                | prepare the    |                |
|                |                          |                |                |                |                |                | published      |                                               |                | PRF, e.g. 502  |                |
|                |                          |                |                |                |                |                | version:**     |                                               |                |                |                |
|                |                          |                |                |                |                |                |                |                                               |                |                |                |
|                |                          |                |                |                |                |                | **502**        |                                               |                |                |                |
+----------------+--------------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------------+----------------+----------------+----------------+
| H.04           | IsRequestedForSameDayACH | AN             | 0, 1           | 1              | 46             | 46             | n/a            | See Section [1.6                              | Is Requested   | 0 = No         | SSF            |
|                |                          |                |                |                |                |                |                | Validation](#validation-for-same-day-ach-sda) | for Same Day   |                |                |
|                |                          |                |                |                |                |                |                | for Same Day ACH                              | ACH            | 1 = Yes        |                |
|                |                          |                |                |                |                |                |                |                                               |                |                |                |
|                |                          |                |                |                |                |                |                |                                               |                | Blank = No     |                |
|                |                          |                |                |                |                |                |                |                                               |                |                |                |
|                |                          |                |                |                |                |                |                |                                               |                | Indicates if   |                |
|                |                          |                |                |                |                |                |                |                                               |                | all payments   |                |
|                |                          |                |                |                |                |                |                |                                               |                | are requested  |                |
|                |                          |                |                |                |                |                |                |                                               |                | for Same Day   |                |
|                |                          |                |                |                |                |                |                |                                               |                | ACH            |                |
|                |                          |                |                |                |                |                |                |                                               |                | processing.    |                |
|                |                          |                |                |                |                |                |                |                                               |                | When "Blank",  |                |
|                |                          |                |                |                |                |                |                |                                               |                | SDA validation |                |
|                |                          |                |                |                |                |                |                |                                               |                | will not       |                |
|                |                          |                |                |                |                |                |                |                                               |                | occur.         |                |
|                |                          |                |                |                |                |                |                |                                               |                |                |                |
|                |                          |                |                |                |                |                |                |                                               |                | Any value      |                |
|                |                          |                |                |                |                |                |                |                                               |                | other than 1   |                |
|                |                          |                |                |                |                |                |                |                                               |                | will be        |                |
|                |                          |                |                |                |                |                |                |                                               |                | defaulted to a |                |
|                |                          |                |                |                |                |                |                |                                               |                | 0.             |                |
+----------------+--------------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------------+----------------+----------------+----------------+
| H.05           | Filler                   |                |                | 804            | 47             | 850            | n/a            |                                               | n/a            |                |                |
+----------------+--------------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------------+----------------+----------------+----------------+

## ACH Schedule Header Record

+---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **ACH Schedule Header Record**                                                                                                                                                                                                              |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| **\#**         | **Field Name**         | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**                                | **Downstream   |
|                |                        |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                                          | Mapping**      |
+================+========================+================+================+================+================+================+================+================+================+==========================================+================+
| 01.01          | Record Code            | AN             | "01"           | 2              | 1              | 2              | If invalid,    | Invalid:       | n/a            | ACH Schedule Header Record Identifier =  |                |
|                |                        |                |                |                |                |                | missing or out | Reason Group 1 |                | "01"                                     |                |
|                |                        |                |                |                |                |                | of order,      | message 6.     |                |                                          |                |
|                |                        |                |                |                |                |                | reject the     | Missing or out |                |                                          |                |
|                |                        |                |                |                |                |                | file           | of order:      |                |                                          |                |
|                |                        |                |                |                |                |                |                | Reason Group 1 |                |                                          |                |
|                |                        |                |                |                |                |                |                | message 4.     |                |                                          |                |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.02          | AgencyACHText          | AN             |                | 4              | 3              | 6              | n/a            |                | Agency Text    | Used to identify the payment to the      | FedACH         |
|                |                        |                |                |                |                |                |                |                |                | recipient. (In FedACH file Company Name  |                |
|                |                        |                |                |                |                |                |                |                |                | field.)                                  | Offset Notices |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | Funds Control  |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.03          | ScheduleNumber         | AN             |                | 14             | 7              | 20             | Right justify  | If dup in this | Schedule       | Any lower-case values received on input  | DNP            |
|                |                        |                |                |                |                |                | zero fill to   | file: Reason   | Number         | are stored as upper-case.                |                |
|                |                        |                |                |                |                |                | the left and   | Group 2        |                |                                          | PACER          |
|                |                        |                |                |                |                |                | remove all     | message 1.     |                |                                          |                |
|                |                        |                |                |                |                |                | spaces. If     |                |                |                                          | PIR            |
|                |                        |                |                |                |                |                | schedule       | If dup in this |                |                                          |                |
|                |                        |                |                |                |                |                | number is a    | fiscal year:   |                |                                          | PPS            |
|                |                        |                |                |                |                |                | duplicate      | Reason Group 2 |                |                                          |                |
|                |                        |                |                |                |                |                | of 1) other    | message 2.     |                |                                          | TOP            |
|                |                        |                |                |                |                |                | schedules on   |                |                |                                          |                |
|                |                        |                |                |                |                |                | this file, or  | If blank or    |                |                                          | IPP            |
|                |                        |                |                |                |                |                | 2) other       | invalid:       |                |                                          |                |
|                |                        |                |                |                |                |                | schedules in   | Reason Group 1 |                |                                          | TCS            |
|                |                        |                |                |                |                |                | PAM during a   | message 6.     |                |                                          |                |
|                |                        |                |                |                |                |                | fiscal year    |                |                |                                          | 1691           |
|                |                        |                |                |                |                |                | for the ALC    |                |                |                                          |                |
|                |                        |                |                |                |                |                | and the        |                |                |                                          | SnAP           |
|                |                        |                |                |                |                |                | schedule's     |                |                |                                          |                |
|                |                        |                |                |                |                |                | processing     |                |                |                                          | DoD-DCAS       |
|                |                        |                |                |                |                |                | status is not  |                |                |                                          |                |
|                |                        |                |                |                |                |                | removed,       |                |                |                                          | SRF            |
|                |                        |                |                |                |                |                | reject the     |                |                |                                          |                |
|                |                        |                |                |                |                |                | schedule.      |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                | If all blank   |                |                |                                          |                |
|                |                        |                |                |                |                |                | or invalid,    |                |                |                                          |                |
|                |                        |                |                |                |                |                | reject the     |                |                |                                          |                |
|                |                        |                |                |                |                |                | schedule.      |                |                |                                          |                |
|                |                        |                |                |                |                |                | Valid          |                |                |                                          |                |
|                |                        |                |                |                |                |                | characters     |                |                |                                          |                |
|                |                        |                |                |                |                |                | include A-Z,   |                |                |                                          |                |
|                |                        |                |                |                |                |                | 0-9, dash (-). |                |                |                                          |                |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.04          | PaymentTypeCode        | AN             |                | 25             | 21             | 45             | If all blank   | Reason Group 1 | Type of        | See [Appendix                            | DNP            |
|                |                        |                |                |                |                |                | or Type of     | message 6      | Payment        | E](#appendix-e---paymenttypecode-values) |                |
|                |                        |                |                |                |                |                | Payment not    |                |                | for a complete listing. for a complete   | PACER          |
|                |                        |                |                |                |                |                | configured in  |                |                | listing.                                 |                |
|                |                        |                |                |                |                |                | PAM, reject    |                |                |                                          | PIR            |
|                |                        |                |                |                |                |                | the schedule.  |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | PPS            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | TOP            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | IPP            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | TCS            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | DoD-DCAS       |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | SRF            |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.05          | StandardEntryClassCode | A              | CCD; PPD;      | 3              | 46             | 48             | If blank or    | Reason Group 1 | Standard Entry |                                          | DNP            |
|                |                        |                |                |                |                |                | invalid,       | message 6      | Class Code     |                                          |                |
|                |                        |                | IAT;           |                |                |                | reject the     |                |                |                                          | FedACH         |
|                |                        |                |                |                |                |                | schedule.      |                |                |                                          |                |
|                |                        |                | CTX            |                |                |                |                |                |                |                                          | IPP            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | PIR            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | PPS            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | DoD-DCAS       |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | SRF            |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.06          | AgencyLocationCode     | N              |                | 8              | 49             | 56             | Must be a      | Reason Group 1 | ALC            | Shown as Payment ALC in SRF.             | DNP            |
|                |                        |                |                |                |                |                | valid agency   | message 6      |                |                                          |                |
|                |                        |                |                |                |                |                | location code. |                |                |                                          | FedACH         |
|                |                        |                |                |                |                |                | If not, reject |                |                |                                          |                |
|                |                        |                |                |                |                |                | the schedule.  |                |                |                                          | PACER          |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | PIR            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | PPS            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | TOP            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | IPP            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | TCS            |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | 1691           |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | Funds Control  |
|                |                        |                |                |                |                |                |                |                |                |                                          | Offset Notices |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | SnAP           |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | DoD-DCAS       |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                |                                          | SRF            |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.07          | Filler                 |                |                | 1              | 57             | 57             | n/a            |                |                |                                          |                |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.08          | Federal Employer       | AN             |                | 10             | 58             | 67             | n/a            |                | Agency EIN     | Position 1-9 is used as the FEIN.        | FedACH         |
|                | Identification Number  |                |                |                |                |                |                |                |                | Position 10 is used by Treasury and      |                |
|                |                        |                |                |                |                |                |                |                |                | should remain blank.                     |                |
|                |                        |                |                |                |                |                |                |                |                |                                          |                |
|                |                        |                |                |                |                |                |                |                |                | This is passed in the Company            |                |
|                |                        |                |                |                |                |                |                |                |                | Identification field FedACH file for     |                |
|                |                        |                |                |                |                |                |                |                |                | child support purposes. It is the FEIN   |                |
|                |                        |                |                |                |                |                |                |                |                | of the agency that is the employer of    |                |
|                |                        |                |                |                |                |                |                |                |                | the payee, If this is not available then |                |
|                |                        |                |                |                |                |                |                |                |                | provides the FEIN of the agency sending  |                |
|                |                        |                |                |                |                |                |                |                |                | the payment.                             |                |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 01.09          | Filler                 |                |                | 783            | 68             | 850            | n/a            |                | n/a            |                                          |                |
+----------------+------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+

##  Check Schedule Header Record

+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **Check Schedule Header Record**                                                                                                                                                                                                               |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| **\#**         | **Field Name**            | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**                                | **Downstream   |
|                |                           |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                                          | Mapping**      |
+================+===========================+================+================+================+================+================+================+================+================+==========================================+================+
| 11.01          | Record Code               | AN             | "11"           | 2              | 1              | 2              | If invalid,    | Invalid:       | n/a            | Check Schedule Header Record Identifier  |                |
|                |                           |                |                |                |                |                | missing or out | Reason Group 1 |                | = "11"                                   |                |
|                |                           |                |                |                |                |                | of order,      | message 6.     |                |                                          |                |
|                |                           |                |                |                |                |                | reject the     | Missing or out |                |                                          |                |
|                |                           |                |                |                |                |                | file           | of order:      |                |                                          |                |
|                |                           |                |                |                |                |                |                | Reason Group 1 |                |                                          |                |
|                |                           |                |                |                |                |                |                | message 4.     |                |                                          |                |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 11.02          | ScheduleNumber            | AN             |                | 14             | 3              | 16             | Right justify  | If dup in this | Schedule       | Any lower-case values received on input  | DNP            |
|                |                           |                |                |                |                |                | zero fill to   | file: Reason   | Number         | are stored as upper-case.                |                |
|                |                           |                |                |                |                |                | the left and   | Group 2        |                |                                          | IPP            |
|                |                           |                |                |                |                |                | remove all     | message 1.     |                |                                          |                |
|                |                           |                |                |                |                |                | spaces. If     |                |                |                                          | PACER          |
|                |                           |                |                |                |                |                | schedule       | If dup in this |                |                                          |                |
|                |                           |                |                |                |                |                | number is a    | fiscal year:   |                |                                          | PIR            |
|                |                           |                |                |                |                |                | duplicate      | Reason Group 2 |                |                                          |                |
|                |                           |                |                |                |                |                | of 1) other    | message 2      |                |                                          | PPS            |
|                |                           |                |                |                |                |                | schedules on   |                |                |                                          |                |
|                |                           |                |                |                |                |                | this file, or  | If blank,:     |                |                                          | TOP            |
|                |                           |                |                |                |                |                | 2) other       | Reason Group 1 |                |                                          |                |
|                |                           |                |                |                |                |                | schedules in   | message 6      |                |                                          | TCS            |
|                |                           |                |                |                |                |                | PAM during a   |                |                |                                          |                |
|                |                           |                |                |                |                |                | fiscal year    |                |                |                                          | PrinCE         |
|                |                           |                |                |                |                |                | for the ALC    |                |                |                                          |                |
|                |                           |                |                |                |                |                | and the        |                |                |                                          | 1691           |
|                |                           |                |                |                |                |                | schedule's     |                |                |                                          |                |
|                |                           |                |                |                |                |                | processing     |                |                |                                          | DoD-DCAS       |
|                |                           |                |                |                |                |                | status is not  |                |                |                                          |                |
|                |                           |                |                |                |                |                | removed,       |                |                |                                          | SRF            |
|                |                           |                |                |                |                |                | reject the     |                |                |                                          |                |
|                |                           |                |                |                |                |                | schedule.      |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                | If schedule    |                |                |                                          |                |
|                |                           |                |                |                |                |                | number is all  |                |                |                                          |                |
|                |                           |                |                |                |                |                | blank or       |                |                |                                          |                |
|                |                           |                |                |                |                |                | invalid,       |                |                |                                          |                |
|                |                           |                |                |                |                |                | reject the     |                |                |                                          |                |
|                |                           |                |                |                |                |                | schedule.      |                |                |                                          |                |
|                |                           |                |                |                |                |                | Valid          |                |                |                                          |                |
|                |                           |                |                |                |                |                | characters     |                |                |                                          |                |
|                |                           |                |                |                |                |                | include A-Z,   |                |                |                                          |                |
|                |                           |                |                |                |                |                | 0-9, dash (-). |                |                |                                          |                |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| > 11.03        | PaymentTypeCode           | AN             |                | 25             | 17             | 41             | If all blank   | Reason Group 1 | Type of        | See [Appendix                            | DNP            |
|                |                           |                |                |                |                |                | or Type of     | message 6      | Payment        | E](#appendix-e---paymenttypecode-values) |                |
|                |                           |                |                |                |                |                | Payment not    |                |                | for a complete listing. for a complete   | IPP            |
|                |                           |                |                |                |                |                | configured in  |                |                | listing.                                 |                |
|                |                           |                |                |                |                |                | PAM, reject    |                |                |                                          | PACER          |
|                |                           |                |                |                |                |                | the schedule.  |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | PIR            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | PPS            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | TOP            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | TCS            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | TCIS           |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | PrinCE         |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | DoD-DCAS       |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | SRF            |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 11.04          | AgencyLocationCode        | N              |                | 8              | 42             | 49             | Must be a      | Reason Group 1 | ALC            | Shown as Payment ALC in SRF.             | DNP            |
|                |                           |                |                |                |                |                | valid agency   | message 6      |                |                                          |                |
|                |                           |                |                |                |                |                | location code. |                |                |                                          | IPP            |
|                |                           |                |                |                |                |                | If not, reject |                |                |                                          |                |
|                |                           |                |                |                |                |                | the schedule.  |                |                |                                          | PrinCE         |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | PACER          |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | PIR            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | PPS            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | TOP            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | TCS            |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | TCIS           |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | 1691           |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | Offset Notices |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | DoD-DCAS       |
|                |                           |                |                |                |                |                |                |                |                |                                          |                |
|                |                           |                |                |                |                |                |                |                |                |                                          | SRF            |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 11.05          | Filler                    | AN             |                | 9              | 50             | 58             | n/a            |                | Filler         | Filler                                   |                |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| 11.06          | CheckPaymentEnclosureCode | A              | "nameonly"     | 10             | 59             | 68             | If invalid     | Reason Group 1 | Enclosure Code | If stub, the stub record should be       | TCS            |
|                |                           |                | "letter",      |                |                |                | reject the     | message 6      |                | provided for each payment in the         |                |
|                |                           |                | "stub"         |                |                |                | schedule. If   |                |                | schedule using the Check Stub Record. If | PACER          |
|                |                           |                | "insert" or    |                |                |                | "stub" and     |                |                | letter, a separate letter file should be |                |
|                |                           |                | blank          |                |                |                | there is no    |                |                | provided.                                | PrinCE         |
|                |                           |                |                |                |                |                | stub record    |                |                |                                          |                |
|                |                           |                |                |                |                |                | for each       |                |                |                                          | TOP            |
|                |                           |                |                |                |                |                | payment,       |                |                |                                          |                |
|                |                           |                |                |                |                |                | reject the     |                |                |                                          |                |
|                |                           |                |                |                |                |                | schedule.      |                |                |                                          |                |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+
| > 11.07        | Filler                    |                |                | 782            | 69             | 850            | n/a            |                | n/a            |                                          |                |
+----------------+---------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+------------------------------------------+----------------+

##  ACH Payment Data Record

+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **ACH Payment Data Record**                                                                                                                                                                                                                        |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| **\#**         | **Field Name**             | **Type**       | **Field Value**    | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**                               | **Downstream   |
|                |                            |                |                    |                | Position**     | Position**     | rules**        |                | Name**         |                                         | Mapping**      |
+:===============+============================+================+====================+================+================+================+================+================+================+=========================================+================+
| 02.01          | Record Code                | AN             | "02"               | 2              | 1              | 2              | If invalid,    | Invalid:       | n/a            | ACH Payment Record Indicator = "02".    |                |
|                |                            |                |                    |                |                |                | missing or out | Reason Group 1 |                |                                         |                |
|                |                            |                |                    |                |                |                | of order,      | message 6.     |                |                                         |                |
|                |                            |                |                    |                |                |                | reject the     | Missing or out |                |                                         |                |
|                |                            |                |                    |                |                |                | file           | of order:      |                |                                         |                |
|                |                            |                |                    |                |                |                |                | Reason Group 1 |                |                                         |                |
|                |                            |                |                    |                |                |                |                | message 4.     |                |                                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.02          | AgencyAccountIdentifier    | AN             |                    | 16             | 3              | 18             | Can be less    |                | Account Number | This is the payee account number used   | DNP            |
|                |                            |                |                    |                |                |                | than 16        |                |                | by the agency if different from TIN.    |                |
|                |                            |                |                    |                |                |                | characters.    |                |                | TIN must be provided in the Payee       | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                | Identifier field. See field below.      |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                | If Custom Agency Rule ID = 'VACP', do   |                |
|                |                            |                |                    |                |                |                |                |                |                | not justify, store value as received.   | PACER          |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                | This field is passed to                 | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                | PACER/PPS and FedACH as the Payee       | PPS            |
|                |                            |                |                    |                |                |                |                |                |                | Account                                 |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                | Number.                                 |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                | It is displayed as the Payee ID in      |                |
|                |                            |                |                    |                |                |                |                |                |                | PACER/PPS and TCIS.                     | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.03          | Amount                     | N              | \$\$\$\$\$\$\$\$cc | 10             | 19             | 28             | Right justify, | If invalid,    | Payment Amount |                                         | DNP            |
|                |                            |                |                    |                |                |                | zero fill. If  | Reason Group 5 |                |                                         |                |
|                |                            |                |                    |                |                |                | invalid or all | message 3. See |                |                                         | PACER          |
|                |                            |                |                    |                |                |                | blank, mark    | Validation for |                |                                         |                |
|                |                            |                |                    |                |                |                | payment as     | Balancing      |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | invalid See    | section for    |                |                                         |                |
|                |                            |                |                    |                |                |                | Section 1.5    | other error    |                |                                         | PPS            |
|                |                            |                |                    |                |                |                | Validation for | codes.         |                |                                         |                |
|                |                            |                |                    |                |                |                | Balancing for  |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                | Amount, SEC,   |                |                |                                         |                |
|                |                            |                |                    |                |                |                | and            |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                | Transaction    |                |                |                                         |                |
|                |                            |                |                    |                |                |                | Code cross     |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                | validation     |                |                |                                         |                |
|                |                            |                |                    |                |                |                | rules.         |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Funds Control  |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.04          | AgencyPaymentTypeCode      | AN             |                    | 1              | 29             | 29             | n/a            |                | Agency Payment |                                         | DNP            |
|                |                            |                |                    |                |                |                |                |                | Code           |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PACER          |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.05          | IsTOP_Offset               | AN             | "0";"1"            | 1              | 30             | 30             | n/a            |                | TOP            | "0" = no                                | TOP            |
|                |                            |                |                    |                |                |                |                |                | Eligibility    |                                         |                |
|                |                            |                |                    |                |                |                |                |                | Indicator      | "1" = yes                               | SRF            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                | Any value other than 0 will be          |                |
|                |                            |                |                    |                |                |                |                |                |                | defaulted to a 1..Agreed upon between   |                |
|                |                            |                |                    |                |                |                |                |                |                | TOP and agency.                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.06          | PayeeName                  | AN             |                    | 35             | 31             | 65             | Left justify,  | Reason group 5 | Payee Name     | This will be truncated to 22 characters | DNP            |
|                |                            |                |                    |                |                |                | blank fill.    | message 3      |                | when sent to FedACH for PPD and CCD. It |                |
|                |                            |                |                    |                |                |                |                |                |                | will be truncated to 16 for CTX and 35  | PACER          |
|                |                            |                |                    |                |                |                | If all blank,  |                |                | for IAT.                                |                |
|                |                            |                |                    |                |                |                | mark payment   |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | as invalid.    |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.07          | PayeeAddressLine_1         | AN             |                    | 35             | 66             | 100            | If IAT and     | Reason group 5 | Payee Address  | If IAT, this is required. This is       | DNP            |
|                |                            |                |                    |                |                |                | this is all    | message 3      | Line 1         | designated as the payee street address  |                |
|                |                            |                |                    |                |                |                | blank, mark    |                |                | for IAT payments.                       | FedACH         |
|                |                            |                |                    |                |                |                | payment as     |                |                |                                         |                |
|                |                            |                |                    |                |                |                | invalid        |                |                | While not required for other entry      | IPP            |
|                |                            |                |                    |                |                |                |                |                |                | class codes, address data is requested  |                |
|                |                            |                |                    |                |                |                |                |                |                | from agencies so RFCs can send offset   | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                | letters if necessary.                   |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.08          | PayeeAddressLine_2         | AN             |                    | 35             | 101            | 135            | n/a            |                | Payee Address  | Optional address line.                  | DNP            |
|                |                            |                |                    |                |                |                |                |                | Line 2         |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.09          | CityName                   | AN             |                    | 27             | 136            | 162            | Left justify,  | Reason group 5 | City           | Foreign city if IAT. Will be truncated  | DNP            |
|                |                            |                |                    |                |                |                | blank fill. If | message 3      |                | to 23 characters when sent to FedACH.   |                |
|                |                            |                |                    |                |                |                | IAT and this   |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                | is all blank,  |                |                | Optional for domestic payments (entry   |                |
|                |                            |                |                    |                |                |                | mark payment   |                |                | class codes other than IAT).            | IPP            |
|                |                            |                |                    |                |                |                | as invalid     |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.10          | StateName                  | AN             |                    | 10             | 163            | 172            | n/a            |                | State Name     | If IAT, the foreign state, territory or | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                | province name is provided in this       |                |
|                |                            |                |                    |                |                |                |                |                |                | field, if applicable.                   | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.11          | StateCodeText              | AN             |                    | 2              | 173            | 174            | n/a            |                | State          | Optional field for domestic state       | DNP            |
|                |                            |                |                    |                |                |                |                |                |                | abbreviation for PPD/CCD entries.       |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.12          | PostalCode                 | AN             |                    | 5              | 175            | 179            | n/a            |                | Zip Code 5 or  | Foreign postal code (IAT) or domestic   | DNP            |
|                |                            |                |                    |                |                |                |                |                | Geo Code       | zip code (all other entry class codes). |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                |                |                | If ~~bb~~nnn,  |                                         |                |
|                |                            |                |                    |                |                |                |                |                | store as Geo   |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                | Code if Geo    |                                         |                |
|                |                            |                |                    |                |                |                |                |                | Code is not    |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                | populated      |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.13          | PostalCodeExtension        | AN             |                    | 5              | 180            | 184            | n/a            |                | Zip Code 4     | Remainder of foreign postal code (IAT)  | DNP            |
|                |                            |                |                    |                |                |                |                |                |                | or domestic 4-digit zip (all other      |                |
|                |                            |                |                    |                |                |                |                |                |                | entry class codes).                     | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.14          | CountryCodeText            | AN             |                    | 2              | 185            | 186            | If IAT and     | Reason group 5 | Country Code   | If IAT the ISO country code is          | DNP            |
|                |                            |                |                    |                |                |                | value is blank | message 3      | treasury.gov   | required.                               |                |
|                |                            |                |                    |                |                |                | and/or         |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                | contains       |                |                | Contact the Treasury Agency Outreach    |                |
|                |                            |                |                    |                |                |                | zeros, mark    |                |                | Team for a complete list of country     | IPP            |
|                |                            |                |                    |                |                |                | payment as     |                |                | codes at                                |                |
|                |                            |                |                    |                |                |                | invalid.       |                |                |                                         | Offset Notices |
|                |                            |                |                    |                |                |                |                |                |                | FS.AgencyOutreach@fiscal.treasury.gov   |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.15          | Routing Number             | N              |                    | 9              | 187            | 195            | If not valid   | Reason Group 5 | Depositor RTN  |                                         | DNP            |
|                |                            |                |                    |                |                |                | RTN, mark      | message 3      |                |                                         |                |
|                |                            |                |                    |                |                |                | payment as     |                |                |                                         | PACER          |
|                |                            |                |                    |                |                |                | invalid.       |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-CAS        |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.16          | Account Number             | AN             |                    | 17             | 196            | 212            | If all blank   | Reason Group 5 | Depositor      |                                         | DNP            |
|                |                            |                |                    |                |                |                | or if value    | message 3      | Account Number |                                         |                |
|                |                            |                |                    |                |                |                | contains only  |                |                |                                         | PACER          |
|                |                            |                |                    |                |                |                | zeroes, mark   |                |                |                                         |                |
|                |                            |                |                    |                |                |                | payment as     |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | invalid.       |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.17          | ACH_TransactionCode        | N              |                    | 2              | 213            | 214            | If 1) not one  | Reason Group 5 | ACH            | See [Appendix                           | DNP            |
|                |                            |                |                    |                |                |                | of the values  | message 3      | Transaction    | A](#appendix-a---ach-transaction-codes) |                |
|                |                            |                |                    |                |                |                | listed 2)      |                | Code +         | for values / definitions / stored       | PACER          |
|                |                            |                |                    |                |                |                | Codes 42, 43,  |                |                | values                                  |                |
|                |                            |                |                    |                |                |                | 52, or 53 are  |                | Account Type   |                                         | PIR            |
|                |                            |                |                    |                |                |                | used for a     |                |                |                                         |                |
|                |                            |                |                    |                |                |                | type of        |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                | payment other  |                |                |                                         |                |
|                |                            |                |                    |                |                |                | than Vendor or |                |                |                                         | FedACH         |
|                |                            |                |                    |                |                |                | 3) all blank,  |                |                |                                         |                |
|                |                            |                |                    |                |                |                | mark payment   |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                | as invalid.    |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.18          | PayeeIdentifier_Additional | AN             |                    | 9              | 215            | 223            | If not 9       | Reason Group 5 | Secondary TIN  | Used to provide TOP with additional     | DNP            |
|                |                            |                |                    |                |                |                | numeric or all | message 3      |                | information.                            |                |
|                |                            |                |                    |                |                |                | blank, mark    |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                | payment as     |                |                |                                         |                |
|                |                            |                |                    |                |                |                | invalid. Blank |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                | is the only    |                |                |                                         |                |
|                |                            |                |                    |                |                |                | non-numeric    |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                | allowed        |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PACER          |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.19          | PayeeName_Additional       | AN             |                    | 35             | 224            | 258            | n/a            |                | Secondary Name | Used to provide TOP with additional     | DNP            |
|                |                            |                |                    |                |                |                |                |                |                | information.                            |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PACER          |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.20          | PaymentID                  | AN             |                    | 20             | 259            | 278            | If all blank   | Reason Group 1 | Payment ID     | Unique payment identifier within the    | DNP            |
|                |                            |                |                    |                |                |                | or duplicate   | message 6      |                | schedule for associating all related    |                |
|                |                            |                |                    |                |                |                | within the     |                |                | records for the payment.                | IPP            |
|                |                            |                |                    |                |                |                | schedule,      |                |                |                                         |                |
|                |                            |                |                    |                |                |                | reject the     |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | schedule.      |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.21          | Reconcilement              | AN             |                    | 100            | 279            | 378            | Store value as |                | Reconcilement  | Limited Use. Requires agency / Treasury | PACER          |
|                |                            |                |                    |                |                |                | received, do   |                |                |                                         |                |
|                |                            |                |                    |                |                |                | not justify    |                |                | collaboration prior to use.             | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                | See [Appendix                           |                |
|                |                            |                |                    |                |                |                |                |                |                | B](#appendix-b-agency-specific-values)  |                |
|                |                            |                |                    |                |                |                |                |                |                | for Agency Specific Values.             |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.22          | TIN                        | AN             |                    | 9              | 379            | 387            | If not 9       | Reason Group 5 | TIN            | This is the payee TIN.                  | DNP            |
|                |                            |                |                    |                |                |                | numeric or all | message 3      |                |                                         |                |
|                |                            |                |                    |                |                |                | blank, mark    |                |                | This field is associated with Payee     | PACER          |
|                |                            |                |                    |                |                |                | payment as     |                |                | Name.                                   |                |
|                |                            |                |                    |                |                |                | invalid. Blank |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | is the only    |                |                |                                         |                |
|                |                            |                |                    |                |                |                | non-numeric    |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                | allowed        |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TCS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | IPP            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SnAP           |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.23          | Payment Recipient TIN      | AN             | "1"=SSN            | 1              | 388            | 388            | If not 1, 2, 3 | Reason Group 5 | TOP Payment    | Indicates the type of payee identifier  | DNP            |
|                | indicator                  |                |                    |                |                |                | or blank, mark | message 3      | Recipient TIN  | provided in the TIN field.              |                |
|                |                            |                | "2"=EIN            |                |                |                | payment as     |                | Indicator      |                                         | IPP            |
|                |                            |                |                    |                |                |                | invalid. Blank |                |                |                                         |                |
|                |                            |                | "3"=ITIN           |                |                |                | is the only    |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | non-numeric    |                |                |                                         |                |
|                |                            |                | or blank           |                |                |                | allowed        |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP NG         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS       |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.24          | Additional Payee TIN       | AN             | "1"=SSN            | 1              | 389            | 389            | If not 1, 2, 3 | Reason Group 5 | TOP Secondary  | Indicates the type of payee identifier  | DNP            |
|                | Indicator                  |                |                    |                |                |                | or blank, mark | message 3      | Payee TIN      | provided in the                         |                |
|                |                            |                | "2"=EIN            |                |                |                | payment as     |                | Indicator      | PayeeIdentifier_Additioanl field.       | IPP            |
|                |                            |                |                    |                |                |                | invalid. Blank |                |                |                                         |                |
|                |                            |                | "3"=ITIN           |                |                |                | is the only    |                |                |                                         | PIR            |
|                |                            |                |                    |                |                |                | non-numeric    |                |                |                                         |                |
|                |                            |                | or blank           |                |                |                | allowed        |                |                |                                         | PPS            |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | TOP NG         |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | DoD-DCAS,      |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         | SRF            |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.25          | Amount eligible for offset | AN             |                    | 10             | 390            | 399            | If not all     | Reason Group 5 | Amount         | The amount of the payment that is       | TOP            |
|                |                            |                |                    |                |                |                | blank or all   | message 3      | eligible for   | eligible for offset. Will be sent to    |                |
|                |                            |                |                    |                |                |                | numeric, mark  |                | offset         | TOP.                                    |                |
|                |                            |                |                    |                |                |                | payment as     |                |                |                                         |                |
|                |                            |                |                    |                |                |                | invalid. Blank |                |                | If the entire payment is eligible for   |                |
|                |                            |                |                    |                |                |                | is the only    |                |                | offset, pass blanks in this field. If a |                |
|                |                            |                |                    |                |                |                | non-numeric    |                |                | portion of the payment is eligible for  |                |
|                |                            |                |                    |                |                |                | allowed.       |                |                | offset, specify the amount in this      |                |
|                |                            |                |                    |                |                |                |                |                |                | field.                                  |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.26          | PayeeAddressLine_3         | AN             |                    | 35             | 400            | 434            | n/a            |                | Payee Address  | Optional address line.                  |                |
|                |                            |                |                    |                |                |                |                |                | Line 3         |                                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.27          | PayeeAddressLine_4         | AN             |                    | 35             | 435            | 469            | n/a            |                | Payee Address  | Optional address line.                  |                |
|                |                            |                |                    |                |                |                |                |                | Line 4         |                                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.28          | CountryName                | AN             |                    | 40             | 470            | 509            | n/a            |                | Country Name   | Full name of country for foreign        |                |
|                |                            |                |                    |                |                |                |                |                |                | addresses                               |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.29          | ConsularCode               | AN             |                    | 3              | 510            | 512            | n/a            |                | Geo Code       | Contact the Agency Outreach Team for a  | SRF            |
|                |                            |                |                    |                |                |                |                |                |                | complete list of consular / geo codes   |                |
|                |                            |                |                    |                |                |                |                |                |                | at                                      |                |
|                |                            |                |                    |                |                |                |                |                |                |                                         |                |
|                |                            |                |                    |                |                |                |                |                |                | FS.AgencyOutreach@fiscal.treasury.gov   |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.30          | SubPaymentTypeCode         | AN             |                    | 32             | 513            | 544            | n/a            |                | Sub Payment    | Value agreed to between Agency and      | SRF            |
|                |                            |                |                    |                |                |                |                |                | Type Code      | Treasury.                               |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.31          | PayerMechanism             | AN             | CreditCard,        | 20             | 545            | 564            | If not one of  | Reason Group 5 | Payer          |                                         | SRF            |
|                |                            |                | DebitCard, SVC,    |                |                |                | the values     | message 3      | Mechanism      |                                         |                |
|                |                            |                | ACH, BookEntry,    |                |                |                | listed, mark   |                |                |                                         |                |
|                |                            |                | EBT, or blank      |                |                |                | invalid.       |                |                |                                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.32          | PaymentDescriptionCode     | AN             |                    | 2              | 565            | 566            | n/a            | n/a            | Payment        | Provides additional description for TOP | TOP            |
|                |                            |                |                    |                |                |                |                |                | Description    | offsetting. Values agreed to between    |                |
|                |                            |                |                    |                |                |                |                |                | Code           | Agency and TOP.                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+
| 02.33          | Filler                     |                |                    | 284            | 567            | 850            | n/a            |                | n/a            |                                         |                |
+----------------+----------------------------+----------------+--------------------+----------------+----------------+----------------+----------------+----------------+----------------+-----------------------------------------+----------------+

##  Check Payment Data Record

For additional clarifications for populating address fields, refer to
[Appendix C](#appendix-c---addressing-reference-information).

+----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **Check Payment Data Record**                                                                                                                                                                                                    |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| **\#**    | **Field Name**             | **Type**  | **Field Value**    | **Length** | **Start    | **End      | **Validation | **Error   | **Stored      | **Notes**                              | **Downstream Mapping**      |
|           |                            |           |                    |            | Position** | Position** | rules**      | Code**    | Name**        |                                        |                             |
+===========+============================+===========+====================+============+============+============+==============+===========+===============+========================================+=================+===========+
| 12.01     | Record Code                | AN        | "12"               | 2          | 1          | 2          | If invalid,  | Invalid:  | n/a           | Check Payment Data Record Indicator =  |                             |
|           |                            |           |                    |            |            |            | missing or   | Reason    |               | "12"                                   |                             |
|           |                            |           |                    |            |            |            | out of       | Group 1   |               |                                        |                             |
|           |                            |           |                    |            |            |            | order,       | message   |               |                                        |                             |
|           |                            |           |                    |            |            |            | reject the   | 6.        |               |                                        |                             |
|           |                            |           |                    |            |            |            | file         | Missing   |               |                                        |                             |
|           |                            |           |                    |            |            |            |              | or out of |               |                                        |                             |
|           |                            |           |                    |            |            |            |              | order:    |               |                                        |                             |
|           |                            |           |                    |            |            |            |              | Reason    |               |                                        |                             |
|           |                            |           |                    |            |            |            |              | Group 1   |               |                                        |                             |
|           |                            |           |                    |            |            |            |              | message   |               |                                        |                             |
|           |                            |           |                    |            |            |            |              | 4.        |               |                                        |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.02     | AgencyAccountIdentifier    | AN        |                    | 16         | 3          | 18         | Can be less  |           | Account       | This is the payee account number used  | DNP                         |
|           |                            |           |                    |            |            |            | than 16      |           | Number        | by the agency if different from TIN.   |                             |
|           |                            |           |                    |            |            |            | characters.  |           |               | TIN must be provided in the Payee      | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               | Identifier field. See field below.     |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               | If Custom Agency Rule ID = 'VACP', do  |                             |
|           |                            |           |                    |            |            |            |              |           |               | not justify, store value as received.  | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               | This field is passed to                | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               | PACER/PPS as the Payee Account Number  | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               | and displayed as the Payee ID in       |                             |
|           |                            |           |                    |            |            |            |              |           |               | PACER/PPS and TCIS.                    | TCIS                        |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.03     | Amount                     | N         | \$\$\$\$\$\$\$\$cc | 10         | 19         | 28         | Right        | If        | Payment       |                                        | DNP                         |
|           |                            |           |                    |            |            |            | justify,     | invalid,  | Amount        |                                        |                             |
|           |                            |           |                    |            |            |            | zero fill.   | Reason    |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              | Group 5   |               |                                        |                             |
|           |                            |           |                    |            |            |            | If all blank | message   |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            | or zero      | 3.        |               |                                        |                             |
|           |                            |           |                    |            |            |            | dollar, mark |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            | payment as   |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | invalid.     |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCIS                        |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.04     | AgencyPaymentTypeCode      | AN        |                    | 1          | 29         | 29         | n/a          |           | Agency        |                                        | DNP                         |
|           |                            |           |                    |            |            |            |              |           | Payment Code  |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCIS                        |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | Offset Notices              |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.05     | IsTOP_Offset               | AN        | "0";"1"            | 1          | 30         | 30         | n/a          |           | TOP           | "0" = no                               | TOP                         |
|           |                            |           |                    |            |            |            |              |           | Eligibility   |                                        |                             |
|           |                            |           |                    |            |            |            |              |           | Indicator     | "1" = yes                              | SRF                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               | Any value other than 0 will be         |                             |
|           |                            |           |                    |            |            |            |              |           |               | defaulted to a 1. Agreed upon between  |                             |
|           |                            |           |                    |            |            |            |              |           |               | TOP and agency.                        |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.06     | PayeeName                  | AN        |                    | 35         | 31         | 65         | Left         | Reason    | Payee Name    |                                        | DNP                         |
|           |                            |           |                    |            |            |            | Justify,     | group 5   |               |                                        |                             |
|           |                            |           |                    |            |            |            | blank fill.  | message 3 |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | If all       |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            | blank, mark  |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | payment as   |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            | invalid      |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCIS                        |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | Offset Notices              |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.07     | PayeeAddressLine_1         | AN        |                    | 35         | 66         | 100        | If blank and |           | Payee Address | Suspect indicates an internal          | DNP                         |
|           |                            |           |                    |            |            |            | enclosure    |           | Line 1        |                                        |                             |
|           |                            |           |                    |            |            |            | code is not  |           |               | Treasury review for mail ability; it   | IPP                         |
|           |                            |           |                    |            |            |            | "nameonly",  |           |               | is not marked invalid.                 |                             |
|           |                            |           |                    |            |            |            | mark payment |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            | as suspect.  |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | Offset Notices              |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.08     | PayeeAddressLine_2         | AN        |                    | 35         | 101        | 135        | n/a          |           | Payee Address |                                        | DNP                         |
|           |                            |           |                    |            |            |            |              |           | Line 2        |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | Offset Notices              |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.09     | PayeeAddressLine_3         | AN        |                    | 35         | 136        | 170        | n/a          |           | Payee Address |                                        | DNP                         |
|           |                            |           |                    |            |            |            |              |           | Line 3        |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.10     | PayeeAddressLine_4         | AN        |                    | 35         | 171        | 205        | n/a          |           | Payee Address | .                                      | DNP                         |
|           |                            |           |                    |            |            |            |              |           | Line 4        |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.11     | CityName                   | AN        |                    | 27         | 206        | 232        |              |           | City          | Will be truncated to 20 characters for | DNP                         |
|           |                            |           |                    |            |            |            |              |           |               | foreign addresses when State Name is   |                             |
|           |                            |           |                    |            |            |            |              |           |               | used.                                  | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.12     | StateName                  | AN        |                    | 10         | 233        | 242        | n/a          |           | State Name    | Use for foreign state, province or     | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               | territory name                         |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.13     | StateCodeText              | AN        |                    | 2          | 243        | 244        |              |           | State         | Use this field for domestic state      | DNP                         |
|           |                            |           |                    |            |            |            |              |           |               | code.                                  |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.14     | PostalCode                 | AN        |                    | 5          | 245        | 249        | If domestic, |           | Zip Code 5 or | If foreign address, use this field for | DNP                         |
|           |                            |           |                    |            |            |            | blank, and   |           | Geo Code      | universal postal code.                 |                             |
|           |                            |           |                    |            |            |            | enclosure    |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            | code is not  |           | If ~~bb~~nnn, | If domestic, use 5 digit zip code.     |                             |
|           |                            |           |                    |            |            |            | "nameonly"   |           | store as Geo  |                                        | PACER                       |
|           |                            |           |                    |            |            |            | mark payment |           | Code if Geo   | Suspect indicates an internal Treasury |                             |
|           |                            |           |                    |            |            |            | as suspect.  |           | Code is not   | review for mail ability; it is not     | PIR                         |
|           |                            |           |                    |            |            |            |              |           | populated     | marked invalid.                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.15     | PostalCodeExtension        | AN        |                    | 5          | 250        | 254        | n/a          |           | Zip Code 4    | If foreign address, use this field for | DNP                         |
|           |                            |           |                    |            |            |            |              |           |               | remainder of universal postal code     |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               | If domestic, use 4 digit zip code.     |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.16     | PostNet BarcodeDelivery    | AN        |                    | 3          | 255        | 257        | n/a          |           | Delivery      | Includes barcode check digit           | PrinCE                      |
|           | Point                      |           |                    |            |            |            |              |           | Point         |                                        |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.17     | Filler                     | AN        |                    | 14         | 258        | 271        | n/a          |           |               |                                        |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.18     | CountryName                | AN        |                    | 40         | 272        | 311        | n/a          |           | Country Name  | Full name of country for foreign       | DNP                         |
|           |                            |           |                    |            |            |            |              |           |               | addresses                              |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.19     | ConsularCode               | AN        |                    | 3          | 312        | 314        | n/a          |           | Geo Code      | Contact the Treasury Agency Outreach   | DNP                         |
|           |                            |           |                    |            |            |            |              |           |               | Team for a complete list of consular / |                             |
|           |                            |           |                    |            |            |            |              |           |               | geo codes at                           | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               | FS.AgencyOutreach@fiscal.treasury.gov  |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.20     | CheckLegendText1           | AN        |                    | 55         | 315        | 369        | n/a          |           | Object Line 1 | Agency provided information that will  | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               | be printed at the bottom of the check  |                             |
|           |                            |           |                    |            |            |            |              |           |               | and provided to PACER for cancellation | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               | notices. This can include allotter     |                             |
|           |                            |           |                    |            |            |            |              |           |               | information, month/year payment is     | SRF                         |
|           |                            |           |                    |            |            |            |              |           |               | for, specific note to the payee, etc.  |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.21     | CheckLegendText2           | AN        |                    | 55         | 370        | 424        | n/a          |           | Object Line 2 | Agency provided information that will  | PrinCE                      |
|           |                            |           |                    |            |            |            |              |           |               | be printed at the bottom of the check  |                             |
|           |                            |           |                    |            |            |            |              |           |               | and provided to PACER for cancellation | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               | notices. This can include allotter     |                             |
|           |                            |           |                    |            |            |            |              |           |               | information, month/year payment is     | SRF                         |
|           |                            |           |                    |            |            |            |              |           |               | for, specific note to the payee, etc.  |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.22     | PayeeIdentifier_Secondary  | AN        |                    | 9          | 425        | 433        | If not 9     | Reason    | Secondary TIN | Used to provide TOP with additional    | DNP                         |
|           |                            |           |                    |            |            |            | numeric or   | Group 5   |               | information.                           |                             |
|           |                            |           |                    |            |            |            | all blank,   | message 3 |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            | mark payment |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | as invalid.  |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            | Blank is the |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | only         |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            | non-numeric  |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | allowed.     |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | sRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.23     | PartyName_Secondary        | AN        |                    | 35         | 434        | 468        | n/a          |           | Secondary     | Used to provide TOP with additional    | DNP                         |
|           |                            |           |                    |            |            |            |              |           | Name          | information.                           |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.24     | PaymentID                  | AN        |                    | 20         | 469        | 488        | If all blank | Reason    | Payment ID    | Unique payment identifier within the   | PrinCE (when check          |
|           |                            |           |                    |            |            |            | or duplicate | Group 1   |               | schedule for associating all related   | enclosure code = "letter")  |
|           |                            |           |                    |            |            |            | within the   | message 6 |               | records for the payment. Also used to  |                             |
|           |                            |           |                    |            |            |            | schedule,    |           |               | match letters to checks during         | DNP                         |
|           |                            |           |                    |            |            |            | reject the   |           |               | printing / enclosing if check          |                             |
|           |                            |           |                    |            |            |            | schedule.    |           |               | enclosure code in schedule header =    | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               | "letter".                              |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.25     | Reconcilement              | AN        |                    | 100        | 489        | 588        | Store value  |           | Reconcilement | Limited Use. Requires agency /         | PACER                       |
|           |                            |           |                    |            |            |            | as received, |           |               | Treasury                               |                             |
|           |                            |           |                    |            |            |            | do not       |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            | justify      |           |               | collaboration prior to use.            |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               | See [Appendix                          |                             |
|           |                            |           |                    |            |            |            |              |           |               | B](#appendix-b-agency-specific-values) |                             |
|           |                            |           |                    |            |            |            |              |           |               | for Agency Specific Values.            |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.26     | SpecialHandling            | AN        |                    | 50         | 589        | 638        | n/a          |           | Special       | If CheckPaymentEnclosureCode=          | Check Listing               |
|           |                            |           |                    |            |            |            |              |           | Handling      | 'nameonly' and this field is blank,    |                             |
|           |                            |           |                    |            |            |            |              |           |               | store 'name only' in this field.       |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               | Requires agency / Treasury             |                             |
|           |                            |           |                    |            |            |            |              |           |               | collaboration prior to use. Used for   |                             |
|           |                            |           |                    |            |            |            |              |           |               | manual handling of checks.             |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.27     | TIN                        | AN        |                    | 9          | 639        | 647        | If not 9     | Reason    | TIN           | This is the payee TIN. This field is   | DNP                         |
|           |                            |           |                    |            |            |            | numeric or   | Group 5   |               | associated with Payee Name.            |                             |
|           |                            |           |                    |            |            |            | all blank,   | message 3 |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            | mark payment |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | as invalid.  |           |               |                                        | PACER                       |
|           |                            |           |                    |            |            |            | Blank is the |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | only         |           |               |                                        | PIR                         |
|           |                            |           |                    |            |            |            | non-numeric  |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | allowed.     |           |               |                                        | PPS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TOP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | TCS                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | IPP                         |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | DoD-DCAS                    |
|           |                            |           |                    |            |            |            |              |           |               |                                        |                             |
|           |                            |           |                    |            |            |            |              |           |               |                                        | SRF                         |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.28     | USPSIntelligentMailBarcode | AN        |                    | 50         | 648        | 697        |              |           | Intelligent   | Agencies that use Intelligent Mail     | Prince                      |
|           |                            |           |                    |            |            |            |              |           | Mail Bar Code | Barcode services from the post office  |                             |
|           |                            |           |                    |            |            |            |              |           | Data          | must work with Treasury to get         |                             |
|           |                            |           |                    |            |            |            |              |           |               | specific formatting of this field.     |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------+-----------+
| 12.29     | Payment Recipient TIN      | AN        | "1"=SSN            | 1          | 698        | 698        | If not 1, 2, | Reason    | TOP Payment   | Indicates the type of payee identifier provided in the   | DNP       |
|           | indicator                  |           |                    |            |            |            | 3 or blank,  | Group 5   | Recipient TIN | (TIN) field.                                             |           |
|           |                            |           | "2"=EIN            |            |            |            | mark payment | message 3 | Indicator     |                                                          | IPP       |
|           |                            |           |                    |            |            |            | as invalid.  |           |               |                                                          |           |
|           |                            |           | "3"=ITIN           |            |            |            | Blank is the |           |               |                                                          | PIR       |
|           |                            |           |                    |            |            |            | only         |           |               |                                                          |           |
|           |                            |           | or blank           |            |            |            | non-numeric  |           |               |                                                          | PPS       |
|           |                            |           |                    |            |            |            | allowed.     |           |               |                                                          |           |
|           |                            |           |                    |            |            |            |              |           |               |                                                          | TOP NG    |
|           |                            |           |                    |            |            |            |              |           |               |                                                          |           |
|           |                            |           |                    |            |            |            |              |           |               |                                                          | DoD-DCAS  |
|           |                            |           |                    |            |            |            |              |           |               |                                                          |           |
|           |                            |           |                    |            |            |            |              |           |               |                                                          | SRF       |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------------------------+-----------+
| 12.30     | Secondary Payee TIN        | AN        | "1"=SSN            | 1          | 699        | 699        | If not 1, 2, | Reason    | TOP Secondary | Indicates the type of payee identifier provided in the   | DNP       |
|           | Indicator                  |           |                    |            |            |            | [3](#v830)   | Group 5   | Payee TIN     | PayeeIdentifier_Secondaryfield                           |           |
|           |                            |           | "2"=EIN            |            |            |            | or blank,    | message 3 | Indicator     |                                                          | IPP       |
|           |                            |           |                    |            |            |            | mark payment |           |               |                                                          |           |
|           |                            |           | "3"=ITIN           |            |            |            | as invalid.  |           |               |                                                          | PIR       |
|           |                            |           |                    |            |            |            | Blank is the |           |               |                                                          |           |
|           |                            |           | or blank           |            |            |            | only         |           |               |                                                          | PPS       |
|           |                            |           |                    |            |            |            | non-numeric  |           |               |                                                          |           |
|           |                            |           |                    |            |            |            | allowed.     |           |               |                                                          | TOP NG    |
|           |                            |           |                    |            |            |            |              |           |               |                                                          |           |
|           |                            |           |                    |            |            |            |              |           |               |                                                          | DoD-DCAS  |
|           |                            |           |                    |            |            |            |              |           |               |                                                          |           |
|           |                            |           |                    |            |            |            |              |           |               |                                                          | SRF       |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------------------------+-----------+
| 12.31     | Amount eligible for offset | AN        |                    | 10         | 700        | 709        | If not all   | Reason    | Amount        | The amount of the payment that is eligible for offset.   | TOP       |
|           |                            |           |                    |            |            |            | blank or all | Group 5   | eligible for  | Will be sent to TOP.                                     |           |
|           |                            |           |                    |            |            |            | numeric,     | message 3 | offset        |                                                          |           |
|           |                            |           |                    |            |            |            | mark payment |           |               | If the entire payment is eligible for offset, pass       |           |
|           |                            |           |                    |            |            |            | as invalid.  |           |               | blanks in this field. If a portion of the payment is     |           |
|           |                            |           |                    |            |            |            | Blank is the |           |               | eligible for offset, specify the amount in this field.   |           |
|           |                            |           |                    |            |            |            | only         |           |               |                                                          |           |
|           |                            |           |                    |            |            |            | non-numeric  |           |               |                                                          |           |
|           |                            |           |                    |            |            |            | allowed.     |           |               |                                                          |           |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------+-----------+
| 12.32     | SubPaymentTypeCode         | AN        |                    | 32         | 710        | 741        | n/a          |           | Sub Payment   | Value agreed to between Agency and     | SRF                         |
|           |                            |           |                    |            |            |            |              |           | Type Code     | Treasury.                              |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.33     | PayerMechanism             | AN        | Cash, Check,       | 20         | 742        | 761        | If not one   | Reason    | Payer         |                                        | SRF                         |
|           |                            |           | BookEntry, or      |            |            |            | of the       | Group 5   | Mechanism     |                                        |                             |
|           |                            |           | blank              |            |            |            | values       | message 3 |               |                                        |                             |
|           |                            |           |                    |            |            |            | listed or    |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | blank, mark  |           |               |                                        |                             |
|           |                            |           |                    |            |            |            | invalid      |           |               |                                        |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.34     | PaymentDescriptionCode     | AN        |                    | 2          | 762        | 763        | n/a          | n/a       | Payment       | Provides additional description for    | TOP                         |
|           |                            |           |                    |            |            |            |              |           | Description   | TOP offsetting. Values agreed to       |                             |
|           |                            |           |                    |            |            |            |              |           | Code          | between Agency and TOP.                |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+
| 12.35     | Filler                     |           |                    | 87         | 764        | 850        | n/a          |           | n/a           |                                        |                             |
+-----------+----------------------------+-----------+--------------------+------------+------------+------------+--------------+-----------+---------------+----------------------------------------+-----------------------------+

##  ACH Addendum Record

This record is optional for PPD, CCD, and IAT Standard Entry Class
Codes.

If the SEC code is "PPD" or "CCD" only one addenda is allowed per
payment.

If the SEC code is "IAT" up to two remittance addenda are allowed per
payment. Additionally, for IAT payments, PAM will build the Mandatory
IAT Addenda Records using data provided in the Payment Data record and
append this remittance addenda if provided by the agency.

+--------------------------------------------------------------------------------------------------------------------------------------+
| **ACH Addendum Record**                                                                                                              |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-----------+
| **\#**   | **Field     | **Type** | **Field  | **Length** | **Start    | **End      | **Validation | **Error  | **Stored | **Notes** |
|          | Name**      |          | Value**  |            | Position** | Position** | rules**      | Code**   | Name**   |           |
+==========+=============+==========+==========+============+============+============+==============+==========+==========+===========+
| 03.01    | Record Code | AN       | "03"     | 2          | 1          | 2          | If invalid,  | Invalid: | n/a      | ACH       |
|          |             |          |          |            |            |            | missing or   | Reason   |          | Addendum  |
|          |             |          |          |            |            |            | out of       | Group 1  |          | Record =  |
|          |             |          |          |            |            |            | order,       | message  |          | "03"      |
|          |             |          |          |            |            |            | reject the   | 6.       |          |           |
|          |             |          |          |            |            |            | file         | Missing  |          |           |
|          |             |          |          |            |            |            |              | or out   |          |           |
|          |             |          |          |            |            |            |              | of       |          |           |
|          |             |          |          |            |            |            |              | order:   |          |           |
|          |             |          |          |            |            |            |              | Reason   |          |           |
|          |             |          |          |            |            |            |              | Group 1  |          |           |
|          |             |          |          |            |            |            |              | message  |          |           |
|          |             |          |          |            |            |            |              | 4.       |          |           |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-----------+
| 03.02    | PaymentID   | AN       |          | 20         | 3          | 22         | If this      | Reason   | n/a      | SRF       |
|          |             |          |          |            |            |            | information  | Group 1  |          |           |
|          |             |          |          |            |            |            | is not the   | message  |          |           |
|          |             |          |          |            |            |            | same as in   | 6        |          |           |
|          |             |          |          |            |            |            | the ACH      |          |          |           |
|          |             |          |          |            |            |            | payment data |          |          |           |
|          |             |          |          |            |            |            | record,      |          |          |           |
|          |             |          |          |            |            |            | reject the   |          |          |           |
|          |             |          |          |            |            |            | schedule.    |          |          |           |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-----------+
| 03.03    | Addenda     | AN       |          | 80         | 23         | 102        | n/a          |          | Addenda  | DNP       |
|          | Information |          |          |            |            |            |              |          | Record   |           |
|          |             |          |          |            |            |            |              |          |          | FedACH    |
|          |             |          |          |            |            |            |              |          |          |           |
|          |             |          |          |            |            |            |              |          |          | IPP       |
|          |             |          |          |            |            |            |              |          |          |           |
|          |             |          |          |            |            |            |              |          |          | PIR       |
|          |             |          |          |            |            |            |              |          |          |           |
|          |             |          |          |            |            |            |              |          |          | PPS       |
|          |             |          |          |            |            |            |              |          |          |           |
|          |             |          |          |            |            |            |              |          |          | SRF       |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-----------+
| 03.04    | Filler      |          |          | 748        | 103        | 850        | n/a          |          | n/a      |           |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-----------+

###  CTX Validation Rules

The validation rules in the table below apply for Standard Entry Class
Code "CTX".

  -----------------------------------------------------------------------
  **Rule**                **Result**              **Error**
  ----------------------- ----------------------- -----------------------
  If the SEC code is      Mark payment invalid    Reason Group 5 message
  "CTX" and the ISA, BPR,                         3
  or SE segments are                              
  missing then mark                               
  payment invalid.                                

  If the first ACH        Mark payment invalid    Reason Group 5 message
  Addendum Record in                              3
  position 23-25 does not                         
  contain "ISA" then mark                         
  payment invalid.                                

  For the ISA segment,    Mark payment invalid    Reason Group 5 message
  the character in                                3
  position 26 must equal                          
  the segment delimiter;                          
  the character in                                
  position 128 must equal                         
  the segment terminator.                         
  If the segment                                  
  delimiter equals the                            
  segment terminator then                         
  mark the payment                                
  invalid.                                        

  For the BPR segment, if Mark payment invalid    Reason Group 5 message
  BPR-02 is non-numeric                           3
  then mark the payment                           
  invalid                                         

  For the SE segment, if  Mark payment invalid    Reason Group 5 message
  SE-01 is non-numeric                            3
  then mark the payment                           
  invalid                                         

  Invalid hex cannot be                           
  used for delimiters (as                         
  defined in Section 1.3)                         
  -----------------------------------------------------------------------

### ACH Addendum Record for CTX payments

Use the following format for Standard Entry Class Code "CTX".

Contact the Treasury Outreach Team for a copy of the CTX Addenda
Supplements: FS.AgencyOutreach@fiscal.treasury.gov

For CTX, 999 ACH Addendum records with 10 Addenda are allowed.

Note: one BPR per payment is expected.

+----------------------------------------------------------------------------------------------------------------------------------------+--------------+
| **ACH Addendum Record**                                                                                                                |              |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-------------+--------------+
| **\#**   | **Field     | **Type** | **Field  | **Length** | **Start    | **End      | **Validation | **Error  | **Stored | **Notes**   | **Downstream |
|          | Name**      |          | Value**  |            | Position** | Position** | rules**      | Code**   | Name**   |             | Mapping**    |
+==========+=============+==========+==========+============+============+============+==============+==========+==========+=============+==============+
| 04.01    | Record Code | AN       | "04"     | 2          | 1          | 2          | If invalid,  | Invalid: | n/a      |             |              |
|          |             |          |          |            |            |            | missing or   | Reason   |          |             |              |
|          |             |          |          |            |            |            | out of       | Group 1  |          |             |              |
|          |             |          |          |            |            |            | order,       | message  |          |             |              |
|          |             |          |          |            |            |            | reject the   | 6.       |          |             |              |
|          |             |          |          |            |            |            | file         | Missing  |          |             |              |
|          |             |          |          |            |            |            |              | or out   |          |             |              |
|          |             |          |          |            |            |            |              | of       |          |             |              |
|          |             |          |          |            |            |            |              | order:   |          |             |              |
|          |             |          |          |            |            |            |              | Reason   |          |             |              |
|          |             |          |          |            |            |            |              | Group 1  |          |             |              |
|          |             |          |          |            |            |            |              | message  |          |             |              |
|          |             |          |          |            |            |            |              | 4.       |          |             |              |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-------------+--------------+
| 04.02    | PaymentID   | AN       |          | 20         | 3          | 22         | If this      | Reason   | n/a      |             | SRF          |
|          |             |          |          |            |            |            | information  | Group 1  |          |             |              |
|          |             |          |          |            |            |            | is not the   | message  |          |             |              |
|          |             |          |          |            |            |            | same as in   | 6        |          |             |              |
|          |             |          |          |            |            |            | the ACH      |          |          |             |              |
|          |             |          |          |            |            |            | payment data |          |          |             |              |
|          |             |          |          |            |            |            | record,      |          |          |             |              |
|          |             |          |          |            |            |            | reject the   |          |          |             |              |
|          |             |          |          |            |            |            | schedule.    |          |          |             |              |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-------------+--------------+
| 04.03    | Addenda     | AN       |          | 800        | 23         | 822        | See above    | If       | Addenda  | Place up to | FedACH       |
|          | Information |          |          |            |            |            | for          | invalid, |          | ten 80      |              |
|          |             |          |          |            |            |            | validation   | Reason   |          | character   | IPP          |
|          |             |          |          |            |            |            | rules        | Group 5  |          | Addenda.    |              |
|          |             |          |          |            |            |            |              | message  |          |             | PACER        |
|          |             |          |          |            |            |            |              | 3.       |          | Position 4  |              |
|          |             |          |          |            |            |            |              |          |          | of the ISA  | PIR          |
|          |             |          |          |            |            |            |              |          |          | segment     |              |
|          |             |          |          |            |            |            |              |          |          | will be     | PPS          |
|          |             |          |          |            |            |            |              |          |          | used as the |              |
|          |             |          |          |            |            |            |              |          |          | data        | DNP          |
|          |             |          |          |            |            |            |              |          |          | element     |              |
|          |             |          |          |            |            |            |              |          |          | separator   | SRF          |
|          |             |          |          |            |            |            |              |          |          | and         |              |
|          |             |          |          |            |            |            |              |          |          | position    |              |
|          |             |          |          |            |            |            |              |          |          | 106 will be |              |
|          |             |          |          |            |            |            |              |          |          | used as the |              |
|          |             |          |          |            |            |            |              |          |          | segment     |              |
|          |             |          |          |            |            |            |              |          |          | terminator. |              |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-------------+--------------+
| 04.04    | Filler      |          |          | 28         | 823        | 850        | n/a          |          | n/a      |             |              |
+----------+-------------+----------+----------+------------+------------+------------+--------------+----------+----------+-------------+--------------+

##  CARS TAS/BETC Record 

This record is optional. However, if TAS/BETC data received in the SPS
certification does not match TAS/BETC data received with the schedule,
the differences will be reported to CARS. Zero to 100 TAS/BETC
recommended per payment ; maximum 1000 unique TAS/BETC recommended per
schedule.

+-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **CARS TAS/BETC Record**                                                                                                                                                                                                      |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name**                     | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                                    |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+====================================+================+================+================+================+================+================+================+================+================+================+
| G.01           | Record Code                        | AN             | "G~~b~~"       | 2              | 1              | 2              | If invalid,    | Invalid:       | n/a            | CARS TAS/BETC  |                |
|                |                                    |                |                |                |                |                | missing or out | Reason Group 1 |                | Record = "G"   |                |
|                |                                    |                |                |                |                |                | of order,      | message 6.     |                | followed by    |                |
|                |                                    |                |                |                |                |                | reject the     | Missing or out |                | space          |                |
|                |                                    |                |                |                |                |                | file           | of order:      |                |                |                |
|                |                                    |                |                |                |                |                |                | Reason Group 1 |                |                |                |
|                |                                    |                |                |                |                |                |                | message 4.     |                |                |                |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.02           | PaymentID                          | AN             |                | 20             | 3              | 22             | If this        | Reason Group 1 | n/a            |                | SRF            |
|                |                                    |                |                |                |                |                | information is | message 6      |                |                |                |
|                |                                    |                |                |                |                |                | not the same   |                |                |                |                |
|                |                                    |                |                |                |                |                | as the payment |                |                |                |                |
|                |                                    |                |                |                |                |                | data record,   |                |                |                |                |
|                |                                    |                |                |                |                |                | reject the     |                |                |                |                |
|                |                                    |                |                |                |                |                | schedule.      |                |                |                |                |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.03           | Sub-levelPrefixCode                | AN             |                | 2              | 23             | 24             | n/a            |                | TAS Sub-level  |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Prefix Code    |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.04           | AllocationTransferAgencyIdentifier | AN             |                | 3              | 25             | 27             | n/a            |                | TAS Allocation |                | PACER          |
|                |                                    |                |                |                |                |                |                |                | Transfer       |                |                |
|                |                                    |                |                |                |                |                |                |                | Agency         |                | PIR            |
|                |                                    |                |                |                |                |                |                |                | Identifier     |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.05           | AgencyIdentifier                   | AN             |                | 3              | 28             | 30             | n/a            |                | TAS Agency     |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Identifier     |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.06           | BeginningPeriodOfAvailability      | AN             |                | 4              | 31             | 34             | n/a            |                | TAS Beginning  |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Period of      |                |                |
|                |                                    |                |                |                |                |                |                |                | Availability   |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.07           | EndingPeriodOfAvailability         | AN             |                | 4              | 35             | 38             | n/a            |                | TAS Ending     |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Period of      |                |                |
|                |                                    |                |                |                |                |                |                |                | Availability   |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.08           | AvailabilityTypeCode               | AN             |                | 1              | 39             | 39             | n/a            |                | TAS            |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Availability   |                |                |
|                |                                    |                |                |                |                |                |                |                | Type Code      |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.09           | MainAccountCode                    | AN             |                | 4              | 40             | 43             | n/a            |                | TAS Main       |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Account Code   |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.10           | Sub-accountCode                    | AN             |                | 3              | 44             | 46             | n/a            |                | TAS            |                | DNP            |
|                |                                    |                |                |                |                |                |                |                | Sub-Account    |                |                |
|                |                                    |                |                |                |                |                |                |                | Code           |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.11           | BusinessEventTypeCode              | AN             |                | 8              | 47             | 54             | n/a            |                | BETC           |                | DNP            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.12           | AccountClassificationAmount        | N              |                | 10             | 55             | 64             | If not numeric |                | TAS/BETC       |                | DNP            |
|                |                                    |                |                |                |                |                | default to     |                | Amount         |                |                |
|                |                                    |                |                |                |                |                | zero.          |                |                |                | PACER          |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PPS            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | PIR            |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                |                |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.13           | IsCredit                           | AN             | "0" or "1"     | 1              | 65             | 65             | If blank,      |                | IsCredit       | Used to        | DNP            |
|                |                                    |                |                |                |                |                | default to "0" |                |                | indicate       |                |
|                |                                    |                |                |                |                |                |                |                | (PAM stores    | whether amount | PIR            |
|                |                                    |                |                |                |                |                |                |                | amount as      | is a debit "0" |                |
|                |                                    |                |                |                |                |                |                |                | signed         | or credit "1". | PPS            |
|                |                                    |                |                |                |                |                |                |                | negative when  |                |                |
|                |                                    |                |                |                |                |                |                |                | IsCredit value |                | DoD-DCAS       |
|                |                                    |                |                |                |                |                |                |                | = "1")         |                |                |
|                |                                    |                |                |                |                |                |                |                |                |                | SRF            |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| G.14           | Filler                             |                |                | 785            | 66             | 850            | n/a            |                | n/a            |                |                |
+----------------+------------------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

##  Check Stub Record 

This record is only required if the CheckEnclosureCode field in the
payment data record is "stub". Agencies must have approval from their
RFC before using this service.

+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+--------+
| **Check Stub Record**                                                                                                                                                                    |        |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+--------------------+----------------+----------------+----------------+--------+
| **\#**         | **Field Name**              | **Type**       | **Field        | **Length**     | **Position**   | **Validation       | **Error Code** | **Stored       | **Notes**      |        |
|                |                             |                | Value**        |                |                | rules**            |                | Name**         |                |        |
+================+=============================+================+================+================+================+====================+================+================+:==============:+:======:+
| 13.01          | Record Code                 | AN             | "13"           | 2              | 1-2            | If                 | Missing or out | n/a            | Check Stub     |        |
|                |                             |                |                |                |                | CheckEnclosureCode | of order:      |                | record code =  |        |
|                |                             |                |                |                |                | = "stub" and this  | Reason Group 1 |                | "13"           |        |
|                |                             |                |                |                |                | record is missing  | message 4.     |                |                |        |
|                |                             |                |                |                |                | or out of order,   | Invalid:       |                |                |        |
|                |                             |                |                |                |                | reject the file.   | Reason Group 1 |                |                |        |
|                |                             |                |                |                |                | If value is        | message 6.     |                |                |        |
|                |                             |                |                |                |                | invalid, reject    |                |                |                |        |
|                |                             |                |                |                |                | the file.          |                |                |                |        |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+--------------------+----------------+----------------+----------------+--------+
| 13.02          | PaymentID                   | AN             |                | 20             | 3-22           | If this            | Reason Group 1 | n/a            | SRF            |        |
|                |                             |                |                |                |                | information is not | message 6      |                |                |        |
|                |                             |                |                |                |                | the same as the    |                |                |                |        |
|                |                             |                |                |                |                | payment data       |                |                |                |        |
|                |                             |                |                |                |                | record, reject the |                |                |                |        |
|                |                             |                |                |                |                | schedule.          |                |                |                |        |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+--------------------+----------------+----------------+----------------+--------+
| 13.03          | PaymentIdentificationLine_1 | AN             |                | 55             | 23-77          | n/a                |                | Payment        | Checks are     |        |
|                | (through 14)                |                |                |                |                |                    |                | Identification | mailed to the  |        |
|                |                             |                |                |                | 78-132         |                    |                | Line 1 -- 14   | address in the |        |
|                |                             |                |                |                |                |                    |                |                | payment data   |        |
|                |                             |                |                |                | 133-187        |                    |                |                | record, not    |        |
|                |                             |                |                |                |                |                    |                |                | the stub       |        |
|                |                             |                |                |                | 188-242        |                    |                |                | address.       |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 243-297        |                    |                |                | SRF            |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 298-352        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 353-407        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 408-462        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 463-517        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 518-572        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 573-627        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 628-682        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 683-737        |                    |                |                |                |        |
|                |                             |                |                |                |                |                    |                |                |                |        |
|                |                             |                |                |                | 738-792        |                    |                |                |                |        |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+--------------------+----------------+----------------+----------------+--------+
| 13.04          | Filler                      |                |                | 58             | 793-850        | n/a                |                | n/a            |                |        |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+--------------------+----------------+----------------+----------------+--------+

##  Procurement Record (deleted)

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
|                                                                                                                                                                                                           |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+================+================+================+================+================+================+================+================+================+================+================+================+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
|                |                |                |                |                |                |                |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

##  DNP Record 

This record is optional. Agencies must work with DNP to determine exact
values to send in DNP Detail field.

+-----------------------------------------------------------------------------------------------------------------------------------------+
| **DNP Record**                                                                                                                          |
+-----------+-----------+-----------+-----------+------------+------------+------------+--------------+-----------+-----------+-----------+
| **\#**    | **Field   | **Type**  | **Field   | **Length** | **Start    | **End      | **Validation | **Error   | **Stored  | **Notes** |
|           | Name**    |           | Value**   |            | Position** | Position** | rules**      | Code**    | Name**    |           |
+===========+===========+===========+===========+============+============+============+==============+===========+===========+===========+
| DD.01     | Record    | AN        | "DD"      | 2          | 1          | 2          | If invalid,  | Invalid:  | n/a       |           |
|           | Code      |           |           |            |            |            | missing or   | Reason    |           |           |
|           |           |           |           |            |            |            | out of       | Group 1   |           |           |
|           |           |           |           |            |            |            | order,       | message   |           |           |
|           |           |           |           |            |            |            | reject the   | 6.        |           |           |
|           |           |           |           |            |            |            | file         | Missing   |           |           |
|           |           |           |           |            |            |            |              | or out of |           |           |
|           |           |           |           |            |            |            |              | order:    |           |           |
|           |           |           |           |            |            |            |              | Reason    |           |           |
|           |           |           |           |            |            |            |              | Group 1   |           |           |
|           |           |           |           |            |            |            |              | message   |           |           |
|           |           |           |           |            |            |            |              | 4.        |           |           |
+-----------+-----------+-----------+-----------+------------+------------+------------+--------------+-----------+-----------+-----------+
| DD.02     | Payment   | AN        |           | 20         | 3          | 22         | If this      | Reason    | n/a       | SRF       |
|           | ID        |           |           |            |            |            | information  | Group 1   |           |           |
|           |           |           |           |            |            |            | is not the   | message 6 |           |           |
|           |           |           |           |            |            |            | same as the  |           |           |           |
|           |           |           |           |            |            |            | payment data |           |           |           |
|           |           |           |           |            |            |            | record,      |           |           |           |
|           |           |           |           |            |            |            | reject the   |           |           |           |
|           |           |           |           |            |            |            | schedule.    |           |           |           |
+-----------+-----------+-----------+-----------+------------+------------+------------+--------------+-----------+-----------+-----------+
| DD.03     | DNP       | AN        |           | 766        | 23         | 788        | Store value  |           | DNP       | DNP       |
|           | Detail    |           |           |            |            |            | as received, |           | Detail    |           |
|           |           |           |           |            |            |            | do not       |           |           | PPS       |
|           |           |           |           |            |            |            | justify      |           |           |           |
|           |           |           |           |            |            |            |              |           |           | DoD-DCAS  |
|           |           |           |           |            |            |            |              |           |           |           |
|           |           |           |           |            |            |            |              |           |           | SRF       |
+-----------+-----------+-----------+-----------+------------+------------+------------+--------------+-----------+-----------+-----------+
| DD.04     | Filler    |           |           | 62         | 789        | 850        | N/A          |           |           |           |
+-----------+-----------+-----------+-----------+------------+------------+------------+--------------+-----------+-----------+-----------+

##  Schedule Trailer Control Record

+----------------------------------------------------------------------------------------------------------------------------------------------------+
| **Schedule Trailer Control Record**                                                                                                                |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+
| **\#**     | **Field Name** | **Type**   | **Field    | **Length** | **Start    | **End      | **Validation | **Error    | **Stored   | **Notes**  |
|            |                |            | Value**    |            | Position** | Position** | rules**      | Code**     | Name**     |            |
+============+================+============+============+============+============+============+==============+============+============+============+
| T.01       | Record Code    | AN         | "T~~b~~"   | 2          | 1          | 2          | If invalid,  | Invalid:   | n/a        | Schedule   |
|            |                |            |            |            |            |            | missing or   | Reason     |            | Trailer    |
|            |                |            |            |            |            |            | out of       | Group 1    |            | Record =   |
|            |                |            |            |            |            |            | order,       | message 6. |            | "T"        |
|            |                |            |            |            |            |            | reject the   | Missing or |            | followed   |
|            |                |            |            |            |            |            | file         | out of     |            | by space   |
|            |                |            |            |            |            |            |              | order:     |            |            |
|            |                |            |            |            |            |            |              | Reason     |            |            |
|            |                |            |            |            |            |            |              | Group 1    |            |            |
|            |                |            |            |            |            |            |              | message 4. |            |            |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+
| T.02       | Filler         | AN         | Blanks     | 10         | 3          | 12         | n/a          |            |            |            |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+
| T.03       | ScheduleCount  | N          |            | 8          | 13         | 20         | If invalid,  | Reason     | For check  | See        |
|            |                |            |            |            |            |            | reject the   | Group 1    | schedule   | Validation |
|            |                |            |            |            |            |            | schedule.    | message 6  | store as   | for        |
|            |                |            |            |            |            |            |              |            | "Total     | Balancing  |
|            |                |            |            |            |            |            |              |            | Check      | for        |
|            |                |            |            |            |            |            |              |            | Payment    | additional |
|            |                |            |            |            |            |            |              |            | Items".    | rules.     |
|            |                |            |            |            |            |            |              |            | For ACH    |            |
|            |                |            |            |            |            |            |              |            | schedule   |            |
|            |                |            |            |            |            |            |              |            | store as   |            |
|            |                |            |            |            |            |            |              |            | "Total ACH |            |
|            |                |            |            |            |            |            |              |            | Payment    |            |
|            |                |            |            |            |            |            |              |            | Items"     |            |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+
| T.04       | Filler         | AN         | Blanks     | 3          | 21         | 23         | n/a          |            |            |            |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+
| T.05       | ScheduleAmount | N          |            | 15         | 24         | 38         | If invalid,  | Reason     | For Check  | See        |
|            |                |            |            |            |            |            | reject the   | Group 1    | schedule,  | Validation |
|            |                |            |            |            |            |            | schedule.    | message 6  | store as   | for        |
|            |                |            |            |            |            |            |              |            | "Total     | Balancing  |
|            |                |            |            |            |            |            |              |            | Check      | for        |
|            |                |            |            |            |            |            |              |            | Payment    | additional |
|            |                |            |            |            |            |            |              |            | Amount".   | rules.     |
|            |                |            |            |            |            |            |              |            | For ACH    |            |
|            |                |            |            |            |            |            |              |            | schedule,  |            |
|            |                |            |            |            |            |            |              |            | store as   |            |
|            |                |            |            |            |            |            |              |            | "Total ACH |            |
|            |                |            |            |            |            |            |              |            | Payment    |            |
|            |                |            |            |            |            |            |              |            | Amount"    |            |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+
| T.06       | Filler         |            |            | 812        | 39         | 850        | n/a          |            | n/a        |            |
+------------+----------------+------------+------------+------------+------------+------------+--------------+------------+------------+------------+

##  File Trailer Control Record

+------------------------------------------------------------------------------------------------------------------------------------------------+
| **File Trailer Control Record**                                                                                                                |
+----------+----------------------+----------+----------+------------+------------+------------+--------------+----------+----------+------------+
| **\#**   | **Field Name**       | **Type** | **Field  | **Length** | **Start    | **End      | **Validation | **Error  | **Stored | **Notes**  |
|          |                      |          | Value**  |            | Position** | Position** | rules**      | Code**   | Name**   |            |
+==========+======================+==========+==========+============+============+============+==============+==========+==========+============+
| E.01     | Record Code          | AN       | "E~~b~~" | 2          | 1          | 2          | If invalid,  | Invalid: |          | File       |
|          |                      |          |          |            |            |            | missing or   | Reason   |          | Trailer    |
|          |                      |          |          |            |            |            | out of       | Group 1  |          | Record =   |
|          |                      |          |          |            |            |            | order,       | message  |          | "E"        |
|          |                      |          |          |            |            |            | reject the   | 6.       |          | followed   |
|          |                      |          |          |            |            |            | file         | Missing  |          | by a space |
|          |                      |          |          |            |            |            |              | or out   |          |            |
|          |                      |          |          |            |            |            |              | of       |          |            |
|          |                      |          |          |            |            |            |              | order:   |          |            |
|          |                      |          |          |            |            |            |              | Reason   |          |            |
|          |                      |          |          |            |            |            |              | Group 1  |          |            |
|          |                      |          |          |            |            |            |              | message  |          |            |
|          |                      |          |          |            |            |            |              | 4.       |          |            |
+----------+----------------------+----------+----------+------------+------------+------------+--------------+----------+----------+------------+
| E.02     | TotalCount_Records   | N        |          | 18         | 3          | 20         | If invalid,  | Reason   |          | Includes   |
|          |                      |          |          |            |            |            | reject the   | Group 1  |          | header and |
|          |                      |          |          |            |            |            | file         | message  |          | trailer    |
|          |                      |          |          |            |            |            |              | 6        |          | records    |
|          |                      |          |          |            |            |            |              |          |          |            |
|          |                      |          |          |            |            |            |              |          |          | See        |
|          |                      |          |          |            |            |            |              |          |          | Validation |
|          |                      |          |          |            |            |            |              |          |          | for        |
|          |                      |          |          |            |            |            |              |          |          | Balancing  |
|          |                      |          |          |            |            |            |              |          |          | for        |
|          |                      |          |          |            |            |            |              |          |          | additional |
|          |                      |          |          |            |            |            |              |          |          | rules.     |
+----------+----------------------+----------+----------+------------+------------+------------+--------------+----------+----------+------------+
| E.03     | TotalCount_Payments  | N        |          | 18         | 21         | 38         | If invalid,  | Reason   | Total    | See        |
|          |                      |          |          |            |            |            | reject the   | Group 1  | File     | Validation |
|          |                      |          |          |            |            |            | file.        | message  | payment  | for        |
|          |                      |          |          |            |            |            |              | 6        | Items    | Balancing  |
|          |                      |          |          |            |            |            |              |          |          | for        |
|          |                      |          |          |            |            |            |              |          |          | additional |
|          |                      |          |          |            |            |            |              |          |          | rules.     |
+----------+----------------------+----------+----------+------------+------------+------------+--------------+----------+----------+------------+
| E.04     | TotalAmount_Payments | N        |          | 18         | 39         | 56         | If invalid,  | Reason   | Total    | See        |
|          |                      |          |          |            |            |            | reject the   | Group 1  | File     | Validation |
|          |                      |          |          |            |            |            | file.        | message  | Payment  | for        |
|          |                      |          |          |            |            |            |              | 6        | Amount   | Balancing  |
|          |                      |          |          |            |            |            |              |          |          | for        |
|          |                      |          |          |            |            |            |              |          |          | additional |
|          |                      |          |          |            |            |            |              |          |          | rules.     |
+----------+----------------------+----------+----------+------------+------------+------------+--------------+----------+----------+------------+
| E.05     | Filler               |          |          | 794        | 57         | 850        | n/a          |          | n/a      |            |
+----------+----------------------+----------+----------+------------+------------+------------+--------------+----------+----------+------------+

# Appendices

## Appendix A - ACH Transaction Codes

The following table defines the Transaction Code values received in an
ACH Detail Payment record and the value PAM should record.

  ---------------------------------------------------------------------------
   **Transaction  **Definition**                          **Account Type (for
      Code**                                              recording in PAM)**
  --------------- --------------------------------------- -------------------
        22        Checking Account Credit                          C

        23        Checking Account Credit Prenote                  C

        24        Checking Account Zero Dollar Credit              C
                  with remittance                         

        32        Savings Account Credit                           S

        33        Savings Account Credit Prenote                   S

        34        Savings Account Zero Dollar Credit with          S
                  remittance                              

        42        General Ledger Credit                            G

        43        General Ledger Credit Prenote                    G

        52        Loan Account Credit                              L

        53        Loan Account Credit Prenote                      L
  ---------------------------------------------------------------------------

For Custom Agency Rule ID = "IRS" and Depositor Account Number =
"BONDS":

  ---------------------------------------------------------------------------
   **Transaction  **Definition**             **Account Type       **Savings
      Code**                                 (for recording in   Bonds Owner
                                             PAM)**                Type**
  --------------- -------------------------- ------------------ -------------
        22        Checking Account Credit    C                      Gift

        32        Savings Account Credit     S                      Self
  ---------------------------------------------------------------------------

## Appendix B -- Agency Specific Values

This appendix includes derivation, parsing and storage rules for those
agencies configured with a Custom Agency Rule ID in the RAT
configuration. These rules are used to parse and store the reconcilement
field in position 279-378 in the ACH Data Record and position 489-588 in
the Check Data Record, and to derive the values for specific data
elements as noted below. If the Custom Agency Rule ID has a value the
following rules apply.

### For Custom Agency Rule ID = "IRS" and Depositor Account Number NOT equal "BONDS"

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **IRS Tax Refunds Reconcilement Field**                                                                                                                                                                   |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1              | Tax Period     | AN             |                | 4              | 1              | 4              | n/a            |                | IRS Tax Yr     |                | PACER, TOP,    |
|                | (Year)         |                |                |                |                |                |                |                | Date           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2              | Tax Period     | AN             |                | 2              | 5              | 6              | n/a            |                | IRS Tax Mo     |                | PACER, TOP,    |
|                | (Month)        |                |                |                |                |                |                |                | Date           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3              | Master File    | AN             |                | 2              | 7              | 8              | n/a            |                | IRS MFT Code   |                | PACER, TOP,    |
|                | Tax (MFT) Code |                |                |                |                |                |                |                |                |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 4              | Service Center | AN             |                | 2              | 9              | 10             | n/a            |                | IRS Serv Cntr  |                | PACER, TOP,    |
|                | Code           |                |                |                |                |                |                |                | Code           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 5              | District       | AN             |                | 2              | 11             | 12             | n/a            |                | IRS Dist Ofc   |                | PACER, TOP,    |
|                | Office Code    |                |                |                |                |                |                |                | Code           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 6              | File TIN Code  | AN             |                | 1              | 13             | 13             | n/a            |                | IRS TIN Code   |                | PACER, TOP,    |
|                |                |                |                |                |                |                |                |                |                |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 7              | Name Control   | AN             |                | 4              | 14             | 17             | n/a            |                | IRS Name Ctrl  |                | PACER, TOP,    |
|                |                |                |                |                |                |                |                |                |                |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 8              | Plan Report    | AN             |                | 3              | 18             | 20             | n/a            |                | IRS Plan Rpt   |                | PACER, TOP,    |
|                | Number         |                |                |                |                |                |                |                | Num            |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 9              | Split Refund   | AN             |                | 1              | 21             | 21             |                |                | IRS Split      |                | PACER, TOP     |
|                | Code           |                |                |                |                |                |                |                | Refund Code    |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 10             | Injured Spouse | AN             |                | 1              | 22             | 22             | n/a            |                | IRS Injured    |                | TOP,           |
|                | Code           |                |                |                |                |                |                |                | Spouse         |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 11             | Debt Bypass    | AN             |                | 1              | 23             | 23             | n/a            |                | TOP            | If there is a  | TOP            |
|                | Indicator Code |                |                |                |                |                |                |                | Eligibility    | value in this  |                |
|                |                |                |                |                |                |                |                |                | Indicator      | field, store   |                |
|                |                |                |                |                |                |                |                |                |                | this value     |                |
|                |                |                |                |                |                |                |                |                |                | instead of the |                |
|                |                |                |                |                |                |                |                |                |                | value received |                |
|                |                |                |                |                |                |                |                |                |                | in the payment |                |
|                |                |                |                |                |                |                |                |                |                | data record.   |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 12             | Document       | AN             |                | 14             | 24             | 37             | n/a            |                | IRS Document   |                | TOP            |
|                | Locator Number |                |                |                |                |                |                |                | Locator Number |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 13             | Check Detail   | A              | "enclose" or   | 10             | 38             | 47             | If the field   |                | Enclosure Code |                |                |
|                | Enclosure Code |                | blank          |                |                |                | value is NOT   |                |                |                |                |
|                |                |                |                |                |                |                | "enclose" or   |                |                |                |                |
|                |                |                |                |                |                |                | blank, or      |                |                |                |                |
|                |                |                |                |                |                |                | Payment is ACH |                |                |                |                |
|                |                |                |                |                |                |                | make the field |                |                |                |                |
|                |                |                |                |                |                |                | value blank.   |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 14             | Filler         |                |                | 53             | 48             | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

#### Payment Legacy Account Symbol Derivation Rules for IRS

> The following table describes the derivation of the Legacy Account
> Symbol value based on listed fields in the IRS Reconcilement Record.
> The Legacy Account Symbol value is a concatenation (with no delimiter)
> of the values in the following fields in the respective order of the
> table below.
>
> Record the derived value in the Check Payment Detail's Legacy Account
> Symbol attribute.

+-------------------------------------+----------------------------------------------+
| **IRS Tax Refunds Reconcilement     | **Field Length**                             |
| Field**                             |                                              |
+:===================================:+:============================================:+
| District Office Code                | 2                                            |
+-------------------------------------+----------------------------------------------+
| File TIN Code                       | 1                                            |
+-------------------------------------+----------------------------------------------+
| Master File Tax (MFT) Code          | 2                                            |
+-------------------------------------+----------------------------------------------+
| Tax Period (Year)                   | 2 (use the last two digits of the Tax Period |
|                                     | (Year) field, which is a length of 4)        |
+-------------------------------------+----------------------------------------------+
| Tax Period (Month)                  | 2                                            |
+-------------------------------------+----------------------------------------------+
| Name Control                        | 4                                            |
+-------------------------------------+----------------------------------------------+
| Plan Report Number                  | 3                                            |
+-------------------------------------+----------------------------------------------+
| Sample Legacy Account Symbol value: 650300912RODR123                               |
+------------------------------------------------------------------------------------+

### For Custom Agency Rule ID = "IRS" and Depositor Account Number = "BONDS"

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **IRS Savings Bonds Orders Reconcilement Field**                                                                                                                                                          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1              | Tax Period     | AN             |                | 4              | 1              | 4              | n/a            |                | IRS Tax Yr     |                | PACER, TOP,    |
|                | (Year)         |                |                |                |                |                |                |                | Date           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2              | Tax Period     | AN             |                | 2              | 5              | 6              | n/a            |                | IRS Tax Mo     |                | PACER, TOP,    |
|                | (Month)        |                |                |                |                |                |                |                | Date           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3              | Master File    | AN             |                | 2              | 7              | 8              | n/a            |                | IRS MFT Code   |                | PACER, TOP,    |
|                | Tax (MFT) Code |                |                |                |                |                |                |                |                |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 4              | Service Center | AN             |                | 2              | 9              | 10             | n/a            |                | IRS Serv Cntr  |                | PACER, TOP,    |
|                | Code           |                |                |                |                |                |                |                | Code           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 5              | District       | AN             |                | 2              | 11             | 12             | n/a            |                | IRS Dist Ofc   |                | PACER, TOP,    |
|                | Office Code    |                |                |                |                |                |                |                | Code           |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 6              | File TIN Code  | AN             |                | 1              | 13             | 13             | n/a            |                | IRS TIN Code   |                | PACER, TOP,    |
|                |                |                |                |                |                |                |                |                |                |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 7              | Name Control   | AN             |                | 4              | 14             | 17             | n/a            |                | IRS Name Ctrl  |                | PACER, TOP,    |
|                |                |                |                |                |                |                |                |                |                |                | TCS            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 8              | Plan Report    | AN             |                | 3              | 18             | 20             | n/a            |                | IRS Plan Rpt   |                | PACER, TCS,    |
|                | Number         |                |                |                |                |                |                |                | Num            |                | TOP            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 9              | Split Refund   | AN             |                | 1              | 21             | 21             |                |                | IRS Split      |                | TOP, PACER     |
|                | Code           |                |                |                |                |                |                |                | Refund Code    |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 10             | Injured Spouse | AN             |                | 1              | 22             | 22             | n/a            |                | IRS Injured    |                | TOP,           |
|                | Code           |                |                |                |                |                |                |                | Spouse         |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 11             | Debt Bypass    | AN             |                | 1              | 23             | 23             | n/a            |                | TOP            | If there is a  | TOP            |
|                | Indicator Code |                |                |                |                |                |                |                | Eligibility    | value in this  |                |
|                |                |                |                |                |                |                |                |                | Indicator      | field, store   |                |
|                |                |                |                |                |                |                |                |                |                | this value     |                |
|                |                |                |                |                |                |                |                |                |                | instead of the |                |
|                |                |                |                |                |                |                |                |                |                | value received |                |
|                |                |                |                |                |                |                |                |                |                | in the payment |                |
|                |                |                |                |                |                |                |                |                |                | data record.   |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 12             | Bond Name 1    | AN             |                | 33             | 24             | 56             | n/a            |                | Bond Name 1    |                | SnAP           |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 13             | Bond REG Code  | AN             | "0", "1",      | 1              | 57             | 57             | n/a            |                | Bond REG Code  | Bond REG Code  | SnAP           |
|                |                |                | Blank          |                |                |                |                |                |                | Indicator      |                |
|                |                |                |                |                |                |                |                |                |                | Values:        |                |
|                |                |                |                |                |                |                |                |                |                |                |                |
|                |                |                |                |                |                |                |                |                |                | - 0 = Co-Owner |                |
|                |                |                |                |                |                |                |                |                |                |                |                |
|                |                |                |                |                |                |                |                |                |                | - 1 = POD      |                |
|                |                |                |                |                |                |                |                |                |                |                |                |
|                |                |                |                |                |                |                |                |                |                | **BLANK =** No |                |
|                |                |                |                |                |                |                |                |                |                | second name    |                |
|                |                |                |                |                |                |                |                |                |                | owner          |                |
|                |                |                |                |                |                |                |                |                |                |                |                |
|                |                |                |                |                |                |                |                |                |                | (this is the   |                |
|                |                |                |                |                |                |                |                |                |                | connector)     |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 14             | Bond Name 2    | AN             |                | 33             | 58             | 90             |                |                | Bond Name 2    |                | SnAP           |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 15             | Filler         | AN             |                | 10             | 91             | 100            |                |                |                |                | SnAP           |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

### For Custom Agency Rule ID = "VA" or "VACP" and Check payments

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **VA Education, C&P, and Insurance Check Reconcilement Field**                                                                                                                                            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1.             | VA Stn Code    | AN             |                | 2              | 1              | 2              | n/a            |                | VA Stn Code    |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | VA Fin Code    | AN             |                | 2              | 3              | 4              | n/a            |                | VA Fin Code    |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3.             | VA Courtesy    | AN             |                | 1              | 5              | 5              | n/a            |                | VA Courtesy    |                | PACER          |
|                | Code           |                |                |                |                |                |                |                | Code           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 4.             | VA Approp Code | AN             |                | 1              | 6              | 6              | n/a            |                | VA Approp Code |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 5.             | VA Address Seq | AN             |                | 1              | 7              | 7              | n/a            |                | VA Address Seq |                | PACER          |
|                | Code           |                |                |                |                |                |                |                | Code           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 6.             | VA Polcy Pre   | AN             |                | 2              | 8              | 9              | n/a            |                | VA Polcy Pre   |                | PACER          |
|                | Code           |                |                |                |                |                |                |                | Code           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 7.             | VA Polcy Num   | AN             |                | 2              | 10             | 11             | n/a            |                | VA Polcy Num   |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 8.             | VA Pay Period  | AN             |                | 12             | 12             | 23             | n/a            |                | VA Pay Period  |                | PACER          |
|                | Info           |                |                |                |                |                |                |                | Info           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 9.             | VA Name Code   | AN             |                | 3              | 24             | 26             | n/a            |                | VA Name Code   |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 10.            | Filler         |                |                | 74             | 27             | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

### For Custom Agency Rule ID = "VA" or "VACP" and ACH payments

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **VA Education, C&P, and Insurance ACH Reconcilement Field**                                                                                                                                              |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1.             | VA Stn Code    | AN             |                | 2              | 1              | 2              | n/a            |                | VA Stn Code    |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | VA Fin Code    | AN             |                | 2              | 3              | 4              | n/a            |                | VA Fin Code    |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3.             | VA Approp Code | AN             |                | 1              | 5              | 5              | n/a            |                | VA Approp Code |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 4.             | VA Address Seq | AN             |                | 1              | 6              | 6              | n/a            |                | VA Address Seq |                | PACER          |
|                | Code           |                |                |                |                |                |                |                | Code           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 5.             | VA Polcy Pre   | AN             |                | 2              | 7              | 8              | n/a            |                | VA Polcy Pre   |                | PACER          |
|                | Code           |                |                |                |                |                |                |                | Code           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 6.             | VA Polcy Num   | AN             |                | 2              | 9              | 10             | n/a            |                | VA Polcy Num   |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 7.             | VA Pay Period  | AN             |                | 12             | 11             | 22             | n/a            |                | VA Pay Period  |                | PACER          |
|                | Info           |                |                |                |                |                |                |                | Info           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 8.             | Filler         |                |                | 78             | 23             | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

###  For Custom Agency Rule ID = "SSA" and "SSA-Daily"

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **SSA Daily and Monthly Benefit Reconcilement Field**                                                                                                                                                     |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1.             | Program        | AN             |                | 1              | 1              | 1              | n/a            |                | SSA Program    | **Refer to the | PACER, TOP,    |
|                | Service Center |                |                |                |                |                |                |                | Service Center | "Payment ALC   | TCS            |
|                | Code (PSC)     |                |                |                |                |                |                |                | Code (PSC)     | Derivation     |                |
|                |                |                |                |                |                |                |                |                |                | Rules" section |                |
|                |                |                |                |                |                |                |                |                |                | below.**       |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | Payment ID     | AN             |                | 2              | 2              | 3              | n/a            |                | SSA Payment ID |                | TOP            |
|                | Code (PIC)     |                |                |                |                |                |                |                | Code (PIC)     |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3.             | TIN Indicator  | AN             |                | 1              | 4              | 4              | n/a            |                | TOP            | If there is a  | TOP            |
|                | Offset         |                |                |                |                |                |                |                | Eligibility    | value in this  |                |
|                |                |                |                |                |                |                |                |                | Indicator      | field, store   |                |
|                |                |                |                |                |                |                |                |                |                | this value     |                |
|                |                |                |                |                |                |                |                |                |                | instead of the |                |
|                |                |                |                |                |                |                |                |                |                | value received |                |
|                |                |                |                |                |                |                |                |                |                | in the payment |                |
|                |                |                |                |                |                |                |                |                |                | data record..  |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 4.             | Filler         |                |                | 96             | 5              | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

### For Custom Agency Rule ID = "SSA-A"

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **SSA Allotments Reconcilement Field**                                                                                                                                                                    |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1.             | Program        | AN             |                | 1              | 1              | 1              | n/a            |                | SSA Program    | **Refer to the | PACER, TOP,    |
|                | Service Center |                |                |                |                |                |                |                | Service Center | "Payment ALC   | TCS            |
|                | Code (PSC)     |                |                |                |                |                |                |                | Code (PSC)     | Derivation     |                |
|                |                |                |                |                |                |                |                |                |                | Rules" section |                |
|                |                |                |                |                |                |                |                |                |                | below.**       |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | Payment ID     | AN             |                | 2              | 2              | 3              | n/a            |                | SSA Payment ID |                | TOP            |
|                | Code (PIC)     |                |                |                |                |                |                |                | Code (PIC)     |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3.             | Filler         |                |                | 97             | 4              | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

### For Custom Agency Rule ID = "SSA", "SSA-A", and "SSA-Daily"

For the above two reconcilement records, the following describe
derivation rules for payment detail attribute values.

#### Payment ALC Derivation Rules

The following table lists the Payment ALC for various Program Service
Center (PSC) values in a Reconcilement Record.

If the SSA PSC Code is not 1-8 or A-H, then use the Schedule's ALC as
the Payment ALC.

Record the derived value in the Payment Detail's ALC attribute \[whether
ACH or Check payment\].

The Gray Shaded area is information only, and does not impact deriving
the ALC.

+---------+---------+-------------+--------------+----------------+-------------+--------------+
| **RSI   | **DI    | **Payment** | **PSC**      | **All          | **Payment   | **RFC**      |
| Code**  | Code**  |             |              | Certifications | Aftermath** |              |
|         |         | **Agency    |              | except Cycle   |             |              |
| **PSC** | **PSC** | Location    |              | 2, 3 and 4**   |             |              |
|         |         | Code        |              |                |             |              |
|         |         | (ALC)**     |              |                |             |              |
+=========+=========+=============+==============+================+=============+==============+
| 1       | A       | 28045200    | North        |                | X           |              |
|         |         |             | Eastern PSC  |                |             |              |
|         |         |             | 1            |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| 2       | B       | 28045300    | Mid-Atlantic | X              | X           | Philadelphia |
|         |         |             | PSC 2        |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| 3       | C       | 28045400    | South        |                | X           |              |
|         |         |             | Eastern PSC  |                |             |              |
|         |         |             | 3            |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| 4       | D       | 28045500    | Great Lakes  |                | X           |              |
|         |         |             | PSC 4        |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| 5       | E       | 28045900    | Western PSC  | X              | X           | San Fran     |
|         |         |             | 5            |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| 6       | F       | 28045600    | Mid-America  | X              | X           | Kansas City  |
|         |         |             | PSC 6        |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| 7 or 8  | G or H  | 28043000    | ODO/OIO PSC  | X              | X           | Philadelphia |
|         |         |             | 7/8          |                |             |              |
+---------+---------+-------------+--------------+----------------+-------------+--------------+
| \--     | \--     | 28040004    | SSI          | X              | X           | Kansas City  |
+---------+---------+-------------+--------------+----------------+-------------+--------------+

#### Payment Legacy Account Symbol Derivation Rules

The following table lists the Legacy Account Symbol value for various
Program Service Center values in a Reconcilement Record.

Map the derived value to the Legacy Account Symbol attribute for ACH and
Check.

  -----------------------------------------------------------------------
   **Program Service Center Code**        **Legacy Account Symbol**
  ---------------------------------- ------------------------------------
                1 -- 8                           2828X8006000

                A - H                            2828X8007000
  -----------------------------------------------------------------------

### For Custom Agency Rule ID = "RRB"

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **RRB Daily and Monthly Benefit Reconcilement Field**                                                                                                                                                     |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1.             | Beneficiary    | AN             |                | 2              | 1              | 2              | n/a            |                | RRB            |                | TOP            |
|                | Symbol         |                |                |                |                |                |                |                | Beneficiary    |                |                |
|                |                |                |                |                |                |                |                |                | Symbol         |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | Prefix Code    | AN             |                | 1              | 3              | 3              | n/a            |                | RRB Prefix     |                | TOP            |
|                |                |                |                |                |                |                |                |                | Code           |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3.             | Payee Code     | AN             |                | 1              | 4              | 4              | n/a            |                | RRB Payee Code |                | TOP            |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 4.             | Object Code    | AN             |                | 1              | 5              | 5              | n/a            |                | Object Code    |                | PACER          |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 5.             | Filler         |                |                | 95             | 6              | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

### For Custom Agency Rule ID = "CCC"

+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **CCC Reconcilement Field**                                                                                                                                                                               |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name** | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+================+================+================+================+================+================+================+================+================+================+================+
| 1.             | TOP Payment    | A              |                | 2              | 1              | 2              | If one payment | n/a            | TOP Payment    |                | TOP            |
|                | Agency ID      |                |                |                |                |                | within a       |                | Agency ID      |                |                |
|                |                |                |                |                |                |                | schedule       |                |                |                |                |
|                |                |                |                |                |                |                | contains this  |                |                |                |                |
|                |                |                |                |                |                |                | value, the     |                |                |                |                |
|                |                |                |                |                |                |                | entire         |                |                |                |                |
|                |                |                |                |                |                |                | schedule will  |                |                |                |                |
|                |                |                |                |                |                |                | be sent to TOP |                |                |                |                |
|                |                |                |                |                |                |                | with this      |                |                |                |                |
|                |                |                |                |                |                |                | value.         |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | TOP Agency     | A              |                | 2              | 3              | 4              | If one payment | n/a            | TOP Agency     |                | TOP            |
|                | Site ID        |                |                |                |                |                | within a       |                | Site ID        |                |                |
|                |                |                |                |                |                |                | schedule       |                |                |                |                |
|                |                |                |                |                |                |                | contains this  |                |                |                |                |
|                |                |                |                |                |                |                | value, the     |                |                |                |                |
|                |                |                |                |                |                |                | entire         |                |                |                |                |
|                |                |                |                |                |                |                | schedule will  |                |                |                |                |
|                |                |                |                |                |                |                | be sent to TOP |                |                |                |                |
|                |                |                |                |                |                |                | with this      |                |                |                |                |
|                |                |                |                |                |                |                | value.         |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 3.             | Filler         | AN             |                | 96             | 5              | 100            |                |                |                |                |                |
+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

### Generic Reconcilement Field

Format should be used by agencies that need to pass the Legacy Treasury
Account Symbol field to TCIS and do not have an existing reconcilement
field layout. PAM will record the value received for both ACH and Check
payments.

+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
| **Generic Reconcilement Field**                                                                                                                                                                                        |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| **\#**         | **Field Name**              | **Type**       | **Field        | **Length**     | **Start        | **End          | **Validation   | **Error Code** | **Stored       | **Notes**      | **Downstream   |
|                |                             |                | Value**        |                | Position**     | Position**     | rules**        |                | Name**         |                | Mapping**      |
+================+=============================+================+================+================+================+================+================+================+================+================+================+
| 1.             | LegacyTreasuryAccountSymbol | AN             |                | 16             | 1              | 16             |                | n/a            | Legacy Account |                | TCIS           |
|                |                             |                |                |                |                |                |                |                | Symbol         |                |                |
|                |                             |                |                |                |                |                |                |                |                |                | PPS            |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+
| 2.             | Filler                      | AN             |                | 84             | 17             | 100            |                |                |                |                |                |
+----------------+-----------------------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+----------------+

##  Appendix C - Addressing Reference Information

The following information describes how address fields should be
populated for check payments in order to ensure correct addressing and
timely delivery to the USPS.

**[Domestic Addresses]{.underline}**

Standard:

- Agency populates City Name, State Code Text, Postal Code and/or Postal
  Code Extension fields separately

Workaround 1:

- Agency populates Postal Code and/or Postal Code Extension fields
  separately, but cannot populate City Name and State Code Text fields
  separately

Workaround 2:

- Agency cannot populate City Name, State Code Text, or Postal Code
  and/or Postal Code Extension fields separately. Note: Check Detail
  Records with blank Postal Codes are marked suspect, requiring manual
  processing. This could delay delivery.

+---------------------+--------------+-----------------+-----------------+
| **Field**           | **Standard** | **Workaround    | **Workaround    |
|                     |              | 1**             | 2**             |
+=====================+==============+=================+=================+
| Payee Name          | -            | -               | -               |
+---------------------+--------------+-----------------+-----------------+
| Address Line1       | -            | -               | -               |
+---------------------+--------------+-----------------+-----------------+
| Address Line 2      | -            | -               | -               |
+---------------------+--------------+-----------------+-----------------+
| Address Line 3      | -            | -               | -               |
+---------------------+--------------+-----------------+-----------------+
| Address Line 4      | -            | -               | -               |
+---------------------+--------------+-----------------+-----------------+
| City Name           | -            |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| State Name          |              |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| State Code Text     | -            |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| Postal Code &Postal | -            | -               |                 |
| Code Extension      |              |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| Barcode             |              |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| Filler              |              |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| Country Name        |              |                 |                 |
+---------------------+--------------+-----------------+-----------------+
| Consular Code       |              |                 |                 |
+---------------------+--------------+-----------------+-----------------+

**[Foreign Addresses]{.underline}**

Payment is considered foreign when Country Name, Geo Code, or Postal
Code (formatted as ~~bb~~nnn) is populated.

Standard:

- Agency populates City Name, State Name, Postal Code and/or Postal Code
  Extension, Country Name, and Consular Code fields separately

Workaround 1:

- Agency populates Country Name and Consular Code fields separately but
  cannot populate City Name, State Name, Postal Code and/or Postal Code
  Extension fields separately

Workaround 2:

- Agency cannot populate City Name, State Name, Postal Code and/or
  Postal Code Extension, Country Name, and Consular Code fields
  separately. Note: payments that are considered domestic will be marked
  suspect when Postal Code field is blank, requiring manual processing.
  This could delay delivery.

Workaround 3:

- Postal Code is assumed to be a Consular Code when ~~bb~~nnn format is
  used

Note: Check Detail Records with Country Name or Consular Code fields
populated are mailed according to foreign mailing rules.

Consular Code can be populated in Consular Code or Postal Code fields.

+-----------------+--------------+--------------+--------------+---------------+
| **Field**       | **Standard** | **Workaround | **Workaround | **Workaround  |
|                 |              | 1**          | 2**          | 3**           |
+=================+==============+==============+==============+===============+
| Payee Name      | -            | -            | -            | -             |
+-----------------+--------------+--------------+--------------+---------------+
| Address Line1   | -            | -            | -            | -             |
+-----------------+--------------+--------------+--------------+---------------+
| Address Line 2  | -            | -            | -            | -             |
+-----------------+--------------+--------------+--------------+---------------+
| Address Line 3  | -            | -            | -            | -             |
+-----------------+--------------+--------------+--------------+---------------+
| Address Line 4  |              | -            | -            | -             |
+-----------------+--------------+--------------+--------------+---------------+
| City Name       | -            |              |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| State Name      | -            |              |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| State Code Text |              |              |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| Postal Code &   | -            |              |              | - (~~bb~~nnn) |
| Postal Code     |              |              |              |               |
| Extension       |              |              |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| Barcode         |              |              |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| Filler          |              |              |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| Country Name    | -            | -            |              |               |
+-----------------+--------------+--------------+--------------+---------------+
| Consular Code   | -            | -            |              |               |
+-----------------+--------------+--------------+--------------+---------------+

## Appendix D - Glossary of Terms

  ----------------------------------------------------------------------------------------------
  Field Name                            Definition
  ------------------------------------- --------------------------------------------------------
  ACH_TransactionCode                   Designates a transaction's account type (checking,
                                        savings, or general ledger) and the transaction type
                                        (credit, or prenote).

  AccountNumber                         The payment recipient's bank account number.

  AccountClassificationAmount           The transaction amount associated with TAS/BETC to be
                                        reported to Treasury for the Central Accounting
                                        Reporting System (CARS) compliancy.

  Addenda Information                   Additional information about the payment to be included
                                        in ACH addenda (PPD+ and CCD+)

  AgencyAccountIdentifier               Account number used internally by an agency to identify
                                        a payee.

  AgencyIdentifier                      Used in conjunction with the main account code,
                                        represents the department, agency, or establishment of
                                        the U.S. government that is responsible for the TAS.

  AgencyLocationCode                    Identifies the accounting office within an agency that
                                        reports disbursements and collections to Treasury.

  AgencyPaymentTypeCode                 Internal agency code used to identify the type of
                                        payment within the specific agency.

  AgencyACHText                         Used to identify the agency responsible for the
                                        payments. First four bytes go to FedACH.

  AllocationTransferAgencyIdentifier    Represents the agency receiving funds through an
                                        allocation transfer.

  Amount                                The amount of the transaction.

  Amount eligible for offset            Indicates the amount of the payment that is eligible for
                                        offset. This is sent to TOP for offsetting purposes.

  AvailabilityTypeCode                  Identifies no-year accounts (X), clearing/suspense
                                        accounts (F), and Treasury central summary general
                                        ledger accounts (A)

  BeginningPeriodOfAvailability         In annual and multiyear accounts, identifies the first
                                        year of availability under law that an account may incur
                                        new obligations.

  BusinessEventTypeCode                 Identifies the type of activity (gross disbursement,
                                        offsetting collection, investment in Treasury
                                        securities, etc.) and the effect of a transaction on the
                                        Fund Balance With Treasury (FBWT). Is used in
                                        combination with the Treasury Account Symbol to classify
                                        transactions reported to Treasury through all
                                        Governmentwide Accounting (CARS-compliant) Financial
                                        Management Systems.

  CheckPaymentEnclosureCode             Code denoting that a check will include an enclosure --
                                        either a stub or a letter.

  CheckLegendText1                      Free-form field for agency's use -- all data in this
                                        field is printed on the check.

  CheckLegendText2                      Free-form field for agency's use -- all data in this
                                        field is printed on the check.

  CityName                              Name of the city in which the payee resides.

  ConsularCode                          Indicates the code for mailing bulk check shipments to
                                        foreign countries.

  CountryCodeText                       The ISO country code for the country in which the
                                        payment recipient resides. Used for IAT payments

  CountryName                           The full name of the country in which the payment
                                        recipient resides.

  EndingPeriodOfAvailability            Identifies the last year of availability under law that
                                        an account (annual and multiyear) may incur new
                                        obligations.

  FederalEmployerIdentificationNumber   A number assigned to businesses by the IRS and used by
                                        DHHS for tracking child support payments.

                                        

                                        

  InputSystem                           Identifies the sending trading partner.

  IsCredit                              Indicates if the item is a debit or credit, for example,
                                        a value of \'1\' would mean that it is a \'credit\'. If
                                        the attribute is not populated, then the default value
                                        is understood to be \'0\'.

  IsSpecialHandling                     Requires agency / Treasury collaboration prior to use.
                                        Used for manual handling of checks such as the
                                        designated agent.

  IsTOP_Offset                          Indicates the payment eligibility of the payment offset

  MainAccountCode                       Identifies the type and purpose of the fund.

  PayeeName                             Contains the full name of a payee. The name of the party
                                        whether individual or organization.

  PayeeName_Secondary                   Contains the full name of an additional payee. The name
                                        of the secondary party whether individual or
                                        organization.

  PayeeAddressLine_1-4                  The mailing address of the payee. Address lines left
                                        blank will not be printed.

  PayeeIdentifier                       The payee\'s Taxpayer Identification Number (TIN),
                                        Vendor ID, Social Security Number (SSN) Employer Tax
                                        Identification Number (EIN), Individual Taxpayer
                                        Identification Number (ITIN), or other valid Payee ID.

  PayeeIdentifier_Secondary             The Secondary party's Taxpayer Identification Number
                                        (TIN), Vendor ID, Social Security Number (SSN), Employer
                                        Tax Identification Number (EIN), Individual Taxpayer
                                        Identification Number (ITIN), or other valid ID.

  PayeeStreetAddress                    Address of the payee for mailing TOP offset letter, if
                                        needed, for ACH payments. For IAT payments, this field
                                        is required.

  PayerMechanism                        Identifies the payment medium for the transaction.

  PaymentDescriptionCode                Identifier to allow TOP to determine the percentage of
                                        offset to be taken during delinquent debt screening.
                                        Examples include, "IN" (Initial), "RE" (Regular), "FI"
                                        (Final", "BA" (Bankruptcy).

  PaymentIdentifier                     Used as a unique identifier for each payment to relate
                                        the records appropriately. Also used for matched
                                        letters.

  PaymentIdentificationLine_1-14        Content to be printed on check stubs.

  Payment Recipient TIN indicator       Indicates whether the TIN (Taxpayer Identification
                                        Number) is an SSN (Social Security Number), Employer
                                        Identification Number (EIN), or Individual Taxpayer
                                        Identification Number (ITIN).

  PaymentTypeCode                       The type of payment as defined in [Appendix
                                        E](#appendix-e---paymenttypecode-values)

  PostalCode                            A code of letters and/or digits added to an address
                                        line. (Zip code)

  PostalCodeExtension                   4-digit extension of postal code.

  PostNetBarcodeDeliveryPoint           Part of the PostNet barcode that is applied to the check
                                        for obtaining postage discounts.

  PostNetBarcodeCheckDigit              Part of the PostNet barcode that is applied to the check
                                        for obtaining postage discounts.

                                        

                                        

  Reconcilement                         Free-form field for agency's use in reconciling the
                                        payment using PACER. Data in this field is passed to
                                        PACER when formatted correctly. Contact Treasury for
                                        additional information.

  Record Code                           A set record identifier used to specify the type of
                                        record in the file.

  RoutingNumber                         The routing number is used synonymously as ABA routing
                                        number and routing transit number. The routing number
                                        consists of 9 digits, for example XXXXYYYYC where XXXX
                                        is Federal Reserve Routing Symbol, YYYY is ABA
                                        Institution Identifier, and C is the Check Digit. The
                                        first two digits of the nine digit routing number must
                                        be in the ranges 00 through 12, 21 through 32, 61
                                        through 72, or 80. The digits are assigned as follows:
                                        00 is used by the United States Government; 01 through
                                        12 are the normal routing numbers; 21 through 32 were
                                        assigned only to thrift institutions (e.g. credit unions
                                        and savings banks) through 1985, currently are still
                                        used by the thrift institutions, or their successors; 61
                                        through 72 are used for electronic transactions; 80 is
                                        used for traveler\'s checks. The first two digits
                                        correspond to the 12 Federal Reserve Banks.

  ScheduleAmount                        The total dollar amount associated with the schedule.

  ScheduleCount                         Total number of payments for each schedule.

  ScheduleNumber                        A number assigned by the agency to identify a group of
                                        payments.

  Standard Payment Request Version      The version of the PAM Standard Payment Request that
  Number                                this file is programmed against.

  StandardEntryClassCode                A three digit code defined by NACHA to characterize the
                                        nature of the ACH transaction. PPD (Prearranged Payment
                                        and Deposit Entry) is an entry initiated by an
                                        organization pursuant to a standing or a single entry
                                        authorization from a receiver to effect a transfer of
                                        funds to impact a consumer account of the receiver. CCD
                                        (Cash Concentration or Disbursement entry) is an entry
                                        initiated by an organization to affect a transfer of
                                        funds to impact the account of that organization or
                                        another organization. IAT (International ACH
                                        Transaction) is a payment that involves a financial
                                        agency's office that is not located in the territorial
                                        jurisdiction of the United States. CTX (Corporate Trade
                                        Exchange) is a corporate format which allows for up to
                                        9,999 addenda records allowing full and complete
                                        remittance information to be transmitted with a single
                                        payment.

  StateCodeText                         The US state code as managed by the USPS (United States
                                        Postal Service) .

  StateName                             The foreign state, province or territory name

  StubLine_1 (through 14)               Agency data to be printed on the stub. Agencies
                                        interested in using this service should contact their
                                        local RFC.

  Sub-accountCode                       Identifies an available receipt or other
                                        Treasury-defined subdivision of the main account.

  Sub-levelPrefixCode                   When populated, represents a programmatic breakdown of
                                        the account for Treasury publication purposes.

  TotalAmount_Payments                  Total amount of the payment request file. Used to verify
                                        the integrity of the file.

  TotalCount_Payments                   Total number of payments in the file. Used to verify the
                                        integrity of the file.

  TotalCount_Records                    Total number of records in the file. Used to verify the
                                        integrity of the file.

  USPSIntelligentMailBarcode            Field used for all intelligent barcode data elements. An
                                        agency must work with their servicing RFC and the US
                                        Postal Service to arrange for this service.
  ----------------------------------------------------------------------------------------------

##  

## Appendix E - PaymentTypeCode Values 

Consult Treasury to determine **PaymentTypeCode** and TOP offset
eligibility.

Of note, the following values represent the current list configured in
PAM and are subject to change.

  -----------------------------------------------------------------------
  **PaymentTypeCode**          **Notes**
  ---------------------------- ------------------------------------------
  Allotment                    

  Annuity                      

  ChildSupport                 

  Daily Benefit                

  Education                    Reserved for use by VA

  Fee                          

  Insurance                    Reserved for use by VA

  Miscellaneous                

  Monthly Benefit              

  Refund                       

  Salary                       

  Thrift                       

  Travel                       

  Vendor                       
  -----------------------------------------------------------------------

# Document History Continued

+-------------+------------+---------------------------------------+------------+
| **Version   | **Author** | **Summary**                           | **Date of  |
| Number**    |            |                                       | Version**  |
+:===========:+============+=======================================+:==========:+
| 7.2.1       | Ashley     | DOC 165 for CR 1399:                  | 2015-7-20  |
|             | Shirk      |                                       |            |
|             |            | Section 3.2.3:                        |            |
|             |            |                                       |            |
|             |            | Updated reconcilement field layout    |            |
|             |            | for VA and VACP check payments by     |            |
|             |            | adding the VA Name Code field. VA is  |            |
|             |            | looking for this field on their       |            |
|             |            | cancellation file layouts from PACER. |            |
+-------------+------------+---------------------------------------+------------+
| 7.2.0       | Ashley     | DOC 148 for CR 1365:                  | 2015-5-28  |
|             | Shirk      |                                       |            |
|             |            | Appendix B Agency Specific Values:    | 2015-06-01 |
|             | Cecelia    |                                       |            |
|             | Walsh      | Added "SSA-Daily" value to account    | 2015-06-08 |
|             |            | for new custom agency rule id created |            |
|             | Susan      |                                       |            |
|             | Santcoeur  | ACH Schedule Header Record and Check  |            |
|             |            | Schedule Header Record**:**           |            |
|             |            |                                       |            |
|             |            | Updated validation rules for schedule |            |
|             |            | number when same combination of       |            |
|             |            | schedule number + ALC + fiscal year)  |            |
|             |            | is submitted for a removed scheduled. |            |
|             |            |                                       |            |
|             |            | PAMBRG-275 Doc-only changes           |            |
|             |            | throughout to sync with Treasury's    |            |
|             |            | requested changes for website.        |            |
+-------------+------------+---------------------------------------+------------+
| 7.1.0       | Ashley     | CR 1054:                              | 2014-12-31 |
|             | Shirk      |                                       |            |
|             |            | Special Handling field in the Check   | 2014-10-20 |
|             | Cecelia    | Payment Data Record:                  |            |
|             | Walsh      |                                       | 2014-10-1  |
|             |            | Added rule to store 'name only' value |            |
|             | Ashley     | when the schedule is marked name only |            |
|             | Shirk      | and the special handling field is     |            |
|             |            | blank. This will ensure the payments  |            |
|             |            | show up on the special handling tab   |            |
|             |            | in UC 3032 for users.                 |            |
|             |            |                                       |            |
|             |            | PAMBRG-150                            |            |
|             |            |                                       |            |
|             |            | Document update only to update        |            |
|             |            | downstream mapping for DNP, PPS, PIR, |            |
|             |            | PACER, and IPP.                       |            |
|             |            |                                       |            |
|             |            | CR 963 10.20.14 CW: Update to IAT     |            |
|             |            | Country Code validation rules to not  |            |
|             |            | allow zeros avoid FedACH rejects.     |            |
|             |            |                                       |            |
|             |            | CR 848:                               |            |
|             |            |                                       |            |
|             |            | Appendix A:                           |            |
|             |            |                                       |            |
|             |            | Added 24, 34 values to the list of    |            |
|             |            | allowable ACH transaction codes. Per  |            |
|             |            | NACHA guidelines, these codes may be  |            |
|             |            | used for \$0 CTX payments.            |            |
+-------------+------------+---------------------------------------+------------+
| 7.0.0       | Ashley     | DOC 0579 for Release 7.0 Changes:     | 2013-12-16 |
|             | Shirk      |                                       |            |
|             |            | CR 4306:                              |            |
|             |            |                                       |            |
|             |            | ACH Payment Data Record:              |            |
|             |            |                                       |            |
|             |            | Added validation to Depositor Acct \# |            |
|             |            | field. Payment will be marked invalid |            |
|             |            | if value contains all zeroes.         |            |
|             |            |                                       |            |
|             |            | CR 4571:                              |            |
|             |            |                                       |            |
|             |            | Check Detail Record:                  |            |
|             |            |                                       |            |
|             |            | Removed validation from the Payment   |            |
|             |            | Amount field. PAM will begin          |            |
|             |            | accepting 10 digit check amounts in   |            |
|             |            | release 7.                            |            |
|             |            |                                       |            |
|             |            | New DNP Record added for agencies to  |            |
|             |            | use to send Agency specific           |            |
|             |            | information to DNP. Contents in the   |            |
|             |            | DNP Detail field will be sent to DNP  |            |
|             |            | in the OFS-DNP Screening and the SRF. |            |
|             |            | Record is optional and sent per       |            |
|             |            | payment.                              |            |
|             |            |                                       |            |
|             |            | CR 4563:                              |            |
|             |            |                                       |            |
|             |            | Added new generic reconcilement field |            |
|             |            | layout to be used by agencies that    |            |
|             |            | need to pass a Legacy Treasury Acct   |            |
|             |            | Symbol to TCIS and receive value on   |            |
|             |            | TRACS file.                           |            |
+-------------+------------+---------------------------------------+------------+
| 5.0.5       | Ashley     | DOC 0609 for CR 4485:                 | 2014-1-30  |
|             | Shirk      |                                       |            |
|             |            | **ACH & Check Payment Data Record:**  |            |
|             |            |                                       |            |
|             |            | AgencyAccountIdentifier:              |            |
|             |            |                                       |            |
|             |            | - added rule to store value as        |            |
|             |            |   received for Custom Agency Rule ID  |            |
|             |            |   'VACP'. Change was made due to      |            |
|             |            |   request from VA that leading spaces |            |
|             |            |   must be maintained in this field    |            |
|             |            |   until they're able to implement     |            |
|             |            |   changes in their system             |            |
|             |            |                                       |            |
|             |            | Reconcilement:                        |            |
|             |            |                                       |            |
|             |            | - Added validation note that value    |            |
|             |            |   should be stored as received. Per   |            |
|             |            |   SDG, all reconcilement field data   |            |
|             |            |   does not get justified before or    |            |
|             |            |   after parsing.                      |            |
|             |            |                                       |            |
|             |            | - Added 'VACP' custom agency rule ID  |            |
|             |            |   to existing 'VA' reconcilement      |            |
|             |            |   field layout. Per feedback from VA  |            |
|             |            |   meeting, VACP payments will use the |            |
|             |            |   same reconcilement field layout.    |            |
+-------------+------------+---------------------------------------+------------+
| 5.0.4       | Toyiah     | DOC 0558 for CR 4409:                 | 2013-10-01 |
|             | Cavole     |                                       |            |
|             |            | 1.5 Validation for Balancing          |            |
|             |            |                                       |            |
|             |            | Added validation rule to accept zero  |            |
|             |            | dollar payment amount, if the SEC     |            |
|             |            | code = CTX as stated in the           |            |
|             |            | Validation for Balancing section.     |            |
|             |            |                                       |            |
|             |            | Updated Section 2.4 ACH Payment Data  |            |
|             |            | Record to include the SEC code as an  |            |
|             |            | additional validation point.          |            |
+-------------+------------+---------------------------------------+------------+
| 5.0.3       | Ashley     | DOC 0503: updated Section 1.4         | 2013-04-05 |
|             | Shirk      | Validation for Balancing to clarify   |            |
|             |            | that the                              |            |
|             |            | FileTrailer.TotalCount_Records should |            |
|             |            | include header and trailer records    |            |
|             |            | (doc change only)                     |            |
|             |            |                                       |            |
|             |            | DOC 0503 for CR 4257: Made the        |            |
|             |            | following changes to the Check        |            |
|             |            | Payment Data Record based on feedback |            |
|             |            | from mail operations staff and agency |            |
|             |            | outreach group:                       |            |
|             |            |                                       |            |
|             |            | Payee AddressLine_4: removed note     |            |
|             |            | that stated agencies shouldn't use    |            |
|             |            | address line for foreign addresses.   |            |
|             |            | If an agency is unable to provide     |            |
|             |            | breakout fields, they can use all     |            |
|             |            | four lines of address for foreign     |            |
|             |            | payments.                             |            |
|             |            |                                       |            |
|             |            | CityName & StateCodeText: removed     |            |
|             |            | validation that marked suspect when   |            |
|             |            | blank.                                |            |
|             |            |                                       |            |
|             |            | PostalCode:                           |            |
|             |            |                                       |            |
|             |            | - Added note that field can be used   |            |
|             |            |   for Geo Code for foreign payments.  |            |
|             |            |                                       |            |
|             |            | Updated validation to mark suspect    |            |
|             |            | when blank for domestic payments and  |            |
|             |            | not foreign                           |            |
+-------------+------------+---------------------------------------+------------+
| 5.0.2       | Carole     | DOC 0481                              | 2013-03-25 |
|             | Kampe      |                                       |            |
|             |            | Document change only: Appendix B --   | 2012-12-14 |
|             | Mike       | added derivation table information to |            |
|             | Lancaster  | the description of the section, to    |            |
|             |            | clarify what is found in the          |            |
|             |            | appendix.                             |            |
|             |            |                                       |            |
|             |            | for CR 4132: In Appendix B -- Agency  |            |
|             |            | Specific Values / Custom Agency Rule  |            |
|             |            | ID = "IRS" when the Depositor Account |            |
|             |            | Number is not equal to "BONDS", added |            |
|             |            | the derivation of Legacy Account      |            |
|             |            | Symbol in the IRS Reconcilement       |            |
|             |            | Record.                               |            |
+-------------+------------+---------------------------------------+------------+
| 5.0.1       | Edward     | DOC 0464 for CR 3890                  | 2012-09-20 |
|             | Perez      |                                       |            |
|             |            | "SSA Daily and Monthly Benefit        |            |
|             |            | Reconcilement Field": added           |            |
|             |            | derivation rule for new Legacy        |            |
|             |            | Account Symbol attribute value.       |            |
+-------------+------------+---------------------------------------+------------+
| 5.0.0       | Brandon    | DOC 0449: In the IRS Reconcilement    | 08-09-2012 |
|             | Mitchell   | field, CheckDetailEnclosureCode       |            |
|             |            | field, changed the value "INSERT" to  | 08-01-2012 |
|             | Ashley     | "ENCLOSE".                            |            |
|             | Shirk      |                                       | 06-11-2012 |
|             |            | DOC 0449:                             |            |
|             | Brandon    |                                       |            |
|             | Mitchell   | **General Structure of the File       |            |
|             |            | Section:**                            |            |
|             |            |                                       |            |
|             |            | Updated to accept multiple addendum   |            |
|             |            | records for CTX payments              |            |
|             |            |                                       |            |
|             |            | Changed GWA TAS/BETC Record limit     |            |
|             |            | from one-100 to one-many because the  |            |
|             |            | limit is a soft limit.                |            |
|             |            |                                       |            |
|             |            | **File Structure Validations:**       |            |
|             |            |                                       |            |
|             |            | Reformatted bullets and put into a    |            |
|             |            | table.                                |            |
|             |            |                                       |            |
|             |            | **Input Management Info:**            |            |
|             |            |                                       |            |
|             |            | Clarified values in the filename. RFC |            |
|             |            | will always be "KFC" and type will    |            |
|             |            | always be "SPR".                      |            |
|             |            |                                       |            |
|             |            | **Hexidecimal Character Validation:** |            |
|             |            |                                       |            |
|             |            | Reformatted bullets and put into a    |            |
|             |            | table.                                |            |
|             |            |                                       |            |
|             |            | Moved section into the File Structure |            |
|             |            | Validation Section                    |            |
|             |            |                                       |            |
|             |            | **Validation for Balancing Section:** |            |
|             |            |                                       |            |
|             |            | Reformatted bullets and put into a    |            |
|             |            | table                                 |            |
|             |            |                                       |            |
|             |            | **Specification Notes:**              |            |
|             |            |                                       |            |
|             |            | Added bullet stating currency values  |            |
|             |            | are not formatted as zoned decimal    |            |
|             |            | based on RAD session feedback.        |            |
|             |            |                                       |            |
|             |            | Added bullet to specify that numeric  |            |
|             |            | fields will be stored as zeros if     |            |
|             |            | blanks are received.                  |            |
|             |            |                                       |            |
|             |            | **StandardPaymentRequestVersion       |            |
|             |            | field**: updated to look for 500 for  |            |
|             |            | the current published version         |            |
|             |            |                                       |            |
|             |            | **IsTOP_Offset field:** updated note  |            |
|             |            | for clarity based on RAD session      |            |
|             |            | feedback.                             |            |
|             |            |                                       |            |
|             |            | **Standard Entry Class Code field:**  |            |
|             |            |                                       |            |
|             |            | Removed IDD as a value, decision was  |            |
|             |            | made that PAM will not process IDD    |            |
|             |            | payments.                             |            |
|             |            |                                       |            |
|             |            | Added CTX as a value in order to      |            |
|             |            | process CTX payments with Release 5   |            |
|             |            |                                       |            |
|             |            | **CountryCodeText**: removed IDD as   |            |
|             |            | value to look for in the notes and    |            |
|             |            | validation rules                      |            |
|             |            |                                       |            |
|             |            | **PartyName**: updated note to        |            |
|             |            | include truncation rules for CTX and  |            |
|             |            | IAT.                                  |            |
|             |            |                                       |            |
|             |            | **ACH Addendum Record:**              |            |
|             |            |                                       |            |
|             |            | Updated notes under the               |            |
|             |            | heading-reformatted for clarity       |            |
|             |            |                                       |            |
|             |            | Created new record format for CTX     |            |
|             |            | payments and added validation rules   |            |
|             |            | for CTX payments. Up to 999 Addendum  |            |
|             |            | records will be accepted per payment  |            |
|             |            | with CTX.                             |            |
|             |            |                                       |            |
|             |            | **GWA TAS/BETC Record**               |            |
|             |            |                                       |            |
|             |            | Updated wording under heading to note |            |
|             |            | that the limit is a recommended       |            |
|             |            | limit.                                |            |
|             |            |                                       |            |
|             |            | **Procurement Record:**               |            |
|             |            |                                       |            |
|             |            | Updated wording under heading to note |            |
|             |            | that the limit is a recommended       |            |
|             |            | limit.                                |            |
|             |            |                                       |            |
|             |            | Added Amount field.                   |            |
|             |            |                                       |            |
|             |            | Replaced all "GWA" text with "CARS".  |            |
|             |            |                                       |            |
|             |            | DOC 0449: Updates from Release 5.0    |            |
|             |            | RAD Session:                          |            |
|             |            |                                       |            |
|             |            | Removed Error Code for the Check      |            |
|             |            | Detail Enclosure Code field in the    |            |
|             |            | IRS Reconcilement.                    |            |
|             |            |                                       |            |
|             |            | Bond Name 1 and Bond Name 2 fields    |            |
|             |            | are changed to alpha-numeric in the   |            |
|             |            | IRS Savings Bonds Orders              |            |
|             |            | Reconcilement.                        |            |
|             |            |                                       |            |
|             |            | Bond REG Code field will have no      |            |
|             |            | validation in the IRS Savings Bonds   |            |
|             |            | Orders Reconcilement.                 |            |
+-------------+------------+---------------------------------------+------------+
| 4.2.1       | Brandon    | DOC 0393: Corrected the IRS           | 05-17-2012 |
|             | Mitchell   | Reconcilement field order             |            |
| 4.2.1       |            |                                       | 05-09-2012 |
|             | Brandon    | DOC 0393: Added a table in Appendix A |            |
| 4.2.1       | Mitchell   | to represent translation of the ACH   | 05-02-2012 |
|             |            | Transaction Code to the Savings Bonds |            |
| 4.2.1       | Brandon    | Owner Type.                           | 4-30-2012  |
|             | Mitchell   |                                       |            |
| 4.2.1       |            | Grammar and punctuation corrections.  | 3-7-2012   |
|             | Brandon    |                                       |            |
|             | Mitchell   | DOC 0393:                             |            |
|             |            |                                       |            |
|             | Ashley     | Removed the Garnishment Indicator     |            |
|             | Shirk      | from the ACH Schedule Header Record   |            |
|             |            | (Field #11) and in Appendix B         |            |
|             |            | Glossary of Terms.                    |            |
|             |            |                                       |            |
|             |            | In Appendix C -- Agency Specific      |            |
|             |            | Values, added the SSA Allotments      |            |
|             |            | Reconcilement File table for Custom   |            |
|             |            | Agency Rule "SSA-A".                  |            |
|             |            |                                       |            |
|             |            | Savings Bonds Orders Updates          |            |
|             |            |                                       |            |
|             |            | Added SnAP in the Downstream System   |            |
|             |            | Column for the following fields:      |            |
|             |            |                                       |            |
|             |            | ACH Header Record                     |            |
|             |            |                                       |            |
|             |            | Agency Location Code                  |            |
|             |            |                                       |            |
|             |            | Schedule Number                       |            |
|             |            |                                       |            |
|             |            | In the ACH Payments Data Record, the  |            |
|             |            | Depositor Account Number will equal   |            |
|             |            | "BONDS". This will be the identified  |            |
|             |            | to indicate the Payment is a Savings  |            |
|             |            | Bonds Order.                          |            |
|             |            |                                       |            |
|             |            | ACH Payment Data Record (added SnAP)  |            |
|             |            |                                       |            |
|             |            | PayeeIdentifer                        |            |
|             |            |                                       |            |
|             |            | AddressLine_1                         |            |
|             |            |                                       |            |
|             |            | AddressLine_2                         |            |
|             |            |                                       |            |
|             |            | Added a table to Appendix B with the  |            |
|             |            | breakdown of the Reconcilement field  |            |
|             |            | rules if Custom Agency Rule ID =      |            |
|             |            | "IRS" and Depositor Account Number =  |            |
|             |            | "BONDS".                              |            |
|             |            |                                       |            |
|             |            | The following are new fields that     |            |
|             |            | will need to be added to PAM:         |            |
|             |            |                                       |            |
|             |            | Bond Name 1                           |            |
|             |            |                                       |            |
|             |            | Bond REG Code                         |            |
|             |            |                                       |            |
|             |            | Bond Name 2                           |            |
|             |            |                                       |            |
|             |            | DOC 0393:                             |            |
|             |            |                                       |            |
|             |            | Updated validation note in            |            |
|             |            | AgencyAccountIdentifier in the Check  |            |
|             |            | Record to match validation note found |            |
|             |            | in AgencyAccountIdentifier in ACH     |            |
|             |            | Record.                               |            |
|             |            |                                       |            |
|             |            | Changes from RAD session on 3/7/12:   |            |
|             |            |                                       |            |
|             |            | In the Check and ACH Records, changed |            |
|             |            | type for                              |            |
|             |            | PaymentRecipientTINIndicator,         |            |
|             |            | SecondaryPayeeTINIndicator, and       |            |
|             |            | AmountEligibleforOffset to AN because |            |
|             |            | blanks are allowed.                   |            |
|             |            |                                       |            |
|             |            | Appendix C - Agency Specific Values:  |            |
|             |            | Updated note, use the Custom Agency   |            |
|             |            |                                       |            |
|             |            | Rule ID to determine when to fill     |            |
|             |            | fields.                               |            |
|             |            |                                       |            |
|             |            | Added separate table in the           |            |
|             |            | Reconcilement field for VA ACH        |            |
|             |            | payments. VA ACH will not contain the |            |
|             |            | Courtesy Code field.                  |            |
|             |            |                                       |            |
|             |            | Map the TOP Eligibility Indicator to  |            |
|             |            | the Debt Bypass Indicator Code for    |            |
|             |            | IRS                                   |            |
|             |            |                                       |            |
|             |            | Payments. Value is the same as the    |            |
|             |            | IRS Bypass Code, so use one field.    |            |
|             |            |                                       |            |
|             |            | Object Code in RRB Reconcilement      |            |
|             |            | field: updated stored name, should be |            |
|             |            | Object Code instead of RRB Object     |            |
|             |            | Code.                                 |            |
|             |            |                                       |            |
|             |            | Clarified note for the TIN Indicator  |            |
|             |            | Offset and Debt Bypass Indicator Code |            |
|             |            |                                       |            |
|             |            | fields to explain that the value in   |            |
|             |            | the reconcilement field should be     |            |
|             |            | used                                  |            |
|             |            |                                       |            |
|             |            | when received instead of the value in |            |
|             |            | the data record.                      |            |
|             |            |                                       |            |
|             |            | Added note to ACH Addendum Record,    |            |
|             |            | the record is required for IDD        |            |
|             |            | payments.                             |            |
|             |            |                                       |            |
|             |            | Added error codes and validation      |            |
|             |            | rules to the following fields:        |            |
|             |            |                                       |            |
|             |            | Amount Eligible for Offset, Secondary |            |
|             |            | Payee TIN Indicator, Payment          |            |
|             |            | Recipient TIN Indicator, Payee        |            |
|             |            | Identifier, PayeeIdentifier_Secondary |            |
|             |            |                                       |            |
|             |            | Updated notes in the General          |            |
|             |            | Structure of File section for         |            |
|             |            | clarification:                        |            |
|             |            |                                       |            |
|             |            | Records within the Check and ACH      |            |
|             |            | payment data records can be received  |            |
|             |            | in any order.                         |            |
|             |            |                                       |            |
|             |            | The addenda record is required for    |            |
|             |            | IDD payments.                         |            |
+-------------+------------+---------------------------------------+------------+
| 4.2.1       | Ashley     | DOC 0351: New release due to changes  | 3-1-2012   |
|             | Shirk      | made in version 3.1.9. Approval for   |            |
|             |            | 3.1.9 carries to 4.2.1.               |            |
+-------------+------------+---------------------------------------+------------+
| 4.2.0       | Dorothy    | DOC 0160:                             | 12-1-2011  |
|             | Carpenter  |                                       |            |
|             |            | Changes for release 4.2 including:    |            |
|             |            |                                       |            |
|             |            | \- Changed field names                |            |
|             |            | (postalcodeextension,                 |            |
|             |            | checkpaymentenclosurecode,            |            |
|             |            | isspecialhandling,                    |            |
|             |            | accountclassificationamount) to be    |            |
|             |            | consistent with the FMS Schema        |            |
|             |            |                                       |            |
|             |            | \- Changed Schedule \# valid          |            |
|             |            | characters to be consistent with SPS  |            |
|             |            | for matching purposes.                |            |
|             |            |                                       |            |
|             |            | \- Combined Account Number and Suffix |            |
|             |            | for Account number in a single 16     |            |
|             |            | character field and changed the       |            |
|             |            | schema name to                        |            |
|             |            | AgencyAccountIdentifier               |            |
|             |            |                                       |            |
|             |            | \- Added TIN (schema name =           |            |
|             |            | payeeidentifier) to both check and    |            |
|             |            | ACH data record to accommodate        |            |
|             |            | downstream system needs for a SSN/TIN |            |
|             |            | in addition to the current legacy     |            |
|             |            | account number.                       |            |
|             |            |                                       |            |
|             |            | \- Added agency specific mapping for  |            |
|             |            | the reconcilement field (Appendix C)  |            |
|             |            | to accommodate larger agencies that   |            |
|             |            | will convert to the SPR in PAM 4.2    |            |
|             |            | and beyond.                           |            |
|             |            |                                       |            |
|             |            | \- Removed all USPS Intelligent       |            |
|             |            | barcode fields and added one field    |            |
|             |            | that will contain all data elements.  |            |
|             |            | Irwin Richmond suggested this         |            |
|             |            | approach due to the variety of data   |            |
|             |            | elements that can be used for the     |            |
|             |            | different intelligent barcode         |            |
|             |            | services.                             |            |
|             |            |                                       |            |
|             |            | \- Added a note to Object Lines 1&2   |            |
|             |            | that the allotter name and claim      |            |
|             |            | number should be mapped to the first  |            |
|             |            | 35 characters and will be passed to   |            |
|             |            | PACER.                                |            |
|             |            |                                       |            |
|             |            | \- Added an indicator for primary and |            |
|             |            | secondary payee identifier to         |            |
|             |            | differentiate between SSN and EIN     |            |
|             |            |                                       |            |
|             |            | \- Added "amount eligible for offset" |            |
|             |            | field to accommodate situations where |            |
|             |            | a subset of the payment amount will   |            |
|             |            | be sent to TOP (RRB and disposable    |            |
|             |            | pay for salary)                       |            |
|             |            |                                       |            |
|             |            | \- Changed the maximum TAS/BETC per   |            |
|             |            | schedule from 300 to 1000.            |            |
|             |            |                                       |            |
|             |            | \- Changed Stored Name of Payee       |            |
|             |            | Identifier Secondary Account Number   |            |
|             |            | to "Secondary TIN"                    |            |
|             |            |                                       |            |
|             |            | \- Changed stored name of Party Name  |            |
|             |            | Secondary to "Secondary Payee"        |            |
|             |            |                                       |            |
|             |            | \- Added a new value to enclosure     |            |
|             |            | code called "insert" to accommodate   |            |
|             |            | generic inserts for selected agencies |            |
|             |            | (IRS error notice, for example)       |            |
|             |            |                                       |            |
|             |            | \- Removed Contracting Office         |            |
|             |            | Procurement Agency ID from the        |            |
|             |            | Procurement Record per Rosa Chan and  |            |
|             |            | Marcel Jemio.                         |            |
|             |            |                                       |            |
|             |            | \- Removed Office Code from the VA    |            |
|             |            | Reconcilement data per VA insurance   |            |
|             |            | meeting on 9/29/2011                  |            |
|             |            |                                       |            |
|             |            | \- Separated Split Refund Code and    |            |
|             |            | Injured Spouse indicator into two     |            |
|             |            | fields to eliminate overuse of this   |            |
|             |            | field. Will combine in downstream     |            |
|             |            | files until those systems can change. |            |
|             |            |                                       |            |
|             |            | Updated validation rules to note that |            |
|             |            | records within the Check Payment Data |            |
|             |            | and ACH Payment Data Records can be   |            |
|             |            | received in any order based on        |            |
|             |            | discussions held with SSI about the   |            |
|             |            | order of the GWA TAS/BETC record.     |            |
|             |            |                                       |            |
|             |            | Removed Payment Code field from       |            |
|             |            | Derived table. Updated the General    |            |
|             |            | Structure of the File section to      |            |
|             |            | state there can be more than one      |            |
|             |            | procurement record per detail         |            |
|             |            | payment.                              |            |
|             |            |                                       |            |
|             |            | Added a note for the Agency EIN       |            |
|             |            | field; it is passed in Company        |            |
|             |            | Identification field and is the FEIN  |            |
|             |            | of the agency that is the employer of |            |
|             |            | the payee. If this is not available   |            |
|             |            | then provide the FEIN of the agency   |            |
|             |            | sending the payment.                  |            |
|             |            |                                       |            |
|             |            | Add IDD as a Standard Entry Class     |            |
|             |            | Code to distinguish between domestic  |            |
|             |            | and foreign payments. PAM will        |            |
|             |            | convert IDD to PPD on output.         |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.9       | Ashley     | DOC 0351: Mapped the TOP Eligibility  | 03-14-2012 |
|             | Shirk      | Indicator to the TIN Indicator Offset |            |
|             |            | field located in the reconcilement    |            |
|             |            | field.                                |            |
|             |            |                                       |            |
|             |            | Updated file structure validations to |            |
|             |            | note that IDD payments need to be in  |            |
|             |            | Country Code and Routing Number       |            |
|             |            | order.                                |            |
|             |            |                                       |            |
|             |            | Added validation to CountryCodeText   |            |
|             |            | field. If Country Code is blank for   |            |
|             |            | IDD payments, the payment should be   |            |
|             |            | marked invalid.                       |            |
|             |            |                                       |            |
|             |            | Added a separate table in Appendix C  |            |
|             |            | for VA ACH payments. ACH payments     |            |
|             |            | will not contain the VA Courtesy Code |            |
|             |            | per RAD Session on 3/6/12.            |            |
|             |            |                                       |            |
|             |            | CR 3545: Added PSC/ALC table to the   |            |
|             |            | Reconcilement field for Custom Agency |            |
|             |            | Rule ID = "SSA". The table will be    |            |
|             |            | used to derive the Payment ALC per    |            |
|             |            | PSC.                                  |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.8       | Ashley     | DOC 0339: For Release 4.1 Generic     | 01-18-2012 |
| Developer   | Shirk      | Types of Payment implementation,      |            |
| version     |            | added Object Line 3 to the Derived    |            |
| only for    |            | Data Elements section. Value will be  |            |
| changes in  |            | stored on the input rather than the   |            |
| Release     |            | output.                               |            |
| 4.1.        |            |                                       |            |
|             |            | Parsed reconcilement into separate    |            |
|             |            | fields for VA and SSA in order to     |            |
|             |            | accommodate for VA and SSA using      |            |
|             |            | version 3.1.4.                        |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.7       | Carole     | DOC 0304:                             | 12-10-2011 |
|             | Kampe      |                                       |            |
| Developer   |            | Section 1.6: Deleted Payment Code     |            |
| version     |            | from the Derived Data Table. It will  |            |
| only for    |            | be stored from the RAT configuration. |            |
| changes in  |            |                                       |            |
| Release     |            | Section 1.7: Added a fourth node for  |            |
| 4.1.        |            | Unique ID so that each dataset        |            |
|             |            | received on a day will not overlay    |            |
|             |            | onto a previously received file.      |            |
|             |            |                                       |            |
|             |            | Added IDD as a value for the SEC code |            |
|             |            | to distinguish between domestic and   |            |
|             |            | foreign ACH payments. SSA is using    |            |
|             |            | v3.1.4 of the SPR so this update is   |            |
|             |            | required in this layout of the SPR.   |            |
|             |            | Also adding to the 4.1.0 version.     |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.6       | Ashley     | DOC 0291: Discussed on 11-4 AMUQ \#   | 11-8-2011  |
|             | Shirk      | 325 Clarified the validation rules to |            |
| Developer   |            | state that the ACH Payment Data       |            |
| version     |            | records should be ordered by routing  |            |
| only. We    |            | number at the schedule level.         |            |
| will not    |            |                                       |            |
| change the  |            | Updated validation rules to note that |            |
| published   |            | records within the Check Payment Data |            |
| version     |            | and ACH Payment Data Records can be   |            |
| number.     |            | received in any order based on        |            |
|             |            | discussions held with SSI about the   |            |
| Chng for    |            | order of the GWA TAS/BETC record.     |            |
| 3.1.6 will  |            |                                       |            |
| be          |            | This is being implemented in the      |            |
| implemented |            | Release 4.0, December 2011.           |            |
| in Release  |            |                                       |            |
| 4.0         |            |                                       |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.5       | Dorothy    | DOC 0210:                             | 5-17-2011  |
|             | Carpenter  |                                       |            |
|             |            | \- CR2847: Release 3.1 workaround to  |            |
|             |            | store type of payment as Daily        |            |
|             |            | Benefit S for SSA so that we can      |            |
|             |            | configure SSA values in 3.1. In 4.0   |            |
|             |            | the type of payment configuration     |            |
|             |            | values will move to the R-A-T to      |            |
|             |            | allow multiple agencies to use        |            |
|             |            | generic type of payment names.        |            |
|             |            |                                       |            |
|             |            | \- CR2789: document change only --    |            |
|             |            | clarified validation rules to include |            |
|             |            | prenote rules discussed in the        |            |
|             |            | validation for balancing section.     |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.4       | Dorothy    | DOC0116:                              | 12-20-2010 |
|             | Carpenter  |                                       |            |
|             |            | \- Added derived data element Payment |            |
|             |            | code.                                 |            |
|             |            |                                       |            |
|             |            | \- Clarified the correct published    |            |
|             |            | version in the validation rules for   |            |
|             |            | the Standard Payment Request Version  |            |
|             |            | field.                                |            |
|             |            |                                       |            |
|             |            | -Changed Input System stored name     |            |
|             |            | from Agency Identification to         |            |
|             |            | Originating Agency Identifier.        |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.3       | Dorothy    | DOC0031:                              | 11-8-2010  |
|             | Carpenter  |                                       |            |
|             |            | \- Added suspect check validation     |            |
|             |            |                                       |            |
|             |            | \- Changed several validation rules   |            |
|             |            | to ensure that right/left             |            |
|             |            | justification of content is only      |            |
|             |            | performed where necessary (field is   |            |
|             |            | truncated or manipulated on output).  |            |
|             |            |                                       |            |
|             |            | \- Changed InputSystem to match       |            |
|             |            | against "agency identifier" instead   |            |
|             |            | of "agency identification".           |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.2       | Dorothy    | DOC0031:                              | 10-7-2010  |
|             | Carpenter  |                                       |            |
|             |            | \- Changed "count" fields in schedule |            |
|             |            | trailer from 18 characters to 15 for  |            |
|             |            | amounts and 8 for \# of items for     |            |
|             |            | proper mapping downstream.            |            |
|             |            |                                       |            |
|             |            | \- Added note to Input Management     |            |
|             |            | section to clarify that CA scheduler  |            |
|             |            | may replace IM                        |            |
|             |            |                                       |            |
|             |            | \- Added clarification to             |            |
|             |            | Specification Notes for proper        |            |
|             |            | formatting of fields based on type.   |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.1       | Dorothy    | DOC0031:                              | 9-14-2010  |
|             | Carpenter  |                                       |            |
|             |            | \- Addressed questions raised by SSA, |            |
|             |            | SDG and SPS.                          |            |
|             |            |                                       |            |
|             |            | \- Removed requirement that ACH files |            |
|             |            | be sorted by account number -- only   |            |
|             |            | sort is RTN.                          |            |
|             |            |                                       |            |
|             |            | \- Stated in check amount field that  |            |
|             |            | only amounts of 9 digits or less are  |            |
|             |            | allowed. Added validation rule.       |            |
|             |            |                                       |            |
|             |            | \- Changed field names and stored     |            |
|             |            | names for check / ACH address lines   |            |
|             |            | to "payee address line" to be         |            |
|             |            | consistent with schema                |            |
|             |            |                                       |            |
|             |            | \- Changed Payment Identification to  |            |
|             |            | Payment ID to be consistent with      |            |
|             |            | schema and distinguish this field     |            |
|             |            | from payment identification lines.    |            |
|             |            |                                       |            |
|             |            | \- Added derived data elements table  |            |
|             |            | to instruct SDG to store method of    |            |
|             |            | payment based on record code.         |            |
|             |            |                                       |            |
|             |            | \- changed contracting office agency  |            |
|             |            | ID to contracting office procurement  |            |
|             |            | agency ID to be consistent with       |            |
|             |            | schema                                |            |
|             |            |                                       |            |
|             |            | \- changed statecode and countrycode  |            |
|             |            | to statecodetext and countrycodetext  |            |
|             |            | to be consistent with the schema      |            |
|             |            |                                       |            |
|             |            | \- Changed the record count for IAT   |            |
|             |            | remittance addenda to allow for two   |            |
|             |            | addenda records.                      |            |
|             |            |                                       |            |
|             |            | \- Added validation to ACH            |            |
|             |            | transaction code to only allow vendor |            |
|             |            | type of payment for codes 42, 43, 52, |            |
|             |            | and 53. This is due to ARM dated      |            |
|             |            | 4/30/2004.                            |            |
+-------------+------------+---------------------------------------+------------+
| 3.1.0       | Dorothy    | DOC 0031:                             | 8-16-2010  |
|             | Carpenter  |                                       |            |
|             |            | \- Changed version to indicate this   |            |
|             |            | is the version that will be           |            |
|             |            | implemented in 3.1                    |            |
|             |            |                                       |            |
|             |            | \- Added procurement record and five  |            |
|             |            | new procurement/PIR related fields.   |            |
|             |            |                                       |            |
|             |            | \- Removed agency text from check     |            |
|             |            | schedule record because               |            |
|             |            | checklegendtext allows for any agency |            |
|             |            | specific check data.                  |            |
|             |            |                                       |            |
|             |            | \- Changed name of agency text in ACH |            |
|             |            | schedule record to AgencyACHtext      |            |
|             |            |                                       |            |
|             |            | Added validation for SPR version      |            |
|             |            | number.                               |            |
|             |            |                                       |            |
|             |            | \- Changed PaymentTypeCode to         |            |
|             |            | validate against configured payments  |            |
|             |            |                                       |            |
|             |            | \- Removed list of Payment Types to   |            |
|             |            | allow for more flexibility for adding |            |
|             |            | new types.                            |            |
|             |            |                                       |            |
|             |            | -Changed field lengths for city,      |            |
|             |            | state, country in both ACH and check  |            |
|             |            | data records                          |            |
|             |            |                                       |            |
|             |            | \- Removed TOPSiteID and TOPAgencyID  |            |
|             |            | because these are used by only one    |            |
|             |            | agency (CCC).                         |            |
|             |            |                                       |            |
|             |            | \- Clarified validation and error     |            |
|             |            | messages throughout                   |            |
|             |            |                                       |            |
|             |            | \- Added IsCredit field to TAS/BETC   |            |
|             |            | record                                |            |
|             |            |                                       |            |
|             |            | \- Removed verbiage about signed zone |            |
|             |            | decimal because the IsCredit          |            |
|             |            | indicator accomplishes the signing of |            |
|             |            | the TAS/BETC amount.                  |            |
|             |            |                                       |            |
|             |            | \- Removed TerritoryCode from ACH     |            |
|             |            | record because the StateName field is |            |
|             |            | intended to handle both foreign       |            |
|             |            | territories and states.               |            |
|             |            |                                       |            |
|             |            | \- Added StateName to account for     |            |
|             |            | foreign addresses                     |            |
|             |            |                                       |            |
|             |            | \- Changed Country to CountryCode to  |            |
|             |            | account for IAT and to CountryName    |            |
|             |            | for foreign checks.                   |            |
|             |            |                                       |            |
|             |            | \- Removed stub address because we    |            |
|             |            | will obtain this from SPS.            |            |
|             |            |                                       |            |
|             |            | \- Removed verbiage about schedule    |            |
|             |            | and file level recovery. Indicated in |            |
|             |            | validation rules the violations that  |            |
|             |            | would result in a schedule rejection  |            |
|             |            | and noted in introduction that        |            |
|             |            | schedule level rejects would result   |            |
|             |            | in file rejection until further       |            |
|             |            | notice.                               |            |
+-------------+------------+---------------------------------------+------------+
| 3.6         | Dorothy    | CR1855:                               | 4-16-2010  |
|             | Carpenter  |                                       |            |
|             |            | \- Added file header and trailer to   |            |
|             |            | introduction                          |            |
|             |            |                                       |            |
|             |            | \- Reworded description of recovery   |            |
|             |            | level based on clarification added to |            |
|             |            | UC101 on 4-7-2010.                    |            |
|             |            |                                       |            |
|             |            | \- Changed record code values based   |            |
|             |            | on input from SDG                     |            |
|             |            |                                       |            |
|             |            | \- Removed standard format version    |            |
|             |            | number from filename because this     |            |
|             |            | would impact IM                       |            |
|             |            |                                       |            |
|             |            | \- Changed the error message for the  |            |
|             |            | input system field.                   |            |
|             |            |                                       |            |
|             |            | \- Changed type description for       |            |
|             |            | PaymentTypeCode to "AN"               |            |
|             |            |                                       |            |
|             |            | \- Removed "blank" from possible      |            |
|             |            | values for Standard Entry Class Code. |            |
|             |            |                                       |            |
|             |            | \- changed stored name for            |            |
|             |            | InputSystem field to Agency           |            |
|             |            | Identification                        |            |
|             |            |                                       |            |
|             |            | \- changed stored name for AgencyText |            |
|             |            | field to Agency Text                  |            |
+-------------+------------+---------------------------------------+------------+
| 3.5         | Dorothy    | CR1822:                               | 3-19-2010  |
|             | Carpenter  |                                       |            |
|             |            | Changed field length of Agency        |            |
|             |            | Address Line 1-5 from 27 to 25 to     |            |
|             |            | coincide with standard check file.    |            |
|             |            |                                       |            |
|             |            | Changed number of Agency address      |            |
|             |            | lines from 5 to 3 and changed the     |            |
|             |            | name to Agency Address_Stub           |            |
|             |            |                                       |            |
|             |            | Added Agency Name_stub field to be    |            |
|             |            | consistent with the standard check    |            |
|             |            | file.                                 |            |
|             |            |                                       |            |
|             |            | Changed validation and error message  |            |
|             |            | for ACH Transaction Code to reject at |            |
|             |            | the payment instead of the file       |            |
|             |            | level.                                |            |
|             |            |                                       |            |
|             |            | Revised the structure of the file to  |            |
|             |            | include a separate schedule and       |            |
|             |            | payment detail based on method of     |            |
|             |            | payment.                              |            |
|             |            |                                       |            |
|             |            | Changed field length for              |            |
|             |            | CheckEnclosureCode from 4 to 8 to     |            |
|             |            | accommodate length of current and     |            |
|             |            | potential values.                     |            |
+-------------+------------+---------------------------------------+------------+
| 3.4         | Dorothy    | CR1822                                | 3-16-2010  |
|             | Carpenter  |                                       |            |
|             |            | Added Territory Name code as an       |            |
|             |            | optional field for IAT payments.      |            |
|             |            |                                       |            |
|             |            | Changed State Code to an optional     |            |
|             |            | field for IAT payments.               |            |
|             |            |                                       |            |
|             |            | Noted that remittance addenda for IAT |            |
|             |            | payments should continue to be sent;  |            |
|             |            | FMS will only format IAT addenda      |            |
|             |            |                                       |            |
|             |            | Corrected file positions based on     |            |
|             |            | AMUQ 2/9/2010                         |            |
|             |            |                                       |            |
|             |            | Added a stored name of "Agency EIN"   |            |
|             |            | to                                    |            |
|             |            | FederalEmployerIdentificationNumber   |            |
|             |            |                                       |            |
|             |            | Changed stored name for               |            |
|             |            | PayeeIdentifierExtension from "n/a"   |            |
|             |            | to "Suffix for Account Number"        |            |
|             |            |                                       |            |
|             |            | Added a stored name of "Payment       |            |
|             |            | Identifier" for the PaymentIdentifier |            |
|             |            | field                                 |            |
|             |            |                                       |            |
|             |            | Added a stored name of                |            |
|             |            | "Reconcilement" to the reconcilement  |            |
|             |            | field                                 |            |
|             |            |                                       |            |
|             |            | Added error code to the InputSystem   |            |
|             |            | field                                 |            |
|             |            |                                       |            |
|             |            | Clarified how AgencyText field is     |            |
|             |            | printed on the check                  |            |
|             |            |                                       |            |
|             |            | Clarified how InputSystem should be   |            |
|             |            | validated against agency profile      |            |
+-------------+------------+---------------------------------------+------------+
| 3.3         | Dorothy    | CR0612 Added glossary and comments    | 1-28-2010  |
|             | Carpenter  | from RFCs.                            |            |
|             |            |                                       |            |
|             |            | removed debits from ACH transaction   |            |
|             |            | code and will add them to a future    |            |
|             |            | release of document.                  |            |
|             |            |                                       |            |
|             |            | added agency address to schedule      |            |
|             |            | record for stubs.                     |            |
|             |            |                                       |            |
|             |            | replaced the key (party name and      |            |
|             |            | amount) with a more distinct key      |            |
|             |            | (payment identifier).                 |            |
|             |            |                                       |            |
|             |            | incorporated comments provided by     |            |
|             |            | Dick Bauder                           |            |
+-------------+------------+---------------------------------------+------------+
| 3.2         | Dorothy    | Updated based on VA, RRB and SSA CPSS | 12-14-2009 |
|             | Carpenter  | review, downstream system review.     |            |
|             |            | Removed secondary records and added a |            |
|             |            | reconcilement field to the common     |            |
|             |            | record to accommodate VA with the     |            |
|             |            | goal of moving more payment types     |            |
|             |            | toward this variable field in lieu of |            |
|             |            | providing distinct data elements in a |            |
|             |            | secondary record. Removed rule that   |            |
|             |            | code 2 checks are only available for  |            |
|             |            | vendor/misc payments. Removed all     |            |
|             |            | balancing except for schedule/file    |            |
|             |            | level.                                |            |
+-------------+------------+---------------------------------------+------------+
| 3.1         | Carole     | Updated with changes after the        | 10-16-2009 |
|             | Kampe      | walkthrough for final approval        |            |
+-------------+------------+---------------------------------------+------------+
| 3.0         | Carole     | Moved to Input File Spec Template and | 5-17-2009  |
|             | Kampe      | added error codes for notification    |            |
+-------------+------------+---------------------------------------+------------+
| 1.0         | Carole     | Initial Draft                         | 12-6-2007  |
|             | Kampe      |                                       |            |
|             |            | Changed payment types                 |            |
+-------------+------------+---------------------------------------+------------+
