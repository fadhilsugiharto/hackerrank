package main

import "fmt"

func palindrome(s string) int32 {
	// Write your code here
	lenS := len(s)
	if lenS == 0 {
		return 0
	}

	var l, r, palindromeCounts int
	palindromeDict := map[string]bool{}

	for r < lenS {
		if !palindromeDict[string(s[l])] {
			palindromeCounts += 1
			palindromeDict[string(s[l])] = true
		}
		for r+1 < lenS && s[l] == s[r+1] {
			if !palindromeDict[s[l:r+2]] {
				palindromeCounts += 1
				palindromeDict[s[l:r+2]] = true
			}
			r++
		}

		for l-1 >= 0 && r+1 < lenS && s[l-1] == s[r+1] {
			if !palindromeDict[s[l-1:r+2]] {
				palindromeCounts += 1
				palindromeDict[s[l-1:r+2]] = true
			}
			l--
			r++
		}

		l = (l+r)/2 + 1
		r = l
	}

	return int32(palindromeCounts)
}

func main() {
	fmt.Println(palindrome("abcddcbabcdcdaadcdcbabcdddcb"))
}
