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

// buat type struct UserBalance
type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

// method untuk struct UserBalance, fungsinya untuk Lock mutex
func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

// method Unlock untuk unlock mutex
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

// buat method change, dengan parameter amount
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

// buat function Transfer dari user1 ke user2
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {

	// kita lock dulu user1
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	// change balance user1 (dikurangi maka negatif amount)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	// lock user2
	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	// change balance user2 (ditambahkan amount)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	// unlock user1
	user1.Unlock()
	// unlock user2
	user2.Unlock()
}

// buat testing deadlock
func TestDeadLock(t *testing.T) {
	// buat user1
	user1 := UserBalance{
		Name:    "Bayu",
		Balance: 1000000,
	}

	// buat user2
	user2 := UserBalance{
		Name:    "Aan",
		Balance: 1000000,
	}

	// kita coba running Transfer dengan goroutine
	go Transfer(&user1, &user2, 100000)
	// di kondisi lain, goroutine lain juga menjalankan transfer dari user2 ke user1
	go Transfer(&user2, &user1, 200000)

	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)

}
