package golanggoroutines

import (
	"fmt"
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
