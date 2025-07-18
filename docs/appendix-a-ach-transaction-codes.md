# 3.1 Appendix A - ACH Transaction Codes

The following table defines the Transaction Code values received in an
ACH Detail Payment record and the value PAM should record.

```
  ---------------------------------------------------------------------------
   **Transaction  **Definition**                          **Account Type (for
      Code**                                              recording in PAM)**
  --------------- --------------------------------------- -------------------
        22        Checking Account Credit                          C

        23        Checking Account Credit Prenote                  C

        24        Checking Account Zero Dollar Credit              C
                  with remittance

        32        Savings Account Credit                           S

        33        Savings Account Credit Prenote                   S

        34        Savings Account Zero Dollar Credit with          S
                  remittance

        42        General Ledger Credit                            G

        43        General Ledger Credit Prenote                    G

        52        Loan Account Credit                              L

        53        Loan Account Credit Prenote                      L
  ---------------------------------------------------------------------------
```

For Custom Agency Rule ID = "IRS" and Depositor Account Number = "BONDS":

```
  ---------------------------------------------------------------------------
   **Transaction  **Definition**             **Account Type       **Savings
      Code**                                 (for recording in   Bonds Owner
                                             PAM)**                Type**
  --------------- -------------------------- ------------------ -------------
        22        Checking Account Credit    C                      Gift

        32        Savings Account Credit     S                      Self
  ---------------------------------------------------------------------------
```
