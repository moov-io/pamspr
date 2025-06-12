# 1.9 Specification Notes

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