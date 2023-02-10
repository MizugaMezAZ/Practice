package main

import "fmt"

func maxRotateFunction(nums []int) int {
	len := len(nums)
	if len == 1 {
		return 0
	}

	f := make([]int, len)

	sum := 0
	f0 := 0

	for i, v := range nums {
		sum += v
		f0 += v * i
	}

	max := f0
	f[0] = f0

	for i := 1; i < len; i++ {
		f[i] = f[i-1] + sum - len*nums[len-i]
		if f[i] > max {
			max = f[i]
		}
	}

	fmt.Println(f)

	return max
}
