// Axandio Biyanatul Lizan 2311102179

package main

import (
	"fmt"
	"strconv"
)

func validasi_voucher(angka_voucher int) bool {

	voucher_str := strconv.Itoa(angka_voucher)
	length := len(voucher_str)

	if length != 5 && length != 6 {
		return false
	}

	two_first_digits, _ := strconv.Atoi(voucher_str[0:2])
	two_last_digits, _ := strconv.Atoi(voucher_str[length-2:])

	if two_first_digits*two_first_digits != two_last_digits*two_last_digits {
		return false
	}

	if length == 5 {
		mid_digit, _ := strconv.Atoi(string(voucher_str[2]))
		if mid_digit%2 != 0 {
			return false
		}
	} else {

		mid_digits := voucher_str[2:5]
		for _, d := range mid_digits {
			digit, _ := strconv.Atoi(string(d))
			if digit%2 != 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	var numVoucher, banyakMahasiswa int
	var i, validCount, invalidCount int

	fmt.Print("Masukan Banyak Mahasiswa : ")
	fmt.Scan(&banyakMahasiswa)

	for i = 0; i < banyakMahasiswa; i++ {
		fmt.Print("Masukan Voucher Mahasiswa Ke-", i+1, " : ")
		fmt.Scan(&numVoucher)

		if validasi_voucher(numVoucher) {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Println("\nBanyak Voucher Valid:", validCount)
	fmt.Println("Banyak Voucher Tidak Valid:", invalidCount)
}
