package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("John")
	if result != "Hello John" {
		t.Error("Expected Hello John, got ", result)
	}
}

func TestHelloWolrdAssert(t *testing.T) {
	result := HelloWorld("John")
	assert.Equal(t, "Hello John", result, "Result must be Hello John")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("John")
	require.Equal(t, "Hello John", result, "Result must be Hello John")
}
