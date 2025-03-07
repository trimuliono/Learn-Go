package main

import "fmt" 

func main() {
	// Comparison Operators

	var fistName = "Tri"
	var lastName = "Muliono"
	var king = "MULIONO"
	var a = 10
	var b = 5
	var c string = "10"
	var result bool = c == "10"


	fmt.Println("1. ", fistName == lastName) // false
	fmt.Println("2. ", fistName != lastName) // true
	fmt.Println("3. ",lastName == king) // false
	fmt.Println("4. ",a > b) // true
	fmt.Println("5. ",a < b) // false
	fmt.Println("6. ",a >= b) // true
	fmt.Println("7. ", a <= b) // false
	fmt.Println("8. ", a != b) // true

	fmt.Println(result) // true
}