package main

import "fmt"

func main() {
	for i := 1; i == 1; {

		fmt.Println("輸入 第一個數字 除以 第二個數字")

		fmt.Print("第一個數字：\t")
		var first int
		fmt.Scan(&first)

		fmt.Print("第二個數字：\t")
		var second int
		fmt.Scan(&second)

		fmt.Println("")
		fmt.Print(first, "除以", second, "是：\t")

		q := first / second
		r := first % second

		fmt.Println(q, "餘", r)

		fmt.Println("")
	}
}
