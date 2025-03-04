package main 

import "fmt"

// Function filter as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

// func as type
type Filter func(string) string

func sayHiWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
		if  name == "Anjing" {
			return "..."
			
		} else {
			return name
		}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("John", spamFilter)
	sayHiWithFilter("Anjing", spamFilter)
	sayHiWithFilter("Doe", spamFilter)
}