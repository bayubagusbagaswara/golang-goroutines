# Membuat Goroutine

- Hanya cukup menambahkan perintah go sebelum memanggil function yang akan kita jalankan dalam goroutine
- `go <nama_functionnya>`
- Saat sebuah function kita jalankan dalam goroutine, function tersebut akan berjalan secara asynchronous, artinya tidak akan ditunggu sampai function tersebut selesai
- Aplikasi akan lanjut berjalan ke kode program selanjutnya tanpa menunggu goroutine yang kita buat selesai
