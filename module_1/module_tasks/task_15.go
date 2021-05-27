package main

import "fmt"

func main() {
	var number, digit int
	fmt.Scan(&number, &digit)
	result, rank := 0, 1
	for number != 0 {
		reminder := number % 10
		if reminder != digit {
			result += reminder * rank
			rank *= 10
		}
		number /= 10
	}
	fmt.Println(result)
}
