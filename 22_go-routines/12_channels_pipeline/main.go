package main

import "fmt"

func main() {
	c := factorial(12)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	list := make(chan int)
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			list <- i
		}
		close(list)
	}()

	go func() {
		var total = 1
		for i := n; i > 0; i-- {
			total *= <-list
		}
		out <- total
		close(out)
	}()
	return out
}
