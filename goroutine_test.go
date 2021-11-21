package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// kita buat function
func RunHelloWorld() {
	fmt.Println("Hello World")
}

// kita buat sebuah Test
func TestHelloWorld(t *testing.T) {
	// misal kita ingin menjalankan function RunHelloWorld di dalam Goroutine (artinya kita jalankan secara Asynchronous), cukup gunakan keyword go
	// jadi disini kita running 2 process yakni Goroutine dan program bisa
	go RunHelloWorld()
	fmt.Println("Program Biasa")

	// kita tunggu sampai program diatas selesai dijalankan, dengan Sleep
	time.Sleep(1 * time.Second)
}
