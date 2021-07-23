package main

import (
	"fmt"
)

// Soal 1
func luasPermukaanTabung() {
	var T, r float32
	fmt.Print("masukkan Tinggi: ")
	fmt.Scan(&T)
	fmt.Print("masukkan radius: ")
	fmt.Scan(&r)
	const phi = 3.14

	var luas float32 = 2 * phi * r * (r + T)
	fmt.Println("Luas Permukaan Tabung adalah", luas)
}

//Soal 2
func scoreManager() {
	var score float32
	fmt.Print("masukkan Nilai: ")
	fmt.Scan(&score)

	switch {
	case score > 0 && score <= 34:
		fmt.Println("Nilai E")
	case score > 34 && score <= 49:
		fmt.Println("Nilai D")
	case score > 49 && score <= 64:
		fmt.Println("Nilai C")
	case score > 64 && score <= 79:
		fmt.Println("Nilai B")
	case score > 79 && score <= 100:
		fmt.Println("Nilai A")
	default:
		fmt.Print("Nilai Invalid")
	}
}

//Soal 3
func faktorBilangan() {
	var input1 int
	fmt.Print("masukkan angka: ")
	fmt.Scan(&input1)

	for i := 1; i <= input1; i++ {
		if input1%i == 0 {
			fmt.Println(i)
		}
	}
}

//Soal 4
func bilanganPrima() {
	var input1 int
	fmt.Print("masukkan angka: ")
	fmt.Scan(&input1)

	//cek apakah input = 2
	if input1 == 2 {
		fmt.Println("Bilangan Prima")
		return
	} else if input1%2 == 0 || input1 < 2 {
		// cek apakah input genap atau kurang dari 2
		fmt.Println("Bukan Bilangan Prima")
		return
	}

	var iterasi = 3

	// lakukan looping untuk mencari angka yang bisa membagi input
	for input1%iterasi != 0 {
		iterasi += 2
	}

	/*
		cek apakah iterasi sudah setara dengan input.
		Jika iya, maka bilangan tersebut prima.
		Jika tidak, maka bilangan tersebut bukan prima.
		Karena bisa dibagi selain angka 1 dan angka itu sendiri.
	*/
	if input1 == iterasi {
		fmt.Println("Bilangan Prima")
	} else {
		fmt.Println("Bukan Bilangan Prima")
	}
}

func palindrome() {
	var word string
	fmt.Print("masukkan kata: ")
	fmt.Scan(&word)
	var wordReversed string

	for i := len(word) - 1; i >= 0; i-- {
		wordReversed = wordReversed + string(word[i])
	}
	if wordReversed == word {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}

func pangkat(base, pangkat int) int {
	var result int = 1
	for i := 1; i <= pangkat; i++ {
		result = result * base
	}
	return result
}

func playWithAsterisk(n int) {
	for i := n; i >= 0; i-- {

		//memberikan jarak awal segitiga
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		//membuat bintang sebanyak iterasi ke-i
		for j := i; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func cetakTabelPerkalian(n int) {

	//looping untuk baris pertama
	for i := 1; i <= n; i++ {

		//looping untuk column pertama
		for j := 1; j <= n; j++ {
			fmt.Print(j * i)

			//jika digit 2, beri spacing lebih sedikit.
			//ini cuma biar rapi ajah.
			if j*i > 9 {
				fmt.Print("   ")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// luasPermukaanTabung()
	// scoreManager()
	// faktorBilangan()
	// bilanganPrima()
	// palindrome()

	// fmt.Println(pangkat(2, 3))  //8
	// fmt.Println(pangkat(5, 3))  // 125
	// fmt.Println(pangkat(10, 2)) // 100
	// fmt.Println(pangkat(2, 5))  // 32
	// fmt.Println(pangkat(7, 3))  // 343

	// playWithAsterisk(5)
	// playWithAsterisk(15)

	// cetakTabelPerkalian(9)

}
