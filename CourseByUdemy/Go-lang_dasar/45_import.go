package main

import (
	"github.com/trimuliono/Learn-Go/tree/main/CourseByUdemy/helper"
	"fmt"
)
 
func main() {
	result := helper.SayHello("Tri")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
}
