package main

import "fmt"

func HoFoG(n int) int {
	g := n - 2
	f := g * g
	h := f + 1
	return h
}

func GoHoF(n int) int {
	f := n * n
	h := f + 1
	g := h - 2
	return g
}

func FoGoH(n int) int {
	h := n + 1
	g := h - 2
	f := g * g
	return f
}

func main() {
	var x, y, z, hasilx, hasily, hasilz int

	fmt.Print("masukan nilai x, y dan z : ")
	fmt.Scan(&x, &y, &z)
	hasilx = FoGoH(x)
	hasily = GoHoF(y)
	hasilz = HoFoG(z)

	fmt.Println("fogoh : ", hasilx)
	fmt.Println("gohof : ", hasily)
	fmt.Println("hofog : ", hasilz)
}
