# Pool

- Pool adalah implementasi design pattern bernama object pool pattern
- Sederhananya, design pattern Pool in digunakan `untuk menyimpan data`, selanjutnya jika ingin menggunakan datanya, kita bisa mengambilnya dari Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali (mengembalikan) ke Poolnya lagi
- Implementasi Pool di Go-Lang ini sudah aman dari problem race condition
- Sering kali digunakan untuk memanage koneksi ke database, karena bikin koneksi di database itu lumayan mahal
- Misal di awal kita akan membuat koneksi sebanyak 10, selanjutnya 10 koneksi tersebut akan kita masukkan ke dalam Pool
- Jika kita butuh koneksi, maka cukup ambil dari Poolnya, setelah selesai menggunakan koneksi, maka kita balikkan ke poolnya
