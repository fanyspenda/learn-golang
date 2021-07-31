package main

func main() {
	/*
		Problem Solving Paradigma
		 1. Complete search (brute force)
		 2. Divide and Conquer
		 3. Greedy
		 4. Dynamic Programming

		 //COMPLETE SEARCH - O(n)
		 - misal ada data di array[5], maka dicari sesuai panjang arraynya
		 - dipakai jika tidak ada solusi lain (last solution)

		 //DIVIDE AND CONQUER
		 - Dibagi, lalu dicari.
		 - Biasanya dipakai di binary search
		 - Syaratnya, harus diurutkan dulu

		 //GREEDY
		 - mencari nilai local optimal, bukan global optimal.
		 - misal dari kota A akan pergi menuju kota Z. dari A, ada bisa melalui 2 kota, C dan D.
		 maka, dari A akan dicari jarak antara A ke C apakah lebih singkat dari A ke D.
		 Kekurangannya, Jika A ke C lebih kecil, tapi dari C ke Z berjarak lebih jauh dari D ke Z,
		 maka itu tidak akan terdeteksi.

		 jika menggunakan global optimal, maka nanti akan dilist dahulu tiap-tiap jalur dari A ke Z.
		 lalu dicari yang paling singkat.
		 kekurangan dari GREEDY, dia memang lebih cepat dari yang menggunakan global optimal, tapi bisa jadi
		 tidak efektif.
		 kekurangan mencari global optimal, dia memang efektif, tapi lama.

		 //DYNAMIC PROGRAMMING
		 1. Top-Down => menggunakan teknik memoization
		 2. Bottom-up => teknik Tabulation




	*/

}
