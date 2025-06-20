package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Eko"
	names[1] = "Kurniawan"
	names[2] = "Khannedy"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// names[3] = "Budi"  error kapasitas cuma 3

	var values = [...]int{ // kalo mau banyak datanya flexible bisa pake [...] (dengan syarat harus langsung diinisialisasi datanya)
		90,
		80,
		95,
		100,
		125, // kalo datanya dibuat perbaris maksudnya di enter gitu nah di akhir datanya wajib pakai , (koma)
	} // kalo kita deklarsinya awalnya misal 3 panjang arraynya dan kita cuma ngisi dua maka yang satu lagi defaultnya 0 (kalo angka) "" (kalo string) begitupun gk diisi semua maka semuanya defaultnya 0

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
}

// di golang gk ada operasi untuk menghapus data array
