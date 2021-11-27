# RWMutex (Read Write Mutex)

- Kadang ada kasus dimana kita ingin melakukan locking tidak hanya pada proses mengubah data, tetapi juga pada proses membaca data
- Sebenarnya kita bisa menggunakan Mutex saja. Namun masalahnya nati akan jadi rebutan antara proses membaca dan mengubah datanya
- Di Go-Lang telah disediakan struct RWMutex)(Read Write Mutex) untuk menangani hal ini. Dimana Mutex jenis ini memiliki dua lock, yakni `lock untuk Read dan lock untuk Write`
