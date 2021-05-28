package main

import "fmt"

func test(x1 *int, x2 *int) {
	res := *x1 * *x2
	fmt.Println(res)
}

func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	test(&x1, &x2)
}
