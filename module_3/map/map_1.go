package main

import "fmt"

func work(x int) int {
	return x + 1
}

func main() {
	var arr [10]int
	m := make(map[int]int)
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
	}
	for _, v := range arr {
		if value, ok := m[v]; ok {
			fmt.Print(value, " ")
		} else {
			m[v] = work(v)
			fmt.Print(m[v], " ")
		}
	}
}
