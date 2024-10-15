package main

import(
	"fmt"
)

func main(){
	var beratKiri, beratKanan float64

	for{
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		if beratKiri >= 9 || beratKanan >= 9{
			fmt.Println("Proses Selesai")
			break
		}

		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		} 

	}
}