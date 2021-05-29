package main

import (
	"fmt"
	"strings"
)

func main() {
	var X, S string
	fmt.Scan(&X, &S)
	fmt.Println(strings.Index(X, S))
}
