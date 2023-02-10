package main

import "sort"

func canReorderDoubled(arr []int) bool {
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] < 0 && arr[j] < 0 {
			return arr[i] > arr[j]
		}

		return arr[i] < arr[j]
	})

	m := make(map[int]int)

	for _, v := range arr {
		if _, exist := m[v/2]; v%2 == 0 && exist {
			m[v/2]--

			if m[v/2] == 0 {
				delete(m, v/2)
			}

		} else {
			m[v]++
		}
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
