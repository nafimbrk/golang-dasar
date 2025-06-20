package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// adddress1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// adddress2 := adddress1 // copy value

	// adddress2.City = "Bandung"
	// fmt.Println(adddress1) // tidak diubah
	// fmt.Println(adddress2) // berubah menjadi bandung

	adddress1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	adddress2 := &adddress1 // pointer

	// var adddress1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}  // yang diatas sebenernya kayak gini tipe datanya
	// var adddress2 *address = &adddress1

	adddress2.City = "Bandung"
	fmt.Println(adddress1) // ikut berubah
	fmt.Println(adddress2) // berubah menjadi bandung
}
