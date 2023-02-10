package main

import "fmt"

// type rec struct{
// 	LT corner
// 	LB corner

// 	RT corner
// 	RB corner
// }

type corner struct {
	X int
	Y int
}

func isRectangleCover(rectangles [][]int) bool {
	cornerCount := make(map[corner]struct{})
	minX, minY, maxX, maxY := rectangles[0][0], rectangles[0][1], rectangles[0][2], rectangles[0][3]

	area := 0
	for _, v := range rectangles {
		x1, y1, x2, y2 := v[0], v[1], v[2], v[3]
		if x1 <= minX && y1 <= minY {
			minX = x1
			minY = y1
		}

		if x2 >= maxX && y2 >= maxY {
			maxX = x2
			maxY = y2
		}
		corners := make([]corner, 0, 4)

		corners = append(corners,
			corner{
				X: x1,
				Y: y2,
			}, corner{
				X: x1,
				Y: y1,
			}, corner{
				X: x2,
				Y: y2,
			}, corner{
				X: x2,
				Y: y1,
			})

		for _, v := range corners {
			if _, exist := cornerCount[v]; exist {
				delete(cornerCount, v)
			} else {
				cornerCount[v] = struct{}{}
			}
		}

		fmt.Println(corners)
		fmt.Println(cornerCount)
		fmt.Println()

		area += (x2 - x1) * (y2 - y1)
	}

	fmt.Println(minX, minY, maxX, maxY)
	fmt.Println(area)

	maxCorners := make([]corner, 0, 4)
	maxCorners = append(maxCorners,
		corner{
			X: minX,
			Y: maxY,
		}, corner{
			X: minX,
			Y: minY,
		}, corner{
			X: maxX,
			Y: maxY,
		}, corner{
			X: maxX,
			Y: minY,
		})

	for _, v := range maxCorners {
		if _, exist := cornerCount[v]; !exist {
			return false
		}
	}

	return area == (maxX-minX)*(maxY-minY) && len(cornerCount) == 4
}
