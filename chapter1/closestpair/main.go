package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []float64{-12, 4, -1, 8, 6, 2, -3, 13}
	fmt.Println(ClosestPair(nums))

}

func ClosestPair(nums []float64) (float64, float64) {
	var a, b, diff float64
	diff = 10000000000000000000000000000000000000
	sort.Float64s(nums)
	for i := 0; i < len(nums)-1; i++ {
		if math.Abs(nums[i]-nums[i+1]) < diff {
			diff = math.Abs(nums[i] - nums[i+1])
			a, b = nums[i], nums[i+1]
		}
	}
	return a, b
}
