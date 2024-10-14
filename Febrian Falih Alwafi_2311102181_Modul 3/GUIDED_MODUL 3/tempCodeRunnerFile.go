package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	
	var a, b, c, d int
	fmt.Println("Masukkan bilangan a, b, c, d (dengan spasi): ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	if a >= c && b >= d {

		permutasiAC := permutation(a, c)
		kombinasiAC := combination(a, c)

		permutasiBD := permutation(b, d)
		kombinasiBD := combination(b, d)

		fmt.Println("Permutasi(a, c) dan Kombinasi(a, c):", permutasiAC, kombinasiAC)
		fmt.Println("Permutasi(b, d) dan Kombinasi(b, d):", permutasiBD, kombinasiBD)
	} else {
		fmt.Println("Syarat a >= c dan b >= d tidak terpenuhi.")
	}
}
