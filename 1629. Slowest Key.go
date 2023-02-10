package main

func slowestKey(releaseTimes []int, keysPressed string) byte {
	key := []byte(keysPressed)

	max := releaseTimes[0]
	ans := key[0]

	for i := 1; i < len(releaseTimes); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]

		if duration > max || (duration == max && key[i] > ans) {
			max = duration
			ans = key[i]
		}

	}

	return ans
}
