package main

import (
	"fmt"
	"time"
)

func main() {
	var inputTime string
	fmt.Scan(&inputTime)
	// layout := "2006-01-02T15:04:05Z07:00"
	parsedTime, _ := time.Parse(time.RFC3339, inputTime)
	fmt.Println(parsedTime.Format(time.UnixDate))
}
