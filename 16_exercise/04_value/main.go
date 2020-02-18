package main

import "fmt"

func main() {
	func() {
		fmt.Println((true && false) || (false && true) || !(false && false))
	}()
}
