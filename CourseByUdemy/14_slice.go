package main

import "fmt"

func main() {
	names := [...]string{"Tri", "Muliono", "Ganteng", "Banget", "Gaada", "Obatnya"}

	slice := names[2:6]

	fmt.Println(slice[0]) // Ganteng

	fmt.Println(slice[1]) // Banget

	fmt.Println(slice) // [Ganteng Banget Gaada Obatnya]
}