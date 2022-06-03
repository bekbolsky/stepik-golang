package main

import (
	"fmt"
)

func task(ch chan int, N int) {
	ch <- N + 1
}

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go task(ch, i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
