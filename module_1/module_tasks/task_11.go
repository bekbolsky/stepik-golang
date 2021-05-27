package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a == 1:
		fmt.Printf("%v korova\n", a)
	case a > 1 && a <= 4:
		fmt.Printf("%v korovy\n", a)
	case a >= 5 && a <= 20:
		fmt.Printf("%v korov\n", a)
	case a >= 21 && a%10 == 1:
		fmt.Printf("%v korova\n", a)
	case a >= 21 && (a%10 == 2 || a%10 == 3 || a%10 == 4):
		fmt.Printf("%v korovy\n", a)
	default:
		fmt.Printf("%v korov\n", a)
	}
}
