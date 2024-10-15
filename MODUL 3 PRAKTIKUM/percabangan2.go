// package main

// import(
	// "fmt"
// )

// func main(){
	// var nam float64
	// var nmk string

	// fmt.Print("Nilai akhir mata kuliah : ")
	// fmt.Scanln(&nam)

	// if nam > 80 {
		// nkm = "A"
	//}
	// if nam > 72.5 {
		// nkm = "AB"
	// }
	// if nam > 65 {
		// nkm = "B"
	// }
	// if nam > 57.5 {
		// nkm = "BC"
	// }
	// if nam > 50 {
		// nkm = "C"
	// }
	// if nam > 40 {
		// nkm = "D"
	// } else if nam <= 40{
		// nkm = "E"
	// }

	// fmt.Println("Nilai mata kuliah: ", nmk)
//


// 1. Jika nam diberikan adalah 80,1. Apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
// 2. Apa saja kesalahan dari program tersebut? mengapa demikian? jelaskan alur program seharusnya!
// 3. Perbaiki program tersebut! ujilah dengan masukan : 93.5, 70.6, dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A','B', dan 'D'

//Jawaban
// 1. Jika dimasukan 80.1 maka keluaran program adalah D. eksekusi program tidak sesuai dengan spesifikasi soal
// 2. Kesalahan pertama adalah pada percabangan yang tidak sesuai dengan spesifikasi soal. Maka jika kita menginputkan angka yang lebih besar dri 40 hingga 100 maka akan terdeteksi D.Mengapa bisa begitu? ini dikarenakan dibagian percabangan tidak ada kondisi yang spesifik untuk terhadap nilai dan grade sesuai yang diminta oleh soal. Seharusnya di percabangan perlu ditambahkan kondisi sampai dengan dengan menggunakan notasi && agar dapat menyesuaikan yang diminta oleh soal. Misal grade AB adalah nilai diatas 72.5 sampai dengan 80, jika diatas 80 maka sudah masuk ke grade A
// 3. Berikut adalah perbaikan dari program tersebut

// Kode setelah diperbaiki
package main

import("fmt")

func main(){
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80{
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5{
		nmk = "B"
	} else if nam > 57.5 && nam <= 65{
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5{
		nmk = "C"
	} else if nam > 40 && nam <= 50{
		nmk = "D"
	} else if nam <= 40{
		nmk = "C"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}	