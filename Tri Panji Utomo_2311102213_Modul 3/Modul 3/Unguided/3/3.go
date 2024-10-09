package main

import (
    "fmt"
    "math"
)

func hitungJarak(x1213, y1213, x2213, y2 float64) float64 {
    return math.Sqrt(math.Pow(x2213-x1213, 2) + math.Pow(y2-y1213, 2))
}

func diDalamLingkaran(cx, cy, r, x, y float64) bool {
    return hitungJarak(cx, cy, x, y) <= r
}

func main() {
    var x1213, y1213, r1, x2213, y2, r2, xt213, yt213 float64

    fmt.Print("lingkaran 1 : ")
    fmt.Scanln(&x1213, &y1213, &r1)

    fmt.Print("lingkaran 2 : ")
    fmt.Scanln(&x2213, &y2, &r2)

    fmt.Print("titik sembarang : ")
    fmt.Scanln(&xt213, &yt213)

    diLingkaran1 := diDalamLingkaran(x1213, y1213, r1, xt213, yt213)
    diLingkaran2 := diDalamLingkaran(x2213, y2, r2, xt213, yt213)

    if diLingkaran1 && diLingkaran2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if diLingkaran1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if diLingkaran2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
