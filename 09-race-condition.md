# Race Condition

## Masalah Dengan Goroutine

- Saat kita menggunakann goroutine, dia tidak hanya berjalan secara concurrent, tetapi bisa parallel juga, karena bisa ada beberapa thread yang berjalan secara parallel
- Hal ini sangat berbahaya ketika kita melakukan maniplasi data variable yang sama `(sharing variable)` oleh beberapa goroutine secara bersamaan
- Hal ini bisa menyebabkan masalah yang namanya `Race Condition`
