# Sync Once

- Once adalah fitur Go-Lang yang bisa kita gunakan untuk memastikan bahwa `sebuah function di eksekusi hanya sekali`
- Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengeksekusi (mengakses)functionnya
- Goroutine yang lain akan dihiraukan, artinya function tersebut tidak akan dieksekusi lagi
