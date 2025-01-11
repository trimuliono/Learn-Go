package main

import "fmt"

func main() {
	var nilaiAkhir byte = 90 // tipe data byte/uint8 = (0-255)
	var absensi = 80 // tipe data default int = (int32/int64) tergantung sistem operasi 32/64 bit


	var LulusNilaiAkhir bool = nilaiAkhir > 80 
	var LulusAbsensi = absensi >= 80 // tipe data default bool (boolean)

	var Lulus = LulusNilaiAkhir && LulusAbsensi
	var Ulangan = ! Lulus

	fmt.Println(Lulus)
	fmt.Println(Ulangan)
}