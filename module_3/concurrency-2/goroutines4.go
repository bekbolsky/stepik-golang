package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)
	sum := 0
	go func(ch chan int) {
		defer close(ch)
		for {
			select {
			case x := <-arguments:
				sum += x
			case <-done:
				output <- sum
				return
			}
		}
	}(output)

	return output
}

func main() {
	args := make(chan int)
	stop := make(chan struct{})

	result := calculator(args, stop)
	args <- 3
	args <- 5
	args <- 2
	// stop <- struct{}{}
	close(stop)

	fmt.Println(<-result)

}
