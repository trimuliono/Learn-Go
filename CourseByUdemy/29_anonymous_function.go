package main 

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// Normal function
func blacklistName(name string) bool {
	return name == "Sialan"
}

func main() {

	// Anonymous function
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("John", blacklist)
	registerUser("Anjing", blacklist)

	registerUser("Tri", blacklistName)
	registerUser("Sialan", blacklistName)
}