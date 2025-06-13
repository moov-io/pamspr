package pamspr

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// BenchmarkReader_vs_TraditionalReader compares streaming vs traditional reader performance
func BenchmarkReader_vs_TraditionalReader(b *testing.B) {
	// Create test file with varying sizes
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		testData := generateTestFileData(size)

		b.Run(fmt.Sprintf("Traditional_Reader_%d_payments", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := NewReader(strings.NewReader(testData))
				file, err := reader.Read()
				if err != nil {
					b.Fatal(err)
				}
				// Force evaluation to prevent optimization
				_ = len(file.Schedules)
			}
		})

		b.Run(fmt.Sprintf("Streaming_Reader_%d_payments", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reader := NewReader(strings.NewReader(testData))

				paymentCount := 0
				err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
					paymentCount++
					return true
				})
				if err != nil {
					b.Fatal(err)
				}
				// Force evaluation
				_ = paymentCount
			}
		})
	}
}

// BenchmarkStreamingWriter_vs_TraditionalWriter compares writer performance
func BenchmarkStreamingWriter_vs_TraditionalWriter(b *testing.B) {
	sizes := []int{100, 1000, 10000}

	for _, size := range sizes {
		testFile := generateTestFile(size)

		b.Run(fmt.Sprintf("Traditional_Writer_%d_payments", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var buf bytes.Buffer
				writer := NewWriter(&buf)
				err := writer.Write(testFile)
				if err != nil {
					b.Fatal(err)
				}
				// Force evaluation
				_ = buf.Len()
			}
		})

		b.Run(fmt.Sprintf("Streaming_Writer_%d_payments", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var buf bytes.Buffer
				writer := NewWriter(&buf)

				// Write header
				err := writer.WriteFileHeader(testFile.Header)
				if err != nil {
					b.Fatal(err)
				}

				// Write schedules and payments
				for _, schedule := range testFile.Schedules {
					err := writer.WriteScheduleHeader(schedule)
					if err != nil {
						b.Fatal(err)
					}

					for _, payment := range schedule.GetPayments() {
						err := writer.WritePayment(payment)
						if err != nil {
							b.Fatal(err)
						}
					}

					err = writer.WriteScheduleTrailer(schedule.GetTrailer())
					if err != nil {
						b.Fatal(err)
					}
				}

				// Write trailer
				err = writer.WriteFileTrailer(testFile.Trailer)
				if err != nil {
					b.Fatal(err)
				}

				// Force evaluation
				_ = buf.Len()
			}
		})
	}
}

// BenchmarkMemoryUsage tests memory usage patterns
func BenchmarkMemoryUsage(b *testing.B) {
	sizes := []int{1000, 10000, 50000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("Traditional_Memory_%d_payments", size), func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				testData := generateTestFileData(size)
				reader := NewReader(strings.NewReader(testData))
				file, err := reader.Read()
				if err != nil {
					b.Fatal(err)
				}

				// Simulate some processing
				totalAmount := int64(0)
				for _, schedule := range file.Schedules {
					for _, payment := range schedule.GetPayments() {
						totalAmount += payment.GetAmount()
					}
				}
				_ = totalAmount
			}
		})

		b.Run(fmt.Sprintf("Streaming_Memory_%d_payments", size), func(b *testing.B) {
			b.ReportAllocs()

			for i := 0; i < b.N; i++ {
				testData := generateTestFileData(size)
				reader := NewReader(strings.NewReader(testData))

				totalAmount := int64(0)
				err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
					totalAmount += payment.GetAmount()
					return true
				})
				if err != nil {
					b.Fatal(err)
				}
				_ = totalAmount
			}
		})
	}
}

// BenchmarkFieldFormatting compares field formatting performance
func BenchmarkFieldFormatting(b *testing.B) {
	payment := &ACHPayment{
		RecordCode:              "02",
		AgencyAccountIdentifier: "ACC123456789012 ",
		Amount:                  150000,
		PayeeName:               "ACME CORPORATION",
		PayeeAddressLine1:       "123 BUSINESS STREET",
		CityName:                "NEW YORK",
		StateName:               "NEW YORK",
		StateCodeText:           "NY",
		PostalCode:              "10001",
		RoutingNumber:           "021000021",
		AccountNumber:           "123456789012345",
		PaymentID:               "PAY000000001",
	}

	b.Run("Traditional_Field_Formatting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			writer := NewWriter(&buf)

			// Format payment using traditional concatenation
			err := writer.WritePayment(payment)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Streaming_Field_Formatting", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			var buf bytes.Buffer
			writer := NewWriter(&buf)

			// Format payment using string builder
			err := writer.WritePayment(payment)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkLargeFileProcessing simulates large file processing
func BenchmarkLargeFileProcessing(b *testing.B) {
	// Skip this test if running short benchmarks
	if testing.Short() {
		b.Skip("Skipping large file benchmark in short mode")
	}

	// Test with very large file (100k payments)
	size := 100000

	b.Run("Streaming_Large_File_Validation", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			testData := generateTestFileData(size)
			reader := NewReader(strings.NewReader(testData))

			err := reader.ValidateFileStructureOnly()
			if err != nil {
				b.Fatal(err)
			}

			stats := reader.GetStats()
			_ = stats.PaymentsProcessed
		}
	})

	b.Run("Streaming_Large_File_Payment_Processing", func(b *testing.B) {
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			testData := generateTestFileData(size)
			reader := NewReader(strings.NewReader(testData))

			paymentCount := 0
			totalAmount := int64(0)

			err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
				paymentCount++
				totalAmount += payment.GetAmount()
				return true
			})
			if err != nil {
				b.Fatal(err)
			}

			_ = paymentCount
			_ = totalAmount
		}
	})
}

