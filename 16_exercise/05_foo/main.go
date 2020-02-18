package main

import "fmt"

func foo(n ...int) {
	fmt.Println("START")
	for _, i := range n {
		fmt.Println(i)
	}
	fmt.Println("END")
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
