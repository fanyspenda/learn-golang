package main

import (
	"day4tasks/module"
	"strings"
)

func compare(string1, string2 string) string {
	if strings.Contains(string1, string2) {
		return string2
	} else if strings.Contains(string2, string1) {
		return string1
	} else {
		return "Neither contains substring"
	}
}

func caesarCipher(offset int, string1 string) string {
	newString := ""

	//cek jika offset besar
	realOffset := offset % 26

	//jika realOffset = 0, berarti ciphernya tidak ada perubahan
	if realOffset == 0 {
		return string1
	}

	for _, v := range string1 {
		cipheredRune := int(v) + realOffset

		// jika lebih besar dari asci 'z',
		// maka kembali ke a
		if cipheredRune > 122 {
			newString += string(cipheredRune - 26)
		} else {
			newString += string(cipheredRune)
		}

	}

	return newString
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp

}

func getMinMax(numbers ...*int) (min int, max int) {
	min = *numbers[0]
	max = *numbers[0]

	for i := 1; i < len(numbers); i++ {
		if min > *numbers[i] {
			min = *numbers[i]
		}
		if max < *numbers[i] {
			max = *numbers[i]
		}
	}

	return min, max

}

func main() {
	// fmt.Println(compare("AKA", "AKASHI"))     // AKA
	// fmt.Println(compare("KANGOORO", "KANG"))  // KANG
	// fmt.Println(compare("KI", "KIJANG"))      // KI
	// fmt.Println(compare("KUPU-KUPU", "KUPU")) // KUPU
	// fmt.Println(compare("ILALANG", "ILA"))    // ILA

	// fmt.Println(caesarCipher(3, "abc"))                           // def
	// fmt.Println(caesarCipher(2, "alta"))                          // cnvc
	// fmt.Println(caesarCipher(10, "alterraacademy"))               // kvdobbkkmknowi
	// fmt.Println(caesarCipher(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	// fmt.Println(caesarCipher(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl

	// a := 1
	// b := 2
	// fmt.Println(a, b)
	// swap(&a, &b)
	// fmt.Println(a, b)

	// var a1, a2, a3, a4, a5, a6, min, max int
	// fmt.Scan(&a1)
	// fmt.Scan(&a2)
	// fmt.Scan(&a3)
	// fmt.Scan(&a4)
	// fmt.Scan(&a5)
	// fmt.Scan(&a6)
	// min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	// fmt.Println("Nilai min ", min)
	// fmt.Println("Nilai max ", max)

	// var a = module.Students{}

	// for i := 0; i < 6; i++ {
	// 	var name string
	// 	fmt.Print("Input " + string(i) + " Student’s Name : ")
	// 	fmt.Scan(&name)
	// 	a.Name = append(a.Name, name)
	// 	var score int
	// 	fmt.Print("Input " + name + " Score : ")
	// 	fmt.Scan(&score)
	// 	a.Score = append(a.Score, score)
	// }

	// fmt.Println("\n\nAvarage Score Students is", a.Average())
	// scoreMax, nameMax := a.Max()
	// fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	// scoreMin, nameMin := a.Min()
	// fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")

	var b = module.Student{NameEncode: "irapb"}

	b.Decode()

}
