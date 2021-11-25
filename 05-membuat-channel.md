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
