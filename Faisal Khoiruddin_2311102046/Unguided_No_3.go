package main

import (
	"fmt"
	"math"
)

func jarak(cx, cy, x, y float64) float64 {
	return math.Sqrt((x-cx)*(x-cx) + (y-cy)*(y-cy))
}

func beradaDalamLingkaran(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Println("Masukkan koordinat titik pusat dan radius dari lingkaran 1 : ")
	_, err := fmt.Scanf("%f %f %f\n", &cx1, &cy1, &r1)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input lingkaran 1: ", err)
		return
	}
	fmt.Println("Masukkan koordinat titik pusat dan radius dari lingkaran 2 : ")
	_, err = fmt.Scanf("%f %f %f\n", &cx2, &cy2, &r2)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input lingkaran 2: ", err)
		return
	}
	fmt.Println("Masukkan koordinat titik sembarang: ")
	_, err = fmt.Scanf("%f %f\n", &x, &y)
	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input titik sembarang: ", err)
		return
	}

	diDalamLingkaran1 := beradaDalamLingkaran(cx1, cy1, r1, x, y)
	diDalamLingkaran2 := beradaDalamLingkaran(cx2, cy2, r2, x, y)

	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
