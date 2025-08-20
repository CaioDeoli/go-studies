package main

import "fmt"

func main() {
	resultado := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}

		return func() int {
			return res * res
		}
	}

	fmt.Println(resultado(1, 2, 3, 4, 5)()) // 225
}
