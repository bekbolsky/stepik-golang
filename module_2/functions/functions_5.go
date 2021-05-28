package main

import "fmt"

func sumInt(nums ...int) (count, sum int) {
	// var count, sum int
	for _, value := range nums {
		sum += value
		count++
	}
	return count, sum
}

func main() {
	a, b := sumInt(1, 0, 3)
	fmt.Println(a, b)
}
