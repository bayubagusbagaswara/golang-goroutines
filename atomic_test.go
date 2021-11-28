package golanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// kita bikin test atomic
func TestAtomic(t *testing.T) {

	// counter
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i <= 1000; i++ {
		// running goroutine
		go func() {
			// ada proses goroutine
			group.Add(1)
			for j := 0; j <= 100; j++ {
				// ubah counternya
				// biasanya kita harus locking
				// tapi dengan atomic kita tidak perlu locking, karena number adalah data primitive
				x = x + 1
				atomic.AddInt64(&x, 1)
			}
			// proses goroutine selesai
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter = ", x)
}
