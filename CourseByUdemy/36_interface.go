package main

import "fmt"

// interface adalah tipe data abstract, dia tidak memiliki implementasi langsung
// sebuah interface berisikan definisi-definisi method
// biasanya interface digunakan sebagai kontrak
// sebuah struct yang memiliki method sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
// sebuah struct boleh memiliki lebih dari 1 interface

// interface adalah tipe data, dia bisa digunakan sebagai parameter sebuah function
// interface tidak bisa langsung diimplementasikan, dia harus diimplementasikan di struct
// interface adalah reference type

type Hasname interface {
	GetName() string
}

// function sayHello menerima parameter value yang bertipe data Hasname
// Hasname adalah interface
// interface adalah tipe data, dia bisa digunakan sebagai parameter sebuah function
func sayHello(value Hasname) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

// implementasi method interface
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Tri"}
	animal := Animal{Name: "Kucing"}

	sayHello(person)
	sayHello(animal)
}