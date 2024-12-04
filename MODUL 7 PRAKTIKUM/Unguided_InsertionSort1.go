// Axandio Biyanatul Lizan - 2311102179
package main

import (
	"fmt"
	"math"
)

func main() {
	var input int
	var data []int
	var statusJarak string

	// Membaca input hingga bilangan negatif ditemukan
	fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	// Melakukan pengurutan
	insertionSort(data)

	// Mencetak data setelah diurutkan
	fmt.Println("Hasil setelah diurutkan:", data)

	// Mengecek jarak data
	statusJarak = periksaJarak(data)
	fmt.Println(statusJarak)
}

// Insertion sort
func insertionSort(arr []int) {
	var i, key, j int

	// Iterasi setiap array
	for i = 1; i < len(arr); i++ {
		key = arr[i]
		j = i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa jarak data
func periksaJarak(arr []int) string {
	var jarak int
	var i int

	if len(arr) < 2 {
		return "Data tidak cukup untuk memeriksa jarak."
	}

	// mengitung jarak awal
	jarak = int(math.Abs(float64(arr[1] - arr[0])))

	// Memeriksa jarak apakah sama
	for i = 1; i < len(arr)-1; i++ {
		if int(math.Abs(float64(arr[i+1]-arr[i]))) != jarak {
			return "Data berjarak tidak tetap"
		}
	}

	// Jika semua jarak sama
	return fmt.Sprintf("Data berjarak %d", jarak)
}
