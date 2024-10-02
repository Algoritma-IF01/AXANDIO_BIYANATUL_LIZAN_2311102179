package main

import (
	"fmt"
)

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
		fmt.Println("Tahun kabisat : true")
	} else {
		fmt.Println("Tahun kabisat : false")
	}
}