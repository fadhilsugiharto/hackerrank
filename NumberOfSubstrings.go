package main

import (
	"fmt"
)

func checkLetterDict(letterDict [26]int, k int) bool {
	for _, letter := range letterDict {
		if letter > 0 && letter != k {
			return false
		}
	}
	return true
}

func countSubstrings(s string, k int) int {
	count := 0

	for i := 0; i < len(s); i++ {
		letterDict := [26]int{}

		for j := i; j < len(s); j++ {
			index := s[j] - 'a'
			letterDict[index] += 1

			if letterDict[index] > k {
				break
			}

			if letterDict[index] == k && checkLetterDict(letterDict, k) {
				count++
			}
		}
	}

	return count
}

func main() {
	s := "aabbcc"
	k := 2
	fmt.Println(countSubstrings(s, k))
}
