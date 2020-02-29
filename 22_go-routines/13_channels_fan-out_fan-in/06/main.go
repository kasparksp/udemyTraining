package main

import (
	"fmt"
)

func main() {

	in := gen()

	f := factorial(in)

	for i := 0; i < 1000; i++ {
		fmt.Println(<-f)
	}
}

func gen() <-chan float64 {
	out := make(chan float64)
	go func() {
		for i := 1; i <= 1000; i++ {
			out <- float64(i)
		}
		close(out)
	}()
	return out
}

func factorial(in <-chan float64) <-chan float64 {
	out := make(chan float64)
	for i := 0; i < 1000; i++ {
		go func() {
			for n := range in {
				out <- fact(n)
			}
		}()
	}
	return out
}

func fact(n float64) float64 {
	var total float64 = 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/
