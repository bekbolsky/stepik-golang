package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var sum int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s := scanner.Text()
		number, _ := strconv.Atoi(s)
		sum += number
	}

	os.Stdout.WriteString(strconv.Itoa(sum))
}
