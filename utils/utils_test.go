package utils

import (
	"os"
	"testing"
)

func TestCheck(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Check() did not panic on error")
		}
	}()

	CheckErr(os.ErrInvalid)
}

func TestReadLines(t *testing.T) {
	content := "line1\nline2\nline3\n"
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	lines, err := ReadLines(tmpfile.Name())
	if err != nil {
		t.Fatalf("ReadLines failed: %v", err)
	}

	expected := []string{"line1", "line2", "line3"}
	for i, line := range lines {
		if line != expected[i] {
			t.Errorf("Expected %q, got %q", expected[i], line)
		}
	}
}
