package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Timer digunakan untuk menghitung waktu yang berlalu dari waktu yang ditentukan.
func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	time := <-timer.C
	fmt.Println(time)
}

// Timer After digunakan untuk menghitung waktu yang berlalu dari waktu yang ditentukan.
func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())
	time := <-channel
	fmt.Println(time)
}

// Timer AfterFunc digunakan untuk menghitung waktu yang berlalu dari waktu yang ditentukan dan menggunakan fungsi callback.
func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Timer")
		group.Done()
	})

	group.Wait()
}

// Timer Ticker digunakan untuk menghitung waktu yang berlalu dari waktu yang ditentukan dan menggunakan fungsi callback.
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
			count++
		}
		if count == 5 {
			ticker.Stop()
			break
		}
	}
}

// Timer Tick digunakan untuk menghitung waktu yang berlalu dari waktu yang ditentukan dan menggunakan fungsi callback.
func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	count := 0
	for time := range channel {
		fmt.Println(time)
		count++
		if count == 5 {
			break
		}
	}
}
