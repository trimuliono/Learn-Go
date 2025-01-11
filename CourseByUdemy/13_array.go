package main

import "fmt"

func main() {
	var names [3]string // deklarasi array dengan panjang index [n] = 3 dan tipe data string, diawal tanpa nilai

	names[0] = "Tri"     // inisialisasi nilai array index ke 0
	names[1] = "Muliono" // inisialisasi nilai array index ke 1
	names[2] = "Ganteng" // inisialisasi nilai array index ke 2

	var values = [5]int{90, 80, 70, 100, 50} // deklarasi array dengan panjang index [n] = 5 dan tipe data int, diawal dengan nilai

	var hobbies = [...]string{ // deklarasi array dengan panjang index [n] bebas (sesuai jumlah elemen) dan tipe data string
		"Ngoding",
		"Baca Komik",
		"rebahan",
		"lari", // jika format pakai enter maka harus ada koma diakhir
	}

	var pets = [...]string{"Kucing", "Cicak", "Semut"} // deklarasi array dengan panjang index [n] bebas (sesuai jumlah elemen) dan tipe data string

	fmt.Println(names)   // [Tri Muliono Ganteng]
	fmt.Println(values)  // [90 80 70 100 50]
	fmt.Println(hobbies) // [Ngoding Baca Komik lari]
	fmt.Println(pets)    // [Kucing Cicak Semut]

	fmt.Println(names[0])   // Tri
	fmt.Println(hobbies[1]) // Baca Komik
	fmt.Println(pets[2])    // Semut

	fmt.Println(len(hobbies)) // 4

	// names[3] = "Gaul" // error karena panjang array hanya 3
	names[2] = "Keren" // bisa karena index ke 2 masih dalam batas panjang array tersebut
	fmt.Println(names) // [Tri Muliono Keren] (Ganteng berubah jadi Keren)

}
