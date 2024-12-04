// Axandio Biyanatul Lizan - 2311102179

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var scanner *bufio.Scanner

	// Membuat scanner untuk input
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan jumlah daerah: ")
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	// Menyimpan data dari semua daerah
	var semuaDaerah [][]int
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan data untuk daerah ke-%d:\n", i+1)
		if scanner.Scan() {
			var line string
			line = scanner.Text()
			daerah := processLine(line) // Proses dan simpan data
			semuaDaerah = append(semuaDaerah, daerah)
		}
	}

	// Setelah semua data dimasukkan, tampilkan output
	fmt.Println("\nHasil setelah diurutkan:")
	for i = 0; i < n; i++ {
		printArray(semuaDaerah[i])
	}
}

// Fungsi memproses input pada satu daerah
func processLine(line string) []int {
	var parts []string
	var m int
	var rumah []int
	var i int

	parts = strings.Fields(line)
	m, _ = strconv.Atoi(parts[0])
	rumah = make([]int, m)
	for i = 0; i < m; i++ {
		rumah[i], _ = strconv.Atoi(parts[i+1])
	}

	selectionSort(rumah) // Urutkan data
	return rumah         // Kembalikan array yang sudah diurutkan
}

// Mengurutkan array menggunakan Selection Sort
func selectionSort(arr []int) {
	var n int
	var i, j, minIdx int

	n = len(arr)
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

// Fungsi mencetak array
func printArray(arr []int) {
	var i, val int

	for i, val = range arr {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(val)
	}
	fmt.Println()
}
