# 2.6 ACH Addendum Record

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

