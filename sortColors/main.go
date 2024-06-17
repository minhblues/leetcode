package main

import (
	"fmt"
)

func sortColors(nums []int) {
	colors := make(map[int]int)
	for _, v := range nums {
		colors[v]++
	}

	k := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < colors[i]; j++ {
			nums[k] = i
			k++
		}

	}
}

func main() {
	nums := []int{1, 1, 2, 0, 2, 0, 2, 0}
	sortColors(nums)

	fmt.Println(nums)
}
