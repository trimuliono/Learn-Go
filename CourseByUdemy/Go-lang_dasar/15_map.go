package main

import "fmt"

func main() {
	var person1 map[string]string = map[string]string{}
	person1["name"] = "Eko"
	person1["address"] = "Subang"

	fmt.Println(person1)

	person := map[string]string{
		"name": "tri",
		"address": "jakarta",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	book := make(map[string]string)
	book["tittle"] = "Belajar Go-Lang"
	book["author"] = "tri"
	book["ups"] = "salah"

	fmt.Println(book)

	book["ups"] = "benar" // update value

	fmt.Println(book)

	delete(book, "ups") // delete pair key-value

	fmt.Println(book)
	fmt.Println(len(book))
	fmt.Println(book["tittle"])
}