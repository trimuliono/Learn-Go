package main // `package main` is the main package of the program

import "fmt" // `import "fmt"` is used to import the fmt package from the standard library

func main() { // `func main()` is the entry point of the program
	fmt.Println("Tri") // `fmt.Println("Tri")` is used to print the string "Tri" to the console
	fmt.Println("Tri Muliono") // `fmt.Println("Tri Muliono")` is used to print the string "Tri Muliono" to the console
	fmt.Println("Tri Muliono Ganteng") // `fmt.Println("Tri Muliono Ganteng")` is used to print the string "Tri Muliono Ganteng" to the console

	fmt.Println(len("Tri Muliono")) // `fmt.Println(len("Tri"))` is used to print the length of the string "Tri" to the console
	fmt.Println("Tri Muliono"[0]) // `fmt.Println("Tri Muliono"[0])` is used to print the first character of the string "Tri Muliono" in ASCII to the console
	fmt.Println("Tri Muliono Ganteng"[1]) // `fmt.Println("Tri Muliono"[1])` is used to print the second character of the string "Tri Muliono" IN ASCII to the console
}