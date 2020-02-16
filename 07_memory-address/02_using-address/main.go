package main

import "fmt"

const cmToInch float64 = 0.393700787

func main() {
	var cm float64
	fmt.Print("用厘米輸入長度：\t")
	fmt.Scan(&cm)
	inch := cm * cmToInch
	fmt.Println(cm, "厘米等於", inch, "吋。")

}
