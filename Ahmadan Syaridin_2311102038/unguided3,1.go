package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1 int
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Scan(&cx2, &cy2, &r2)

	var x, y int
	fmt.Scan(&x, &y) //2311102038

	distance1 := math.Sqrt(float64((x-cx1)*(x-cx1) + (y-cy1)*(y-cy1)))
	distance2 := math.Sqrt(float64((x-cx2)*(x-cx2) + (y-cy2)*(y-cy2)))
     //2311102038
	if distance1 <= float64(r1) && distance2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if distance1 <= float64(r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if distance2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else { //2311102038
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
} //2311102038
