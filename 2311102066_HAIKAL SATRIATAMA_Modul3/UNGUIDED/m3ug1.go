package main

import (
	"fmt"
)

func checkPointPosition(cx1, cy1, r1, cx2, cy2, r2, x, y int) string {

	d1 := (x-cx1)*(x-cx1) + (y-cy1)*(y-cy1)
	d2 := (x-cx2)*(x-cx2) + (y-cy2)*(y-cy2)

	inCircle1 := d1 <= r1*r1
	inCircle2 := d2 <= r2*r2

	if inCircle1 && inCircle2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if inCircle1 {
		return "Titik di dalam lingkaran 1"
	} else if inCircle2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Scan(&x, &y)

	result := checkPointPosition(cx1, cy1, r1, cx2, cy2, r2, x, y)
	fmt.Println(result)
}
