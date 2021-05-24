package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	a = a * 2 + 100

	fmt.Println(a)
}
