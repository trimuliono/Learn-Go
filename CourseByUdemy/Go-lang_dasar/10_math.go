package main

import "fmt" 

func main() {
	// Math operations

	var a = 20
	var b = 5
	var c = 4
	var d = a + b
	var e = b - c
	var f = b * c
	var g = a / c
	var h = a % c

	fmt.Println(a, b, c, d, e, f, g, h)

	// Augmented assignment
	fmt.Println("Origin value; ", a, b, c, d, e)
	a += 10 // a = a + 10
	b -= 2 // b = b - 2
	c *= 3 // c = c * 3
	d /= 4 // d = d / 4
	e %= 2 // e = e % 2
	fmt.Println("Value Augmented assignment; ", a, b, c, d, e)

	// Unary operator
	fmt.Println("Origin value; ", a, b)

	a++ // a = a + 1
	b-- // b = b - 1

	fmt.Println("Unary operator; ", a, b)

	a++ // a = a + 1
	b-- // b = b - 1

	fmt.Println("Unary operator; ", a, b)


}