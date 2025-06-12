# 1.6 Validation for Same Day ACH (SDA)

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