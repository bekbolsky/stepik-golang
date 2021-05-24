package main

import "fmt"

func main() {
	var x, count, max int
	max = 0

	for fmt.Scan(&x); x != 0; fmt.Scan(&x) {
		if x > max {
			max = x
			count = 1
		} else if x == max {
			count++
		}
	}
	fmt.Println(count)
}
