package main

import "fmt"

// tipe data nil adalah tipe data yang tidak memiliki nilai
// tipe data nil biasanya digunakan untuk menandakan bahwa suatu variabel tidak memiliki nilai
// dalam golang, tipe data nil bisa digunakan untuk beberapa tipe data seperti interface, function, map, slice, pointer, dan channel
// tipe data nil tidak bisa digunakan untuk tipe data numerik seperti int, float, bool, string, rune, dan sebagainya	


func NewMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string] string {
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")

	if data == nil {
		fmt.Println("Data kosong")

	} else {
		fmt.Println(data)
	}

	kode := NewMap("Kode")

	if kode == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(kode)
		fmt.Println(kode["name"])
	}

}