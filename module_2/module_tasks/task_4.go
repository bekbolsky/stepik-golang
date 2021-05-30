package main

import "fmt"

func main() {
	// var num string
	// fmt.Scan(&num)
	// for i := 0; i < len(num); i++ {
	// 	fmt.Print((num[i] - '0') * (num[i] - '0'))
	// }
	var num, result int
	fmt.Scan(&num)

	rank := 1
	for num != 0 {
		reminder := num % 10
		square := reminder * reminder
		result += square * rank
		if square > 9 {
			rank *= 100
		} else {
			rank *= 10
		}
		num /= 10
	}
	fmt.Println(result)
}
