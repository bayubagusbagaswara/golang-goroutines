package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

// bikin counter
var counter = 0

// bikin function OnlyOnce
func OnlyOnce() {
	counter++
}

// bikin test untuk once
func TestOnce(t *testing.T) {
	// bikin once nya
	once := sync.Once{}
	// bikin wait group
	group := sync.WaitGroup{}

	// bikin perulangan
	for i := 0; i < 100; i++ {
		// runnung goroutine
		go func() {
			// tambahkan groupnya 1 per 1
			group.Add(1)
			once.Do(OnlyOnce) // masukan function yang tidak ada parameternya

			group.Done()
		}()
	}
	group.Wait()
	// print counternya ada berapa
	fmt.Println("Counter", counter)
}
