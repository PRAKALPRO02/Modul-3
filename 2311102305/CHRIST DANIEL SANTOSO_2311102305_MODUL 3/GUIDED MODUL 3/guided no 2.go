package main

import "fmt"

func hitungFaktorial(x int) int {
	if x == 0 || x == 1 {
		return 1
	}
	return x * hitungFaktorial(x-1)
}
func hitungPermutasi(x, y int) int {
	return hitungFaktorial(x) / hitungFaktorial(x-y)
}
func hitungKombinasi(x, y int) int {
	return hitungFaktorial(x) / (hitungFaktorial(y) * hitungFaktorial(x-y))
}
func main() {
	var p, q, r, s int
	fmt.Println("Masukkan empat bilangan (pisahkan dengan spasi): ")
	fmt.Scanf("%d %d %d %d", &p, &q, &r, &s)
	if p >= r && q >= s {
		perm1 := hitungPermutasi(p, r)
		komb1 := hitungKombinasi(p, r)

		perm2 := hitungPermutasi(q, s)
		komb2 := hitungKombinasi(q, s)

		fmt.Printf("Permutasi(%d, %d) dan Kombinasi(%d, %d): %d, %d\n", p, r, p, r, perm1, komb1)
		fmt.Printf("Permutasi(%d, %d) dan Kombinasi(%d, %d): %d, %d\n", q, s, q, s, perm2, komb2)
	} else {
		fmt.Println("Kondisi p >= r dan q >= s tidak dipenuhi.")
	}
}
