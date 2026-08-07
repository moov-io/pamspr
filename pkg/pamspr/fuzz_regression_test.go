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
