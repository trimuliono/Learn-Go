package main

import "fmt"

func main() {
	names := [...]string{"Tri", "Muliono", "Ganteng", "Banget", "Gaada", "Obatnya"}

	slice1 := names[2:6]
	fmt.Println(slice1) // [Ganteng Banget Gaada Obatnya]
	fmt.Println(slice1[0]) // Ganteng
	fmt.Println(slice1[1]) // Banget

	slice2 := names[:3] // Deklarasi slice tipe data string secara implisit karena array reference berupa string
	fmt.Println(slice2) // [Tri Muliono Ganteng]

	var slice3 []string = names[3:] // Deklarasi slice tipe data string secara eksplisit
	fmt.Println(slice3) // [Banget Gaada Obatnya]

	var slice4 []string = names[:] // Deklarasi slice tipe data string secara eksplisit
	fmt.Println(slice4) // [Tri Muliono Ganteng Banget Gaada Obatnya]

}