package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// buat function, parameter group dan tipe datanya WaitGroup
func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done() // menambahkan tapi -1 (artinya mengurangi)

	// add untuk menmabahkan 1 proses asynchronous
	fmt.Println("Hello")
	// lalu sleep 1 second
	time.Sleep(1 * time.Second)
}

// testng WaitGroup, kita running function RunAsynchronous menggunakan goroutine sebanyak 100, setelah selesai nanti kita wait
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		// running goroutine
		go RunAsynchronous(group)
	}

	// tungguin sampai semua goroutine selesai
	group.Wait()
	fmt.Print("Selesai semua goroutine")

}
