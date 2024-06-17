package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		nums1[0] = nums2[0]
		return
	}

	length := m + n

	for i := 0; i < length; i++ {
		// nums1[i] = nums2[i-m]
		j := i - m

		if nums1[i] > nums2[j] {

		}
	}

	// sort.Ints(nums1)

}

func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// nums2 := []int{2, 5, 6}

	// merge(nums1, 3, nums2, 3)

	// fmt.Println(nums1)

	num := 10_000.3
	fmt.Println(num)
}
