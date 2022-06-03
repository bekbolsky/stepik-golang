package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	nr := bufio.NewReader(os.Stdin)
	inputTime, _ := nr.ReadString('\n')
	inputTime = inputTime[:len(inputTime)-1]
	layout := "2006-01-02 15:05:05"
	parsedTime, _ := time.Parse(layout, inputTime)
	if parsedTime.Hour() >= 13 {
		parsedTime = parsedTime.AddDate(0, 0, 1)
		fmt.Print(parsedTime.Format(layout))
	} else {
		fmt.Print(parsedTime.Format(layout))
	}
}
