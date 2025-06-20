package main

import "fmt"

func main() {
	// var name string (kalo tidak langsung di inisialisasi harus pakai tipe data)
	// tanpa var (asalkan pake := dan itu hanya berlaku pas awal inisialisasi, saat mau mengubah tetep =)
	name := "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var (
		firstName = "Eko"
		lastName  = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}

// di golang varible kalo gk dipake akan error
