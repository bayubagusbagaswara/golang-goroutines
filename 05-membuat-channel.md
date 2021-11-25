# Membuat Channel

- Channel di Go-Lang direpresentasikan dengan tipe data `chan`
- Untuk membuat channel sangat mudah, kita bisa menggunakan function `make()`, mirip ketika membuat map
- Namun saat pembuatan channel, kita harus tentukan tipe data apa yang bisa dimasukkan kedalam channel tersebut

## Mengirim dan Menerima Data dari Channel

- Untuk mengirim data, kita bisa gunakan kode : `channel <- data`
- Sedangkan untuk menerima data, bisa gunakan kode : `data <- channel`
- Jika selesai menggunakan channel, jangan lupa untuk menutup channelnya menggunakan function `close()`

## Channel Sebagai Parameter

- Dalam kenyataan pembuatan aplikasi, seringnya kita akan `mengirim channel ke function lain via parameter`
- Sebelumnya kita tahu bahkan di Go-Lang by default, `parameter adalah pass by value`, artinya value akan diduplikasi kemudian baru dikirim ke function parameter (tidak merubah value aslinya atau reference). Sehingga jika kita ingin mengirim data asli, maka kita biasa gunakan `pointer (agar pass by reference)`

## Channel In dan Out

- Saat kita mengirim channel sebagai parameter, isi function tersebut bisa mengirim dan menerima data dari channel tersebut
- Kadang kita ingin memberi tahu terhadap function, misalnya bahwa channel tersebut hanya digunakan untuk mengirim data, atau hanya dapat digunakan untuk menerima data
- Hal ini bisa kita lakukan di parameter dengan cara `menandai` apakah channel ini digunakan untuk `in (mengirim data)` atau `out (menerima data)`

## Buffered Channel

- Secara default, channel itu hanya bisa menerima 1 data
- Artinya jika kita ingin menambah data yang ke-2, maka kita akan diminta menunggu sampai data yang ke-1 ada yang mengambil terlebih dahulu
- Kadang-kadang ada kasus dimana pengirim lebih cepat dibanding penerima, dalam hal ini jika kita menggunakan channel, maka otomatis pengirim akan ikut lambat juga
- Untungnya ada Buffered Channel, yaitu `buffer yang bisa digunakan untuk menampung data antrian di Channel`

## Buffer Capacity

- Kita bebas memasukkan berapa jumlah kapasitas antrian di dalam buffer
- Jika kita set misal 5, artinya kita bisa menerima 5 data di buffer
- Jika kita mengirim data ke-6, maka kita diminta untuk menunggu sampai buffer ada yang kosong
- Ini cocok sekali ketika memang goroutine yang menerima data lebih lambat dari yang mengirim data
