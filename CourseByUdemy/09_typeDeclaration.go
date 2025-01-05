package main

import "fmt"

func main() {
	// Type Declaration in Go Lang is used to create a new type from an existing type
	// this is useful when we want to create a new type from an existing type
	// so that we can use the new type in a more readable way
	// and also to avoid the confusion of using the same type in multiple places
	// and prevent the misuse of the type in the code base by restricting the type to a specific type

	// example: we have age and quantity which are both int but we want to restrict the use of age and quantity
	// so we can prevent the misuse of age and quantity in the code base

	type noKTP string // create a new type noKTP from string
	type Age byte // create a new type Age from byte/uint8
	type Quantity float32 // create a new type quantity from float32
	type BMI float64 // create a new type Weight from float64

	var ktpTri noKTP = "3333333333333333" // assign value to noKTP
	var contoh string = "1111111111111111"
	var contohKTP noKTP = noKTP(contoh) // convert string to noKTP

	var age Age = 25 // assign value to Age
	var quantity Quantity = 10.0 // assign value to Quantity
	var weight BMI = 65.0 // assign value to Weight
	var height BMI = 157.0 // assign value to Height
	var bmi = weight / height * 100// calculate BMI
	// var bmi = weight / age {Error: invalid operation: weight / age (mismatched types BMI and Age)}
	// var bmi = weight / quantity {Error: invalid operation: weight / quantity (mismatched types BMI and Quantity)}

	fmt.Println(ktpTri)
	fmt.Println(contohKTP)

	fmt.Println(age)
	fmt.Println(weight)
	fmt.Println(height)
	fmt.Println(quantity)

	fmt.Println(bmi)
}