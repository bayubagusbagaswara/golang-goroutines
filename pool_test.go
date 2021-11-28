package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// kita bikin test pool
func TestPool(t *testing.T) {
	pool := sync.Pool{

		// bikin default value, jika poolnya dalam keadaan kosong datanya
		New: func() interface{} {
			return "New"
		},
	}

	// masukkan data ke pool
	pool.Put("Bayu")
	pool.Put("Bagus")
	pool.Put("Bagaswara")

	// running menggunakan goroutine sebanyak 10 goroutine
	for i := 0; i < 10; i++ {
		go func() {
			// ambil datanya dari pool
			data := pool.Get()

			fmt.Println(data)
			// ketika ada sleep, maka ada kemungkinan goroutine yang lain tidak kebagian data
			time.Sleep(1 * time.Second)

			// setelah selesai, kemalikan data ke poolnya
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}
