package main

import (
	"fmt"
	"errors" // package bawaan go untuk error handling
)


func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		// error merupakan sebuah interface, jadi bisa diisi dengan apapun
		// biasanya diisi dengan string
		// errors.New() merupakan function untuk membuat error
		// errors.New() membutuhkan parameter string
		// errors.New() akan mengembalikan sebuah error
		return nilai / pembagi, nil // nil adalah tipe data kosong, digunakan untuk interface dan pointer
	}
}

func main() {
	hasil, err := pembagian(10, 0)
	if err == nil {
		fmt.Println("Hasil", hasil, "Error", err)
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Hasil", hasil, "Error:", err.Error()) // mencetah kedua nilai variabel hasil dan err
		fmt.Println("Error:", err.Error())
	}
}