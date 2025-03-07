package main

import "fmt" 

// membuat function dengan parameter
func sayHello (firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

func main() {
	// memanggil function yg ada parameter harus mengirimkan semua parameter sesuai dengan yg di tentukan
	sayHello("Tri", "Muliono")
	sayHello("Budi", "Santoso")
}