package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {

	// bikin timernya
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	// kirim ke channel
	time := <-timer.C

	// ambil datanya
	fmt.Println(time)
}
func TestAftre(t *testing.T) {

	// bikin timernya, balikannya langsung berupa channel
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(1)

	/// masukkan function yang akan dieksekusi setelah 5 detik
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
	})
	fmt.Println(time.Now())

	group.Wait()
}
