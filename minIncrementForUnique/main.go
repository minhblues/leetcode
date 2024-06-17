package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	res := 0
	old := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			old = nums[i]
			nums[i] = nums[i-1] + 1
			res += nums[i] - old
		}
	}

	return res

}

func climbStairs(n int) int {

	maps := map[int]int{
		1:  1,
		2:  2,
		3:  3,
		4:  5,
		5:  8,
		6:  13,
		7:  21,
		8:  34,
		9:  55,
		10: 89,
		11: 144,
		12: 233,
		13: 377,
		14: 610,
		15: 987,
		16: 1597,
		17: 2584,
		18: 4181,
		19: 6765,
		20: 10946,
		21: 17711,
		22: 28657,
		23: 46368,
		24: 75025,
		25: 121393,
		26: 196418,
		27: 317811,
		28: 514229,
		29: 832040,
		30: 1346269,
		31: 2178309,
		32: 3524578,
		33: 5702887,
		34: 9227465,
		35: 14930352,
		36: 24157817,
		37: 39088169,
		38: 63245986,
		39: 102334155,
		40: 165580141,
		41: 267914296,
		42: 433494437,
		43: 701408733,
		44: 1134903170,
		45: 1836311903,
	}

	return maps[n]
}

func twoSum(nums []int, target int) []int {
	sums := make(map[int]int)

	for i, num := range nums {
		if index, ok := sums[target-num]; ok {
			return []int{i, index}
		}

		sums[nums[i]] = i
	}

	return nil
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	median := 0.0

	total := 0

	for _, num := range nums1 {
		total += num
	}
	for _, num := range nums2 {
		total += num
	}

	median = float64(total) / float64(len(nums1)+len(nums2))

	return median
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strReverse := Reverse(str)

	if str == strReverse {
		return true
	}

	return false

}

func main() {
	// nums := []int{3, 2, 1, 2, 1, 7}
	// nums := []int{1, 4, 4, 5, 7}

	// res := minIncrementForUnique(nums)
	// fmt.Println(res)

	nums := []int{3, 2, 4}

	fmt.Println(twoSum(nums, 6))

	num := 123421
	fmt.Println(isPalindrome(num))

}
