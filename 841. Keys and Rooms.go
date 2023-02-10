package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true

	for _, v := range rooms[0]{
		visitRoom(v, visited, rooms)
	}

	for _, v := range visited {
		if !v {
			return false
		}
	}

    return true
}

func visitRoom(n int, visited []bool, rooms[][]int) {
	if visited[n] {
		return
	}

	visited[n] = true

	for _, v := range rooms[n] {
		visitRoom(v, visited, rooms)
	}
}