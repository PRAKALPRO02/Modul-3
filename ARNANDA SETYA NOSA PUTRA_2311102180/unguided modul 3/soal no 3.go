package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	if cekDalamLingkaran(x, y, cx1, cy1, r1) && cekDalamLingkaran(x, y, cx2, cy2, r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cekDalamLingkaran(x, y, cx1, cy1, r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cekDalamLingkaran(x, y, cx2, cy2, r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func cekDalamLingkaran(x, y, cx, cy, r int) bool {
	var jarak float64
	jarak = math.Sqrt(float64((x-cx)*(x-cx) + (y-cy)*(y-cy)))
	return jarak <= float64(r)
}
