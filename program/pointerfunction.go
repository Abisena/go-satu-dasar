package program

type Alamats struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(alamat *Alamats) {
	alamat.Country = "Indonesia"
}
