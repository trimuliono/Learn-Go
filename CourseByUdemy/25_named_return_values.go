package main 

import "fmt"

func getcompleteName() (firstName, middleName, lastName string) {

	firstName = "John"
	middleName = "Doe"
	lastName = "Smith"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getcompleteName()
	fmt.Println(a, b, c)
}