package main

import "fmt"

func slowestKey(keyTimes [][]int32) string {
	var longestKey int32
	var longestDuration int32

	for i := 0; i < len(keyTimes); i++ {
		key := keyTimes[i][0] + 'a'
		duration := keyTimes[i][1]

		if i > 0 {
			prevDuration := keyTimes[i-1][1]
			interval := duration - prevDuration

			if interval > longestDuration {
				longestDuration = interval
				longestKey = key
			}
		} else {
			longestDuration = duration
			longestKey = key
		}
	}

	return string(longestKey)
}

func main() {
	keyTimes := [][]int32{{0, 10}, {0, 13}, {4, 15}, {5, 16}, {4, 20}}
	fmt.Println(slowestKey(keyTimes))
}
