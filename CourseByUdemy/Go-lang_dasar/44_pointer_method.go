package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { // pointer method to change the value of Name
	man.Name = "Mr. " + man.Name
}

func main() {
	tri := Man{"Tri"}
	tri.Married()
	fmt.Println(tri.Name)
}