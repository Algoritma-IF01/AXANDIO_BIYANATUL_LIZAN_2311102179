// Axandio Biyanatul Lizan (2311102179)

package main

import (
	"fmt"
)

func kelipatan_rekursif(kelipatan int) int {
	var bilanganBulat int
	fmt.Print("Masukkan bilangan bulat (negatif untuk berhenti): ")
	fmt.Scan(&bilanganBulat)

	if bilanganBulat < 0 {
		return kelipatan
	}

	if bilanganBulat%4 == 0 {
		kelipatan += bilanganBulat
	}

	return kelipatan_rekursif(kelipatan)
}

func main() {
	var jumlah int
	jumlah = kelipatan_rekursif(0)
	fmt.Println("Hasil dari kelipatan 4 adalah", jumlah)
}
