package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ntkjer/sedgewick/chapter1/stopwatch"
)

func main() {
	nums := readInts()
	var stopwatch *stopwatch.Stopwatch = new(stopwatch.Stopwatch)
	stopwatch.Start()
	c := FourSumCount(nums)
	stopwatch.Stop()
	time := stopwatch.ElapsedTime()
	fmt.Printf("there are %d pairs that sum to 0....  elapsed time %v\n", c, time)
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
func FourSumCount(nums []int) int {
	sort.Ints(nums)
	N := len(nums)
	count := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				for h := k + 1; h < N; h++ {
					if nums[i]+nums[j]+nums[k] == 0 {
						count++
					}
				}
			}
		}
	}

	return count
}