// BenchmarkConcurrentProcessing tests concurrent file processing
func BenchmarkConcurrentProcessing(b *testing.B) {
	size := 1000
	testData := generateTestFileData(size)

	b.Run("Sequential_Processing", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			reader := NewReader(strings.NewReader(testData))

			paymentCount := 0
			err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
				// Simulate some processing work
				_ = len(payment.GetPayeeName())
				paymentCount++
				return true
			})
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Concurrent_Processing", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				reader := NewReader(strings.NewReader(testData))

				paymentCount := 0
				err := reader.ProcessPaymentsOnly(func(payment Payment, scheduleIndex, paymentIndex int) bool {
					// Simulate some processing work
					_ = len(payment.GetPayeeName())
					paymentCount++
					return true
				})
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}

// Helper functions for benchmark data generation

func generateTestFileData(numPayments int) string {
	var builder strings.Builder

	// File header
	builder.WriteString("H BENCHMARK TEST SYSTEM                 5020")
	builder.WriteString(strings.Repeat(" ", 804))
	builder.WriteByte('\n')

	// Calculate payments per schedule (max 1000 per schedule)
	paymentsPerSchedule := 1000
	if numPayments < 1000 {
		paymentsPerSchedule = numPayments
	}

	scheduleCount := (numPayments + paymentsPerSchedule - 1) / paymentsPerSchedule

	for s := 0; s < scheduleCount; s++ {
		// ACH Schedule header
		builder.WriteString("01AGCY")
		builder.WriteString(fmt.Sprintf("%014d", s+1))
		builder.WriteString("Vendor                   CCD12345678 1234567890")
		builder.WriteString(strings.Repeat(" ", 783))
		builder.WriteByte('\n')

		// Payments in this schedule
		paymentsInSchedule := paymentsPerSchedule
		remaining := numPayments - (s * paymentsPerSchedule)
		if remaining < paymentsPerSchedule {
			paymentsInSchedule = remaining
		}

		scheduleAmount := int64(0)
		for p := 0; p < paymentsInSchedule; p++ {
			paymentAmount := int64(100000) // $1000.00
			scheduleAmount += paymentAmount

			// ACH Payment
			builder.WriteString("02")
			builder.WriteString(fmt.Sprintf("ACC%013d ", p))
			builder.WriteString(fmt.Sprintf("%010d", paymentAmount))
			builder.WriteString("V0")
			builder.WriteString(fmt.Sprintf("PAYEE NAME %05d              ", p))
			builder.WriteString("123 MAIN STREET                    ")
			builder.WriteString("SUITE 100                          ")
			builder.WriteString("ANYTOWN                    ")
			builder.WriteString("STATE     ")
			builder.WriteString("ST")
			builder.WriteString("12345")
			builder.WriteString("1234 ")
			builder.WriteString("US")
			builder.WriteString("021000021")
			builder.WriteString(fmt.Sprintf("ACCT%012d    ", p))
			builder.WriteString("22")
			builder.WriteString("         ")
			builder.WriteString("                                   ")
			builder.WriteString(fmt.Sprintf("PAY%017d ", p))
			builder.WriteString(strings.Repeat(" ", 100)) // Reconcilement
			builder.WriteString("123456789")
			builder.WriteString("1")
			builder.WriteString("0")
			builder.WriteString(fmt.Sprintf("%010d", paymentAmount))
			builder.WriteString("                                   ") // Address line 3
			builder.WriteString("                                   ") // Address line 4
			builder.WriteString("UNITED STATES                           ")
			builder.WriteString("   ")
			builder.WriteString("VENDOR PAYMENT                  ")
			builder.WriteString("ACH CREDIT          ")
			builder.WriteString("  ")
			builder.WriteString(strings.Repeat(" ", 284)) // Filler
			builder.WriteByte('\n')
		}

		// Schedule trailer
		builder.WriteString("T ")
		builder.WriteString(strings.Repeat(" ", 10))
		builder.WriteString(fmt.Sprintf("%08d", paymentsInSchedule))
		builder.WriteString("   ")
		builder.WriteString(fmt.Sprintf("%015d", scheduleAmount))
		builder.WriteString(strings.Repeat(" ", 812))
		builder.WriteByte('\n')
	}

	// File trailer
	totalRecords := 2 + (scheduleCount * 2) + numPayments // H + E + (schedule headers + trailers) + payments
	totalAmount := int64(numPayments) * 100000            // $1000.00 per payment

	builder.WriteString("E ")
	builder.WriteString(fmt.Sprintf("%018d", totalRecords))
	builder.WriteString(fmt.Sprintf("%018d", numPayments))
	builder.WriteString(fmt.Sprintf("%018d", totalAmount))
	builder.WriteString(strings.Repeat(" ", 794))
	builder.WriteByte('\n')

	return builder.String()
}

