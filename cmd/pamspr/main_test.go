package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestConvertToJSONNotImplemented(t *testing.T) {
	dir := t.TempDir()
	input := filepath.Join(dir, "in.spr")
	output := filepath.Join(dir, "out.json")
	if err := os.WriteFile(input, []byte("placeholder"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	err := convertToJSON(input, output)
	if err == nil {
		t.Fatal("convertToJSON() error = nil, want not implemented")
	}
	if !strings.Contains(err.Error(), "not implemented") {
		t.Errorf("convertToJSON() error = %q, want it to contain %q", err.Error(), "not implemented")
	}
	if _, statErr := os.Stat(output); !os.IsNotExist(statErr) {
		t.Errorf("convertToJSON() wrote %s, want no output file", output)
	}
}
