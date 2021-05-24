package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Printf("%v ", j)
		}
		fmt.Println()
	}
}
