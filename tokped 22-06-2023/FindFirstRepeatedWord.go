package main

import (
	"fmt"
	"strings"
)

func firstRepeatedWord(sentence string) string {
	splitFunc := func(c rune) bool {
		separatorChars := " .,"
		return strings.ContainsRune(separatorChars, c)
	}

	slice := strings.FieldsFunc(sentence, splitFunc)

	dictRepeat := map[string]bool{}
	for _, s := range slice {
		if !dictRepeat[s] {
			dictRepeat[s] = true
		} else {
			return s
		}
	}
	return ""
}

func main() {
	fmt.Println(firstRepeatedWord("He had had quite enough of this nonsense."))
}
