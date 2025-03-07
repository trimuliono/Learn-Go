package main 

import "fmt"

// struct adalah kumpulan dari field
// field adalah tempat untuk menyimpan data
// struct mirip seperti class di OOP
// struct adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
// struct biasanya digunakan untuk membuat tipe data baru
type Customer struct {
	Name, Address string
	Age int
}

// struct method
// struct method adalah function yang menempel pada struct
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {

	// cara 1 mennggunakan struct
	var tri Customer
	fmt.Println(tri)

	tri.Name = "Tri"
	tri.Address = "Jakarta"
	tri.Age = 26
	fmt.Println(tri)
	fmt.Println(tri.Address)
	fmt.Println(tri.Name)
	fmt.Println(tri.Age)

	// cara 2 menggunakan struct
	joko := Customer{
		Name: "Joko",
		Address: "Bandung",
		Age: 30,
	}
	fmt.Println(joko)

	// cara 3 menggunakan struct
	budi := Customer{"Budi", "Bekasi", 25}
	fmt.Println(budi)

	tri.sayHello("Budi")
	joko.sayHello("Andi")
}