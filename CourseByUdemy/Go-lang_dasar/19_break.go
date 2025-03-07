package main 

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // digunakan untuk menghentikan semua perulangan
		}
		fmt.Println("Perulangan ke ", i)
	}
}