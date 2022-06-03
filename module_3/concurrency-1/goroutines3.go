package main

import (
	"fmt"
	// "time"
)

// removeDuplicates accepts two channels - inputStream and outputStream,
// on the first one it receive strings, on the second it send values without repeats.
// As a result, the outputStream should contain values that are not repeated in a row.
func removeDuplicates(inputStream, outputStream chan string) {
	var previousValue string
	for v := range inputStream {
		if previousValue != v {
			outputStream <- v
			previousValue = v
		}
	}
	close(outputStream)
}

// func printer(c chan string) {
// 	for {
// 		msg := <-c
// 		fmt.Println(msg)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)
	// go printer(outputStream)
	var input string
	fmt.Printf("Enter some text with duplicates: ")
	fmt.Scanln(&input)

	go func() {
		defer close(inputStream)

		for _, v := range input {
			inputStream <- string(v)
		}
	}()

	fmt.Printf("Output without duplicates: ")
	for v := range outputStream {
		fmt.Printf("%s", v)
	}
	fmt.Printf("\n")

}
