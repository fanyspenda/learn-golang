package main

import (
	"fmt"
	"math"
)

func simpleEquation(A, B, C int) (x, y, z int, error string) {
	factor := []int{}

	//cari faktor dari B dahulu
	for i := 1; i <= B; i++ {
		if B%i == 0 {
			factor = append(factor, i)
		}
	}

	//brute force
	//lakukan looping 3 kali di slice factor untuk cek apakah perhitungan dengan
	//index saat ini.
	for i := 0; i < len(factor); i++ {
		for j := 0; j < len(factor); j++ {
			for k := 0; k < len(factor); k++ {
				if factor[i]+factor[j]+factor[k] == A {
					if factor[i]*factor[j]*factor[k] == B {
						if math.Pow(float64(factor[i]), 2)+math.Pow(float64(factor[j]), 2)+math.Pow(float64(factor[k]), 2) == float64(C) {
							return factor[i], factor[j], factor[k], ""
						}
					}
				}
			}
		}
	}

	return 0, 0, 0, "No Solution"

}

//soal 2
func moneyCoins(money int) []int {
	change := []int{}
	moneyTemp := money

	changeList := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}

	for moneyTemp > 0 {
		for i := len(changeList) - 1; i >= 0; i-- {
			if changeList[i] <= moneyTemp {
				change = append(change, changeList[i])
				moneyTemp = moneyTemp - changeList[i]
				break
			}
		}
	}

	return change

}

//soal 3
func dragonOfLoowater(dragons, knights []int) (int, string) {

	total := 0
	isKnightFall := false

	for i := 0; i < len(dragons); i++ {
		min := knights[0]
		for j := 0; j < len(knights); j++ {
			if dragons[i] <= knights[j] {
				isKnightFall = false
				if min > knights[j] {
					min = knights[j]
				}
			} else {
				isKnightFall = true
			}
		}
		total += min
	}

	if isKnightFall {
		return 0, "Knights fall"
	}
	return total, ""

}

//soal 4
//berhasil memotong array, tapi nggk bisa return indexnya
func binarySearch(numbers []int, key int) int {

	index := -1
	currentArray := numbers

	for i := 0; i < len(currentArray); i++ {
		midIndex := 0
		if len(currentArray)%2 == 0 {
			// fmt.Println("im even")
			midIndex = (len(currentArray) / 2) - 1
		} else {
			// fmt.Println("im odd")
			midIndex = ((len(currentArray) + 1) / 2) - 1
		}

		if currentArray[midIndex] == key {
			// fmt.Println("i finally got here...")
			index = midIndex
			break
		} else if len(currentArray) == 1 {
			// fmt.Println("not found")
			index = -1
			break
		} else if currentArray[midIndex] < key {
			currentArray = currentArray[midIndex+1 : len(currentArray)]
			i = -1
		} else if currentArray[midIndex] > key {
			currentArray = currentArray[0 : midIndex+1]
			i = -1
		}
		// fmt.Println("currentArray", currentArray)
		// fmt.Println("midIndex", midIndex)
	}
	return index
}

func memoizedFibonacci() func(int) int {
	cache := map[string]int{}
	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		key := string(n)
		_, isExist := cache[key]
		if isExist {
			return cache[key]
		}
		result := 0
		if 0 == n {
			result = 0
		} else if 1 == n {
			result = 1
		} else {
			result = fibonacci(n-1) + fibonacci(n-2)
		}
		cache[key] = result
		return result
	}

	return fibonacci
}

func main() {
	// fmt.Println(simpleEquation(1, 2, 3))
	// fmt.Println(moneyCoins(543))
	// fmt.Println(dragonOfLoowater([]int{5, 10}, []int{5}))
	// fmt.Println(binarySearch([]int{1, 2, 3}, 3))
	// fmt.Println(binarySearch([]int{3, 4, 5, 6, 7, 9, 11, 12, 15, 17, 19}, 5))

	newFibonacci := memoizedFibonacci()
	fmt.Println(newFibonacci(10))

}
