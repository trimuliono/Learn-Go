package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32) // output 32768 karna nilai masih dalam batas int32
	fmt.Println(nilai64) // output 32768 karna nilai masih dalam batas int64
	fmt.Println(nilai16) // output -32768 karna nilai sudah melebihi batas int16 sehingga nilai akan kembali berputar dari nilai terkecil

	var name = "TrI Muliono"
	var T = name[0] // default tipe datanya byte/uint8
	var Tstring = string(T) // conver byte ke string
	var R = string(name[1]) // cara lain untuk mengambil string dari byte
	var I = name[2]
	var i = name[8]

	fmt.Println(T) // output 84 karna huruf T dalam ASCII adalah 84
	fmt.Println(I) // output 73 karna huruf I dalam ASCII adalah 73
	fmt.Println(i) // output 111 karna huruf i dalam ASCII adalah 111

	fmt.Println(Tstring) // output T
	fmt.Println(R) // output r

}