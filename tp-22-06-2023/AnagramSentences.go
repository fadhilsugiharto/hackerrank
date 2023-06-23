package main

import (
	"fmt"
	"sort"
	"strings"
)

func countSentences(wordSet []string, sentences []string) []int64 {
	// Write your code here
	counts := make([]int64, len(sentences))
	anagramCounts := make(map[string]int)

	for _, word := range wordSet {
		sortedWord := sortString(strings.ToLower(word))
		anagramCounts[sortedWord]++
	}

	for i, sentence := range sentences {
		sentenceWords := strings.Split(sentence, " ")
		sentenceCounts := make(map[string]int)

		for _, word := range sentenceWords {
			sortedWord := sortString(strings.ToLower(word))
			sentenceCounts[sortedWord]++
		}

		count := int64(1)
		for word, freq := range sentenceCounts {
			count *= int64(anagramCounts[word]) * int64(freq)
		}

		counts[i] = count
	}

	return counts
}

func sortString(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// jal dhewe
//func countAnagram(wordSet []string, sentences []string) []int {
//	dictSentence := map[string]int{}
//	for _, word := range wordSet {
//		word = sortString(word)
//		dictSentence[word]++
//	}
//
//	result := make([]int, len(sentences))
//	for i, sentence := range sentences {
//		count := 1
//		words := strings.Split(sentence, " ")
//
//		for _, word := range words {
//			word = sortString(word)
//			count *= dictSentence[word]
//		}
//		result[i] = count
//	}
//
//	return result
//}

func main() {
	wordSet := []string{"the", "bats", "tabs", "in", "cat", "act", "tac", "cta", "sbat"}
	sentences := []string{"cat the bats", "in the act", "act tabs in"}

	counts := countSentences(wordSet, sentences)
	fmt.Println(counts)
	//fmt.Println(countAnagram(wordSet, sentences))
}
