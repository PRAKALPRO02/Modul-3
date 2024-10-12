package main
import "fmt"

// Fungsi buat menghitung faktorial 
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

// Fungsi buat menghitung permutasi
func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

// Fungsi buat menghitung kombinasi
func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	// Input 4 bilangan 
	var a, b, c, d int
	fmt.Println("Masukkan bilangan a, b, c, d: ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	// Cek syarat a >= c dan b >= d
	if a >= c && b >= d {
		// Menghitung permutasi dan kombinasi a dan c
		permutasiAC := permutation(a, c)
		kombinasiAC := combination(a, c)

		// Menghitung permutasi dan kombinasi b dan d
		permutasiBD := permutation(b, d)
		kombinasiBD := combination(b, d)

		// Output hasil
		fmt.Println("Permutasi(a, c) dan Kombinasi(a, c):", permutasiAC, kombinasiAC)
		fmt.Println("Permutasi(b, d) dan Kombinasi(b, d):", permutasiBD, kombinasiBD)
	} else {
		fmt.Println("Syarat a >= c dan b >= d tidak terpenuhi.")
	}
}