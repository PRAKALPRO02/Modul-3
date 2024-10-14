package main

import "fmt"

func factorial(n int) {
	fmt.Print(n, " ")
	for {
		if n%2==0 {
			n = n/2
			fmt.Print(n, " ")
		} else if n==1 {
			break
		}else {
			n = n*3+1
			fmt.Print(n, " ")
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	if n<1000000 {
		factorial(n)
	}
}