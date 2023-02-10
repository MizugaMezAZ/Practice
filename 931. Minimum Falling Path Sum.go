package main

import "math"

var totalMin int

func minFallingPathSum(A [][]int) int {
	if len(A) == 0 {
		return 0
	}

	dp := make([]int, len(A[0]))

	for i := 0; i < len(A); i++ {
		dpCopy := make([]int, len(dp))
		copy(dpCopy, dp)
		for j := 0; j < len(dp); j++ {
			dp[j] = A[i][j] + MinInMany(getValue(dpCopy, j-1), getValue(dpCopy, j), getValue(dpCopy, j+1))
		}
	}
	return MinInMany(dp...)
}

func getValue(dp []int, i int) int {
	if i < 0 || i >= len(dp) {
		return math.MaxInt32
	}
	return dp[i]
}

// func minFallingPathSum(matrix [][]int) int {
// 	totalMin = math.MaxInt

// 	if len(matrix) == 1 {
// 		return matrix[0][0]
// 	}

// 	min := math.MaxInt

// 	for i := range matrix[0] {
// 		_, n := fallPath(matrix, 0, 0, i)
// 		if n < min {
// 			min = n
// 		}
// 	}

// 	return min
// }

// func fallPath(matrix [][]int, count, x, y int) (bool, int) {
// 	if x >= len(matrix) || x < 0 || y >= len(matrix) || y < 0 {
// 		return false, count
// 	}

// 	count += matrix[x][y]

// 	if totalMin < count+(len(matrix)-y-1)*-100 {
// 		return false, count
// 	}

// 	min := []int{}

// 	for i := y - 1; i <= y+1; i++ {
// 		if exist, n := fallPath(matrix, count, x+1, i); exist {
// 			min = append(min, n)
// 		}
// 	}

// 	if len(min) == 0 {
// 		return true, count
// 	}

// 	return true, MinInMany(min...)
// }

func MinInMany(n ...int) int {
	if len(n) == 1 {
		return n[0]
	}

	min := n[0]

	for _, v := range n {
		if v < min {
			min = v
		}
	}

	return min
}
