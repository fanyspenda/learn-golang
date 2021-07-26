package main

import (
	"fmt"
	"strconv"
)

//Soal 2
func fasterExponent(n, pow int) int {

	// Konsep
	// kita bisa bagi jadi 2 bagian jika pangkatnya genap. Contoh:
	// 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2
	// = (2 * 2 * 2 * 2) * (2 * 2 * 2 * 2)
	// sehingga, looping cukup setengahnya saja,
	// kemudian dikalikan hasil perpangkatan tadi.

	// Jika ganjil, tinggal buang dulu 1 angka. Contoh:
	// 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2
	// = (2 * 2 * 2 * 2) * (2 * 2 * 2 * 2) * 2
	result := 1

	if pow == 0 {
		result = 1
	} else if pow%2 == 0 {
		for i := 1; i <= pow/2; i++ {
			fmt.Println("perulangan ke-" + fmt.Sprint(i))
			result *= n
		}
		result *= result
	} else {
		for i := 1; i <= (pow-1)/2; i++ {
			result *= n
			fmt.Println("perulangan ke-" + fmt.Sprint(i))
		}
		result *= result
		result *= n
	}

	return result
}

//Soal 3
func arrayMerge(string1, string2 []string) []string {
	result := []string{}
	combinedSlice := []string{}

	//menggabung semua array terlebih dahulu
	combinedSlice = append(string1, string2...)

	//mengambil nilai index pertama dari array yang sudah digabung
	//(karena index pertama pasti unik dan tidak mungkin kembar dengan result)
	result = append(result, string1[0])

	//perulangan dimulai dari index ke-1
	//karena index ke-0 sudah masuk ke result
	for i := 1; i < len(combinedSlice); i++ {
		fmt.Println("perulangan ke", i)

		//cek apakah value dari combinedSlice sudah ada di result
		for j := 0; j < len(result); j++ {
			fmt.Println("perulangan dalam ke", j+1)
			if combinedSlice[i] != result[j] && j == (len(result)-1) {
				result = append(result, combinedSlice[i])
			} else if combinedSlice[i] == result[j] {
				break
			}
		}
	}

	return result
}

//Soal 4
func munculSekali(angka string) []int {
	result := []int{}
	combinedSlice := []int{}

	for _, v := range angka {
		intValue, _ := strconv.Atoi(string(v))
		combinedSlice = append(combinedSlice, intValue)
	}

	fmt.Println("combinedSlice", combinedSlice)

	for i := 0; i < len(combinedSlice); i++ {

		for j := 0; j < len(combinedSlice); j++ {

			if combinedSlice[i] == combinedSlice[j] && j != i {
				break
			} else if j == len(combinedSlice)-1 {
				result = append(result, combinedSlice[i])
			}
		}
	}

	return result
}

func pairSum(arr []int, target int) []int {
	result := []int{}

	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				result = append(result, i)
				result = append(result, j)
			}
		}
	}

	return result
}

func main() {
	// fmt.Println(fasterExponent(2, 6))
	// fmt.Println(arrayMerge([]string{"1", "2", "2"}, []string{"3", "2", "4", "5"}))

	// fmt.Println(munculSekali("1234123"))    // [4]
	// fmt.Println(munculSekali("76523752"))   // [6, 3]
	// fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	// fmt.Println(munculSekali("1122334455")) // []
	// fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]

	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]

}
