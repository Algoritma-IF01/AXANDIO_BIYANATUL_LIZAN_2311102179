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
	var n, i int
	var scanner *bufio.Scanner

	// Membuat scanner untuk membaca input
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan jumlah daerah: ")
	if scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())
	}

	// Menyimpan hasil data yang dimasukkan
	var semuaDaerah [][]int

	// Memproses setiap daerah
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

// Proses input pada satu daerah
func processLine(line string) []int {
	var parts []string
	var ganjil, genap []int
	var m, i, rumah int

	parts = strings.Fields(line)

	m, _ = strconv.Atoi(parts[0])
	ganjil = []int{}
	genap = []int{}
	for i = 0; i < m; i++ {
		rumah, _ = strconv.Atoi(parts[i+1])
		if rumah%2 == 0 {
			genap = append(genap, rumah)
		} else {
			ganjil = append(ganjil, rumah)
		}
	}
	selectionSort(ganjil, true)  // Ascending (bil ganjil)
	selectionSort(genap, false) // Descending (bil genap)

	// Gabungkan hasil dan kembalikan
	ganjil = append(ganjil, genap...)
	return ganjil
}

// Mengurutkan array (selection sort)
func selectionSort(arr []int, ascending bool) {
	var n, i, j, extremeIdx int

	n = len(arr)
	for i = 0; i < n-1; i++ {
		extremeIdx = i
		for j = i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[extremeIdx]) || (!ascending && arr[j] > arr[extremeIdx]) {
				extremeIdx = j
			}
		}
		arr[i], arr[extremeIdx] = arr[extremeIdx], arr[i]
	}
}

// Cetak Array
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
