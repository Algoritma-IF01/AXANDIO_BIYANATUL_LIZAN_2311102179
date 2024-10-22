package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi hitungSkor untuk menghitung jumlah soal yang diselesaikan dan total skor (waktu)
func hitungSkor(waktuSoal []int) (int, int) {
	soal := 0
	skor := 0

	for _, waktu := range waktuSoal {
		if waktu <= 300 {
			soal++
			skor += waktu
		}
	}
	return soal, skor
}

// Fungsi untuk mencari pemenang dari daftar peserta
func cariPemenang() {
	var pemenang string
	maxSoal := 0
	minWaktu := 9999999

	// Scanner untuk membaca input baris penuh
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan nama peserta dan waktu soal (atau 'Selesai' untuk mengakhiri): ")

	for {
		

		// Baca input dari user
		scanner.Scan()
		input := scanner.Text()

		// Memeriksa apakah input adalah 'selesai'
		if strings.ToLower(input) == "selesai" {
			fmt.Println(" ")
			break
		}

		var nama string
		waktuSoal := make([]int, 8)

		// Parsing input
		fmt.Sscanf(input, "%s %d %d %d %d %d %d %d %d", &nama, &waktuSoal[0], &waktuSoal[1], &waktuSoal[2], &waktuSoal[3], &waktuSoal[4], &waktuSoal[5], &waktuSoal[6], &waktuSoal[7])

		// Menghitung jumlah soal yang diselesaikan dan total waktu
		soal, skor := hitungSkor(waktuSoal)

		// Menentukan pemenang berdasarkan jumlah soal dan waktu
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}

	// Menampilkan pemenang
	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}

func main() {
	cariPemenang()
}
