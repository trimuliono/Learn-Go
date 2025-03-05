package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1  := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer
	address3 := &address1 // pointer
	address2.city = "Bandung"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address3) // ikut berubah
	fmt.Println(address2) // berubah menjadi Bandung

	address2 = &Address{"Bali", "DKI Jakarta", "Indonesia"} // merubah alamat memori pointer
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // asterisk operator digunakan untuk mengubah nilai pointer
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	*address3 = Address{"Medan", "SUMUT", "Indonesia"} // asterisk operator digunakan untuk mengubah nilai pointer
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	
}