package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	textRunes := []rune(text)
	length := len(textRunes)
	var revText string
	for i := length - 1; i >= 0; i-- {
		revText += string(textRunes[i])
	}
	if text == revText {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