func generateTestFile(numPayments int) *File {
	file := &File{
		Header: &FileHeader{
			RecordCode:               "H ",
			InputSystem:              "BENCHMARK TEST SYSTEM",
			StandardPaymentVersion:   "502",
			IsRequestedForSameDayACH: "0",
		},
		Schedules: make([]Schedule, 0),
	}

	// Calculate payments per schedule
	paymentsPerSchedule := 1000
	if numPayments < 1000 {
		paymentsPerSchedule = numPayments
	}

	scheduleCount := (numPayments + paymentsPerSchedule - 1) / paymentsPerSchedule
	totalAmount := int64(0)

	for s := 0; s < scheduleCount; s++ {
		schedule := &ACHSchedule{
			Header: &ACHScheduleHeader{
				RecordCode:              "01",
				AgencyACHText:           "AGCY",
				ScheduleNumber:          fmt.Sprintf("%014d", s+1),
				PaymentTypeCode:         "Vendor",
				StandardEntryClassCode:  "CCD",
				AgencyLocationCode:      "12345678",
				FederalEmployerIDNumber: "1234567890",
			},
			BaseSchedule: BaseSchedule{
				Payments: make([]Payment, 0),
			},
		}

		// Payments in this schedule
		paymentsInSchedule := paymentsPerSchedule
		remaining := numPayments - (s * paymentsPerSchedule)
		if remaining < paymentsPerSchedule {
			paymentsInSchedule = remaining
		}

		scheduleAmount := int64(0)
		for p := 0; p < paymentsInSchedule; p++ {
			paymentAmount := int64(100000) // $1000.00
			scheduleAmount += paymentAmount
			totalAmount += paymentAmount

			payment := &ACHPayment{
				RecordCode:                   "02",
				AgencyAccountIdentifier:      fmt.Sprintf("ACC%013d ", p),
				Amount:                       paymentAmount,
				AgencyPaymentTypeCode:        "V",
				IsTOP_Offset:                 "0",
				PayeeName:                    fmt.Sprintf("PAYEE NAME %05d", p),
				PayeeAddressLine1:            "123 MAIN STREET",
				PayeeAddressLine2:            "SUITE 100",
				CityName:                     "ANYTOWN",
				StateName:                    "STATE",
				StateCodeText:                "ST",
				PostalCode:                   "12345",
				PostalCodeExtension:          "1234",
				CountryCodeText:              "US",
				RoutingNumber:                "021000021",
				AccountNumber:                fmt.Sprintf("ACCT%012d", p),
				ACH_TransactionCode:          "22",
				PaymentID:                    fmt.Sprintf("PAY%017d", p),
				TIN:                          "123456789",
				PaymentRecipientTINIndicator: "1",
				AdditionalPayeeTINIndicator:  "0",
				AmountEligibleForOffset:      fmt.Sprintf("%010d", paymentAmount),
				CountryName:                  "UNITED STATES",
				SubPaymentTypeCode:           "VENDOR PAYMENT",
				PayerMechanism:               "ACH CREDIT",
			}

			schedule.Payments = append(schedule.Payments, payment)
		}

		schedule.Trailer = &ScheduleTrailer{
			RecordCode:     "T ",
			ScheduleCount:  int64(paymentsInSchedule),
			ScheduleAmount: scheduleAmount,
		}

		file.Schedules = append(file.Schedules, schedule)
	}

	// Calculate total records
	totalRecords := 2 + (scheduleCount * 2) + numPayments // H + E + (schedule headers + trailers) + payments

	file.Trailer = &FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   int64(totalRecords),
		TotalCountPayments:  int64(numPayments),
		TotalAmountPayments: totalAmount,
	}

	return file
}
