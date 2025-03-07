package main

import "fmt"

type Address struct {
	city, province, country string
}

func changeCountryToIndonesia(address *Address) { // pointer function
	address.country = "Indonesia"
}

func main() {
	// memangggil pointer function dengan new operator
	address := new(Address)
	changeCountryToIndonesia(address)

	// memanggil pointer function dengan pointer literal struct
	var alamat Address = Address{} //ini sama dengan => alamat := Address{}
	changeCountryToIndonesia(&alamat)

	fmt.Println(address)
	fmt.Println(alamat)
}