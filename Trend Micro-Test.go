package main

// bdaaadadb
// abacb
// zthtzh
//   0 0 0
//   a b d
// b 0 1 0 2
// d 0 1 1 6
// a 1 1 1 7
// a 0 1 1 6
// a 1 1 1 7
// d 1 1 0 3
// a 0 1 0 2
// d 0 1 1 6
// b 0 0 1 4

// n +1
// n + 3 +1 - n -1
// 0 1  2  3 4 5 6 7
// 0 5 -1 -1 1 4 2 3

func Solution(s string) int {
	if len(s) == 1 {
		return 0
	}

	// 計算出現多少字母
	count := make([]int, 26)
	n := 0
	for _, v := range s {
		count[v-'a']++
	}

	// 照字母順序儲存推移量 (減少空間消耗)
	indexMap := make(map[int]int)
	for i, v := range count {
		if v > 0 {
			indexMap[i] = n
			n++
		}
	}

	leng := len(s)
	pos := make([]int, 1<<n)

	// 初始化 index 0為0 其餘-1
	for i := 1; i < len(pos); i++ {
		pos[i] = -1
	}

	status := 0
	maxLen := 0

	for i := 0; i < leng; i++ {
		ch := s[i]

		status ^= 1 << indexMap[int(ch-'a')]

		if pos[status] >= 0 {
			maxLen = max(maxLen, i+1-pos[status])
		} else {
			pos[status] = i + 1
		}

	}

	return maxLen
}
