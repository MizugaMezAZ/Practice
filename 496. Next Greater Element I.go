package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var s []int
	var m = make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums2[i] {
			s = s[:len(s)-1] // pop stack elements small than current
		}
		v := -1
		if len(s) > 0 {
			v = s[len(s)-1]
		}
		m[nums2[i]] = v

		s = append(s, nums2[i])

		fmt.Println(s)
	}
	var rst []int
	for _, v := range nums1 {
		rst = append(rst, m[v])
	}
	return rst
}

// func nextGreaterElement(nums1 []int, nums2 []int) []int {
// 	ans := make([]int, 0, len(nums1))
// 	// key = 值 val = great
// 	great := make(map[int]int)
// 	// 存放還未找到great的值
// 	not := make(map[int]struct{})

// 	for _, v := range nums2 {
// 		not[v] = struct{}{}
// 		for key := range not {
// 			if v > key {
// 				great[key] = v
// 				delete(not, key)
// 			}
// 		}
// 	}

// 	for key := range not {
// 		great[key] = -1
// 	}

// 	for _, v := range nums1 {
// 		ans = append(ans, great[v])
// 	}

// 	return ans
// }
