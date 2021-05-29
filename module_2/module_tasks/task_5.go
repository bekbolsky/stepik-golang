package main

import (
	"fmt"
	"math"
)

var k, v, p float64

func M() float64 {
	m := p * v
	return m
}

func W() float64 {
	m := M()
	w := math.Sqrt(k / m)
	return w

}

func T() float64 {
	w := W()
	t := 6 / w
	return t
}

func main() {
	fmt.Scan(&k, &v, &p)
	// k, v, p = 1296, 6, 6

	fmt.Println(T())
}
