package main

import (
	"fmt"
	"math"
)

func volumeBall(jariJari float64) float64 {
	return (4.0 / 3.0) * 3.1415926535 * math.Pow(jariJari, 3)
}

func luasBall(jariJari float64) float64 {
	return 4 * 3.1415926535 * math.Pow(jariJari, 2)
}

func main() {
	var jariJari float64
	fmt.Print("Jejari: ")
	fmt.Scanln(&jariJari)

	volume := volumeBall(jariJari)
	luas := luasBall(jariJari)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}
