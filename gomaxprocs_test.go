package golanggoroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// bikin test
func TestGetGomaxprocs(t *testing.T) {

	// kita coba running goroutine untuk mengetahui jumlah goroutine yang dirunning oleh golang saat ini
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	// ambil jumlahCpu di laptop kita
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// mengetahui jumalh thread di Golang
	// -1 artinya kita tidak mengubah jumlah threadnya
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	// tunggu proses goroutine selesai semua
	group.Wait()
}

// test menambah thread
func TestChangeThreadNumber(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// ubah jumlah thread
	// misal kita ubah dari 4 menjadi 10
	runtime.GOMAXPROCS(10)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}
