package main

import (
	"fmt"
	"go-satu-dasar/program"
)

func main() {
	program.Result()
	program.Slice()
	program.Array()
	program.Map()
	// function as value
	data := program.SetName
	fmt.Println(data("Muhamad Abisena Putrawan"))
	// function as parameter
	program.GreetWithFilter("Muhamad Abisena", program.NameFilter)
	program.GreetWithFilter("Anjing", program.NameFilter)
	// function anonymus
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	program.RegisterUser("Jauhara", blacklist)
	// defer panic recover, ketika tidak true makan akan menghasilkan nil padasaat di recover
	program.RunApp(false)
	// struct menggunakan pointer
	program.InsertCostumer(&program.Costumer{
		Name:    "Muhamad Abisena",
		Address: "Bojong",
		Age:     18,
	})
	// struct method menggunakan pointer
	pelanggan := &program.Pelanggan{
		Name:    "Muhamad Abisena",
		Address: "Bojong",
		Age:     18,
	}
	pelanggan.InsertPelanggan()

	named := &program.Named{Name: "Abisena"}
	program.SayHello(named)
}

// func greet(prefix string, names ...string) {
// 	for _, name := range names {
// 		fmt.Printf("%s, %s!\n", prefix, name)
// 	}
// }

// func main() {
// 	greet("Hello", "John", "Mary", "David")
// }
