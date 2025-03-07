package main 

import "fmt" 

// function factorial using loop
func factorailLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

// function factorial using recursive
func factorialRecursion(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursion(value - 1)
	}
}

func main() {
	fmt.Println(factorailLoop(5))
	result := 5*4*3*2*1
	fmt.Println(result)

	fmt.Println(factorialRecursion(5))
}