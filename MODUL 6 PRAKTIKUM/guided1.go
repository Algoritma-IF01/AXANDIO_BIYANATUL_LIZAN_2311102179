// Axandio Biyanatul Lizan - 2311102179

package main

import "fmt"

type arrInt [2023]int


func terkecil(tabInt arrInt, n int) int {
	idx := 0 
	for j := 1; j < n; j++ { 
		if tabInt[j] < tabInt[idx] { 
			idx = j 
		}
	}
	return idx 
}

func main() {
	var n int
	var data arrInt

	fmt.Print("Masukkan jumlah elemen (N <= 2023): ")
	fmt.Scan(&n)

	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023")
		return
	}

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}
	idxTerkecil := terkecil(data, n)
	fmt.Printf("\nIndeks nilai terkecil: %d\n", idxTerkecil)
	fmt.Printf("Nilai terkecil: %d\n", data[idxTerkecil])
}
