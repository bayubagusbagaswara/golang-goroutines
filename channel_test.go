package golanggoroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// kita akan membuat channel
func TestCreateChannel(t *testing.T) {

	// gunakan function make(), dan tipe data channelnya adalah string
	channel := make(chan string)

	// misal kita bikin go routine dengan anonymous function
	// sleep dulu 2 detik, lalu kirim sebuah data ke channel
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Bayu Bagus Bagaswara"
		fmt.Println("Selesai mengirim data ke channel")
	}()
	// INGAT! PASTIKAN JIKA MENGGUNAKAN CHANNEL ITU, JIKA MENGIRIM DATA MAKA HARUS ADA YANG MENERIMA DATA
	// kemudian kita tunggu datanya, dan masukkan data dari channel ke variable data
	data := <-channel
	// lalu print, apakah data yang diterima itu sama dengan data yang dikirim
	fmt.Println(data)

	// tunggu sampai go routine berhasile mengirim data ke channel
	time.Sleep(5 * time.Second)

	// setelah menggunakan channel, maka jangan lupa untuk menutupnya
	close(channel)

}

// kita bikin function dimana parameternya adalah channel
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	// kirimkan data ke channel
	channel <- "Bayu Bagus Bagaswara"

}

// kita buat test untuk channel
func TestChannelAsParameter(t *testing.T) {

	// buat channel
	channel := make(chan string)
	defer close(channel)

	// kita kirimkan channelnya ke function GiveMeResponse
	go GiveMeResponse(channel)

	// kita terima data dari channel yang dihasilkan GiveMeResponse
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// kita buat function yang khusus untuk mengirim data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Bayu Bagus Bagaswara"
}

// kita buat function yang khusus untuk menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

// kita buat testing untuk In dan Out channel
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	// buat go routine OnlyIn dengan parameter channel
	go OnlyIn(channel)
	// buat go routine OnlyOut dengan parameter channel
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// kita buat function buffered channel
func TestBufferedChannel(t *testing.T) {

	channel := make(chan string, 3) // kita tentukan kapasitas antrian untuk buffered channel

	defer close(channel)

	// kita kirimkan data ke channel, kita bisa mengirim 3 data ke dalam buffer
	channel <- "Bayu"
	channel <- "Bagus"
	channel <- "Bagaswara"

	// kita coba ambil semua data dari channel
	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("Selesai")
}

// kita buat channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// kita kirim data ke channel dari goroutine
	go func() {
		// misal kita mengirim data ke channel sebanyak 10 kali
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		// setelah 10 kali, maka kita close channel nya (agar perulangannya juga berhenti)
		close(channel)
	}()

	// ambil data dari channel, menggunakan range, karena kita tidak tau jumlah data yang dikirimkan ke channel, sehingga kita pake range, agar secara otomatis semua data yang dikirim ke channel akan kita ambil
	for data := range channel {
		fmt.Println("Menerima data", data)
	}
	fmt.Println("Selesai")
}

// kita bikin test channel select
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	// running goroutine
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// select ambil data dari channel1 dan channel2
	// kita lakukan perulangan, karena kita ingin mengambil semua data di channelnya, tapi harus pastikan kapan perulangan itu berhenti
	// misalnya kita bikin manual counter untuk menghentikan perulangannya
	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}

}
