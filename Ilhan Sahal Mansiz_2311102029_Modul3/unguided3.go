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
	var x1, y1, r1 float64
	fmt.Scan(&x1, &y1, &r1)

	var x2, y2, r2 float64
	fmt.Scan(&x2, &y2, &r2)

	var xt, yt float64
	fmt.Scan(&xt, &yt)

	dalamLingkaran1 := didalam(x1, y1, r1, xt, yt)
	dalamLingkaran2 := didalam(x2, y2, r2, xt, yt)

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