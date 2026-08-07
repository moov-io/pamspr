package pamspr_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/pamspr/pkg/pamspr"
)

func FuzzReader(f *testing.F) {
	populateCorpus(f)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		r := pamspr.NewReader(strings.NewReader(contents))
		file, err := r.Read()
		if err != nil {
			// Also try structure-only validation path
			r2 := pamspr.NewReader(strings.NewReader(contents))
			_ = r2.ValidateFileStructureOnly()
			return
		}
		if file != nil {
			for _, s := range file.Schedules {
				_ = s.Validate()
				_ = s.GetPayments()
			}
		}
	})
}

func populateCorpus(f *testing.F) {
	f.Helper()

	f.Add("")
	f.Add("H")

	_ = filepath.Walk("testdata", func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(path), ".txt") {
			bs, err := os.ReadFile(path)
			if err != nil || len(bs) > 512*1024 {
				return nil
			}
			f.Add(string(bs))
		}
		return nil
	})
}
