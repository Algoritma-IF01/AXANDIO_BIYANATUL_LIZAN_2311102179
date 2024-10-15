package main

import (
	"fmt"
)

func biayaPengiriman(beratParsel int){
	beratKG := beratParsel / 1000
	sisaGram := beratParsel % 1000

	biayaKg := beratKG * 10000

	var biayaSisa int

	if beratKG >= 10 {
		biayaSisa = 0
	} else {
		if sisaGram >= 500 {
			biayaSisa = sisaGram * 5
		} else {
			biayaSisa = sisaGram * 15
		}
	}

	totalBiaya := biayaKg + biayaSisa

	fmt.Printf("Detail berat: %d kg + %d gram\n", beratKG, sisaGram)
	fmt.Printf("Detail Biaya: Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total Biaya: Rp. %d\n", totalBiaya)
}

func main(){
	var beratParsel int

	fmt.Print("Berat Persel (dalam gram) : ")
	fmt.Scan(&beratParsel)

	biayaPengiriman(beratParsel)
}