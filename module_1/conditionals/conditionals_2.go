package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	first := a / 100
	second := a / 10 % 10
	third := a % 100 % 10

	if first != second && second != third && first != third {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
