package main

func permute(nums []int) [][]int {
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))
	ans := make([][]int, 0)

	back(len(nums), nums, path, used, &ans)

	return ans
}

func back(targetLen int, nums []int, path []int, used []bool, ans *[][]int) {
	if len(path) == targetLen {
		temp := make([]int, targetLen)
		copy(temp, path)

		*ans = append(*ans, temp)
		return
	}

	for i := 0; i < targetLen; i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])

		used[i] = true
		back(targetLen, nums, path, used, ans)
		used[i] = false

		path = path[:len(path)-1]
	}
}
