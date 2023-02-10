package main

func numIslands(grid [][]byte) int {
	var n int

	for i, j := range grid {
		for k, l := range j {
			if l == '1' {
				findisland(grid, i, k)
				n++
			}
		}
	}

	return n
}

func findisland(m [][]byte, i, j int) {
	// g := m

	if i < 0 || i > len(m)-1 || j < 0 || j > len(m[0])-1 {
		return
	}

	if m[i][j] != '1' {
		return
	}

	m[i][j] = '0'

	findisland(m, i-1, j)
	findisland(m, i+1, j)
	findisland(m, i, j-1)
	findisland(m, i, j+1)
}
