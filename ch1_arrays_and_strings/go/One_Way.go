// loop through first string and add the number of letters into hash table
// string1 = "pale" => {p: 1, a: 1, l: 1, e: 1}
// then loop through second string and abstract the number of letters
// string2 = "ple" => {p: 0, a: 1, l: 0, e: 0}
// string2 = "bale" => {p: 1, a: 0, l: 0, e: 0}
// string1 = "pales" & string2 = "pale" => {p: 0, a: 0, l: 0, e: 0, s :1}
// string1 = "pale" & string2 = "bake" => {p: 1, a: 0, l: 1, e: 0}

// if the sum of items > 1, then return false

package main

import "fmt"

func OneWay(string1 string, string2 string) bool {
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
	result := 0
	for _, value := range seenSummary {
		result += value
	}
	if result <= 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(OneWay("pale", "rabe"))
}
