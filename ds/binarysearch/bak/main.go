package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(nums)

	arg := os.Args[1]
	toFind, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}

	if rank(toFind, nums) == -1 {
		fmt.Println(toFind)
	}

}

func rank(key int, a []int) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < a[mid] {
			hi = mid - 1
		} else if key > a[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
