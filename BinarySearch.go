package main

import (
	"fmt"
	"strconv"
)

func sortir(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				index = j
			}
		}
		if min != nums[i] {
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	return nums
}

func binarySearch(nums []int, target int) string {
	first := 0
	last := len(nums) - 1

	for first <= last {
		middle := (first + last) / 2
		if nums[middle] == target {
			return "found on index " + strconv.Itoa(middle)
		}
		if nums[middle] > target {
			last = middle
		} else {
			first = middle + 1
		}
	}

	return "not found"
}

func main() {
	arr := []int{6, 8, 7, 5, 4, 3, 2, 1}
	arr = sortir(arr)
	fmt.Println(binarySearch(arr, 2))
}
