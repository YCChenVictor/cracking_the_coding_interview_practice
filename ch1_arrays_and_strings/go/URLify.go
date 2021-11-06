// split the string by space
// concatenate the string with "%20"
package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields("Mr John Smith ")
	fmt.Println(strings.Join(words[:], "%20"))
}
