package main

import "fmt"

func main() {
	var hello = make(map[int]string)

	hello[1] = "Hi"
	hello[2] = "Hell~"
	hello[3] = "Hey"
	hello[4] = "Bello"
	hello[5] = "Hello"

	for key, val := range hello {
		fmt.Println(key, "-", val)
	}
}
