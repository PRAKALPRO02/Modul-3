// Febrian Falih Alwafi
// 2311102181
//S1F-11-02

package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung jarak antara dua titik (a, b) dan (c, d)
func hitungJarak(x1, y1, x2, y2 float64) float64 {
    return math.Hypot(x1-x2, y1-y2) // Menggunakan fungsi math.Hypot untuk menghitung jarak Euclidean
}

// Fungsi untuk memeriksa apakah titik (x, y) berada di dalam lingkaran dengan pusat (cx, cy) dan radius r
func diDalamLingkaran(cx, cy, r, x, y float64) bool {
    return hitungJarak(cx, cy, x, y) <= r
}

func main() {
    // Array untuk menyimpan data dua lingkaran
    var lingkaran [2]struct {
        cx, cy, r float64
    }

    fmt.Println("Masukkan data dua lingkaran (koordinat pusat dan radius):")
    for i := 0; i < 2; i++ {
        fmt.Printf("Lingkaran %d - masukkan (cx, cy, r): ", i+1)
        fmt.Scan(&lingkaran[i].cx, &lingkaran[i].cy, &lingkaran[i].r)
    }

    // Meminta input koordinat titik
    var x, y float64
    fmt.Print("Masukkan koordinat titik (x, y): ")
    fmt.Scan(&x, &y)

    // Memeriksa apakah titik berada di dalam salah satu atau kedua lingkaran
    var beradaDiLingkaran [2]bool
    for i := 0; i < 2; i++ {
        beradaDiLingkaran[i] = diDalamLingkaran(lingkaran[i].cx, lingkaran[i].cy, lingkaran[i].r, x, y)
    }

    // Menampilkan hasil
    switch {
    case beradaDiLingkaran[0] && beradaDiLingkaran[1]:
        fmt.Println("Titik berada di dalam lingkaran 1 dan 2.")
    case beradaDiLingkaran[0]:
        fmt.Println("Titik berada di dalam lingkaran 1.")
    case beradaDiLingkaran[1]:
        fmt.Println("Titik berada di dalam lingkaran 2.")
    default:
        fmt.Println("Titik berada di luar lingkaran 1 dan 2.")
    }
}


