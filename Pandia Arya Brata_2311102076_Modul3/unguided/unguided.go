package main

import (
	"fmt"
	"math"
)

func priksa_titik(cx1, cx2, r1, cy1, cy2, r2, x, y float64) string {
	titik1 := math.Sqrt(math.Pow(x-cx1, 2) + math.Pow(y-cy1, 2))
	titik2 := math.Sqrt(math.Pow(x-cx2, 2) + math.Pow(y-cy2, 2))

	if titik1 <= r1 && titik2 <= r2 {
		return ("Titik di dalam lingkaran 1 dan 2")
	} else if titik1 <= r1 {
		return ("Titik di dalam lingkaran 1")
	} else if titik2 <= r2 {
		return ("Titik di dalam lingkaran 2")
	} else {
		return ("Titik di luar lingkaran 1 dan 2")
	}
}

func main() {
	var x, y, cx1, cx2, r1, cy1, cy2, r2 float64

	fmt.Print("Masukan kordinat titik pusat dan radius lingkaran : ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Masukan kordinat titik pusat dan radius lingkaran : ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukan kordinat titik x dan y : ")
	fmt.Scan(&x, &y)

	hasil := priksa_titik(x, y, cx1, cy1, r1, cx2, cy2, r2)
	fmt.Print(hasil)
}
