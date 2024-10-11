package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y int

    // Input data lingkaran pertama
	fmt.Scan(&cx1, &cy1, &r1) // 2311102002

    // Input data lingkaran kedua
	fmt.Scan(&cx2, &cy2, &r2) // 2311102002

	fmt.Scan(&x, &y) // 2311102002

	distance1 := calcDistance(x, y, cx1, cy1)
	distance2 := calcDistance(x, y, cx2, cy2)

	checkPosition(distance1, distance2, r1, r2)
}


func calcDistance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func checkPosition(d1, d2 float64, r1, r2 int) {
	if d1 <= float64(r1) && d2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 <= float64(r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
} // 2311102002