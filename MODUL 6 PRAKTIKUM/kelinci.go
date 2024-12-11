// Axandio Biyanatul Lizan - 2311102179

package main

import(
	"fmt"
)

func main(){
	var N int
	var weights []float64
	var i, j int
	var minWeight, maxWeight float64

	fmt.Print("Masukan jumlah anak kelinci : ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
		return
	}

	weights = make([]float64, N)
	fmt.Println("Masukan berat anak kelinci : ")
	for i = 0; i < N; i++{
		fmt.Scan(&weights[i])
	}

	minWeight, maxWeight = weights[0], weights[0]

	for j = 1; j < N; j++ {
		if weights[j] < minWeight {
			minWeight = weights[j]
		}
		if weights[j] > maxWeight {
			maxWeight = weights[j]
		}
	}

	fmt.Printf("Berat Kelinci Terkecil : %.2f\n", minWeight)
	fmt.Printf("Berat Kelinci Terbesar : %.2f\n", maxWeight)
}