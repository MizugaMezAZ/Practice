package main

func canCompleteCircuit(gas []int, cost []int) int {
	for i := range gas {
		if gas[i] < cost[i] {
			continue
		}

		if find(i, gas, cost) {
			return i
		}
	}

	return -1
}

func find(first int, gas, cost []int) bool {
	tank := 0
	index := first

	for {
		tank += gas[index] - cost[index]

		if tank < 0 {
			return false
		}

		index++

		if index > len(gas) {
			index = 0
		}

		if index == first {
			break
		}
	}

	return true
}
