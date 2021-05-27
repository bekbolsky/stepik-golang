package main

import "fmt"

func main() {
	var n, num, min, minCount int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		nums[i] = num
	}
	min = nums[0]
	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
			minCount = 1
		} else if min == nums[i] {
			minCount++
		}
	}
	fmt.Println(minCount)
}
