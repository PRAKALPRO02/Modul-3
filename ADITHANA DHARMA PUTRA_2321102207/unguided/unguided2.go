package main

import "fmt"
import "math"


func cek(titikX, titikY, cx1, cy1, r1, cx2, cy2, r2 float64) string {
    jarakLingkaran1 := math.Sqrt(math.Pow(titikX-cx1, 2) + math.Pow(titikY-cy1, 2))
    JarakLingkaran2 := math.Sqrt(math.Pow(titikX-cx2, 2) + math.Pow(titikY-cy2, 2))

    if jarakLingkaran1 <= r1 && JarakLingkaran2 <= r2 {
        return "titik didalam lingkaran 1 dan 2"
    } else if jarakLingkaran1 <= r1 {
        return "titik didalam lingkaran 1"
    } else if JarakLingkaran2 <= r2 {
        return "titik didalam lingkaran 2"
    } else {
        return "titik diluar lingkaraan 1 dan 2"
    }
}

func main() {
    var titikX, titikY, cx1, cy1, r1, cx2, cy2, r2 float64

    fmt.Print("Masukkan pusaat koordinat dan radius lingkaran 1 (x y radius): ")
    fmt.Scan(&cx1, &cy1, &r1)

    fmt.Print("Masukkan pusat koordinat dan radius lingkaran 2(x y radius): ")
    fmt.Scan(&cx2, &cy2, &r2)

    fmt.Print("masukkan titik koordinat (x y): ")
    fmt.Scan(&titikX, &titikY)

    Hasil := cek(titikX, titikY, cx1, cy1, r1, cx2, cy2, r2)
    fmt.Println(Hasil)
}