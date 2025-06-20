package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Eko"
	// person["address"] = "Subang"

	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	// fmt.Println(person["salah"])  // kalo coba akses key yang tidak ada maka secara otomatis menggunakan default value tergantung tipenya, kalo tipenya string default value nya "" kalo number 0

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Eko"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups") // kalo map dia bisa delete beda dengan array dan slice

	fmt.Println(book)

}
