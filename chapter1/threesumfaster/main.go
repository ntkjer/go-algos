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
	sort.Ints(nums)
	var stopwatch *stopwatch.Stopwatch = new(stopwatch.Stopwatch)
	stopwatch.Start()
	c := ThreeSumCount(nums)
	stopwatch.Stop()
	time := stopwatch.ElapsedTime()
	fmt.Printf("3sum : there are %d pairs that sum to 0....  elapsed time %v\n", c, time)
	stopwatch.Start()
	c = ThreeSumFaster(nums)
	stopwatch.Stop()
	time = stopwatch.ElapsedTime()
	fmt.Printf("3sumFaster : there are %d pairs that sum to 0....  elapsed time %v\n", c, time)
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
func ThreeSumCount(nums []int) int {
	N := len(nums)
	count := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if binarysearch.Rank(-nums[i]-nums[j], nums) > j {
				count++
			}
		}
	}
	return count
}

func ThreeSumFaster(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			_, ok := m[-sum]
			if ok {
				count++
			}

		}
	}
	return count / 2
}
