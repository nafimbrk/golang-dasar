package main

import "fmt"

func main() {
	counter := 0

	// anggap aja kode yang sangat panjang disini

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)

}
