package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	adddress1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	adddress2 := &adddress1 // pointer

	adddress2.City = "Bandung"
	fmt.Println(adddress1) // ikut berubah
	fmt.Println(adddress2) // berubah menjadi bandung

	// adddress2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*adddress2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(adddress1)
	fmt.Println(adddress2)
}
