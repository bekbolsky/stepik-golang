package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	var revX, revY int
	rankX, rankY := 1, 1

	copyX := x
	for {
		if copyX > 0 && copyX < 10 {
			break
		}
		copyX = copyX / 10
		rankX *= 10
	}

	copyY := y
	for {
		if copyY > 0 && copyY < 10 {
			break
		}
		copyY = copyY / 10
		rankY *= 10
	}

	for i := x; i != 0; i = x {
		reminder := x % 10
		x = x / 10
		revX += reminder * rankX
		rankX /= 10
	}
	for ii := y; ii != 0; ii = y {
		reminder := y % 10
		y = y / 10
		revY += reminder * rankY
		rankY /= 10
	}

	for i := revX; i != 0; i /= 10 {
		digitX := i % 10
		for j := revY; j != 0; j /= 10 {
			digitY := j % 10
			if digitX == digitY {
				fmt.Print(digitX, " ")
			}
		}
	}

}
