package main

import "fmt"

func main() {
	var N, positives int
	fmt.Scan(&N)
	slice := make([]int, N)

	for i := range slice {
		fmt.Scan(&slice[i])
		if slice[i] > 0 {
			positives++
		}
	}
	fmt.Println(positives)
}
