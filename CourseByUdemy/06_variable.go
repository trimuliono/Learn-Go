package main

import "fmt"

func main() {
	
	
	var name = "Tri Muliono" // tipe data string tidak perlu diisi karena sudah diisi value
	fmt.Println("Hello, my name is", name)

	name = "Tri Ganteng" // value bisa diubah
	fmt.Println("Hello, my name is", name)

	age := 25 // deklarasi variabel tanpa menggunakan var
	fmt.Println("I am", age, "years old")

	age = 26 // value bisa diubah
	fmt.Println("I am", age, "years old")

	var gender string // deklarasi variabel dengan tipe data karena tidak diisi value
	gender = "male" // inisialisasi variabel
	fmt.Println("I'm ", gender)

	// variabel multiple
	var (
		firstName = "Tri"
		lastName = "Muliono"
	)
	fmt.Println("My name is", firstName, lastName)

}