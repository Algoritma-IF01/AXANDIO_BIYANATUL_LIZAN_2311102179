// Axandio Biyanatul Lizan - 2311102179

package main

import (
	"fmt"
)

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, banyakData int, bMin, bMax *float64) {
	var i int
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for i = 1; i < banyakData; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, banyakData int) float64 {
	var total float64
	var i int
	for i = 0; i < banyakData; i++ {
		total += arrBerat[i]
	}
	return total / float64(banyakData)
}

func main() {
	var arrBerat arrBalita
	var banyakData int
	var i int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scanln(&banyakData)

	for i = 0; i < banyakData; i++ {
		fmt.Printf("Masukan berat balita ke-%d (dalam kg): ", i+1)
		fmt.Scanln(&arrBerat[i])
	}

	var bMin, bMax float64

	hitungMinMax(arrBerat, banyakData, &bMin, &bMax)
	var rataRata float64
	rataRata = rerata(arrBerat, banyakData)

	fmt.Printf("\nBerat balita minimum : %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum : %.2f kg\n", bMax)
	fmt.Printf("Rata-rata berat balita : %.2f kg\n", rataRata)
}
