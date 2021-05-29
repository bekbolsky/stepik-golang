package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, another string
	fmt.Scan(&s)
	sRunes := []rune(s)
	for i := range sRunes {
		if strings.Count(s, string(sRunes[i])) > 1 {
			continue
		} else {
			another += string(sRunes[i])
		}
	}
	fmt.Println(another)
}
