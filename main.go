package main

import (
	"fmt"
)

func main() {
	nums := []int{0, 4, 3, 0}
	target := 0

	fmt.Println(twoSums(nums, target))
}

func twoSums(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}
