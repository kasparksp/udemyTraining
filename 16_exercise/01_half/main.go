package main

import "fmt"

var n int

func half(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func main() {
	fmt.Print("Enter an integer.\t")
	fmt.Scan(&n)
	fmt.Println(half(n))
}
