package main

import (
	"fmt"
	"math"
)

// Struct untuk mendefinisikan sebuah Lingkaran
type Lingkaran struct {
	x, y, radius float64
}

// Fungsi untuk menghitung jarak antara dua titik (x1, y1) dan (x2, y2)
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

// Fungsi untuk menentukan apakah titik berada di dalam, di luar, atau di tepi lingkaran
func posisiTitik(l Lingkaran, x, y float64) string {
	dist := jarak(l.x, l.y, x, y)
	if dist < l.radius {
		return "di dalam"
	} else if dist == l.radius {
		return "di tepi"
	} else {
		return "di luar"
	}
}

// Fungsi utama untuk memproses titik terhadap beberapa lingkaran
func main() {
	// Mendefinisikan dua lingkaran
	lingkaran1 := Lingkaran{0, 0, 5}
	lingkaran2 := Lingkaran{2, 2, 3}

	// Memasukkan titik
	var x, y float64
	fmt.Print("Masukkan koordinat titik (x y): ")
	fmt.Scanf("%f %f", &x, &y)

	// Memeriksa posisi titik terhadap kedua lingkaran
	hasil1 := posisiTitik(lingkaran1, x, y)
	hasil2 := posisiTitik(lingkaran2, x, y)

	// Menampilkan hasilnya
	if hasil1 == "di dalam" && hasil2 == "di dalam" {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if hasil1 == "di dalam" {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if hasil2 == "di dalam" {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
