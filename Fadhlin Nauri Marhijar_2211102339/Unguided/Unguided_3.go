package main

import (
	"fmt"
	"math"
)

type Lingkaran struct {
	pusatX   float64
	pusatY   float64
	jariJari float64
}

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func (l Lingkaran) berisiTitik(x, y float64) bool {
	return jarak(l.pusatX, l.pusatY, x, y) <= l.jariJari
}

func posisiTitik(l1, l2 Lingkaran, x, y float64) string {
	dalamL1 := l1.berisiTitik(x, y)
	dalamL2 := l2.berisiTitik(x, y)

	if dalamL1 && dalamL2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if dalamL1 {
		return "Titik di dalam lingkaran 1"
	} else if dalamL2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	lingkaran1 := Lingkaran{1, 2, 3}
	lingkaran2 := Lingkaran{4, 5, 6}
	titik := Lingkaran{7, 8, 0}

	hasil := posisiTitik(lingkaran1, lingkaran2, titik.pusatX, titik.pusatY)
	fmt.Println(hasil)
}
