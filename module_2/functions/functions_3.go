package main

import "fmt"

func vote(x int, y int, z int) int {
	if x == 0 && y == 0 {
		return 0
	} else if x == 0 && z == 0 {
		return 0
	} else if y == 0 && z == 0 {
		return 0
	}
	return 1
}

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))
}
