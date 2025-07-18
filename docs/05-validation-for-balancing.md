# 1.5 Validation for Balancing

| Rule | Result | Error Code |
|------|--------|------------|
| If the ACH_Transaction Code indicates "Prenote" all of the payments in the schedule must be zero dollar amounts. | If not, reject the file. | Error Reason Group 4 Message 5 |
| If zero dollar payment amounts are received in a schedule, the ACH Transaction Code must indicate "Prenote" or the SEC code = CTX. | If not, reject the file | Error Reason Group 4 Message 4 |
| If greater than zero dollar payment amounts are received in the schedule, the ACH Transaction Code must **not** indicate "Zero Dollar Credit" when SEC code = CTX. | If not, reject the file. | Error Reason Group 4 Message 3 |
| The ScheduleTrailer.ScheduleCount must equal the accumulated number of Payments in the schedule. | If not, reject the schedule | Error Reason Group 3 Message 6 for ACH schedules.<br>Error Reason Group 3 Message 4 for check schedules. |
| The ScheduleTrailer.ScheduleAmount must equal the accumulated amount of Payments in the schedule. | If not, reject the schedule | Error Reason Group 3 Message 5 for ACH schedules.<br>Error Reason Group 3 Message 3 for check schedules. |
| The FileTrailer.TotalCount_Records must equal the accumulated number of records on the file (this includes header and trailer records). | If not, reject the file | Error Reason Group 3 Message 2 |
| The FileTrailer.TotalCount_Payments must equal the accumulated number of payments on the file. | If not, reject the file | Error Reason Group 3 Message 2 |
| The FileTrailer.TotalAmount_Payments must equal the accumulated amount of payments on the file. | If not, reject the file | Error Reason Group 3 Message 1 |
