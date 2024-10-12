package main

import (
	"fmt"
)

func main() {
	var x1, y1, radius1, x2, y2, radius2, px, py int
	fmt.Println("Masukkan koordinat dan radius lingkaran 1, lingkaran 2, serta titik x, y:")
	fmt.Scan(&x1, &y1, &radius1, &x2, &y2, &radius2, &px, &py)

	hasilCekLingkaran(px, py, x1, y1, radius1, x2, y2, radius2)
}
func apakahDalamLingkaran(px, py, pusatX, pusatY, radius int) bool {
	jarakSquared := (px-pusatX)*(px-pusatX) + (py-pusatY)*(py-pusatY)
	return jarakSquared <= radius*radius
}
func hasilCekLingkaran(px, py, cx1, cy1, r1, cx2, cy2, r2 int) {
	diLingkaran1 := apakahDalamLingkaran(px, py, cx1, cy1, r1)
	diLingkaran2 := apakahDalamLingkaran(px, py, cx2, cy2, r2)

	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik berada di dalam kedua lingkaran")
	} else if diLingkaran1 {
		fmt.Println("Titik berada di dalam lingkaran pertama")
	} else if diLingkaran2 {
		fmt.Println("Titik berada di dalam lingkaran kedua")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran")
	}
}
