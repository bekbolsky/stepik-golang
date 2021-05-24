package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	hours := 360 / 12
	fmt.Println("It is", a/hours, "hours", 2*(a%hours), "minutes.")
}
