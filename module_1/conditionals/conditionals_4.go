package main

import "fmt"

func main() {
	var bilet int
	fmt.Scan(&bilet)

	first_half := bilet / 1000
	second_half := bilet % 1000

	first := first_half / 100
	second := first_half / 10 % 10
	third := first_half % 100 % 10
	fourth := second_half / 100
	fivth := second_half / 10 % 10
	sixth := second_half % 100 % 10

	if first + second + third == fourth + fivth + sixth {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
