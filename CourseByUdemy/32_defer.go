package main 

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

// dalam pemograman lain kita kenal try catch, pada golang kita kenal dengan defer,panic,recover
func runApplication() {
	defer logging() // defer akan dijalankan terakhir meskipun ada error
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}