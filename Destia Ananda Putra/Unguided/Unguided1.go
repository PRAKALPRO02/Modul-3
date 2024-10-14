package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func di_dalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah lingkaran: ")
	fmt.Scan(&n)

	// Menggunakan slice untuk menyimpan data lingkaran
	circles := make([]struct {
		cx, cy, r float64
	}, n)

	// Input data lingkaran
	fmt.Println("Masukkan data lingkaran:")
	for i := 0; i < n; i++ {
		fmt.Printf("Lingkaran %d (cx, cy, r): ", i+1)
		fmt.Scan(&circles[i].cx, &circles[i].cy, &circles[i].r)
	}

	// Input koordinat titik
	var x, y float64
	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	// Mengecek apakah titik berada di dalam masing-masing lingkaran
	dalam_Lingkaran := make([]bool, n)
	for i := 0; i < n; i++ {
		dalam_Lingkaran[i] = di_dalam(circles[i].cx, circles[i].cy, circles[i].r, x, y)
	}

	// Menampilkan hasil
	result := ""
	for i := 0; i < n; i++ {
		if dalam_Lingkaran[i] {
			result += fmt.Sprintf("Titik di dalam lingkaran %d ", i+1)
		}
	}

	if result == "" {
		fmt.Println("Titik di luar semua lingkaran")
	} else {
		fmt.Println(result)
	}
}
