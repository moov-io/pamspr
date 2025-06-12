# 3.3 Appendix C - Addressing Reference Information

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

