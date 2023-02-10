package main

type Rectangle struct {
	X1, X2 int
	Y1, Y2 int
}

// leetcode 223
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - (max(min(ax2, bx2)-max(ax1, bx1), 0) * max(min(ay2, by2)-max(ay1, by1), 0))
}

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}

// 	return y
// }

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
