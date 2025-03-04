package main

import "fmt"

// membuat sebuah function
func sayHello() {
	fmt.Println("Hello") // blok kode di dalam function
}

// tidak ada batasan dalam memanggil function
func main() {
	sayHello() // memanggil function sayHello
	sayHello() // memanggil function sayHello
	sayHello() // memanggil function sayHello
}