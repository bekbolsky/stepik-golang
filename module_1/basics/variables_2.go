package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	c = a + b
	fmt.Println(c)
}
