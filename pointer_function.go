package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func cangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}              // bisa kasih & nya disini
	cangeCountryToIndonesia(&address) // atau kalo sudah terlanjur dibuat bisa ditaruh disini

	fmt.Println(address)
}
