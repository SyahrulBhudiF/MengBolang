package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// RunHelloWorld is a function that prints "Hello World" to the console
func RunHelloWorld() {
	fmt.Println("Hello World")
}

// TestCreateGoroutines is a test function that creates a goroutine and prints "Ups" to the console
func TestCreateGoroutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

// TestManyGoroutine is a test function that creates 100000 goroutines and prints the number of each goroutine, Goroutines are so lightweight that they can be created in large numbers
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(1 * time.Second)
}
