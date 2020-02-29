package main

import "fmt"

var final = make(chan int)

func main() {
	one := count(incrementor("1"))
	two := count(incrementor("2"))

	total := <-one + <-two

	fmt.Println("Final Counter:", total)
}

func incrementor(s string) chan int {
	var c = make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println("Process: "+s+" printing:", i)
			c <- 1
		}
		close(c)
	}()
	return c
}

func count(c chan int) chan int {
	var sum int
	var m = make(chan int)
	go func() {
		for n := range c {
			sum += n
		}
		m <- sum
	}()
	return m
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
