package golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Atomic digunakan untuk mengurangi kinerja pada proses pemrosesan data
// yang bersifat bersamaan.
// Atomic digunakan untuk memanipulasi data primitive di concurrent safe.
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Hasil ", x)
}
