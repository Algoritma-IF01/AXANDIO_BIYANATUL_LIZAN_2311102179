package main

import (
	"fmt"
)

func cariFaktor(b int) []int {
	var faktor []int

	// Loop untuk mencari faktor dari 1 sampai b
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}

	// Mengembalikan slice berisi faktor-faktor
	return faktor
}

func cekPrima(b int) bool {
	// Sebuah bilangan prima hanya memiliki 2 faktor: 1 dan dirinya sendiri
	faktor := cariFaktor(b)

	// Jika faktor yang ditemukan hanya 2, berarti bilangan tersebut prima
	return len(faktor) == 2
}

func main() {
	var bilangan int

	// Input bilangan bulat
	fmt.Print("Masukkan bilangan bulat (lebih besar dari 0): ")
	fmt.Scan(&bilangan)

	// Validasi input agar lebih dari 0
	if bilangan > 0 {
		// Mencari faktor-faktor bilangan
		faktor := cariFaktor(bilangan)

		// Menampilkan faktor-faktor bilangan
		fmt.Printf("Faktor dari %d: ", bilangan)
		for _, f := range faktor {
			fmt.Printf("%d ", f)
		}
		fmt.Println()

		// Menentukan apakah bilangan prima
		if cekPrima(bilangan) {
			fmt.Println("Prima: true")
		} else {
			fmt.Println("Prima: false")
		}
	} else {
		fmt.Println("Bilangan harus lebih besar dari 0.")
	}
}
