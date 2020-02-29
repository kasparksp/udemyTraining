package main

import "fmt"

var c = make(chan int)

func main() {
	f := <-factorial(100)
	fmt.Println("Total:", f)
}

func factorial(n int) chan int {
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
		}
		c <- total
		close(c)
	}()
	return c
}
