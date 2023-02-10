package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// fmt.Println(PointCal([]int{0, 10, 0, 10, 0, 10, 0, 10, 0, 10, 0, 10, 0, 10, 0, 10, 0, 10, 10, 10, 10}))
	fmt.Println(PointCal([]int{6, 4, 0, 10, 1, 0, 0, 10, 0, 10, 0, 8, 9, 1, 3, 5, 4, 0, 10, 3, 4}))
	// fmt.Println(Solution("zthtzh"))
}

func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}

	r := make([]int, 0)
	m := 1
	for {
		r = append(r, x%10)

		x /= 10

		if x == 0 {
			break
		}

		m *= 10
	}

	ans := 0

	for _, v := range r {
		ans += v * m
		m /= 10
	}

	return ans
}

func heapSort(data []int) {

}

func binarysearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func isRever(i int) bool {
	x := strconv.Itoa(i)
	ru := []rune(x)

	left := 0
	right := len(ru) - 1
	for left < right {
		ru[left], ru[right] = ru[right], ru[left]

		if string(ru) != x {
			return false
		}

		left++
		right--
	}

	return true
}
