package helper // nama package harus sama dengan nama folder

import "fmt"

var version = "1.0.0" // variabel ini tidak bisa diakses dari luar package
var Application = "golang"

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