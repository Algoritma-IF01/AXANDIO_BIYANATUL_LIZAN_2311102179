// Axandio Biyanatul Lizan (2311102179)

package main

import (
	"fmt"
)

func main() {
	var banyakRombongan, jumlahMenu, jumlahOrang, biayaTotal, biayaTambahan int
	fmt.Print("Masukan Jumlah Rombongan: ")
	fmt.Scan(&banyakRombongan)
	var sisaMakanan bool
	var i int

	for i = 0; i < banyakRombongan; i++ {
		fmt.Print("Masukan Jumlah Menu : ")
		fmt.Scan(&jumlahMenu)
		fmt.Print("Jumlah Orang : ")
		fmt.Scan(&jumlahOrang)
		fmt.Print("Apakah makanan habis? (1 untuk true, 0 untuk false) : ")
		fmt.Scan(&sisaMakanan)

		var biaya int

		if jumlahMenu > 50 {
			biaya = 100000
		} else if jumlahMenu > 3 {
			biaya = 10000 + (jumlahMenu-3)*2500
		} else {
			biaya = 10000
		}

		if sisaMakanan == true {
			biayaTambahan = 0
		} else {
			biayaTambahan = biaya * jumlahOrang
		}

		fmt.Printf("Total biaya untuk rombongan ke-%d: Rp %d\n", i+1, biayaTotal)
	}

}
