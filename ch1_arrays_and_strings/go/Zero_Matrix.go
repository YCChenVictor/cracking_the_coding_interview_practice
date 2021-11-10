// Given, a matrix:
// 1 2 3    7 0 9
// 4 0 6 => 0 0 0
// 7 8 9    1 0 3
// 1 2 3    1 0 3
// 1. set the elements in the same row into zero
// 2. set the elements with same index in other rows into zero

package main

import "fmt"

func zeroMatrix(matrix [][]int) [][]int {
	// record the i, j of zero
	zeroLocations := [][2]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				zeroLocation := [2]int{i, j}
				zeroLocations = append(zeroLocations, zeroLocation)
			}
		}
	}
	for i := 0; i < len(zeroLocations); i++ {
		// set the elements in the same row into zero
		replacementRow := make([]int, len(matrix)-1)
		matrix[zeroLocations[i][0]] = replacementRow
		// set the elements with same index in other rows into zero
		for j := 0; j < len(matrix); j++ {
			if matrix[j][zeroLocations[i][1]] == 0 {
				continue
			}
			matrix[j][zeroLocations[i][1]] = 0
		}
	}
	return matrix
}

func main() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 0, 6},
		[]int{7, 8, 9},
		[]int{1, 2, 0},
	}
	fmt.Println(zeroMatrix(matrix))
}
