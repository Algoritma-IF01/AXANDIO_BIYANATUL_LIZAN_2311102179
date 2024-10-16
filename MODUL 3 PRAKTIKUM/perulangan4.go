package main

import (
	"fmt"
)

// Fungsi untuk menghitung operasi F(k) dan mengalikan hasilnya ke variabel hasil
func operasiF(k float64) float64 {
	up := (4*k + 2) * (4*k + 2)  // Menghitung bagian atas dari rumus
	down := (4*k + 1) * (4*k + 3) // Menghitung bagian bawah dari rumus
	return up / down               // Mengembalikan hasil bagi
}

// Fungsi untuk menghitung hampiran akar 2
func operasiAkar2(k int) float64 {
	hasil := 1.0 // Inisialisasi hasil dengan 1.0
	for i := 0; i <= k; i++ {
		hasil *= operasiF(float64(i)) // Mengalikan hasil dari operasi F untuk setiap i
	}
	return hasil
}

func main() {
	var k int

	// Input nilai k dari pengguna
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Menghitung hampiran akar 2 dan menampilkan hasilnya
	hasil := operasiAkar2(k)
	fmt.Printf("Nilai hampiran akar 2 = %.10f\n", hasil)
}
