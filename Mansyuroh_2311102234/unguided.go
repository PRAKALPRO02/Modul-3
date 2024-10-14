package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int
	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

	dalamLingkaran1 := titikDalamLingkaran_234(cx1, cy1, r1, x, y)
	dalamLingkaran2 := titikDalamLingkaran_234(cx2, cy2, r2, x, y)

	switch {
	case dalamLingkaran1 && dalamLingkaran2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case dalamLingkaran1:
		fmt.Println("Titik di dalam lingkaran 1")
	case dalamLingkaran2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func hitungJarak_234(cx1, cy1, cx2, cy2 int) float64 {
	x := cx1 - cx2
	y := cy1 - cy2
	return math.Sqrt(float64(x*x + y*y))
}

func titikDalamLingkaran_234(tengahX, tengahY, jariJari, x, y int) bool {
	return hitungJarak_234(tengahX, tengahY, x, y) < float64(jariJari)
}
