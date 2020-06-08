package main

import (
	"fmt"
)

func main() {
	inputs := []int{10, -9, 20, 25, 21, 40, 50, -20}
	inputsB := []int{5, -3, -5, -6, -7, -8}
	fmt.Println("get local val is ", getLocalMin(inputs))
	fmt.Println("get local val is for iter ", getLocalMinIter(inputs))
	fmt.Println("get local val is ", getLocalMin(inputsB))
	fmt.Println("get local val is for iter ", getLocalMinIter(inputsB))
	fmt.Println(inputs, inputsB)
}

func checkLocalMin(nums []int) int {
	mid := len(nums) / 2
	if nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1] {
		return mid
	}
	return checkLocalMin(nums[0 : mid-1])
}

func getLocalMin(nums []int) int {
	mid := len(nums) / 2
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[0]
		}
		return nums[1]
	} else {
		if nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1] {
			return nums[mid]
		}
		return getLocalMin(nums[:mid])
	}
}

func getLocalMinIter(nums []int) int {
	low, mid, high := 0, len(nums)/2, len(nums)-1
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	for low < high {
		mid = low + (high-low)/2
		if mid == 0 {
			if nums[mid] < nums[mid+1] {
				return nums[mid]
			}
			//return -9999999999
		}
		if mid == len(nums)-1 {
			if nums[mid] < nums[mid+1] {
				return nums[mid]
			}
			//return -9999999999
		}
		if nums[mid-1] > nums[mid] && nums[mid+1] > nums[mid] {
			return nums[mid]
		} else if nums[mid-1] < nums[mid] {
			high = mid - 1
		} else if nums[mid+1] < nums[mid] {
			low = mid + 1
		}
	}
	return -99999999
}
