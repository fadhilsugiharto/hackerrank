package main

import (
	"fmt"
	"sort"
	"strings"
)

func concatenateSubstrings(substrings []string) string {
	sort.Slice(substrings, func(i, j int) bool {
		return substrings[i]+substrings[j] < substrings[j]+substrings[i]
	})

	return strings.Join(substrings, "")
}

func main() {
	substrings := []string{"a", "bd", "ac", "cd"}
	fmt.Println(concatenateSubstrings(substrings))
	// Output: aacbdcd

	substrings = []string{"c", "cc", "cca", "cccb"}
	fmt.Println(concatenateSubstrings(substrings))
	// Output: ccacccbccc
}
