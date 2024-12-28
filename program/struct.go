package program

import "fmt"

type Costumer struct {
	Name, Address string
	Age           int
}

// struct adalah tipe data yang dapat menyimpan koleksi nilai dengan tipe yang sama, tetapi dengan menggunakan kunci (key) untuk mengakses nilai. Struct tidak memiliki urutan yang tetap
func InsertCostumer(v *Costumer) {
	fmt.Println(v.Name, v.Address, v.Age)
}
