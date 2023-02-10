package main

func productQueries(n int, queries [][]int) []int {
	power := []int{}
	xx := 1
	for n != 0 {
		if n&1 == 1 {
			power = append(power, xx)
		}

		n = n >> 1
		xx *= 2
	}

	ans := make([]int, 0)
	for _, v := range queries {
		n := 1

		if v[0] == v[1] {
			n = power[v[0]] % 1000000007

		} else {
			for i := v[0]; i <= v[1]; i++ {
				n = n * power[i] % 1000000007

			}
		}

		ans = append(ans, n)
	}

	return ans
}
