package pamspr

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewFieldFormatter(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	if f == nil || f.validator == nil {
		t.Fatal("expected formatter with validator")
	}
}

func TestFormatRecord_FileHeader(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	header := &FileHeader{
		RecordCode:               "H ",
		InputSystem:              "TEST_SYSTEM",
		StandardPaymentVersion:   "502",
		IsRequestedForSameDayACH: "0",
	}

	out, err := f.FormatRecord(header, "H ")
	if err != nil {
		t.Fatalf("FormatRecord: %v", err)
	}
	if len(out) != RecordLength {
		t.Fatalf("expected length %d, got %d", RecordLength, len(out))
	}
	if out[:2] != "H " {
		t.Fatalf("expected record code H , got %q", out[:2])
	}
	if !strings.Contains(out[:50], "TEST_SYSTEM") {
		t.Fatalf("expected InputSystem in output, got %q", out[:50])
	}
	if out[42:45] != "502" {
		t.Fatalf("expected version 502 at cols 43-45, got %q", out[42:45])
	}
}

func TestFormatRecord_FileTrailerNumeric(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	// Non-pointer path + numeric format tags
	trailer := FileTrailer{
		RecordCode:          "E ",
		TotalCountRecords:   12,
		TotalCountPayments:  5,
		TotalAmountPayments: 12345,
	}

	out, err := f.FormatRecord(trailer, "E ")
	if err != nil {
		t.Fatalf("FormatRecord: %v", err)
	}
	if len(out) != RecordLength {
		t.Fatalf("expected length %d, got %d", RecordLength, len(out))
	}
	if out[:2] != "E " {
		t.Fatalf("expected E  prefix, got %q", out[:2])
	}
}

func TestFormatRecord_ACHScheduleHeader(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	header := &ACHScheduleHeader{
		RecordCode:             "01",
		AgencyACHText:          "TEST",
		ScheduleNumber:         "123",
		PaymentTypeCode:        "SALARY",
		StandardEntryClassCode: "PPD",
		AgencyLocationCode:     "12345678",
	}

	out, err := f.FormatRecord(header, "01")
	if err != nil {
		t.Fatalf("FormatRecord: %v", err)
	}
	if len(out) != RecordLength {
		t.Fatalf("expected length %d, got %d", RecordLength, len(out))
	}
	if out[:2] != "01" {
		t.Fatalf("expected 01 prefix, got %q", out[:2])
	}
}

func TestFormatRecord_ACHPayment(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	pay := &ACHPayment{
		RecordCode:              "02",
		AgencyAccountIdentifier: "ACC001",
		Amount:                  9999,
		PayeeName:               "JANE DOE",
		RoutingNumber:           "021000021",
		AccountNumber:           "123456789",
		ACH_TransactionCode:     "22",
		PaymentID:               "PAY1",
	}
	out, err := f.FormatRecord(pay, "02")
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != RecordLength {
		t.Fatalf("len=%d", len(out))
	}
	if out[:2] != "02" {
		t.Fatalf("prefix %q", out[:2])
	}
}

func TestFormatRecord_UnknownRecordCode(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	_, err := f.FormatRecord(&FileHeader{RecordCode: "XX"}, "ZZ")
	if err == nil {
		t.Fatal("expected error for unknown record code")
	}
}

func TestValueToString(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	cases := []struct {
		in   interface{}
		want string
	}{
		{nil, ""},
		{"abc", "abc"},
		{int(42), "42"},
		{int32(7), "7"},
		{int64(99), "99"},
		{float64(1.5), "1.5"},
		{true, "1"},
		{false, "0"},
		{struct{ X int }{3}, "{3}"},
	}
	for _, tc := range cases {
		got := f.valueToString(tc.in)
		if got != tc.want {
			t.Errorf("valueToString(%v)=%q want %q", tc.in, got, tc.want)
		}
	}
}

