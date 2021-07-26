# Complexity

-   kompleksitas mempengaruhi performa.
-   2 detik itu sudah termasuk kurang efisien.

## Time Complexity (Big-0 complexity)

## Time Complexity

-   menentukan time complexity berdasarkan jumlah operasi yang paling dominan
-   Misal ada fungsi menerima input angka dan melakukan looping sebanyak input tersebut. Maka, perulangan tersebut termasuk dominant operation, dalam hal ini, big-O(n) -> linear operation

### 1. Constant Time - O(1)

-   Merupakan algoritma tercepat
-   menerima input berapapun, operasi tetap hanya dilakukan sekali.

### 2. Linear Time - O(n)

-   operasi berbanding lurus dengan input
-   contoh: fungsi dengan perulangan yang bergantung pada input.

### 3. Linear Time - O(n+m)

-   ada 2 perulangan sejajar, dan fungsi menerima 2 input
-   O(n) + O(m)

### 4. Logarithmic Time - O(log n)

-   Lebih cepat dari Linear Time, tapi lebih lambat dari Constant Time.
-   Biasanya dipakai di searching paradigma
-   jika dalam perulangan, biasanya dibagi 2 setiap operasi. Misal: input 100. 100 dibagi 2 terus menerus hingga > 1. Dengan begitu, perulangan tidak sampai 100 kali.

### 5. Quadratic Time - O(n<sup>2</sup>)

-   **HINDARI !!**
-   contoh paling sering adalah nested looping.

### 6 Exponential (O(2<sup>n</sup>)) dan Factorial Time (O(n!))

-   Lambat juga
-   katanya, lebih lambat dari Quadratic Time. Tapi, nggk tau lagi.

contoh soal:

```go
func function1 (input1 int) int {
    for i:= 0; i < input1; i++ {
        for j:= 0; i < 5; j++ {
            //operasi
        }
    }
}
```

operasi di atas adalah O(input1 \* 5)

## Space Complexity

Jumlah alokasi memori yang digunakan untuk menyimpan variable.

contoh:

```go
func function1 (input1 int) int {
    //membuat array sebanyak input1
}
```

space complexity fungsi di atas adalah O(n)

# Struktur Data

## ARRAY

jumlah array di golang tidak bisa diubah, tidak seperti javascript.

```go
func function1 (input1 int) int {
    var arrayNumber [5] int

    //mengisi array
    arrayNumber[0] = 7
    arrayNumber[1] = 117

    //langsung mengisi 2 index di depan
    var arrayNumber2 [5] int = [5]int{1,2}

    //mengisi di index tertentu
    arrayNumber2 := [5]int{1: 2 ,2: 15}
}
```

## SLICE

-   Struktur data yang bisa menyimpan banyak nilai, tapi hanya bisa menyimpan 1 tipe data dan mempunyai dinamic allocation size.
-   Slice mereferensi dari Array. Sehingga, jika Array diubah, maka Slice juga berubah.
-   Slice juga bisa tanpa mereferensi ke Array.

contoh:

```go
func function1 (input1 int) int {
    var arrayNumber [5] int = [5]int{1,2,3}

    //slice mengambil index 0 sampai 1
    var sliceArr []int = arrayNumber[0:2]
}
```

-   menambah data menggunakan append()

## MAP

-   struktur data yang menyimpan key dan value
-   syntax: `var dataBulan map[int]string`
-   `var dataBulan map[int]string = [int]string{7:"juni"}`
