package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func cangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}              // bisa kasih & nya disini, contoh:  address := &Address{}   dan dia juga sebernnya kayak gini kal pake var yang ada tipe datanya: var address *Address = &Address{}
	cangeCountryToIndonesia(&address) // atau kalo sudah terlanjur bikinnya bukan pointer bisa ditambahkan & nya disini

	fmt.Println(address)
}
