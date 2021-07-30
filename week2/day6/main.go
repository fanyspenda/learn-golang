package main

func main() {
	//CONCURRENT
	// task berjalan secara sendiri-sendiri (independent)

	//Kapan lebih baik pakai concurrency?
	//ketika ada
	// fmt.Println("CONCURRENT")

	//SEQUENTIAL
	//task baru tidak boleh dimulai sebelum task sebelumnya selesai

	//PARALLEL
	//multiple tasks dijalankan secara bersamaan
	//syaratnya, harus punya multi-core machines

	//FEATURE GO
	// Concurrent execution (Goroutines)
	// Synchronization and messaging (Channels)
	// Multi-way concurrent control (Select)

	//GOROUTINES
	//merupakan fungsi yang berjalan secara concurrent (mirip asynchronous)

	//akhirnya, muncul masalah racing time dimana goroutines sedang memproses data
	//tapi proses di bawahnya berjalan dahulu, padahal pengennya synchronous

	// handlingnya dengan 3 cara, yang intinya sama. yaitu
	//1. Sync.waitGroup

}
