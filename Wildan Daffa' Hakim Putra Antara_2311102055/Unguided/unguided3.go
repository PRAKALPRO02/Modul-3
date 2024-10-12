package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func main() {
	var segiR = [2][3]float64{}
	var x, y float64
	var cekR = [2]bool{}
	for i := 0; i < 2; i++ {
		fmt.Scan(&segiR[i][0], &segiR[i][1], &segiR[i][2])
	}
	fmt.Scan(&x, &y)
	for z := 0; z < 2; z++ {
		cekR[z] = didalam(segiR[z][0], segiR[z][1], segiR[z][2], x, y)
		cekR[z] = didalam(segiR[z][0], segiR[z][1], segiR[z][2], x, y)
	}
	if cekR[0] && cekR[1] {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if cekR[0] {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if cekR[1] {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran")
	}
}
