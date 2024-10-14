package main

import (
	"fmt"
	"math"
)

type Circle struct {
	cx, cy, r float64
}

func isPointInside(c Circle, x, y float64) bool {
	distance := math.Sqrt(math.Pow(x-c.cx, 2) + math.Pow(y-c.cy, 2))
	return distance <= c.r
}

func determinePosition(c1, c2 Circle, x, y float64) string {
	in1 := isPointInside(c1, x, y)
	in2 := isPointInside(c2, x, y)

	if in1 && in2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if in1 {
		return "Titik di dalam lingkaran 1"
	} else if in2 {
		return "Titik di dalam lingkaran 2"
	}
	return "Titik di luar lingkaran 1 dan 2"
}

func main() {
	var c1, c2 Circle
	var x, y float64

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&c1.cx, &c1.cy, &c1.r)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&c2.cx, &c2.cy, &c2.r)

	fmt.Println("Masukkan koordinat titik yang akan dicek (x y):")
	fmt.Scan(&x, &y)

	result := determinePosition(c1, c2, x, y)
	fmt.Println(result)
}
