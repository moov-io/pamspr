# 1.2 General Structure of File

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

- DNP Record RC=DD (none- one for each Check Payment Data Record)

<!-- -->

- **Check** Schedule Trailer Control Record RC=T (one for each Schedule
  Header Record)

- **File Trailer Control Record** RC=E (There can only be one File
  Trailer Control Record)
