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
}

// func greet(prefix string, names ...string) {
// 	for _, name := range names {
// 		fmt.Printf("%s, %s!\n", prefix, name)
// 	}
// }

// func main() {
// 	greet("Hello", "John", "Mary", "David")
// }
