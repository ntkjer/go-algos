package main

import (
	"fmt"
)

func main() {
	arrA := [10]int{0, 1, 3, 5, 7, 9, 10, 11, 12, 13}
	arrB := [5]int{0, 1, 2, 3, 4}
	arrAindex, arrBindex := 0, 0
	var matched int
	for arrAindex < len(arrA) && arrBindex < len(arrB) {
		if arrA[arrAindex] > arrB[arrBindex] {
			arrBindex++
		} else if arrA[arrAindex] < arrB[arrBindex] {
			arrAindex++
		} else {
			if matched != arrA[arrAindex] {
				fmt.Println(arrA[arrAindex])
				matched = arrA[arrAindex]
			}
			arrAindex++
			arrBindex++
		}
	}
}
