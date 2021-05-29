package main

import "fmt"

func main() {
	var s, another string
	fmt.Scan(&s)
	sRunes := []rune(s)
	length := len(sRunes)
	for i := 0; i < length; i++ {
		if i%2 == 1 {
			another += string(sRunes[i])
		}
	}
	fmt.Println(another)
}
