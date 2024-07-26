package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronously(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("RunAsynchronously ", i)
	time.Sleep(1 * time.Second)
}

// WaitGroup adalah sebuah struktur yang digunakan untuk mengatur jumlah
// goroutine yang sedang berjalan. Jika jumlah goroutine yang sedang berjalan
// melebihi jumlah goroutine yang diinginkan, maka goroutine yang lebih
// sedikit berjalan akan dihentikan dan di-wait.
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 10; i++ {
		go RunAsynchronously(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
}
