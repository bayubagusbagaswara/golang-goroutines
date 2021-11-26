# Range Channel

- Kadang-kadang ada kasus sebuah channel `dikirim data secara terus menerus` oleh pengirim
- Dan kadang tidak jelas kapan channel tersebut akan berhenti menerima data
- Salah satu yang bisa kita lakukan adalah dengan menggunakan perulangan `range` ketika menerima data dari channel
- Ketika sebuah channel di `close()`, maka secara otomatis perulangan tersebut akan berhenti
- Ini lebih sederhana daripada kita melakukan pengecekan channel secara manual
