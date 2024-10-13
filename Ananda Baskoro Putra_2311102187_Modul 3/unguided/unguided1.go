package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt((x1 - x2) * (x1 - x2) + (y1 - y2) * (y1 - y2))
}

func diDalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Println("Masukkan koordinat pusat dan jari-jari lingkaran 1:")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Println("Masukkan koordinat pusat dan jari-jari lingkaran 2:")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Println("Masukkan koordinat titik sembarang:")
	fmt.Scan(&x, &y)

	diLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	diLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}
}