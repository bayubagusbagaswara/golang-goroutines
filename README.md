# Golang Gouroutines

## Agenda

- Concurrency & Parallel Programming
- Goroutines
- Channel
- Buffered Channel
- Mutex
- WaitGroup
- Atomic
- Ticker
- dan lain-lain

## Pengenalan Parallel Programming

- Dimana sekarang kita bisa dengan mudah membuat proses parallel di aplikasi
- Parallel programming sederhananya adalah memecahkan suatu masalah dengan cara membaginya menjadi yang lebih kecil. Dan dijalankan secara `bersamaan pada waktu yang bersamaan pula`

### Contoh Parallel

- Menjalankan beberapa aplikasi sekaligus di sistem operasi kita (office, editor, browser, dan lain-lain)
- Beberapa koki menyiapkan makanan di restoran, dimana tiap koki membuat makanannya masing-masing
- Antrian di Bank, dimana tiap teller melayani nasabahnya masing-masing

### Process vs Thread

- Process adalah sebuah eksekusi program
- Process mengkonsumsi memory yang besar
- Process saling terisolasi dengan process lain
- Process lama untuk dijalankan dihentikan
- Thread adalah segmen dari process
- Thread menggunakan memory kecil
- Thread bisa saling berhubungan jika dalam process yang sama
- Thread cepat untuk dijalankan dan dihentikan

## Parallel vs Concurrency

- Berbeda dengan parallel (menjalankan beberapa pekerjaan secara bersamaan)
- Concurrency adalah menjalankan beberapa pekerjaan secara bergantian
- Dalam parallel kita biasanya membutuhkan banyak Thread, sedangkan dalam concurrency kita hanya membutuhkan sedikit Thread

### Contoh Concurrency

- Saat kita makan di cafe, kita bisa makan, lalu ngobrol, lalu minum, makan lagi, ngobrol lagi, minum lagi, dan seterusnya
- Tetapi kita tidak bisa pada saat yang bersamaan minum, makan dan ngobrol, hanya bisa melakukan satu hal pada satu waktu. Namun, bisa bergantian kapanpun kita mau
