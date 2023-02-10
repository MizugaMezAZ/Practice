package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0

	str := []rune(s)
	rep := make(map[rune]int)

	max := 0

	for right < len(str) {
		fmt.Println(left, right)

		if index, exist := rep[str[right]]; exist {
			if right-left > max {
				max = right - left
			}

			left = index + 1
			for k := range rep {
				delete(rep, k)
			}

			for i := left; i <= right; i++ {
				rep[str[i]] = i
			}

		} else {
			rep[str[right]] = right
		}

		right++
	}

	if right-left > max {
		max = right - left
	}

	return max
}
