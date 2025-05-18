package main

import "fmt"

func main() {
	//===============================
	//  Hello world !! 😁
	// Menampilkan tulisan Hello dan baris baru
	//============OUTPUTAN===============
	// fmt.Printf("Hello, Afif!\nD4 Teknik Informatika\n")
	// ===============================

	//===============================
	// 🧪 Variabel !!
	// Membuat variabel dengan tipe data string, int, dan float
	// var nama string = "Afif"
	// var umur = 21
	// kota := "Bandung"
	//============OUTPUTAN===============
	// fmt.Println("Hello, nama saya", nama, ", umur saya", umur, ", saya tinggal di", kota)
	//===============================

	//===============================
	// 🧪 Type Data !!
	// Mmebuat sebuah teks beserta typedatannya
	//============OUTPUTAN===============
	// var nama string = "Afif"
	// var umur int = 21
	// var tinggi float64 = 170.5
	// var mahasiswa bool = true

	// fmt.Println("Nama saya:", nama)
	// fmt.Println("Umur saya:", umur)
	// fmt.Println("Tinggi badan:", tinggi)
	// fmt.Println("Status Mahasiswa:", mahasiswa)
	//===============================

	//===============================
	// 🧪 Array!!
	// Membuat array di golang
	//============OUTPUTAN===============
	// var angka [3]int = [3]int{10,20,30}
	//  fmt.Println(angka)
	//  fmt.Println(angka[1])
	//  fmt.Println(len(angka))
	//===============================

	//===============================
	// 🧪 Slice!
	// Membuat slice di golang (penambhana isi dalam array
	//============OUTPUTAN===============
	// buah := []string{"apel", "jeruk", "mangga"}
	// buah = append(buah, "pisang")
	// fmt.Println(buah)
	//===============================
	var angka [3]int = [3]int{11, 200, 31}
	fmt.Println("isi array ", angka)
	fmt.Println("elemen pertama", angka[0])
	fmt.Println("jumlah elmen", len(angka))

}
