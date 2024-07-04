package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before tests")
	m.Run()
	fmt.Println("After tests")
}

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

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on linux")
	}

	result := HelloWorld("John")
	require.Equal(t, "Hello John", result, "Result must be Hello John")
}

func TestSubTest(t *testing.T) {
	t.Run("john", func(t *testing.T) {
		result := HelloWorld("John")
		require.Equal(t, "Hello John", result, "Result must be Hello John")
	})
	t.Run("doe", func(t *testing.T) {
		result := HelloWorld("Doe")
		require.Equal(t, "Hello Doe", result, "Result must be Hello Doe")
	})
}
