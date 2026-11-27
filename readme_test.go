package pamspr_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var readmeGoFence = regexp.MustCompile("(?s)```go\r?\n(.*?)```")

func TestREADMEExamplesCompile(t *testing.T) {
	root, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	source, err := os.ReadFile(filepath.Join(root, "README.md"))
	if err != nil {
		t.Fatal(err)
	}

	var programs []string
	for _, match := range readmeGoFence.FindAllStringSubmatch(string(source), -1) {
		block := match[1]
		if strings.Contains(block, "package main") {
			programs = append(programs, block)
		}
	}
	if len(programs) == 0 {
		t.Fatal("no package main examples found in README.md")
	}

	for i, src := range programs {
		i, src := i, src
		t.Run(exampleName(src, i), func(t *testing.T) {
			t.Parallel()
			dir := t.TempDir()
			if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte(src), 0o644); err != nil {
				t.Fatal(err)
			}
			mod := "module readme-example\n\ngo 1.24.4\n\nrequire github.com/moov-io/pamspr v0.0.0\n\nreplace github.com/moov-io/pamspr => " + strconv.Quote(root) + "\n"
			if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte(mod), 0o644); err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command("go", "build", "-o", "example", ".")
			cmd.Dir = dir
			cmd.Env = append(os.Environ(), "GOWORK=off")
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("go build failed: %v\n%s\n----- source -----\n%s", err, out, src)
			}
		})
	}
}

func exampleName(src string, i int) string {
	switch {
	case strings.Contains(src, "StartACHSchedule"):
		return "create-ach"
	case strings.Contains(src, "StartCheckSchedule"):
		return "create-check"
	case strings.Contains(src, "ValidateFileStructure"):
		return "validate"
	case strings.Contains(src, "NewReader"):
		return "read"
	default:
		return "example-" + strconv.Itoa(i)
	}
}
