package binarysearch

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
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

func TestBinarySearchInts(t *testing.T) {
	f, err := os.Open("input.txt")
	if err != nil {
		t.Error(err)
	}
	ints, err := ReadInts(strings.NewReader(f.Name()))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ints)
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	//scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
