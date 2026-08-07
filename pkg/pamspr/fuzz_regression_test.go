package pamspr

import (
	"strings"
	"testing"
)

func TestRead_ShortLinesNoPanic(t *testing.T) {
	inputs := []string{"", "H", "H ", "01", "X"}
	for _, in := range inputs {
		func(in string) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("panic on input %q: %v", in, r)
				}
			}()
			r := NewReader(strings.NewReader(in + "\n"))
			_, _ = r.Read()
		}(in)
	}
}

func TestParseSchedule_ShortLine(t *testing.T) {
	r := NewReader(strings.NewReader(""))
	if _, err := r.parseSchedule(""); err == nil {
		t.Fatal("expected error for empty schedule line")
	}
	if _, err := r.parseSchedule("0"); err == nil {
		t.Fatal("expected error for 1-char schedule line")
	}
	if _, err := r.parseScheduleHeader("X"); err == nil {
		t.Fatal("expected error for short schedule header")
	}
}

func TestReader_ProcessFileEmpty(t *testing.T) {
	r := NewReader(strings.NewReader(""))
	err := r.ProcessFile(nil, nil, nil)
	if err == nil {
		t.Fatal("expected error for empty file")
	}
}

func TestReader_ValidateStructureShortLines(t *testing.T) {
	// Valid-length lines that are not a complete file
	pad := func(s string) string {
		if len(s) >= RecordLength {
			return s[:RecordLength]
		}
		return s + strings.Repeat(" ", RecordLength-len(s))
	}
	content := pad("H TEST") + "\n" + pad("01") + "\n"
	r := NewReader(strings.NewReader(content))
	_ = r.ValidateFileStructureOnly() // must not panic
}
