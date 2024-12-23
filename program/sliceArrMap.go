package program

import "fmt"

// Array

// Array adalah tipe data yang dapat menyimpan koleksi nilai dengan tipe yang sama. Array memiliki ukuran yang tetap dan tidak dapat diubah setelah dibuat. Contoh:

func Array() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr[3]) // 1
}

// Slice

// Slice adalah tipe data yang dapat menyimpan koleksi nilai dengan tipe yang sama, tetapi tidak memiliki ukuran yang tetap. Slice dapat diubah ukurannya setelah dibuat. Slice adalah referensi ke array yang mendasarinya. Contoh:
func Slice() {
	slice := make([]int, 5, 10)
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	fmt.Println(len(slice)) // 11
	fmt.Println(cap(slice)) // 20
}

// Perbedaan utama antara array dan slice adalah:

// Array memiliki ukuran yang tetap, sedangkan slice tidak.
// Array dapat diakses secara langsung, sedangkan slice harus diakses melalui indeks.

// Map

// Map adalah tipe data yang dapat menyimpan koleksi nilai dengan tipe yang sama, tetapi dengan menggunakan kunci (key) untuk mengakses nilai. Map tidak memiliki urutan yang tetap. Contoh:

var m map[string]int

func Map() {
	m = make(map[string]int)
	m["satu"] = 1
	m["dua"] = 2
	m["tiga"] = 3
	m["empat"] = 4
	m["lima"] = 5
	m["enam"] = 6

	fmt.Println(m["enam"]) // 1
}

// Perbedaan utama antara array, slice, dan map adalah:

// Array dan slice memiliki urutan yang tetap, sedangkan map tidak.
// Array dan slice dapat diakses secara langsung, sedangkan map harus diakses melalui kunci.
