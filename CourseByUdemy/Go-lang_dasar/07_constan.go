package main 

import "fmt" // Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

func main() {
	const job string = "Programmer" // Constant variable (variable that can't be changed)

	// multiple constant

	const (

		firstName = "Tri"
		lastName string = "Muliono"
	)

	fmt.Println(firstName)

	fmt.Println(lastName)
}