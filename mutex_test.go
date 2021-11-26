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
