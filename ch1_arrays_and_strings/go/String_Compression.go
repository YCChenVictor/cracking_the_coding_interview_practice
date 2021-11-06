// loop through the string with the concept of benchmark letter
// string = "aabcccccaaa"
// benchmark letter = "a"
// if current letter != "a" => insert benchmark letter and the number of this letter into array
// and change the benchmark letter into next letter, "b"
// after the whole process, concatenate the elements in array

// 看別人怎麼寫吧

package main

import (
	"fmt"
	"strconv"
)

func StringCompression(String string) string {
	runeSample := []rune(String)
	benchmark := ""
	var counter int
	result := ""
	for i := 0; i < len(runeSample); i++ {
		if benchmark != string(runeSample[i]) {
			result += benchmark + strconv.Itoa(counter)
			benchmark = string(runeSample[i])
			counter = 1
		} else {
			counter += 1
		}
	}
	result += benchmark + strconv.Itoa(counter)
	if len(result[1:]) < len(String) {
		return result[1:]
	} else {
		return String
	}
}

func main() {
	fmt.Println(StringCompression("aaaaaaa"))
}
