package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	sArr := strings.Split(s, "")
	result := strings.Join(sArr, "*")
	fmt.Println(result)
}
