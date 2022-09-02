package main

import (
	"fmt"
	"sync"
)

func work() {
	fmt.Println("work function is working from goroutine")
}

func main() {
	var wg sync.WaitGroup
	// wg := &sync.WaitGroup{}
	// wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			work()
		}()
	}
	wg.Wait()
}
