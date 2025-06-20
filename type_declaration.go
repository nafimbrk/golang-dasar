package main

import "fmt"

func main() {
	type NoKTP string

	var ktpEko NoKTP = "9834748"

	var contoh string = "232273"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohKtp)
}
