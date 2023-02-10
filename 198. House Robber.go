package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	if len(nums) < 3 {
		return max(nums...)
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]

	for i := 2; i < len(dp); i++ {
		dp[i] = nums[i] + max(dp[:i-1]...)
	}

	fmt.Println(dp)

	return max(dp...)
}

func max(n ...int) int {
	max := math.MinInt
	for _, v := range n {
		if v > max {
			max = v
		}
	}

	return max
}