func TestFormatHelpers(t *testing.T) {
	f := NewFieldFormatter(NewValidator())

	if got := f.formatField("hi", 5); got != "hi   " {
		t.Fatalf("formatField pad: %q", got)
	}
	if got := f.formatField("toolong", 4); got != "tool" {
		t.Fatalf("formatField truncate: %q", got)
	}

	if got := f.formatFieldRightJustified("42", 5, '0'); got != "00042" {
		t.Fatalf("right justified: %q", got)
	}
	if got := f.formatFieldRightJustified("  x  ", 3, ' '); got != "  x" {
		t.Fatalf("right justified trim: %q", got)
	}
	if got := f.formatFieldRightJustified("abcdef", 3, '0'); got != "abc" {
		t.Fatalf("right justified truncate: %q", got)
	}

	if got := f.formatFieldNoJustify("ab", 4); got != "ab  " {
		t.Fatalf("nojust pad: %q", got)
	}
	if got := f.formatFieldNoJustify("abcd", 4); got != "abcd" {
		t.Fatalf("nojust exact: %q", got)
	}
	if got := f.formatFieldNoJustify("abcdef", 3); got != "abc" {
		t.Fatalf("nojust truncate: %q", got)
	}

	if got := f.formatNumeric("12a3", 6); got != "000123" {
		t.Fatalf("numeric: %q", got)
	}
	if got := f.formatNumeric("1234567", 4); got != "1234" {
		t.Fatalf("numeric truncate: %q", got)
	}

	if got := f.formatAmount(150, 6); got != "000150" {
		t.Fatalf("amount: %q", got)
	}
}

func TestGetFormatterConfig_Tags(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	def := NewFieldDef(1, 10, true)

	type sample struct {
		A string `format:"numeric"`
		B string `format:"amount"`
		C string `format:"filler"`
		D string `format:"nojust"`
		E string `format:"right,pad=*"`
		F string `format:"left"`
		G string // default text
	}
	st := reflect.TypeOf(sample{})

	checks := []struct {
		idx  int
		ft   FieldFormatType
		just FieldJustification
		pad  rune
	}{
		{0, FormatNumeric, JustifyRight, '0'},
		{1, FormatAmount, JustifyRight, '0'},
		{2, FormatFiller, JustifyLeft, ' '},
		{3, FormatNoJustify, JustifyLeft, ' '},
		{4, FormatText, JustifyRight, '*'},
		{5, FormatText, JustifyLeft, ' '},
		{6, FormatText, JustifyLeft, ' '},
	}
	for _, c := range checks {
		cfg := f.getFormatterConfig(def, st.Field(c.idx))
		if cfg.FormatType != c.ft {
			t.Errorf("field %d format=%s want %s", c.idx, cfg.FormatType, c.ft)
		}
		if cfg.Justification != c.just {
			t.Errorf("field %d just=%s want %s", c.idx, cfg.Justification, c.just)
		}
		if cfg.PadChar != c.pad {
			t.Errorf("field %d pad=%q want %q", c.idx, cfg.PadChar, c.pad)
		}
		if cfg.Length != 10 {
			t.Errorf("field %d length=%d", c.idx, cfg.Length)
		}
	}
}

func TestFormatFieldValue_Direct(t *testing.T) {
	f := NewFieldFormatter(NewValidator())
	def := NewFieldDef(1, 8, true)

	type tagged struct {
		Num  string `format:"numeric"`
		Amt  int64  `format:"amount"`
		Fill string `format:"filler"`
		NJ   string `format:"nojust"`
		R    string `format:"right,pad=0"`
		T    string // text left
	}
	st := reflect.TypeOf(tagged{})

	// numeric
	got, err := f.formatFieldValue("42x", def, st.Field(0))
	if err != nil || got != "00000042" {
		t.Fatalf("numeric: %q err=%v", got, err)
	}
	// amount int64
	got, err = f.formatFieldValue(int64(99), def, st.Field(1))
	if err != nil || got != "00000099" {
		t.Fatalf("amount: %q err=%v", got, err)
	}
	// amount non-int falls back to numeric string path
	got, err = f.formatFieldValue("12", def, st.Field(1))
	if err != nil || got != "00000012" {
		t.Fatalf("amount fallback: %q err=%v", got, err)
	}
	// filler
	got, err = f.formatFieldValue("ignored", def, st.Field(2))
	if err != nil || got != "        " {
		t.Fatalf("filler: %q err=%v", got, err)
	}
	// nojust
	got, err = f.formatFieldValue("ab", def, st.Field(3))
	if err != nil || got != "ab      " {
		t.Fatalf("nojust: %q err=%v", got, err)
	}
	// right pad
	got, err = f.formatFieldValue("7", def, st.Field(4))
	if err != nil || got != "00000007" {
		t.Fatalf("right: %q err=%v", got, err)
	}
	// text left
	got, err = f.formatFieldValue("hi", def, st.Field(5))
	if err != nil || got != "hi      " {
		t.Fatalf("text: %q err=%v", got, err)
	}
}
