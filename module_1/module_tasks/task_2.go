package main

import "fmt"

func main() {
	var a, rev int
	fmt.Scan(&a)
	if a >= 100 && a <= 999 {
		first := a / 100
		second := a / 10 % 10
		third := a % 10
		rev = third*100 + second*10 + first
	}
	fmt.Println(rev)
}
