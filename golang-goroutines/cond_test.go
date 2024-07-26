package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group sync.WaitGroup

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)
	cond.L.Lock()
	cond.Wait()
	fmt.Println("WaitCondition ", value)
	cond.L.Unlock()
}

// Cond adalah sebuah goroutine yang berjalan secara bersamaan,
// yang memiliki sebuah kondisi yang dapat dikontrol oleh goroutine lain,
// yang memiliki kondisi ini, maka goroutine ini akan menunggu kondisi tersebut,
// ketika kondisi tersebut terpenuhi, maka goroutine ini akan menerima signal dari kondisi tersebut.
func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		time.Sleep(1 * time.Second)
	//		cond.Signal()
	//	}
	//}()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
