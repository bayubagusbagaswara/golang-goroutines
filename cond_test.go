package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// bikin variable cond dan masukkan lockernya
var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

// bikin function WaitCondition
func WaitCondition(value int) {

	defer group.Done()
	group.Add(1)

	// pertama kita lakukan lock kondisinya
	cond.L.Lock()

	// kita nanya dulu ke cond, apakah dia boleh jalan atau tidak setelah nglakukin locking
	// Wait artinya setelah saya melakukan locking apakah saya boleh lanjut atau tidak
	// Untuk mengijinkan gorouitne melewati Waitnya, maka kita harus kirim Signal di goroutine

	cond.Wait()

	fmt.Println("Done", value)

	// unlock kondisinya
	cond.L.Unlock()
}

// bikin testing
func TestCond(t *testing.T) {

	for i := 0; i < 10; i++ {
		// jalankan 10 goroutine, dan panggil function WaitCondition
		go WaitCondition(i)
	}

	go func() {
		// perulangan untuk tiap goroutine akan mengirimkan signal
		for i := 0; i < 10; i++ {
			// sleep 1 detik
			time.Sleep(1 * time.Second)
			// tiap 1 detik akan mengirimkan signal ke condition
			// dan hanya 1 goroutine untuk 1 Signal
			cond.Signal()
		}
	}()

	// jika tidak ingin satu persat mengirim Signal, maka bisa menggunakan broadcast
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	// tunggu semua goroutine selesai
	group.Wait()
}
