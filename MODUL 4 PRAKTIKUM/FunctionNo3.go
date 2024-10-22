package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x1-x2), 2) + math.Pow(float64(y1-y2), 2))
}

// Fungsi untuk mengecek apakah titik berada di dalam lingkaran
func isInsideCircle(x, y, cx, cy, r int) bool {
	return distance(x, y, cx, cy) <= float64(r)
}

func main() {
	//deklarasi variabel
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

	// Input untuk lingkaran 1 (titik pusat & radius lingkaran)
	fmt.Scan(&cx1, &cy1, &r1)

	// Input untuk lingkaran 2 (titik pusat & radius lingkaran)
	fmt.Scan(&cx2, &cy2, &r2)

	// Input untuk titik sembarang
	fmt.Scan(&x, &y)

	// Cek posisi titik terhadap lingkaran 1 dan 2
	inCircle1 := isInsideCircle(x, y, cx1, cy1, r1)
	inCircle2 := isInsideCircle(x, y, cx2, cy2, r2)

	// Menentukan output berdasarkan posisi titik
	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
