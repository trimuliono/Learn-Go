package main

import "fmt"

// dalam pemograman lain kita kenal try catch, pada golang kita kenal dengan defer,panic,recover
// defer akan dijalankan terakhir meskipun ada error
// panic akan menghentikan program
// recover akan menangkap panic
// cara yang benar untuk menjalankan recover adalah dengan menggunakan defer
// recover hanya bisa digunakan didalam function yang menggunakan defer, jika tidak maka recover tidak akan berfungsi karena panic akan menghentikan program

func endApp() {
	fmt.Println("End App")
	message := recover() // recover berada di dalam function yang dipanggil oleh defer
	fmt.Println("Terjadi Panic", message)
}

func runApp(error bool) {
	defer endApp() // defer akan dijalankan terakhir meskipun ada error
	if error {
		panic("Ups Error")
	}
}

func main() {
	runApp(true)
	fmt.Println("Tri Muliono")
}