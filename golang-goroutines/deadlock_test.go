package golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance += amount
}

func Transfer(User1 *UserBalance, User2 *UserBalance, amount int) {
	User1.Lock()
	fmt.Println("lock User1", User1.Name)
	User1.Change(amount)

	time.Sleep(1 * time.Second)

	User2.Lock()
	fmt.Println("lock User2", User2.Name)
	User2.Change(amount)

	time.Sleep(1 * time.Second)

	User1.Unlock()
	User2.Unlock()
}

// Deadlock adalah ketika 2 thread bergantung pada 1 thread yang sedang
// menggunakan resource yang lain (lock) dan tidak bisa mengakses resource
// yang lain (unlock) karena resource yang lain sedang di-lock.
// Dibawah ini adalah contoh program yang menyebabkan deadlock.
func TestDeadlock(t *testing.T) {
	user1 := &UserBalance{Name: "Budi", Balance: 1000000}
	user2 := &UserBalance{Name: "Ferdi", Balance: 1000000}

	go Transfer(user1, user2, 100000)
	go Transfer(user2, user1, 10000)

	time.Sleep(10 * time.Second)
	fmt.Println("amount User1", user1.Balance)
	fmt.Println("amount User2", user2.Balance)
}
