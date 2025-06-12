# 1.3 File Structure Validations

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