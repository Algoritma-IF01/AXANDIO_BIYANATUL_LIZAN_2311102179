package main

import (
	"fmt"
)

func cariFaktor(b int) {
	fmt.Printf("Faktor dari %d: ", b)

	// Loop untuk mencari faktor dari 1 sampai b
	for i := 1; i <= b; i++ {
		// Jika b habis dibagi i, maka i adalah faktor dari b
		if b%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println() // Baris baru setelah selesai menampilkan faktor
}

func main() {
	var bilangan int

	// Input bilangan bulat
	fmt.Print("Masukkan bilangan bulat (> 1): ")
	fmt.Scan(&bilangan)

	// Validasi input agar lebih dari 1
	if bilangan > 1 {
		cariFaktor(bilangan)
	} else {
		fmt.Println("Bilangan harus > 1.")
	}
}
