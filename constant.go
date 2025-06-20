package main

import "fmt"

func main() {
	// constant wajib langsung menginisialisasi datanya
	// const firstName string = "Eko"
	// const lastName = "Khannedy"

	const (
		firstName string = "Eko"
		lastName         = "Khannedy"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Herman"
	// lastName = "Kurniansyah"
}

// di golang constant kalo gk dipake gk masalah gk akan error
