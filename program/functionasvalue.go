package program

// function as value adalah sebuah fungsi yang mana di fungsi yang akan kita ambil dan nanti kita masukan data kedalam fungsi nya tetapi kita buat fungsi nya itu disimpan di variabel data yang nanti kita berikan value nya, contoh nya ada di file main.go, perbedaan yang mendasar dengan fungsi lain ialah, kalau kita ingin menggunakan fungsi as value kita tidak perlu memanggil tanda () jadi kita langsung ambil fungsi nya contoh jadi program.SetName saja, tidak perlu program.SetName()
func SetName(name string) string {
	return "Hallo " + name
}
