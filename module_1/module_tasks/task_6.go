package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	sum := a + b
	if sum%2 == 0 {
		mean := sum / 2
		fmt.Println(mean)
	} else {
		mean := float64(sum) / 2
		fmt.Println(mean)
	}
}
