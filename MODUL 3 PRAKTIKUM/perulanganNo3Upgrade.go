package main

import(
	"fmt"
)

func main(){
	var beratKiri, beratKanan float64

	for{
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		totalBerat := beratKiri + beratKanan
		if  beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses Selesai.")
			break	
		}

		if totalBerat > 150 {
			fmt.Println("Proses Selesai.")
			break
		}

		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		} 

		if selisih >= 9{
			fmt.Println("Sepeda motor pak andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak andi akan oleng: false")
		}

	}
}