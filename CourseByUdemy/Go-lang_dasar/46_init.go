package main

import (
	"fmt"

	"github.com/trimuliono/Learn-Go/tree/main/CourseByUdemy/Go-lang_dasar/database"
	"github.com/trimuliono/Learn-Go/tree/main/CourseByUdemy/Go-lang_dasar/helper"

	// Blank Identifier
	// Blank Identifier digunakan untuk mengakses package yang tidak digunakan
	_ "github.com/trimuliono/Learn-Go/tree/main/CourseByUdemy/Go-lang_dasar/internal" // Blank Identifier ditandai dengan underscore ( _ )
)

func main() {
	fmt.Println(helper.SayHello("Tri"))
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error karena version tidak bisa diakses dari luar package
	fmt.Println(database.GetDatabase())
}