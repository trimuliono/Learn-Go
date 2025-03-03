package main 

import "fmt"

func main() {
	name := "tri"

	if name == "Muliono" {
		fmt.Println("Hello Muliono")
	} else if name == "Tri" {
		fmt.Println("Hello Tri")
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	if Length := len(name); Length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}