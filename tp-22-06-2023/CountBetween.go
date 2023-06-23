package main

import (
	"fmt"
	"sort"
)

// brute force
func countBetween(arr []int32, low []int32, high []int32) []int32 {
	// Write your code here
	length := len(low)
	result := make([]int32, length)

	for i := 0; i < length; i++ {
		l := low[i]
		h := high[i]
		count := int32(0)

		for _, num := range arr {
			if num >= l && num <= h {
				count++
			}
		}

		result[i] = count
	}

	return result
}

// through binary search
func countBetweenBinarySearch(arr []int32, low []int32, high []int32) []int32 {
	result := make([]int32, len(low))
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(low); i++ {
		l := low[i]
		h := high[i]

		count := binarySearchCount(arr, h) - binarySearchCount(arr, l-1)
		result[i] = count
	}

	return result
}

// Binary search to find the count of elements less than or equal to target
func binarySearchCount(arr []int32, target int32) int32 {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] <= target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return int32(low)
}

func main() {
	//arr := []int32{4, 8, 7}
	//low := []int32{2, 4}
	//high := []int32{8, 4}

	arr := []int32{100, 200, 300, 500, 400}
	low := []int32{0, 200}
	high := []int32{200, 400}

	fmt.Println(countBetween(arr, low, high))
	fmt.Println(countBetweenBinarySearch(arr, low, high))
}
