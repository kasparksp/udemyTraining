package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("Bar:", i)
	}
}

// nothing happens because func main foo and bar runs at the same time
// main exited sooner than foo and bar run
