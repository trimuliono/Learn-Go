package helper // nama package harus sama dengan nama folder

import "fmt"

// Access Modifier
// Accesss modifier di golang diatur dengan cara penulisan nama variabel atau function
// jika diawali huruf besar maka bisa diakses dari luar package
// jika diawali huruf kecil maka hanya bisa diakses dari dalam package

var version = "1.0.0" // variabel ini tidak bisa diakses dari luar package
var Application = "golang" // variabel ini bisa diakses dari luar package

func 
SayGoodBye(name string) string{
	return "Good bye " + name
}

func SayHello(name string) string{
	return "Hello " + name
}

func contoh() {
	
	SayGoodBye("Tri")
	fmt.Println(version)
}