package main

import "fmt"

func main() {
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			if m < 10 {
				if h < 10 {
					fmt.Println("0", h, ":", "0", m)
				} else {
					fmt.Println(h, ":", "0", m)
				}
			} else {
				if h < 10 {
					fmt.Println("0", h, ":", m)
				} else {
					fmt.Println(h, ":", m)
				}
			}
		}
	}
}
