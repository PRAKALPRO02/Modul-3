package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		cx1, cy1, r1 int
		cx2, cy2, r2 int
		x, y         int
	)

	fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

	switch {
	case cekDalamLingkaran(x, y, cx1, cy1, r1) && cekDalamLingkaran(x, y, cx2, cy2, r2):
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case cekDalamLingkaran(x, y, cx1, cy1, r1):
		fmt.Println("Titik di dalam lingkaran 1")
	case cekDalamLingkaran(x, y, cx2, cy2, r2):
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func cekDalamLingkaran(x, y, cx, cy, r int) bool {
	jarak := math.Hypot(float64(x-cx), float64(y-cy))
	return jarak <= float64(r)
}
