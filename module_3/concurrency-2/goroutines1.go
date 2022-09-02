package main

import (
	"fmt"
)

func work() {
	fmt.Println("work function is working from goroutine")
}

func main() {
	done := make(chan struct{})
	go func() {
		work()
		// done <- struct{}{}
		close(done)
	}()
	<-done
}
