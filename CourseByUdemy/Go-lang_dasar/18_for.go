package  main

import "fmt"

func main() {
	counter := 1

	// for loop sintaks panjang
	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++
	}

	fmt.Println("Telah selesai perulangan")

	// for loop sintaks pendek
	for i :=0; i <= 10; i++ {
		fmt.Println("Perulangan ke -", i)
	}

	fmt.Println("Selesai")

	// for loop bisa untuk tipe data apa saja
	// for loop untuk slice
	names := []string{"John", "Wick", "Ethan", "Hunt"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for loop pakai index dan value
	for index, value := range names {
		fmt.Println("index", index, "=", value)
	}

	// for loop pakai value saja
	for _, name := range names {
		fmt.Println(name)
	}

}