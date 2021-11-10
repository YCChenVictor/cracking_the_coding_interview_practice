// (rotate) turn right, the steps: reverse up to down, then swap the symmetry
// 1 2 3    7 8 9    7 4 1
// 4 5 6 => 4 5 6 => 8 5 2
// 7 8 9    1 2 3    9 6 3
// (anti_rotate) turn left, the steps: reverse left to right, then swap the symmetry
// 1 2 3     3 2 1     3 6 9
// 4 5 6  => 6 5 4  => 2 5 8
// 7 8 9     9 8 7     1 4 7

package main

import (
	"fmt"
)

func rotate(matrix [][]int) [][]int {
	// reverse the matrix
	for i, j := 0, len(matrix)-1; i<j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{
		[]int{1,2,3}, 
		[]int{4,5,6}, 
		[]int{7,8,9},
	}
	fmt.Println(rotate(matrix))
}
