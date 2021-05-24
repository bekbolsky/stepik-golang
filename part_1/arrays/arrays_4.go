package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)

	for i := range slice {
		fmt.Scan(&slice[i])
		if i%2 == 0 {
			fmt.Print(slice[i], " ")
		}
	}
}
