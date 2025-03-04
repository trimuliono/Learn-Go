package main 

import "fmt"

// Variadic function is a function that can accept any number of arguments
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {

		total += number
	}
	return total
}

func main() {
	// Passing multiple arguments to a variadic function
	total := sumAll(10, 10, 10, 10)
	fmt.Println(total)

	//
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10, 10))

	// Passing a slice to a variadic function
	numbers := []int{10, 10, 10}
	fmt.Println(sumAll(numbers...))
}