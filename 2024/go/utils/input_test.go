package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create a temporary test file
func createTempFile(content string) (string, error) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "testfile_")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	// Write content to the temporary file
	_, err = tmpFile.WriteString(content)
	if err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

// TestLazyReadLines tests the LazyReadLines function
func TestLazyReadLines(t *testing.T) {
	// Test 1: Regular file reading
	t.Run("ReadFile", func(t *testing.T) {
		content := "line 1\nline 2\nline 3\n"
		tmpFile, err := createTempFile(content)
		if err != nil {
			t.Fatal("Failed to create temp file:", err)
		}
		defer os.Remove(tmpFile) // Clean up the file after test

		// Call LazyReadLines
		linesChan := LazyReadLines(tmpFile)

		expected := []string{"line 1", "line 2", "line 3"}
		i := 0
		for line := range linesChan {
			if line.Error != nil {
				t.Error("unexpected error reading file", line.Error)
			}
			if line.Line != expected[i] {
				t.Errorf("Expected %s, got %s", expected[i], line)
			}
			i++
		}
	})

	// Test 2: Empty file
	t.Run("EmptyFile", func(t *testing.T) {
		tmpFile, err := createTempFile("")
		if err != nil {
			t.Fatal("Failed to create temp file:", err)
		}
		defer os.Remove(tmpFile)

		// Call LazyReadLines
		linesChan := LazyReadLines(tmpFile)

		// Ensure that no lines are read
		for line := range linesChan {
			t.Errorf("Expected no lines, but got %s", line)
		}
	})

	// Test 3: File with error (non-existent file)
	t.Run("NonExistentFile", func(t *testing.T) {
		// Pass a non-existent file path
		linesChan := LazyReadLines("non_existent_file.txt")

		// Ensure that the channel is closed
		select {
		case line, ok := <-linesChan:
			if ok {
				t.Errorf("Expected channel to be closed, but got line: %s", line)
			}
		default:
			// No lines, as expected
		}
	})
}

func Test_ReadLines(t *testing.T) {
	data, err := ReadLines("./testdata/input_test_1.txt")

	if err != nil {
		t.Error("failed to read input", err)
	}

	assert.Equal(t, data, []string{"123", "456", "789"})

}
