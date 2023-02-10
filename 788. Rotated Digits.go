package main

func rotatedDigits(n int) int {
	dp := make([]int8, n+1)
	count := 0

	for i := 0; i <= n; i++ {
		if i < 10 {
			dp[i] = isGood(i)
			if dp[i] == 1 {
				count++
			}
			continue
		}

		if dp[i%10] == 0 || dp[i/10] == 0 {
			dp[i] = 0
			continue
		}

		if dp[i%10] == -1 && dp[i/10] == -1 {
			dp[i] = -1
			continue
		}

		dp[i] = 1
		count++
	}

	return count
}

func isGood(n int) int8 {
	switch n {
	case 0, 1, 8:
		return -1

	case 2, 5, 6, 9:
		return 1

	default:
		return 0
	}

}
