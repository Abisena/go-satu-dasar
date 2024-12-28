package program

import "fmt"

type Pelanggan struct {
	Name, Address string
	Age           int
}

// struct method adalah fungsi yang ada di dalam struct yang nanti kita buat method nya dengan menggunakan pointer struct itu akan mempermudah kita ketika kita ingin memanggil method tersebut dari struct itu sendiri
func (pelanggan *Pelanggan) InsertPelanggan() {
	fmt.Printf("Hallo %s, alamat kamu di %s dan umur kamu %d\n", pelanggan.Name, pelanggan.Address, pelanggan.Age)
}
