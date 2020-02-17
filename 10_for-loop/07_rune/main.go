package main

import "fmt"

func main() {
	for i := 1; i < 256; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
