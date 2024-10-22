package main

import (
	"fmt"
)

// f(x) = x^2
func f(x int) int {
	return x * x
}

// g(x) = x - 2
func g(x int) int {
	return x - 2
}

// h(x) = x + 1
func h(x int) int {
	return x + 1
}

// Fungsi (fogoh)(x)
func fogoh(x int) int {
	return f(g(h(x)))
}

// Fungsi (gohof)(x)
func gohof(x int) int {
	return g(h(f(x)))
}

// Fungsi (hofog)(x)
func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	// deklarasi variabel
	var a, b, c int
	// input nilai a,b,c
	fmt.Scanln(&a, &b, &c)

	// melakukan transformasi data
	fogohResult := fogoh(a)
	gohofResult := gohof(b)
	hofogResult := hofog(c)

	fmt.Println("Keluaran : ")
	// Hitung dan tampilkan hasil fogoh(a), gohof(b), hofog(c)
	fmt.Print(fogoh(a))
	fmt.Print("\n", gohof(b))
	fmt.Print("\n", hofog(c), "\n")

	fmt.Print("\nPenjelasan :\n")
	fmt.Printf("(fogoh) (%d) = %d\n", a, fogohResult)
	fmt.Printf("(gohof) (%d) = %d\n", b, gohofResult)
	fmt.Printf("(hofog) (%d) = %d\n", c, hofogResult)

}
