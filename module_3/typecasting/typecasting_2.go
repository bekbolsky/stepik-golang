package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	numbers := strings.Split(text, ";")
	for i, number := range numbers {
		number = strings.ReplaceAll(number, " ", "")
		number = strings.ReplaceAll(number, ",", ".")
		numbers[i] = number
	}
	number1, _ := strconv.ParseFloat(numbers[0], 64)
	number2, _ := strconv.ParseFloat(numbers[1], 64)
	result := number1 / number2
	fmt.Printf("%.4f\n", result)
}
