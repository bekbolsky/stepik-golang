package main

import (
	"fmt"
)

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/"
}

func printError(value interface{}) {
	fmt.Printf("value=%v: %T\n", value, value)
}

func main() {
	value1, value2, operation := readTask()
	// var v1, v2 float64
	v1, ok := value1.(float64)
	if !ok {
		printError(value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		printError(value2)
		return
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			result := v1 + v2
			fmt.Printf("%.4f\n", result)
		case "-":
			result := v1 - v2
			fmt.Printf("%.4f\n", result)
		case "*":
			result := v1 * v2
			fmt.Printf("%.4f\n", result)
		case "/":
			result := v1 / v2
			fmt.Printf("%.4f\n", result)
		default:
			fmt.Println("неизвестная операция")
			return
		}
	} else {
		fmt.Println("неизвестная операция")
		return
	}
}
