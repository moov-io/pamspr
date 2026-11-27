package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/pamspr/pkg/pamspr"
)

func TestExamplesRun(t *testing.T) {
	dir := t.TempDir()
	outputDir = dir

	CreateACHFile()
	achPath := outputPath("ach_payments.txt")
	if _, err := os.Stat(achPath); err != nil {
		t.Fatalf("ACH example did not write %s: %v", achPath, err)
	}
	ParseAndValidateFile(achPath)

	CreateCheckFile()
	if _, err := os.Stat(outputPath("check_payments.txt")); err != nil {
		t.Fatalf("check example did not write output: %v", err)
	}

	// Capture Same Day ACH example output
	// ValidateSameDayACH prints to stdout; just run it and require no panic.
	ValidateSameDayACH()

	// The example files must stay in the temp directory, not the repo root.
	if _, err := os.Stat(filepath.Join("..", "ach_payments.txt")); err == nil {
		t.Error("example wrote ach_payments.txt to the repo root")
	}
}

func TestBuildSameDayACHFile(t *testing.T) {
	valid, err := buildSameDayACHFile(50000)
	if err != nil {
		t.Fatalf("buildSameDayACHFile(50000) error = %v", err)
	}
	if err := pamspr.NewValidator().ValidateFileStructure(valid); err != nil {
		t.Fatalf("valid file structure error = %v", err)
	}
	if err := pamspr.NewValidator().ValidateSameDayACH(valid); err != nil {
		t.Fatalf("valid Same Day ACH error = %v", err)
	}

	invalid, err := buildSameDayACHFile(200000000)
	if err != nil {
		t.Fatalf("buildSameDayACHFile(200000000) error = %v", err)
	}
	err = pamspr.NewValidator().ValidateSameDayACH(invalid)
	if err == nil {
		t.Fatal("ValidateSameDayACH() on $2,000,000 payment = nil, want sda_max_amount")
	}
	if !strings.Contains(err.Error(), "sda_max_amount") && !strings.Contains(err.Error(), "Same Day ACH limit") {
		t.Errorf("ValidateSameDayACH() error = %q, want Same Day ACH limit", err.Error())
	}
}
