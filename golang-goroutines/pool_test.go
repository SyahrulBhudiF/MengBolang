package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool adalah sebuah struktur yang digunakan untuk menyimpan data yang
// berbeda-beda. Pool adalah struktur yang menggunakan pooling untuk
// menyimpan data.
func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "Hello World"
		},
	}

	pool.Put("Syahrul")
	pool.Put("Bhudi")
	pool.Put("Ferdiansyah")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("pool put end")
}
