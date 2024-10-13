package main

import (
	"fmt"
	"math"
)

func hitungjarak(x, y, cx, cy float64) float64 {
	hasil := math.Sqrt(math.Pow(x-cx, 2) + math.Pow(y-cy, 2))
	return hasil
}

func posisi(x, y, cx, cy, r float64) bool {
	hasil := hitungjarak(x, y, cx, cy) <= r
	return hasil
}

func main() {
	var x, y, cx1, cy1, cx2, cy2, r1, r2 float64

	fmt.Print("Lingkaran 1 : ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Print("Lingkaran 2 : ")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Print("Titik : ")
	fmt.Scan(&x, &y)

	if posisi(x, y, cx1, cy1, r1) == true && posisi(x, y, cx2, cy2, r2) == true {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if posisi(x, y, cx1, cy1, r1) == true {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if posisi(x, y, cx2, cy2, r2) == true {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}

}
