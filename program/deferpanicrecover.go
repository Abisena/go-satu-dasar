package program

import (
	"fmt"
)

func Logging() {
	fmt.Println("Aplikasi sudah selesai berjalan")

	// recover adalah fungsi bawaan dari golang ketika terjadi panic di dalam program kita akan langsung dihentikan dan dianggap error yang tidak terdeteksi dari program biasanya akan disimpan didalam function terpisah antara panic dan recover
	message := recover()
	fmt.Println(message)
}

func RunApp(error bool) {
	// defer adalah fungsi bawaan dari golang ketika eksekusi program dianggap selesai
	defer Logging()
	if error {
		// panic adalah fungsi bawaan dari golang ketika terjadi error di dalam program kita akan langsung dihentikan
		panic("ERROR")
	}
	fmt.Println("Aplikasi sedang berjalan")
}
