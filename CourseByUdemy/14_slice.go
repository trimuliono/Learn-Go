package main

import "fmt"

func main() {
	names := [...]string{"Tri", "Muliono", "Ganteng", "Banget", "Gaada", "Obatnya"}

	slice1 := names[2:6]
	fmt.Println(slice1) // [Ganteng Banget Gaada Obatnya]
	fmt.Println(slice1[0]) // Ganteng
	fmt.Println(slice1[1]) // Banget

	slice2 := names[:3] // Deklarasi slice tipe data string secara implisit karena array reference berupa string
	fmt.Println(slice2) // [Tri Muliono Ganteng]

	var slice3 []string = names[3:] // Deklarasi slice tipe data string secara eksplisit
	fmt.Println(slice3) // [Banget Gaada Obatnya]

	var slice4 []string = names[:] // Deklarasi slice tipe data string secara eksplisit
	fmt.Println(slice4) // [Tri Muliono Ganteng Banget Gaada Obatnya]

	days := [...]string{"Ahad", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

	daySlice1 := days[5:] // Deklarasi slice tipe data string secara implisit karena array reference berupa string
	fmt.Println(daySlice1) // [Jumat Sabtu]

	daySlice1[0] = "Jum'at baru" // Mengubah nilai array reference
	daySlice1[1] = "Sabtu baru" // Mengubah nilai array reference
	fmt.Println(daySlice1) // [Jum'at baru Sabtu baru]
	fmt.Println(days) // [Ahad Senin Selasa Rabu Kamis Jum'at baru Sabtu baru]
	
	daySlice2 := append(daySlice1, "Hari libur") // Menambahkan elemen baru ke slice && append membuat slice baru yang terlepaskan dari array reference
	fmt.Println(daySlice2) // [Jum'at baru Sabtu baru Hari libur]
	fmt.Println(daySlice1) // [Jum'at baru Sabtu baru] (tidak berubah karena slice baru)
	fmt.Println(days) // [Ahad Senin Selasa Rabu Kamis Jum'at baru Sabtu baru] (tidak berubah karena slice baru)

	daySlice2[0] = "Jumat lama" // Mengubah nilai array reference
	fmt.Println(daySlice2) // [Jumat lama Sabtu baru Hari libur]
	fmt.Println(daySlice1) // [Jumat baru Sabtu baru] (tidak berubah karena slice baru)
	fmt.Println(days) // [Ahad Senin Selasa Rabu Kamis Jum'at baru Sabtu baru] (tidak berubah karena slice baru)

	var newslice []string = make([]string, 2, 5) // Membuat slice baru dengan panjang 2 dan kapasitas 5
	// Menampilkan nilai awal
    fmt.Println("Initial slice:", newslice) // Output: ["", ""]
    fmt.Println("Length:", len(newslice))   // Output: 2
    fmt.Println("Capacity:", cap(newslice)) // Output: 5

	newslice[0] = "Tri"
	newslice[1] = "Muliono"
	// newslice[2] = "Ganteng" // error karena panjang slice hanya 2
	fmt.Println(newslice) // [Tri Muliono]

	newslice2 := append(newslice, "Ganteng") // Menambahkan elemen baru ke slice && append membuat slice baru yang terlepaskan dari array reference
	// Menampilkan nilai awal
    fmt.Println("Initial slice:", newslice2) // Output: ["", ""]
    fmt.Println("Length:", len(newslice2))   // Output: 2
    fmt.Println("Capacity:", cap(newslice2)) // Output: 5

	newslice2[0] = "Tri Baru"
	fmt.Println(newslice2) // [Tri Baru Muliono Ganteng]
	fmt.Println(newslice) // [Tri Muliono]

	fromslice := days[:] // membuat slice baru dengan isi, panjang dan kapasitas sama dengan array days
	toslice := make([]string, len(fromslice), cap(fromslice)) // Membuat slice baru dengan panjang dan kapasitas sama dengan slice fromslice
	fmt.Println("fromslice:", fromslice) // [Ahad Senin Selasa Rabu Kamis Jum'at baru Sabtu baru]
	fmt.Println("toslice:", toslice) // [  ]
	fmt.Println("Length:", len(toslice)) // 7
	fmt.Println("Capacity:", cap(toslice)) // 7

	copy(toslice, fromslice) // Menyalin nilai dari slice fromslice ke slice toslice
	fmt.Println("toslice:", toslice) // [Ahad Senin Selasa Rabu Kamis Jum'at baru Sabtu baru]

	iniArray := [...]int{1, 2, 3, 4, 5} // kita tahu ini array karena panjang index [n] di deklarasi meskipun secara implisit
	iniSlice := []int{1, 2, 3, 4, 5} // kita tahu ini slice karena panjang index [n] tidak di deklarasi

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	fmt.Println(len(iniSlice)) // 5
	fmt.Println(cap(iniSlice)) // 5

	iniSlice = append(iniSlice, 6) // Menambahkan elemen baru ke slice
	fmt.Println(iniSlice)
	fmt.Println("Length:", len(iniSlice)) // 6
	fmt.Println("Capacity:", cap(iniSlice)) // 10 (cap otomatis bertambah 2x lipat dari cap sebelumnya agar tidak terjadi realokasi memori berulang)
}