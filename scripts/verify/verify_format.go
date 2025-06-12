package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	testDir := "./testdata/synthetic"

	// Find all .spr files
	var sprFiles []string

	filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, ".spr") {
			sprFiles = append(sprFiles, path)
		}
		return nil
	})

	fmt.Printf("Found %d SPR files to verify:\n\n", len(sprFiles))

	allValid := true

	for _, file := range sprFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("❌ ERROR reading %s: %v\n", file, err)
			allValid = false
			continue
		}

		// Split into lines
		lines := strings.Split(strings.TrimRight(string(content), "\n"), "\n")

		fmt.Printf("📄 %s:\n", filepath.Base(file))
		fmt.Printf("   Records: %d\n", len(lines))

		valid := true
		for i, line := range lines {
			length := len(line)
			if length != 850 {
				fmt.Printf("   ❌ Record %d: %d characters (expected 850)\n", i+1, length)
				valid = false
				allValid = false
			}
		}

		if valid {
			fmt.Printf("   ✅ All records exactly 850 characters\n")
		}
		fmt.Printf("\n")
	}

	if allValid {
		fmt.Println("🎉 All synthetic SPR files are properly formatted!")
	} else {
		fmt.Println("❌ Some files have formatting issues")
		os.Exit(1)
	}
}
