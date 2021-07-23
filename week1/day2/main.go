package main

import ("fmt")

func main()  {
	//static type = ada tipe data

	//error karena variable tidak terpakai
	// var age int8


	// var name, address string = "Fany", "Probolinggo"
	// var isAlive bool = true

	//dynamic type = tidak ada tipe data
	// var age2 = 3
	// age3 := 3

	// fmt.Println(age, name, address, isAlive, age2, age3)

	//konstanta
	// const (
	// 	phi float32 = 3.14
	// 	phi2
	// )

	//ZERO VALUE
	// boolean = false
	// int = 0
	// float = 0.0

	//phi2 akan otomatis mengikuti phi yang pertama
	//cuma berlaku di const, tidak di var
	// fmt.Println(phi2)

	//CONTROL STRUCTURE = mengatur alur algoritma kita (Branching dan Looping)
	// branching => if dan branch
	// looping => for

	//IF
	// var myAge uint8 = 18
	// var myRole string = "admin"

	// if myAge >= 17 && myRole=="admin" {
	// 	fmt.Println("allowed")
	// }else {
	// 	fmt.Println("not allowed")
	// }

	//Switch-Case
	// switch myAge {
	// case 18:
	// 	fmt.Println("umur 18")
	// default:
	// 	fmt.Println("not found")
	// }

	//switch-case dengan kondisi
	// switch {
	// case myAge > 17:
	// 	fmt.Println("umur Mencukupi")
	// case myAge > 18, myRole=="admin":
	// 	fmt.Println("umur Mencukupi ATAU role adalah admin")
	// default:
	// 	fmt.Println("Umur tidak mencukupi")
	// }

	//LOOPING
	//for
	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	//for (mirip while)
	// i := 2
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i *= 2
	// }

	//forEach
	name := "Fany Ervansyah"

	// for i := 0; i < len(name); i++ {
	// 	fmt.Println(string(name[i]))
	// }

	for _, v := range name {
		fmt.Print(string(v))
	}
}