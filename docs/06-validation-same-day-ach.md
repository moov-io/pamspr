# 1.6 Validation for Same Day ACH (SDA)

| Rule | Result | Error Code |
|------|--------|------------|
| The SPR can only contain schedules with method of payment ACH | If not, and the IsRequestedForSDA value is "1", reject the file. | Error Reason Group 4 Message 7 |
| All individual payment amounts must be less than or equal to the MAX SDA Amount ($1,000,000) | If not, and the IsRequestedForSDA value is "1", reject the file. | Error Reason Group 4 Message 8 |
| Payment Type values must be allowed for SDA. Restricted Payment Types: None | If not, and the IsRequestedForSDA value is "1", reject the file. | Error Reason Group 4 Message 9 |
| SEC code must be a value other than IAT | If not, and the IsRequestedForSDA value is "1", reject the file. | Error Reason Group 4 Message 10 |
