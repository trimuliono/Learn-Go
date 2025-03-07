package main

import "fmt"


type Address struct {
	city, province, country string
}

func main() {
	// pointer bekerja seperti reference di C++
	// pointer adalah variabel yang menyimpan alamat memori dari variabel lain
	// pointer digunakan untuk mengakses variabel lain
	// pointer digunakan untuk mengakses variabel lain tanpa perlu mengakses nilai variabel tersebut
	// pointer digunakan untuk mengubah nilai variabel lain

	

	// deklarasi pointer panjang
	var address3 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	var address4 *Address = &address3 // pointer

	address4.city = "Bandung"
	fmt.Println(address3) // ikut berubah
	fmt.Println(address4) // berubah menjadi Bandung

	// deklarasi pointer singkat
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // pointer ditandai dengan tanda &

	fmt.Println("Address 1:", address1)
	fmt.Println("Address 2:", *address2) // mengakses nilai pointer dengan tanda *
	address2.city = "Bandung" // mengubah nilai address1 melalui pointer address2
	fmt.Println("Address 1:", address1)
	fmt.Println("Address 2:", address2)
}