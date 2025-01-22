package main

import (
	"fmt"
	"go-satu-dasar/database"
	_ "go-satu-dasar/internal"
	"go-satu-dasar/program"
)

func main() {
	// Init
	fmt.Print(database.GetDatabase())
	// function
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
	// Function return value
	fmt.Println(program.VarName())
	fmt.Println(program.VarNamed())
	// Return Multiple Value
	var1, _ := program.Multiple()
	fmt.Println(var1)
	dataNewMap := program.NewMap("")
	if dataNewMap == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(dataNewMap["name"])
	}
	// Type Assertion
	// getData := program.Assertions()
	// result := getData.(string)
	// fmt.Println(result)

	// Type Assertion Switch
	getData2 := program.Assertions()
	switch value := getData2.(type) {
	case string:
		fmt.Println("Data String", value)
	case int:
		fmt.Println("Data Integer", value)
	default:
		fmt.Println("Data Tidak Diketahui")
	}

	// Pointer
	dataAddress := program.Address{City: "Bandung", Province: "Jawa Barat", Country: "Indonesia"}
	dataAddress2 := &dataAddress
	dataAddress2.City = "Jakarta"
	fmt.Println(dataAddress, dataAddress2)

	// Pointer Asterisk
	addressData := new(program.Alamat)
	addressData2 := addressData
	addressData2.City = "Jakarta"
	*addressData2 = program.Alamat{City: "Bogor", Province: "Jawa Barat", Country: "Indonesia"}

	fmt.Println(addressData, addressData2)

	// Pointer Function
	address := program.Alamats{City: "Bogor", Province: "Jawa Barat", Country: ""}
	program.ChangeCountryToIndonesia(&address)
	fmt.Println(address)

	// Pointer Method
	abi := program.Man{Name: "Abisena"}
	abi.Married()
	fmt.Println(abi.Name)
}

// func greet(prefix string, names ...string) {
// 	for _, name := range names {
// 		fmt.Printf("%s, %s!\n", prefix, name)
// 	}
// }

// func main() {
// 	greet("Hello", "John", "Mary", "David")
// }
