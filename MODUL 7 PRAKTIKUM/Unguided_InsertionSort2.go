// Axandio Biyanatul Lizan - 2311102179

package main

import (
	"fmt"
)

type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Rating     int
}
// Maksimal jumlah buku
const nMax int = 7919
type DaftarBuku []Buku

func main() {
	var pustaka DaftarBuku
	var n, rating int

	// Membaca banyaknya data buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// add book
	DaftarkanBuku(&pustaka, n)

	// Cetak buku fav
	fmt.Println("\nBuku Terfavorit:")
	CetakTerfavorit(pustaka, n)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka, n)

	// Cetak 5 buku dengan rating tertinggi
	fmt.Println("\n5 Buku dengan Rating Tertinggi:")
	Cetak5Terbaru(pustaka, n)

	// Cari buku berdasarkan rating
	fmt.Print("\nMasukkan rating untuk mencari buku: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}

// Add Book
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var i int
	var buku Buku

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID Judul Penulis Penerbit Eksemplar Tahun Rating):\n", i+1)
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		*pustaka = append(*pustaka, buku)
	}
}

// Cetak Buku Fav
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var i int
	var terfavorit Buku

	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	terfavorit = pustaka[0]
	for i = 0; i < n; i++ {
		if pustaka[i].Rating > terfavorit.Rating {
			terfavorit = pustaka[i]
		}
	}

	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
		terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun)
}

// Mengurutkan buku Insetion Sort
func UrutBuku(pustaka *DaftarBuku, n int) {
	var i, j int
	var key Buku

	for i = 1; i < n; i++ {
		key = (*pustaka)[i]
		j = i - 1

		for j >= 0 && (*pustaka)[j].Rating < key.Rating {
			(*pustaka)[j+1] = (*pustaka)[j]
			j--
		}
		(*pustaka)[j+1] = key
	}
}

// Cetak 5 buku rating tinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	var i, batas int

	batas = 5
	if n < 5 {
		batas = n
	}

	for i = 0; i < batas; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].Judul)
	}
}

// Binary Searching
func CariBuku(pustaka DaftarBuku, n int, r int) {
	var low, high, mid int
	var buku Buku

	low = 0
	high = n - 1

	for low <= high {
		mid = (low + high) / 2
		if pustaka[mid].Rating == r {
			buku = pustaka[mid]
			fmt.Printf("Buku ditemukan: Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Eksemplar, buku.Rating)
			return
		} else if pustaka[mid].Rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}
