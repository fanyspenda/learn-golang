package main

import "fmt"

func main() {
	//MAP
	var dataBulan map[int]string = map[int]string{7: "juni"}
	fmt.Println(dataBulan[7])

	dataBulan2 := map[int]string{}
	dataBulan2[7] = "July"
	fmt.Println(dataBulan2[7])
}
