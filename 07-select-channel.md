# Select Channel

- Kadang ada kasus dimana kita membuat beberapa channel dan menjalankan beberapa goroutine
- Lalu kita ingin mendapatkan data dari semua channel tersebut
- Untuk melakukan hal tersebut, kita bisa menggunakan select channel di Go-Lang
- Dengan select channel, kita bisa memilih data tercepat dari beberapa channel. Jika data datang secara bersamaan di beberapa channel, maka akan dipilih secara random (data mana yang akan diambil)
