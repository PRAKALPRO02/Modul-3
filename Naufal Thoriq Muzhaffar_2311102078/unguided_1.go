package main
import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func di_dalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var circles [2]struct {
		cx, cy, r float64
	}

	fmt.Println("Masukkan data lingkaran:")
	for i := 0; i < 2; i++ {
		fmt.Printf("Lingkaran %d (cx, cy, r): ", i+1)
		fmt.Scan(&circles[i].cx, &circles[i].cy, &circles[i].r)
	}

	var x, y float64
	fmt.Print("Masukkan koordinat titik (x, y): ")
	fmt.Scan(&x, &y)

	dalam_Lingkaran := [2]bool{}
	for i := 0; i < 2; i++ {
		dalam_Lingkaran[i] = di_dalam(circles[i].cx, circles[i].cy, circles[i].r, x, y)
	}

	result := ""
	for i := 0; i < 2; i++ {
		if dalam_Lingkaran[i] {
			result += fmt.Sprintf("Titik di dalam lingkaran %d ", i+1)
		}
	}
	if result == "" {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	} else {
		fmt.Println(result)
	}
}
