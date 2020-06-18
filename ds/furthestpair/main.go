package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []float64{-5.2, 9.4, 20, -10, 21.1, 40, 50, -20}
	fmt.Println("expected -20 50")
	fmt.Println(ClosestPair(nums))

}

func ClosestPair(nums []float64) (float64, float64) {
	var currMax, currMin float64
	currMin = 100000000000000000000000.0
	sort.Float64s(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] < currMin {
			currMin = nums[i]
		} else if nums[i] > currMax {
			currMax = nums[i]
		}
	}
	return currMin, currMax
}
