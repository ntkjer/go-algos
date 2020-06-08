package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("")
}

func ClosestPair(nums []int) (int, int) {
	startPos := rand.Intn(len(nums))
	x, y := nums[startPos], nums[len(nums)-startPos]
	for _, num := range nums {
		if num < x {
			x = num
		} else if num < y && num > x {
			y = num
		}
	}
	return x, y
}
