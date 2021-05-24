package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	slice := make([]int, N)

	for i := 0; i < N; i++ {
		var number int
		fmt.Scan(&number)
		slice[i] = number
	}
	fmt.Println(slice[3])
}
