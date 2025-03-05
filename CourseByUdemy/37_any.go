package main

import "fmt"


// interface kosong adalah interface yang tidak memiliki deklarasi method satupun
// interface kosong digunakan ketika kita ingin membuat function yang menerima argument sembarang
// interface kosong mirip dengan object di javascript
// interface kosong biasanya digunakan ketika kita tidak tahu tipe data apa yang akan kita terima
// interface kosong bisa digunakan sebagai tipe data return function yang bisa mereturn berbagai tipe data
func Ups() any {
	//return 1
	//return true
	return "Ups"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}