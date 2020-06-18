package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ntkjer/sedgewick/chapter1/binarysearch"
	"github.com/ntkjer/sedgewick/chapter1/stopwatch"
)

func main() {
	nums := readInts()
	var stopwatch *stopwatch.Stopwatch = new(stopwatch.Stopwatch)
	stopwatch.Start()
	c := TwoSumCount(nums)
	stopwatch.Stop()
	time := stopwatch.ElapsedTime()
	fmt.Printf("two sum: there are %d pairs that sum to 0. Elapsed time %v\n", c, time)
	stopwatch.Start()
	c = TwoSumFaster(nums)
	stopwatch.Stop()
	time = stopwatch.ElapsedTime()
	fmt.Printf("two sum faster: there are %d pairs that sum to 0. Elapsed time %v\n", c, time)
}

func readInts() []int {
	nums := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		item := scanner.Text()
		item = strings.TrimSpace(item)
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return nums
}

//Return number of distrint pairs (i, j) s.t. a[i] + a[j] = 0
func TwoSumCount(nums []int) int {
	sort.Ints(nums)
	N := len(nums)
	count := 0
	for i := 0; i < N; i++ {
		if binarysearch.Rank(-nums[i], nums) > i {
			count++
		}
	}
	return count
}

func TwoSumFaster(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	count := 0

	for _, num := range nums {
		_, ok := m[-num]
		if ok {
			count++
		}
		if num == 0 {
			count--
		}
	}
	return count / 2
}
