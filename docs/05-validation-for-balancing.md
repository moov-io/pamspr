# 1.5 Validation for Balancing

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