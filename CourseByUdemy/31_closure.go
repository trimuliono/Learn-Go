package main 

import "fmt" 
 

func main() {

	counter := 0

	// anonymous function
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	// fitur closure ini jika tidak dipakai dengan bijak akan membaut program menjadi sulit dimengerti karena tidak jelas dari mana nilai variable tersebut berubah
	// closure adalah kemampuan sebuah function untuk menyimpan referensi ke variable yang berada di luar function tersebut
	// closure adalah function yang memiliki akses ke variable di luar function tersebut
	increment()
	increment()
	increment()

	fmt.Println(counter)
}

