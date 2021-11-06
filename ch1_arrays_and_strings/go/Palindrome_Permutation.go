// The concept: (unoptimized)
// The complexity = ?
// 1. remove space
// 2. loop through the string and add letters into hash map (key = letter; item = number of letter)
// 3. if only one or no item is odd => return true
package main

import (
	"fmt"
	"strings"
)

func CheckPermutation(String string) bool {
	String = strings.ToLower(String)
	seenSummary := make(map[rune]int)
	for _, rune := range String {
		if rune == 32 {
			continue
		}
		if _, ok := seenSummary[rune]; ok {
			seenSummary[rune] += 1
		} else {
			seenSummary[rune] = 1
		}
	}
	oddAndEvenSummary := make(map[string]int)
	for _, value := range seenSummary {
		if value%2 != 0 {
			oddAndEvenSummary["odd"] += 1
		}
	}
	if oddAndEvenSummary["odd"] > 1 {
		return false
	}
	return true
}

func main() {
	result := CheckPermutation("Tact Coaab")
	fmt.Println(result)
}
