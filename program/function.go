package program

import "fmt"

// Dan ingat, jika di return value datanya membalikan 2 type data, maka perlua dua var untuk menangkap hasil data return nya

// function return value, gunanya membalikan nilai dengan return type datanya
func VarName() (string, string) {
	return "Muhamad	Abisena", "Putrawan"
}

// function named return value, ini membuat variable data dengan type data return nya
func VarNamed() (data1, data2 string) {
	data1 = "Muhamad Abisena"
	data2 = "Putrawan"
	return data1, data2
}

// function Variadic adalah function yang memiliki parameter dengan tipe data ...string, dan juga memiki kemampuan untuk menampung data lebih dari satu inputan type data
func VariadicString(data ...string) string {
	return data[0]
}

func VariadicInt(data ...int) int {
	r := 0

	for _, v := range data {
		r += v
	}

	return r
}

func Result() {
	_, dua := VarName()
	data, _ := VarNamed()
	r := VariadicString("Muhamad", "Abisena", "Putrawan")
	v := VariadicInt(1, 2, 3, 4, 5)
	// tanda garis bawah untuk menghiraukan data atau ignore
	fmt.Println(data, dua, r, v)
}
