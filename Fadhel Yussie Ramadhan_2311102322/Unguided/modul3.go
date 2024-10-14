package main

import "fmt"
import "math"

func jarak(a float64, b float64, c float64, d float64) float64 {
	return math.Sqrt(math.Pow((a-c), 2) + math.Pow((b-d), 2))
}

func didalam(cx float64,cy float64,r float64,x float64,y float64) bool {
	var jarak_titik = jarak(x, y, cx, cy)
	if jarak_titik < r {
		return true
	} else {
		return false
	}
}

func main() {
	var cx1, cx3, cy1, cy2, r1, r2, x, y float64
	var result1, result2 bool
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx3, &cy2, &r2)
	fmt.Scanln(&x, &y)
	result1 = didalam(cx1, cy1, r1, x, y)
	result2 = didalam(cx3, cy2, r2, x, y)
	if result1 && result2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if result1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if result2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}