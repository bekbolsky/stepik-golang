package main

import "fmt"

func main() {
	var number string
	fmt.Scan(&number)
	runes := []rune(number)
	max := string(runes[0])
	for i := 0; i < len(runes); i++ {
		if max < string(runes[i]) {
			max = string(runes[i])
		}
	}
	fmt.Println(max)
}
