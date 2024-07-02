package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("John")
	if result != "Hello John" {
		t.Error("Expected Hello John, got ", result)
	}
}
