package main 

import "fmt"

// Type Assertion adalah cara untuk mengecek tipe data dari sebuah interface
// Type Assertion juga bisa digunakan untuk mengambil value dari sebuah interface	
// Type Assertion juga bisa digunakan untuk mengecek apakah sebuah interface kosong atau tidak

// interface kosong dapat ditulis sebagai interface{} atau any
func random() any {
	return true
}

func main() {
	var data interface{} = "Hello, Golang!"

    // Type assertion tanpa pengecekan (bisa panic jika salah tipe)
    str := data.(string)
    fmt.Println(str) // Output: Hello, Golang!

	var result any = random()
	
	switch value := result.(type) { // type assertion dalam switch case untuk menghindari panic disebebkan oleh type assertion yang salah
	case string:
		fmt.Println("Value is string:", value)
	case int:
		fmt.Println("Value is int:", value)
	default:
		fmt.Println("Unknown type:", value)
	}
}