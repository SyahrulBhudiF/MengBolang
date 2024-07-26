package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

// Once adalah sebuah struktur yang digunakan untuk menjalankan sebuah
// fungsi hanya sekali. Jika fungsi yang dijalankan sudah dijalankan sebelumnya,
// maka tidak akan dijalankan lagi.
// Gourotine yang akan menggunakan Once akan dihentikan jika fungsi yang dijalankan
// sudah dijalankan sebelumnya.
func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println(counter)
}
