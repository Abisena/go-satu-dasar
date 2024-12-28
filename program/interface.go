package program

import "fmt"

// interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
// sebuh interface berisikan definisi definisi method
// dan biasanya interface digunakan untuk membuat kontrak

type HasName interface {
	GetName() string
}

type Named struct {
	Name string
}

func (n *Named) GetName() string {
	return n.Name
}

func SayHello(h HasName) {
	fmt.Println("Hello", h.GetName())
}
