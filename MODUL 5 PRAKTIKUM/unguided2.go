package main

import (
	"fmt"
	"math"
)

const MAX int = 100

type array [MAX]int

// Mengisi array dengan N elemen
func isiArray(arr *array, n *int) {
	var i int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}
}

// Menampilkan seluruh elemen array
func tampilkanArray(arr array, n int) {
	var i int
	fmt.Print("Isi array: ")
	for i = 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks ganjil
func tampilkanIndeksGanjil(arr array, n int) {
	var i int
	fmt.Print("Indeks ganjil: ")
	for i = 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks genap
func tampilkanIndeksGenap(arr array, n int) {
	var i int
	fmt.Print("Indeks genap: ")
	for i = 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// Menampilkan elemen dengan indeks kelipatan x
func tampilkanIndeksKelipatan(arr array, n, x int) {
	var i int
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i = 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
}

// Menghapus elemen pada indeks tertentu
func hapusElemen(arr *array, n *int, indeks int) {
	var i int
	for i = indeks; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
}

// Menghitung rata-rata elemen array
func hitungRataRata(arr array, n int) float64 {
	var sum, i int
	for i = 0; i < n; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n)
}

// Menghitung standar deviasi elemen array
func hitungStandarDeviasi(arr array, n int) float64 {
	var rata, sum float64
	var i int
	rata = hitungRataRata(arr, n)
	for i = 0; i < n; i++ {
		sum += math.Pow(float64(arr[i])-rata, 2)
	}
	return math.Sqrt(sum / float64(n))
}

// Menghitung frekuensi suatu bilangan
func hitungFrekuensi(arr array, n, bilangan int) int {
	var frekuensi, i int
	for i = 0; i < n; i++ {
		if arr[i] == bilangan {
			frekuensi++
		}
	}
	return frekuensi
}

func main() {
	var arr array
	var n, x, indeks, bilangan int
	var rata, sd float64
	var frekuensi int

	isiArray(&arr, &n)
	tampilkanArray(arr, n)
	tampilkanIndeksGanjil(arr, n)
	tampilkanIndeksGenap(arr, n)

	fmt.Print("Masukkan bilangan x untuk indeks kelipatan: ")
	fmt.Scan(&x)
	tampilkanIndeksKelipatan(arr, n, x)

	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&indeks)
	hapusElemen(&arr, &n, indeks)
	tampilkanArray(arr, n)

	rata = hitungRataRata(arr, n)
	fmt.Printf("Rata-rata: %.2f\n", rata)

	sd = hitungStandarDeviasi(arr, n)
	fmt.Printf("Standar deviasi: %.2f\n", sd)

	fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&bilangan)
	frekuensi = hitungFrekuensi(arr, n, bilangan)
	fmt.Printf("Frekuensi bilangan %d: %d\n", bilangan, frekuensi)
}