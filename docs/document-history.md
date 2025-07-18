# Document History

This document provides a comprehensive history of changes made to the Payment Automation Manager (PAM) Standard Payment Request (SPR) file format specification over time.

```
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
```
