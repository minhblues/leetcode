package main

import (
	"fmt"
	"math"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	var res int
	for i := range seats {
		res += int(math.Abs(float64(seats[i] - students[i])))
	}

	return res
}

func main() {

	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}

	res := minMovesToSeat(seats, students)

	fmt.Println(res)
}

// [3,2,1,2,1,7]
// [1,1,2,2,3,7]2
// {}
