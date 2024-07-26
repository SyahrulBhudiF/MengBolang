package golang_goroutines

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

// Race Condition adalah sebuah kondisi yang terjadi ketika ada
// beberapa goroutine yang berjalan secara bersamaan, dan salah satu
// dari goroutine itu menyebabkan perubahan pada variabel global yang
// digunakan oleh goroutine lain.
func TestRaceCondition(t *testing.T) {
	x := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasil " + strconv.Itoa(x))
}

// Solusi untuk Race Condition adalah dengan menggunakan lock yang disebut mutex.
// Lock ini digunakan untuk memastikan bahwa hanya satu goroutine yang berjalan
// secara bersamaan pada waktu yang sama.
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex // lock
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasil " + strconv.Itoa(x))
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

// RWMutex digunakan untuk memastikan bahwa hanya satu goroutine yang berjalan
// secara bersamaan pada waktu yang sama.
// Pada contoh di atas, goroutine akan menambahkan nilai 1 ke akun bank
// dan mencetak nilai akun bank.
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Hasil " + strconv.Itoa(account.GetBalance()))
}
