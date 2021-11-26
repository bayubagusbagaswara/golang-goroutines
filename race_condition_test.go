package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// kita buat test untuk race condition
func TestRaceCondition(t *testing.T) {
	// variable x ini disharing ke beberapa goroutine secara bersamaan, jadinya terjadi Race Condition
	x := 0

	// bikin perulangan sebanyak 1000 kali untuk goroutine
	// artinya akan ada 1000 goroutine
	for i := 1; i <= 1000; i++ {
		go func() {
			// bikin perulangan untuk x, yakni tugas goroutine
			// artinya 1 goroutine akan menambahkan 100
			for j := 1; j <= 100; j++ {
				// increment x
				x = x + 1
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
