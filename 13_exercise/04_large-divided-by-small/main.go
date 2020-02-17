package main

import "fmt"

func main() {
	var large int
	var small int
	fmt.Println("Please enter a large number...")
	fmt.Scan(&large)
	fmt.Println("Please enter a small number...")
	fmt.Scan(&small)
	fmt.Println("The remainder of", large, "divided by", small, "is ...")
	fmt.Println(large % small)
}
