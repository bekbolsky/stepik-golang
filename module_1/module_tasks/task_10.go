package main

import "fmt"

func main() {
	var a, b, res, count int
	fmt.Scan(&a, &b)

	for i := b; i >= a; i-- {
		if i%7 == 0 {
			res = i
			fmt.Println(res)
			count++
			break
		} else {
			continue
		}
	}
	if count == 0 {
		fmt.Println("NO")
	}
}
