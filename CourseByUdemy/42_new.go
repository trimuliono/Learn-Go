package main

import "fmt"

type Address struct {
	city, province, country string
}

// operator new adalah operator yang digunakan untuk membuat pointer kosong
// new(Address) mengembalikan pointer kosong
// new(Address) == &Address{}

func main() {
	address1 := new(Address) // new operator mengembalikan pointer kosong
	address2 := address1

	address2.city = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
}