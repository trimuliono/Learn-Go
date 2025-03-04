package main 

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue // digunakan untuk skip dan melanjutkan ke perulangan berikutnya
		}
		fmt.Println("Perulangan ke ", i)
	}
}