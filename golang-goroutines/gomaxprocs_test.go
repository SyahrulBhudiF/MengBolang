package golang_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

// Gomaxprocs digunakan untuk menghitung jumlah CPU yang tersedia
// di sistem operasi.
func TestPrintCpu(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			fmt.Println("CPU: ", runtime.NumCPU())
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU: ", totalCpu)

	totalThread := runtime.GOMAXPROCS(0)
	fmt.Println("Total thread: ", totalThread)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total goroutines: ", totalGoroutines)

	group.Wait()
}
