package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
)

// bikin functon AddToMap, untuk nambah data ke Map
func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {

	// close group
	defer group.Done()

	// Add untuk menandakan terjadi proses goroutine (group karena disini ada proses goroutine yang mengirimkan data ke Map)
	group.Add(1)

	// simpan datanya
	data.Store(value, value)
}

// bikin test untuk Map
func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	// perulangan
	for i := 0; i < 100; i++ {
		// running goroutine untuk function AddToMap
		go AddToMap(data, i, group)
	}
	group.Wait()

	// iterasi data satu persatu
	data.Range(func(key, value interface{}) bool {

		fmt.Println(key, ":", value)
		return true
	})
}
