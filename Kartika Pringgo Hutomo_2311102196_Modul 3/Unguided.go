package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func dalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64

	fmt.Print("Masukkan pusat dan radius lingkaran 1 (cx1, cy1, r1): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukkan pusat dan radius lingkaran 2 (cx2, cy2, r2): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	dalamL1 := dalamLingkaran(cx1, cy1, r1, x, y)
	dalamL2 := dalamLingkaran(cx2, cy2, r2, x, y)

	if dalamL1 && dalamL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
