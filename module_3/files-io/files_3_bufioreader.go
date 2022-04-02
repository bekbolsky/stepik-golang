package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAndSplit(inputPath string) {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("Error occured reading file - %v", err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	var n int
	for {
		s, err := br.ReadString(';')
		if err != nil {
			return
		}
		n++
		if s == "0;" {
			fmt.Println(n)
			return
		}
	}
}

/* func ReadAndSplit(inputPath string) {
	file, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Printf("Error occured reading file - %v", err)
	}
	splitedText := strings.Split(string(file), ";")
	// fmt.Println(splited)
	for i, value := range splitedText {
		if value == "0" {
			fmt.Printf("index: %v, value: %v, nomer: %v", i, value, i+1)
		}
	}
} */

func main() {
	ReadAndSplit("module_3/files-io/task.data")
}
