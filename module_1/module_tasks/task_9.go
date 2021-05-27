package main

import "fmt"

func main() {
	var n, dr int
	fmt.Scan(&n)

	if n > 10 {
		dr = 1 + ((n - 1) % 9)
	}
	fmt.Println(dr)
}
