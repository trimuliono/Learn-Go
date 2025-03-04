package main 

import "fmt"

func main() {
	name := "John"

	switch name {
	case "John":
		fmt.Println("Hello John")
	case "Doe":
		fmt.Println("Hello Doe")
	default:
		fmt.Println("Hello Stranger")
	}

	switch Length := len(name); Length > 5 {
	case true:
		fmt.Println("Long Name")
	case false:
		fmt.Println("Short Name")
	}

	name = "Doe sdmfsoffmmo"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama cukup panjang")
	case length < 5:
		fmt.Println("nama terlalu pendek")
	}
}