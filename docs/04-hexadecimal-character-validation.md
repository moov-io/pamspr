# 1.4 Hexadecimal Character Validation

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