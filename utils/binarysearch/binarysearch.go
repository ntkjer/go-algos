package binarysearch

import (
	"reflect"
)

func Rank(key interface{}, a []interface{}) int {
	lo := 0
	hi := len(a) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if less(key, a[mid]) {
			hi = mid - 1
		} else if less(key, a[mid]) {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func less(x, y interface{}) bool {
	if reflect.TypeOf(x) != reflect.TypeOf(y) {
		panic("can't compare mismatched input types for x and y")
	}
	switch x.(type) {
	case int:
		a, b := x.(int), y.(int)
		if a > b {
			return false
		}
	case string:
		a, b := x.(string), y.(string)
		if a > b {
			return false
		}
	case float64:
		a, b := x.(float64), y.(float64)
		if a > b {
			return false
		}

	default:
		panic("unhandled type.")
	}
	return true
}
