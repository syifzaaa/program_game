package main

import "fmt"

func menuShow() {
	fmt.Println("Menu:")
	fmt.Println("1. Tambahkan Rating")
	fmt.Println("2. Hapus Rating")
	fmt.Println("3. Lihat Semua Rating")
	fmt.Println("4. Lihat Semua Rating  Berdasarkan Game")
	fmt.Println("5. Top 3 Game Favorit")
	fmt.Println("6. Lihat Semua Game Berdasarkan Rating Terbaik")
	fmt.Println("0. Keluar")
}

func menuOne() {
	fmt.Print("Masukan Judul Game: ")
	var game string
	fmt.Scanln(&game)

	fmt.Pritn("Masukan Rating: ")
	var rating float64
	fmt.Scanln(&rating)

	return game, rating
}

func menuTwo() {
	fmt.Print("Masukan ID Game: ")
	var game string
	fmt.Scanln(&game)

}

func menuTwo() {
	fmt.Print("Menampilkan Seluruh : ")
	var game string
	fmt.Scanln(&game)

}

func main() {
	x := 25
	y := 10

	z := x + y

	fmt.Printf("Sum of x + y = %d", z)
}
