package main

import "fmt"

func main() {
	x := func(fn func(i int) int, i int) func(int) int {
		return fn
	}(func(i int) int {
		return i + 1
	}, 5)
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x(10))
}
