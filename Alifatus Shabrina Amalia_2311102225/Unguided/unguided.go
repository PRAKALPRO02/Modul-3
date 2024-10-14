package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {

	var cx1, cy1, r1 float64
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 float64
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y float64
	fmt.Scan(&x, &y)

	diDalamLingkaran1 := didalam(cx1, cy1, r1, x, y)
	diDalamLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
