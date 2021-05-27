package main

import "fmt"

func main() {
	var n, num, zeroCounter int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&num)
		if num == 0 {
			zeroCounter++
		}
	}
	fmt.Println(zeroCounter)
}
