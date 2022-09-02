package main

import (
	"fmt"
)

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	output := make(chan int)
	go func(ch chan int) {
		defer close(ch)

		select {
		case x := <-firstChan:
			ch <- x * x
		case x := <-secondChan:
			ch <- x * 3
		case <-stopChan:
		}
	}(output)

	return output
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	stop := make(chan struct{})

	result := calculator(ch1, ch2, stop)
	
	ch1 <- 3
	// ch2 <- 4
	// stop <- struct{}{}
	// close(stop)

	fmt.Println(<-result)
}
