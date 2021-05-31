package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func CleanFromTrash(s string) int64 {
	sArr := []rune(s)
	for _, r := range sArr {
		if !unicode.IsDigit(r) {
			s = strings.ReplaceAll(s, string(r), "")
		}
	}
	result, _ := strconv.ParseInt(s, 10, 64)
	return result
}

func adding(s1, s2 string) int64 {
	first := CleanFromTrash(s1)
	second := CleanFromTrash(s2)

	return first + second
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	fmt.Println(adding(a, b))
}
