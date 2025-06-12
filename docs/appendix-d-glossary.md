# 3.4 Appendix D - Glossary of Terms

  ----------------------------------------------------------------------------------------------
  Field Name                            Definition
  ------------------------------------- --------------------------------------------------------
  ACH_TransactionCode                   Designates a transaction's account type (checking,
                                        savings, or general ledger) and the transaction type
                                        (credit, or prenote).

  AccountNumber                         The payment recipient's bank account number.

  AccountClassificationAmount           The transaction amount associated with TAS/BETC to be
                                        reported to Treasury for the Central Accounting
                                        Reporting System (CARS) compliancy.

  Addenda Information                   Additional information about the payment to be included
                                        in ACH addenda (PPD+ and CCD+)

  AgencyAccountIdentifier               Account number used internally by an agency to identify
                                        a payee.

  AgencyIdentifier                      Used in conjunction with the main account code,
                                        represents the department, agency, or establishment of
                                        the U.S. government that is responsible for the TAS.

  AgencyLocationCode                    Identifies the accounting office within an agency that
                                        reports disbursements and collections to Treasury.

  AgencyPaymentTypeCode                 Internal agency code used to identify the type of
                                        payment within the specific agency.

  AgencyACHText                         Used to identify the agency responsible for the
                                        payments. First four bytes go to FedACH.

  AllocationTransferAgencyIdentifier    Represents the agency receiving funds through an
                                        allocation transfer.

  Amount                                The amount of the transaction.

  Amount eligible for offset            Indicates the amount of the payment that is eligible for
                                        offset. This is sent to TOP for offsetting purposes.

  AvailabilityTypeCode                  Identifies no-year accounts (X), clearing/suspense
                                        accounts (F), and Treasury central summary general
                                        ledger accounts (A)

  BeginningPeriodOfAvailability         In annual and multiyear accounts, identifies the first
                                        year of availability under law that an account may incur
                                        new obligations.

  BusinessEventTypeCode                 Identifies the type of activity (gross disbursement,
                                        offsetting collection, investment in Treasury
                                        securities, etc.) and the effect of a transaction on the
                                        Fund Balance With Treasury (FBWT). Is used in
                                        combination with the Treasury Account Symbol to classify
                                        transactions reported to Treasury through all
                                        Governmentwide Accounting (CARS-compliant) Financial
                                        Management Systems.

  CheckPaymentEnclosureCode             Code denoting that a check will include an enclosure --
                                        either a stub or a letter.

  CheckLegendText1                      Free-form field for agency's use -- all data in this
                                        field is printed on the check.

  CheckLegendText2                      Free-form field for agency's use -- all data in this
                                        field is printed on the check.

  CityName                              Name of the city in which the payee resides.

  ConsularCode                          Indicates the code for mailing bulk check shipments to
                                        foreign countries.

  CountryCodeText                       The ISO country code for the country in which the
                                        payment recipient resides. Used for IAT payments

  CountryName                           The full name of the country in which the payment
                                        recipient resides.

  EndingPeriodOfAvailability            Identifies the last year of availability under law that
                                        an account (annual and multiyear) may incur new
                                        obligations.

  FederalEmployerIdentificationNumber   A number assigned to businesses by the IRS and used by
                                        DHHS for tracking child support payments.

                                        

                                        

  InputSystem                           Identifies the sending trading partner.

  IsCredit                              Indicates if the item is a debit or credit, for example,
                                        a value of \'1\' would mean that it is a \'credit\'. If
                                        the attribute is not populated, then the default value
                                        is understood to be \'0\'.

  IsSpecialHandling                     Requires agency / Treasury collaboration prior to use.
                                        Used for manual handling of checks such as the
                                        designated agent.

  IsTOP_Offset                          Indicates the payment eligibility of the payment offset

  MainAccountCode                       Identifies the type and purpose of the fund.

  PayeeName                             Contains the full name of a payee. The name of the party
                                        whether individual or organization.

  PayeeName_Secondary                   Contains the full name of an additional payee. The name
                                        of the secondary party whether individual or
                                        organization.

  PayeeAddressLine_1-4                  The mailing address of the payee. Address lines left
                                        blank will not be printed.

  PayeeIdentifier                       The payee\'s Taxpayer Identification Number (TIN),
                                        Vendor ID, Social Security Number (SSN) Employer Tax
                                        Identification Number (EIN), Individual Taxpayer
                                        Identification Number (ITIN), or other valid Payee ID.

  PayeeIdentifier_Secondary             The Secondary party's Taxpayer Identification Number
                                        (TIN), Vendor ID, Social Security Number (SSN), Employer
                                        Tax Identification Number (EIN), Individual Taxpayer
                                        Identification Number (ITIN), or other valid ID.

  PayeeStreetAddress                    Address of the payee for mailing TOP offset letter, if
                                        needed, for ACH payments. For IAT payments, this field
                                        is required.

  PayerMechanism                        Identifies the payment medium for the transaction.

  PaymentDescriptionCode                Identifier to allow TOP to determine the percentage of
                                        offset to be taken during delinquent debt screening.
                                        Examples include, "IN" (Initial), "RE" (Regular), "FI"
                                        (Final", "BA" (Bankruptcy).

  PaymentIdentifier                     Used as a unique identifier for each payment to relate
                                        the records appropriately. Also used for matched
                                        letters.

  PaymentIdentificationLine_1-14        Content to be printed on check stubs.

  Payment Recipient TIN indicator       Indicates whether the TIN (Taxpayer Identification
                                        Number) is an SSN (Social Security Number), Employer
                                        Identification Number (EIN), or Individual Taxpayer
                                        Identification Number (ITIN).

  PaymentTypeCode                       The type of payment as defined in [Appendix
                                        E](#appendix-e---paymenttypecode-values)

  PostalCode                            A code of letters and/or digits added to an address
                                        line. (Zip code)

  PostalCodeExtension                   4-digit extension of postal code.

  PostNetBarcodeDeliveryPoint           Part of the PostNet barcode that is applied to the check
                                        for obtaining postage discounts.

  PostNetBarcodeCheckDigit              Part of the PostNet barcode that is applied to the check
                                        for obtaining postage discounts.

                                        

                                        

  Reconcilement                         Free-form field for agency's use in reconciling the
                                        payment using PACER. Data in this field is passed to
                                        PACER when formatted correctly. Contact Treasury for
                                        additional information.

  Record Code                           A set record identifier used to specify the type of
                                        record in the file.

  RoutingNumber                         The routing number is used synonymously as ABA routing
                                        number and routing transit number. The routing number
                                        consists of 9 digits, for example XXXXYYYYC where XXXX
                                        is Federal Reserve Routing Symbol, YYYY is ABA
                                        Institution Identifier, and C is the Check Digit. The
                                        first two digits of the nine digit routing number must
                                        be in the ranges 00 through 12, 21 through 32, 61
                                        through 72, or 80. The digits are assigned as follows:
                                        00 is used by the United States Government; 01 through
                                        12 are the normal routing numbers; 21 through 32 were
                                        assigned only to thrift institutions (e.g. credit unions
                                        and savings banks) through 1985, currently are still
                                        used by the thrift institutions, or their successors; 61
                                        through 72 are used for electronic transactions; 80 is
                                        used for traveler\'s checks. The first two digits
                                        correspond to the 12 Federal Reserve Banks.

  ScheduleAmount                        The total dollar amount associated with the schedule.

  ScheduleCount                         Total number of payments for each schedule.

  ScheduleNumber                        A number assigned by the agency to identify a group of
                                        payments.

  Standard Payment Request Version      The version of the PAM Standard Payment Request that
  Number                                this file is programmed against.

  StandardEntryClassCode                A three digit code defined by NACHA to characterize the
                                        nature of the ACH transaction. PPD (Prearranged Payment
                                        and Deposit Entry) is an entry initiated by an
                                        organization pursuant to a standing or a single entry
                                        authorization from a receiver to effect a transfer of
                                        funds to impact a consumer account of the receiver. CCD
                                        (Cash Concentration or Disbursement entry) is an entry
                                        initiated by an organization to affect a transfer of
                                        funds to impact the account of that organization or
                                        another organization. IAT (International ACH
                                        Transaction) is a payment that involves a financial
                                        agency's office that is not located in the territorial
                                        jurisdiction of the United States. CTX (Corporate Trade
                                        Exchange) is a corporate format which allows for up to
                                        9,999 addenda records allowing full and complete
                                        remittance information to be transmitted with a single
                                        payment.

  StateCodeText                         The US state code as managed by the USPS (United States
                                        Postal Service) .

  StateName                             The foreign state, province or territory name

  StubLine_1 (through 14)               Agency data to be printed on the stub. Agencies
                                        interested in using this service should contact their
                                        local RFC.

  Sub-accountCode                       Identifies an available receipt or other
                                        Treasury-defined subdivision of the main account.

  Sub-levelPrefixCode                   When populated, represents a programmatic breakdown of
                                        the account for Treasury publication purposes.

  TotalAmount_Payments                  Total amount of the payment request file. Used to verify
                                        the integrity of the file.

  TotalCount_Payments                   Total number of payments in the file. Used to verify the
                                        integrity of the file.

  TotalCount_Records                    Total number of records in the file. Used to verify the
                                        integrity of the file.

  USPSIntelligentMailBarcode            Field used for all intelligent barcode data elements. An
                                        agency must work with their servicing RFC and the US
                                        Postal Service to arrange for this service.
  ----------------------------------------------------------------------------------------------

##  

