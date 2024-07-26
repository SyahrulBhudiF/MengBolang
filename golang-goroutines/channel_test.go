package golang_goroutines

import (
	"fmt"
	"strconv"
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

// TestBufferedChannel is a test function that creates a channel with a buffer size of 3 and sends 3 messages to the channel
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

// TestRangeChannel is a test function that creates a channel and sends 10 messages to the channel using a for loop
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case message := <-channel1:
			fmt.Println(message + " - channel1")
			counter++
		case message := <-channel2:
			fmt.Println(message + " - channel2")
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
