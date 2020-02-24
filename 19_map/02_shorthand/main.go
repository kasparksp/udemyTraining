package main

import "fmt"

func main() {
	hello := map[int]string{
		1: "Hi",
		2: "Hell~",
		3: "Hey",
		4: "Bello",
		5: "Hello"}

	for key, val := range hello {
		fmt.Println(key, "-", val)
	}
}
