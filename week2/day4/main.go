package main

import (
	"day4/features"
	"fmt"
	"strings"
)

//Variadic FUnction
func penjumlahan(data ...int) (sum int) {
	for _, v := range data {
		sum += v
	}
	return
}

func fungsiDefer() {
	defer func() {
		fmt.Println("jalan terakhir")
	}()
	fmt.Println("jalan pertama")
}

//Car ....
type Car struct {
	Name string
	Tire int
}

//METHOD
//fungsi yang menempel pada tipe (struct ataupun tipe lain)

//SayHello ...
func (car1 Car) SayHello() {
	fmt.Println("Hello" + car1.Name)
}

//INTERFACE
// untuk menggabungkan method-method agar lebih rapi

//CarInterface ...
type CarInterface interface {
	SayHello()
}

func main() {
	//String
	//string memiliki fungsi
	str1 := "hello world"
	str2 := "hello"

	fmt.Println(strings.Contains(str1, str2))

	//substring
	str3 := "cat:dog"

	//cat
	fmt.Println(str3[0:3])
	fmt.Println(str3[:3])

	//dog
	fmt.Println(str3[4:])

	penjumlahan(1, 3, 4, 15)

	//function bisa anonymous dan bisa disimpan di variable

	fungsiAnonim := func(a, b int) {
		fmt.Println(a + b)
	}

	fungsiAnonim(2, 5)

	//CLOSURE

	// fmt.Println(newCounter())

	fungsiDefer()

	//RAM hanya menyimpan alamat dan value dari variable
	//Golang mengatur nama variable

	//POINTER
	var name = "Fany Ervansyah"
	var p *string = &name

	fmt.Println(name)

	//mereturn alamat memeori dari "name"
	fmt.Println(p)

	//mengakses value dari "name" menggunakan alamat memori
	fmt.Println(*p)

	//POINTER TANPA REFERENSI ke VARIABLE
	// var newPointer = new(int)

	//STRUCT, METHOD, INTERFACE
	//STRUCT

	//nama dan properti struct huruf besar jika ingin diakses di module lain

	var mobil1 Car = Car{"BMW", 4}
	var mobil2 Car = Car{}
	var mobil3 Car = Car{Tire: 3}
	mobil3.Name = "Nozomi"

	fmt.Println(mobil1)
	fmt.Println(mobil2)

	//METHOD
	//fungsi yang menempel pada tipe (struct atau tipe lain)

	mobil3.SayHello()

	//interface
	CarInterface.SayHello(mobil1)

	//PACKAGE
	features.FeatureA()
}
