package golanggoroutines

import (
	"fmt"
	"testing"
	"time"
)

// bikin test ticker
func TestTicker(t *testing.T) {

	ticker := time.NewTicker(1 * time.Second)

	// kita akan stop tickernya menggunakan goroutine
	go func() {
		time.Sleep(5 * time.Second)
		// stop ticker
		ticker.Stop()
	}()

	// tiap detik kita akan menerima event dari channelnya
	// ini akan berjalan terus
	for time := range ticker.C {
		fmt.Println(time)
	}
}

// test Tick, artinya kita hanya mengembalikan channelnya saja, tidak dengan Tickernya

func TestTick(t *testing.T) {

	// ini balikannya adalah channel
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}
