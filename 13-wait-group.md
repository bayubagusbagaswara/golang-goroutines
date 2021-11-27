# WaitGroup

- WaitGroup adalah fitur yang bisa digunakan untuk `menunggu sebuah proses selesai dilakukan atau dikerjakan`
- Hal ini kadang diperlukan, misal kita ingin menjalankan beberapa proses menggunakan goroutine, tetapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi kita selesai (mati/shutdown)
- Kasus seperti ini bisa menggunakan WaitGroup
- Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method `Add(int)`, setelah proses goroutine selesai, kita bisa gunakan method `Done()`
- Untuk menunggu semua proses selesai, kita bisa menggunakan method `Wait()`
