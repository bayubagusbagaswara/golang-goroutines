package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// kita buat test mutex
func TestMutex(t *testing.T) {

	x := 0

	// bikin variable mutex nya
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				// disini akan ada 999 goroutine yang akan menunggu, dan hanya boleh satu persatu yang mengakses atau mengubah value
				// karena berdasarkan cara lock dan unlock
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

// buat struct BankAccount
type BankAccount struct {
	// bikin variable RWMutex
	RWMutex sync.RWMutex
	Balance int
}

// buat method struct AddBalance
func (account *BankAccount) AddBalance(amount int) {
	// kita lock dulu untuk Write
	account.RWMutex.Lock()
	// ubah data balance
	account.Balance = account.Balance + amount
	// unlock untuk Write
	account.RWMutex.Unlock()
}

// buat method struct GetBalance
func (account *BankAccount) GetBalance() int {
	// lock untuk Read
	account.RWMutex.RLock()
	// baca data balance di account
	balance := account.Balance
	// unlock untuk Read
	account.RWMutex.RUnlock()
	return balance
}

// buat test untuk RWMutex
func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		// running goroutine
		go func() {
			for j := 0; j < 100; j++ {
				// harusnya totalnya 10000
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}
