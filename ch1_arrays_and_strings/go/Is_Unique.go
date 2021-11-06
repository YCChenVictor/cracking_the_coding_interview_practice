// loop through a string and add them into hash table
// check whether the letter is in the hash table
// if yes, break; else continue
// The complexity = O(n)
package main

import "fmt"

func IsUnique(input string) bool {
	seen := make(map[rune]struct{})
	for _, r := range input {
		if _, ok := seen[r]; ok {
			return false
		} else {
			seen[r] = struct{}{}
		}
	}
	return true
}

func main() {
	result := IsUnique("abcd")
	fmt.Println(result)
}
