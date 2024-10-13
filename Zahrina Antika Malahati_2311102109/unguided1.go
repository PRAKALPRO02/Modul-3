package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi untuk mengecek apakah titik (x, y) berada di dalam lingkaran
func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	// Masukan untuk dua lingkaran
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	// Input dari pengguna
	fmt.Println("Masukkan pusat dan radius lingkaran 1 (cx1, cy1, r1):")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Println("Masukkan pusat dan radius lingkaran 2 (cx2, cy2, r2):")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Println("Masukkan koordinat titik sembarang (x, y):")
	fmt.Scan(&x, &y)

	// Mengecek posisi titik (x, y) relatif terhadap lingkaran 1 dan 2
	diLingkaran1 := didalam(cx1, cy1, r1, x, y)
	diLingkaran2 := didalam(cx2, cy2, r2, x, y)

	// Menentukan output berdasarkan posisi titik
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

//Zahrina Antika Malahati_2311102109
