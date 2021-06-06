package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(num uint) uint {
		var s string
		a := strconv.FormatUint(uint64(num), 10)
		for _, v := range a {
			if (v-'0')%2 == 0 && (v-'0') != 0 {
				s += string(v)
			}
		}
		result, _ := strconv.ParseUint(s, 10, 64)
		if result == 0 {
			return uint(100)
		}
		return uint(result)
	}

	fmt.Println(fn(727178))
	fmt.Println(fn(1234))
	fmt.Println(fn(301))
	fmt.Println(fn(210))
	fmt.Println(fn(200))
	fmt.Println(fn(000))
}
