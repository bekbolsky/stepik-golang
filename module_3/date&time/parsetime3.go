package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()

	splited := strings.Split(text, ",")
	firstDate := splited[0]
	secondDate := splited[1]
	// 13.03.2018 14:00:15
	layout := "02.01.2006 15:04:05"
	parsedFirst, _ := time.Parse(layout, firstDate)
	parsedSecond, _ := time.Parse(layout, secondDate)
	difference := parsedFirst.Sub(parsedSecond)
	if difference.Hours() < 0 {
		difference = parsedSecond.Sub(parsedFirst)
		fmt.Printf("%v", difference)
	} else {
		fmt.Printf("%v", difference)
	}
}
