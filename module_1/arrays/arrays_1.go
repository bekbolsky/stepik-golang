package main

import "fmt"

func main() {
	var workArray [10]uint8

	for i := 0; i < 10; i++ {
		var a uint8
		fmt.Scan(&a)
		workArray[i] = a
	}
	for j := 0; j < 3; j++ {
		var i1, i2 uint8
		fmt.Scan(&i1, &i2)
		workArray[i1], workArray[i2] = workArray[i2], workArray[i1]
	}
	for _, value := range workArray {
		fmt.Print(value, " ")
	}
}
