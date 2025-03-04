package main

import "fmt"

func fullName() (string, string) {
	return "tri", "muliono"
}

func main() {
	// memanggil function multiple return value dengan 2 variable
	firstname, lastname := fullName()
	fmt.Println(firstname, lastname)

	// memanggil function multiple return value dengan mengabaikan salah satu value
	firstName, _ := fullName()
	fmt.Println(firstName)
}