package main

import (
	"fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

// Prosedur untuk mengisi array dengan karakter hingga mencapai tanda titik atau maksimal NMAX
func isiArray(t *tabel, n *int) {
	var input rune
	fmt.Println("Masukkan teks (akhiri dengan TITIK):")
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' {
			break
		}
		t[*n] = input
		*n++
	}
}

// Prosedur untuk mencetak isi array
func cetakArray(t tabel, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

// Prosedur untuk membalikkan urutan isi array
func balikanArray(t *tabel, n int) {
	var i int
	for i = 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

// Fungsi untuk memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	var i int
	for i = 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	var isPalindrom bool

	// Isi array tab
	isiArray(&tab, &m)

	// Cetak teks asli
	fmt.Print("Teks : ")
	cetakArray(tab, m)

	// Balikkan isi array dan cetak
	balikanArray(&tab, m)
	fmt.Print("Reverse Teks : ")
	cetakArray(tab, m)

	// Periksa apakah palindrom
	isPalindrom = palindrom(tab, m)
	fmt.Printf("Palindrom : %v\n", isPalindrom)
}
