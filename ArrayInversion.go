package main

import (
	"fmt"
)

func arrayInvestion(nums []int) int {
	var ans [][]int

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] < nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[j] < nums[k] {
					continue
				}
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}

	return len(ans)
}

func main() {
	fmt.Println(arrayInvestion([]int{5, 3, 4, 2, 1}))
}

//func sortNggawe(nums []int) []int {
//	for i := 0; i < len(nums); i++ {
//		min := nums[i]
//		index := i
//		for j := i + 1; j < len(nums); j++ {
//			if nums[j] < min {
//				min = nums[j]
//				index = j
//			}
//		}
//		nums[i], nums[index] = nums[index], nums[i]
//	}
//	return nums
//}
