# Cond

- Cond adalah implementasi locking berbasis kondisi
- Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi lockingnya. Namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk menunggu apakah perlu menuggu atau tidak
- Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
- Untuk membuat Cond, kita bsia menggunakan function sync.NewCond(Locker)
