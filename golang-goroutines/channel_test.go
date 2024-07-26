package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

// Channel is a way to communicate between goroutines, it is a pipe that connects goroutines, goroutines can send messages to the channel and receive messages from the channel

// TestCreateChannel is a test function that creates a channel and sends a message to the channel
func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
	}()

	message := <-channel
	fmt.Println(message)
}

// GiveMeResponse is a function that sends a message to the channel and have parameter channel
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

// TestChannelAsParameter is a test function that creates a channel and sends a message to the channel using a function that has a channel as a parameter
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	message := <-channel
	fmt.Println(message)
}

// OnlyIn is a function that sends a message to the channel and channel is a write-only channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello World"
}

// OnlyOut is a function that receives a message from the channel and channel is a read-only channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// TestInOutChannel is a test function that creates a channel and sends a message to the channel using a function that has a write-only channel as a parameter and receives a message from the channel using a function that has a read-only channel as a parameter
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Hello World"
		channel <- "Hello World"
		channel <- "Hello World"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}
