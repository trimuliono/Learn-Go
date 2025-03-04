package main

import "fmt"

// Function with return value
func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	// method 1 to call function with return value and store it in variable and then print it 
	result := getHello("World")
	fmt.Println(result)

	// Call function with return value
	fmt.Println(getHello("John"))
	fmt.Println(getHello("Doe"))

}