package main

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	dir := true
	matx := len(mat) - 1
	maty := len(mat[0]) - 1
	x, y := 0, 0
	fmt.Printf("matx: %v maty: %v\n", matx, maty)

	ans := make([]int, 0, matx*maty)

	for x != matx || y != maty {
		fmt.Println(x, y, dir)
		ans = append(ans, mat[x][y])

		if dir {
			if y == maty {
				x++
				dir = false
				fmt.Println("y == matx")
				continue
			}

			if x == 0 {
				y++
				dir = false
				fmt.Println("x == 0")
				fmt.Println()
				continue
			}

			x--
			y++
			fmt.Println("x--, y++")
		} else {

			if x == matx {
				y++
				dir = true
				fmt.Println("x == matx")
				fmt.Println()
				continue
			}

			if y == 0 {
				x++
				dir = true
				fmt.Println("y == 0")
				fmt.Println()
				continue
			}

			x++
			y--
			fmt.Println("x++ y-- ")
			fmt.Println()
		}

		fmt.Println()

	}

	ans = append(ans, mat[matx][maty])

	return ans
}
