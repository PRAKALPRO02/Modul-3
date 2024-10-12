package main

import (
	"fmt"
)

func jarak(cx1, cy1, r1, cx2, cy2, r2, x, y int) string {
	// Menghitung jarak kuadrat titik ke pusat lingkaran 1 dan 2
	d1 := (x-cx1)*(x-cx1) + (y-cy1)*(y-cy1)
	d2 := (x-cx2)*(x-cx2) + (y-cy2)*(y-cy2)

	// Menentukan posisi titik berdasarkan jarak terhadap lingkaran 1 dan 2
	if d1 <= r1*r1 && d2 <= r2*r2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if d1 <= r1*r1 {
		return "Titik di dalam lingkaran 1"
	} else if d2 <= r2*r2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int

	// Input koordinat pusat dan radius lingkaran 1
	fmt.Println("Masukkan cx1, cy1, dan r1:")
	fmt.Scan(&cx1, &cy1, &r1)

	// Input koordinat pusat dan radius lingkaran 2
	fmt.Println
	fmt.Scan(&cx2, &cy2, &r2)

	// Input koordinat titik sembarang
	fmt.Println
	fmt.Scan(&x, &y)

	// Output hasil posisi titik
	fmt.Println(jarak(cx1, cy1, r1, cx2, cy2, r2, x, y))
}
