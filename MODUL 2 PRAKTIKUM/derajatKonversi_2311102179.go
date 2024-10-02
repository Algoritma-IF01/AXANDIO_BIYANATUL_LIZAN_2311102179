package main

import "fmt"

func main(){
	var celcius, reamur, kelvin, fahrenreit float64
	fmt.Print("Temperatur Celcius : ")
	fmt.Scanln(&celcius)
	
	// FORMULA
	reamur = (celcius * 4/5)
	fahrenreit = (celcius * 9/5) + 32
	kelvin = (fahrenreit + 459.67) * 5/9

	fmt.Println("Derajat Reamur : ", reamur)
	fmt.Println("Derajat Fahrenreit : ", fahrenreit)
	fmt.Println("Derajat Kelvin : ", kelvin)
}