package program

import "fmt"

type Blacklist func(string) bool

func RegisterUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Kamu di Blokir", name)
	} else {
		fmt.Println("Selamat Datang", name)
	}
}
