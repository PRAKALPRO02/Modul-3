package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func diDalamLingkaran(x, y, cx, cy, r float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Println("Masukkan koordinat pusat lingkaran 1 (cx1 cy1) dan radius r1:")
	fmt.Scanf("%f %f %f", &cx1, &cy1, &r1)

	fmt.Println("Masukkan koordinat pusat lingkaran 2 (cx2 cy2) dan radius r2:")
	fmt.Scanf("%f %f %f", &cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik (x y):")
	fmt.Scanf("%f %f", &x, &y)

	dalamLingkaran1 := diDalamLingkaran(x, y, cx1, cy1, r1)
	dalamLingkaran2 := diDalamLingkaran(x, y, cx2, cy2, r2)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
