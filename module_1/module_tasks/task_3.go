package main

import "fmt"

const (
	secHour = 3600
	secMin  = 60
)

func main() {
	var seconds, hours, minutes int
	fmt.Scan(&seconds)
	hours = seconds / secHour
	minutes = (seconds - hours*secHour) / secMin

	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
