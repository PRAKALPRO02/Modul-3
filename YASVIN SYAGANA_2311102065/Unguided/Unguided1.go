package main

import (
	"fmt"
	"math"
)

func main() {
	var cx, cxx, cy, cyy, r, rr, x, y float64

	fmt.Println("Lingkaran 1")
	fmt.Print("(cx)(cy)(r) : ")
	fmt.Scanln(&cx, &cy, &r)

	fmt.Println("Lingkaran 2")
	fmt.Print("(cx)(cy)(r) : ")
	fmt.Scanln(&cxx, &cyy, &rr)

	fmt.Print("Input koordinat sembarang (x)(y) : ")
	fmt.Scanln(&x, &y)

	//mari kita cek posisi Lingkaran1 dang Lingkaran2
	posisi1 := didalam(cx, cy, x, y, r)
	posisi2 := didalam(cxx, cyy, x, y, rr)

	if posisi1 && posisi2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if posisi1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if posisi2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}

}

func jarak(a, b, c, d float64) float64 {
	//pengecekan Jarak
	kiri := (a - c) * (a - c)
	kanan := (b - d) * (b - d)
	return math.Sqrt(kiri + kanan)
	//math.Sqrt digunakan untuk pemanggilan akar
}

func didalam(cx, cy, x, y, r float64) bool {
	//pengecekan apakah x, y berada dalam lingkaran
	return jarak(cx, cy, x, y) <= r
}
