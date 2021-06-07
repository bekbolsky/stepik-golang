package main

import (
	"fmt"
	"strings"
)

type Battery struct {
	charge string
}

func (b *Battery) String() string {
	zeros := strings.Count(b.charge, "0")
	ones := strings.Count(b.charge, "1")
	result := "[" + strings.Repeat(" ", zeros) + strings.Repeat("X", ones) + "]"
	return result
}

func main() {
	var s string
	fmt.Scan(&s)
	batteryForTest := &Battery{charge: s}
	fmt.Println(batteryForTest)
}
