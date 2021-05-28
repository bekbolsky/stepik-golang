package main

import "fmt"

func test(x1 *int, x2 *int) {
	tmp := *x1
	*x1 = *x2
	*x2 = tmp
	// *x1, *x2 = *x2, *x1
	fmt.Println(*x1, *x2)
}

func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	test(&x1, &x2)
}
