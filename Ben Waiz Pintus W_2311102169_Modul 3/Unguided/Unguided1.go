package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func isInsideCircle(x, y, cx, cy, r int) bool {
	return distance(x, y, cx, cy) <= float64(r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int
	fmt.Print("Masukkan angka lingkaran pertama (cx1, cy1, r1) : ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Print("Masukkan angka lingkaran kedua (cx2, cy2, r2) : ")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Print("Masukkan angka titik koordinat sembarang (x, y) : ")
	fmt.Scan(&x, &y)

	isInCircle1 := isInsideCircle(x, y, cx1, cy1, r1)
	isInCircle2 := isInsideCircle(x, y, cx2, cy2, r2)

	if isInCircle1 && isInCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if isInCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if isInCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
