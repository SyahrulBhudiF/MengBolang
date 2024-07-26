package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

// Map adalah sebuah struktur yang digunakan untuk menyimpan data yang
// berbeda-beda. Map adalah struktur yang menggunakan hash table untuk
// menyimpan data.
func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go addToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
