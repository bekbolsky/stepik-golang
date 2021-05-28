package main

import "fmt"

func minimumFromFour() int {
	var a, b, c, d, min int
	fmt.Scan(&a, &b, &c, &d)
	if a < b && a < c && a < d {
		min = a
	} else if b < a && b < c && b < d {
		min = b
	} else if c < a && c < b && c < d {
		min = c
	} else {
		min = d
	}
	return min
}

func main() {
	fmt.Println(minimumFromFour())
}
