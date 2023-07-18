package file_utils

import "testing"

func TestReadFileReturnsError(t *testing.T) {
	_, err := ReadFile("non-existent-file")
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func TestReadFileReturnsFileContents(t *testing.T) {
	expected := "Hello, World!"
	actual, err := ReadFile("testdata/testfile.txt")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if string(actual) != expected {
		t.Errorf("Expected %s but got %s", expected, string(actual))
	}
}
