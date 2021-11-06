// The concept: (unoptimized)
// The complexity = O(n)
// loop through a string and add letters into hash table with key = letter, item = the number of this letter in this string
// then loop through another string and abstract the letters from the hash table
package main

import "fmt"

func CheckPermutation(string1 string, string2 string) bool {
	seenSummary := make(map[rune]int)
	for _, rune := range string1 {
		if _, ok := seenSummary[rune]; ok {
			seenSummary[rune] += 1
		} else {
			seenSummary[rune] = 1
		}
	}
	for _, rune := range string2 {
		if _, ok := seenSummary[rune]; ok {
			seenSummary[rune] -= 1
		}
	}
	for _, value := range seenSummary {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	result := CheckPermutation("abcd", "dcba")
	fmt.Println(result)
}
