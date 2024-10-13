package main

import (
	"fmt"
	"math"
)

func distance(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}


func isInside(cx, cy, r, x, y float64) bool {
	return distance(cx, cy, x, y) <= r
}

func main() {
	
	var cx1, cy1, r1, cx2, cy2, r2 float64
	fmt.Println("Masukkan data lingkaran pertama (cx1, cy1, r1):")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Println("Masukkan data lingkaran kedua (cx2, cy2, r2):")
	fmt.Scan(&cx2, &cy2, &r2)

	
	var x, y float64
	fmt.Println("Masukkan koordinat titik (x, y):")
	fmt.Scan(&x, &y)

	
	inCircle1 := isInside(cx1, cy1, r1, x, y)
	inCircle2 := isInside(cx2, cy2, r2, x, y)

	
	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}