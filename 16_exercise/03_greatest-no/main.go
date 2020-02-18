package main

import "fmt"

func test(n []float64) {
	var output float64
	for _, i := range n {
		if i > output {
			output = i
		} else {
			continue
		}
	}
	fmt.Println("The greatest number among", n, "is", output)
}

func main() {
	data := []float64{1, 9, 20, 888, 2, 5, 999, 0}
	test(data)

}
