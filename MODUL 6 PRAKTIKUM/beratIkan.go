// Axandio Biyanatul Lizan - 2311102179

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	// Array untuk menampung berat ikan
	var beratIkan [1000]float64

	// Input berat ikan
	for a := 0; a < x; a++ {
		fmt.Scan(&beratIkan[a])
	}

	// Menghitung total berat di setiap wadah
	jumlahWadah := int(math.Ceil(float64(x) / float64(y))) // Jumlah wadah
	beratIkanPerWadah := make([]float64, jumlahWadah)

	for a := 0; a < x; a++ {
		wadahIndeks := a / y
		beratIkanPerWadah[wadahIndeks] += beratIkan[a]
	}

	// Output total berat per wadah
	for _, beratTotal := range beratIkanPerWadah {
		fmt.Printf("%.2f ", beratTotal)
	}
	fmt.Println()

	// Menghitung rata-rata berat ikan per wadah
	var beratKeseluruhan float64
	for _, beratTotal := range beratIkanPerWadah {
		beratKeseluruhan += beratTotal
	}
	rataRataBerat := beratKeseluruhan / float64(jumlahWadah)
	fmt.Printf("Berat Rata Rata : %.2f\n", rataRataBerat)
}
