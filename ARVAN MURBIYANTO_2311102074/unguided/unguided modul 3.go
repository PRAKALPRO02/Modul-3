package main

import (
	"fmt"
	"math"
)

func hitungJarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func diDalamLingkaran(cx, cy, r, x, y float64) bool {
	return hitungJarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Print("Lingkaran 1 : ")
	fmt.Scanln(&cx1, &cy1, &r1)

	fmt.Print("Lingkaran 2 : ")
	fmt.Scanln(&cx2, &cy2, &r2)

	fmt.Print("Titik sembarang : ")
	fmt.Scanln(&x, &y)

	diLingkaran1 := diDalamLingkaran(cx1, cy1, r1, x, y)
	diLingkaran2 := diDalamLingkaran(cx2, cy2, r2, x, y)

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
