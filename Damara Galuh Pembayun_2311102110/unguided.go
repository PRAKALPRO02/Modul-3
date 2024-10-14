package main
//Damara Galuh Pembayun 2311102110
import (
        "fmt"
        "math"
)

// Fungsi untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
        return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// Fungsi untuk memeriksa apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan jari-jari r
func didalam(cx, cy, r, x, y float64) bool {
        return jarak(cx, cy, x, y) <= r
}

func main() {
        var t int
        fmt.Scan(&t)

        for i := 0; i < t; i++ {
                var cx1, cy1, r1, cx2, cy2, r2, x, y float64
                fmt.Scan(&cx1, &cy1, &r1, &cx2, &cy2, &r2, &x, &y)

                // Cek posisi titik terhadap kedua lingkaran
                if didalam(cx1, cy1, r1, x, y) && didalam(cx2, cy2, r2, x, y) {
                        fmt.Println("Titik di dalam lingkaran 1 dan 2")
                } else if didalam(cx1, cy1, r1, x, y) {
                        fmt.Println("Titik di dalam lingkaran 1")
                } else if didalam(cx2, cy2, r2, x, y) {
                        fmt.Println("Titik di dalam lingkaran 2")
                } else {
                        fmt.Println("Titik di luar lingkaran 1 dan 2")
                }
        }
}