package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(hitungPermutasi(x, y))
}

func hitungFaktorial(m int) int {
	factorial := 1
	for j := 1; j <= m; j++ {
		factorial *= j
	}
	return factorial
}

func hitungPermutasi(a, b int) int {
	if a < b {
		a, b = b, a
	}
	return hitungFaktorial(a) / hitungFaktorial(a-b)
}
