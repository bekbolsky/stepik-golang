package main

import "fmt"

func main() {
	var x, p, y, count int
	count = 0
	fmt.Scan(&x, &p, &y)
	for {
		percent := x * p / 100
		x += percent
		count++
		if x > y {
			break
		}
	}
	fmt.Println(count)
}
