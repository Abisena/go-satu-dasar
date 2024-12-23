package program

import "fmt"

// function itu bukan hanya function as value atau function biasa tapi juga ada yang namanya function as parameter yang mana, function as parameter ini dia terletak di sebuah parameter function yang nanti kita gunakan

// disini kita buat function yang kita masukan kedalam parameter

// digolang juga ada yang namanya type declarations atau semacam penga aliasan data atau funcstion yang nantinya kita aliaskan data atau fungsi tersebut

type Filter func(string) string

func GreetWithFilter(name string, filter Filter) {
	fmt.Println("Hallo " + filter(name))
}

// kita buat fungsi filter
func NameFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
