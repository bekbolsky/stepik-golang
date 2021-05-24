package main

import "fmt"

func main() {
	var x int
	for {
		fmt.Scan(&x)
		if x < 10 {
			continue
		} else if x > 100 {
			break
		} else {
			fmt.Println(x)
		}
	}
}
