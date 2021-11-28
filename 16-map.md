# Map

- Go-Lang memiliki sebuah struct bernama sync.Map
- Map ini mirip Go-LanG map. Namun yang membedakan adalah Map ini aman untuk penggunaan Concurrent (dimana menggunakan goroutine)
- Ada beberapa function yang bisa kita gunakan di Map :
  - Store(key, value) untuk menyimpan data ke Map
  - Load(key) untuk mengambil data dari Map menggunakan key
  - Delete(key) untuk menghapus data di Map menggunakan key
  - Range(function(key, value)) digunakan untuk melakukan iterasi seluruh data di Map
