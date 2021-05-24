package main

import "fmt"

// находит сумму двузначных чисел, кратных 8
func main() {
	var n, sum int
	fmt.Scan(&n)
	sum = 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if x < 100 && x >= 10 && x%8 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
