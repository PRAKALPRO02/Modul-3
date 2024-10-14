package main

import (
    "fmt"
    "math"
)

// Fungsi jarak menghitung jarak antara dua titik (a,b) dan (c,d) dalam koordinat kartesian
// menggunakan rumus jarak Euclidean: √((a-c)² + (b-d)²)
func jarak(a, b, c, d float64) float64 {
    return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

// Fungsi didalam memeriksa apakah suatu titik (x,y) berada di dalam lingkaran
// dengan pusat (cx,cy) dan radius r
func didalam(cx, cy, r, x, y float64) bool {
    return jarak(cx, cy, x, y) <= r
}

func main() {
    // Deklarasi variabel untuk koordinat dan radius dua lingkaran, serta titik yang akan diperiksa
    var cx1, cy1, r1, cx2, cy2, r2, x, y float64

    // Meminta input dari pengguna untuk lingkaran 1
    fmt.Println("Masukkan koordinat titik pusat lingkaran 1 (cx, cy) dan radius (r):")
    fmt.Scanln(&cx1, &cy1, &r1)

    // Meminta input dari pengguna untuk lingkaran 2
    fmt.Println("Masukkan koordinat titik pusat lingkaran 2 (cx, cy) dan radius (r):")
    fmt.Scanln(&cx2, &cy2, &r2)

    // Meminta input koordinat titik yang akan diperiksa
    fmt.Println("Masukkan koordinat titik (x, y):")
    fmt.Scanln(&x, &y)

    // Memeriksa posisi titik terhadap kedua lingkaran dan mencetak hasilnya
    if didalam(cx1, cy1, r1, x, y) && didalam(cx2, cy2, r2, x, y) {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if didalam(cx1, cy1, r1, x, y) {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if didalam(cx2, cy2, r2, x, y) {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran ")
    }
}