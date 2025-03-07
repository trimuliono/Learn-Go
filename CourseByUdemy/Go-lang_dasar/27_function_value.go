package main

import "fmt"

func goodbye(name string) string {
	return "Goodbye " + name
}

func main() {

	// Function value adalah function yang bisa dijadikan sebagai value
	// ini bisa digunakan sebagai parameter atau dijadikan sebagai value dari variabel lain
	contoh := goodbye // memasukkan function goodbye ke dalam variabel contoh
	misal := goodbye

	fmt.Println(contoh("Tri"))
	fmt.Println(misal("Muliono"))
}