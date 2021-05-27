package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	if a >= 100 && a <= 999 {
		first := a / 100
		second := a / 10 % 10
		third := a % 10
		sum = first + second + third
	}
	fmt.Println(sum)
}
