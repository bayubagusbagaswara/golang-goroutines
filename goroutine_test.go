package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Program Biasa")
	time.Sleep(1 * time.Second)
}

// kita akan mencoba membuat goroutine sebanyak 100000
func DisplayNumber(number int) {
	fmt.Println("This Display Number", number)
}

// bikin test untuk membuat banyak goroutine
func TestManyGoroutine(t *testing.T) {

	// kita buat goroutine sebanyak perulangan
	// ingat jika running di goroutine maka secara asynchronous
	// dan karena multicore, maka dia running secara concurrent dan parallel
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
