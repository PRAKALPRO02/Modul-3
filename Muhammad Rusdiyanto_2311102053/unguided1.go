package main

import (
	"fmt"
	"math"
)

func akar(x float64) float64 {
	return math.Sqrt(x)
}

func jarak(a int, b int, c int, d int) float64 {
	return math.Pow(float64(a-c), 2) + math.Pow(float64(b-d), 2)
}

func didalam(pusat_x int, pusat_y int, r int, sembarang_x int, sembarang_y int) bool {
	var jarak_titik = akar(jarak(sembarang_x, sembarang_y, pusat_x, pusat_y))
	if jarak_titik < float64(r) {
		return true
	}
	return false
}

func main() {
	var pusat_x1, pusat_x2, pusat_y1, pusat_y2, r1, r2, sembarang_x, sembarang_y int
	var hasil1, hasil2 bool
	fmt.Print("Titik pusat (x,y) dan jejari lingkaran 1 : ")
	fmt.Scanln(&pusat_x1, &pusat_y1, &r1)
	fmt.Print("Titik pusat (x,y) dan jejari lingkaran 2 : ")
	fmt.Scanln(&pusat_x2, &pusat_y2, &r2)
	fmt.Print("Titik sembarang (x,y) : ")
	fmt.Scanln(&sembarang_x, &sembarang_y)
	hasil1 = didalam(pusat_x1, pusat_y1, r1, sembarang_x, sembarang_y)
	hasil2 = didalam(pusat_x2, pusat_y2, r2, sembarang_x, sembarang_y)
	if hasil1 && hasil2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if hasil1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if hasil2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
