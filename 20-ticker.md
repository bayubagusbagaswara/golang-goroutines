# time.Ticker

- Ticker adalah representasi kejadian yang berulang
- Ketika `waktu ticker sudah expired`, maka `event akan dikirim` kedalam channel
- Untuk membuat ticker, kita bisa menggunakan `time.NewTicker(duration)`
- Untuk menghentikan ticker, kisa bisa menggunakan `Ticker.Stop()`

## time.Tick()

- Kadang kita tidak butuh data Ticker nya, kita `hanya butuh channel nya saja`
- Jika demikian, kita bisa menggunakan function timer.Tick(duration), function ini tidak akan mengembalikan Ticker, hanya mengembalikan channel timer nya saja
