package main

import (
	"fmt"
	"math"
)

func maxDistance(grid [][]int) int {
	n := len(grid)

	landCoordinite := make(map[int][]int)
	waterCoordinite := make(map[int][]int)

	landCount := 0
	waterCount := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				landCoordinite[landCount] = []int{i, j}
				landCount++
			} else {
				waterCoordinite[waterCount] = []int{i, j}
				waterCount++
			}
		}
	}

	if waterCount == 0 || landCount == 0 {
		return -1
	}

	fmt.Println(landCoordinite, waterCoordinite)

	ans := 0
	for _, water := range waterCoordinite {
		r := math.MaxInt64
		for _, land := range landCoordinite {
			r = min(r, abs(water[0]-land[0])+abs(water[1]-land[1]))
		}

		ans = max(ans, r)
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
