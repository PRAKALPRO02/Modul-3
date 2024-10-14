package main

import (
	"fmt"
	"math"
)

// Menghitung jarak antara dua titik
func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

// Memeriksa apakah titik (x, y) berada di dalam lingkaran dengan pusat (dx, dy) dan jari-jari r
func isInsideCircle(x, y, cx, cy, r int) bool {
	return distance(x, y, cx, cy) <= float64(r)
}

func main() {
	var dx1, dy1, r1 int
	var dx2, dy2, r2 int
	var x, y int

	fmt.Println("Masukkan angka lingkaran pertama (dx1, dy1, r1):")
	fmt.Scan(&dx1, &dy1, &r1)

	fmt.Print("Masukkan angka lingkaran kedua (dx2, dy2, r2):")
	fmt.Scan(&dx2, &dy2, &r2)

	fmt.Print("Masukkan angka titik koordinat sembarang (x, y):")
	fmt.Scan(&x, &y)

	isInCircle1 := isInsideCircle(x, y, dx1, dy1, r1)
	isInCircle2 := isInsideCircle(x, y, dx2, dy2, r2)

	if isInCircle1 && isInCircle2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if isInCircle1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if isInCircle2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}
}
