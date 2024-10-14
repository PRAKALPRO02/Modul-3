package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) float64{
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func main(){

	var cx1, cy1, r1 int
	fmt.Println("Masukkan koordinat x, y, dan radius untuk lingkaran 1 : ")
	fmt.Scan(&cx1, &cy1, &r1)

	var cx2, cy2, r2 int
	fmt.Println("Masukkan koordinat x, y, dan radius untuk lingkaran 2 : ")
	fmt.Scan(&cx2, &cy2, &r2)

	// menginput titik sembarang 
	var x, y int
	fmt.Println("Masukkan koordinat x dan y untuk titik sembarang : ")
	fmt.Scan(&x, &y)


	// menghitung jaraj titik sembarang ke pusat lingkaran 1 dan 2 
	d1 := jarak(x, y, cx1, cy1)
	d2 := jarak(x, y, cx2, cy2)


	if d1 <= float64(r1) && d2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 <= float64(r1){
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 <= float64(r2){
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}


}