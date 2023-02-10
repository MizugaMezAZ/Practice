package main

import "fmt"

func PointCal(points []int) int {
	total := 0

	round := 1
	for {
		total += points[(round-1)*2] + points[(round-1)*2+1]

		// 第十局單獨處理
		if round == 10 {
			break
		}

		if isStrike(points, round) {
			total += points[(round-1+1)*2] + points[(round-1+1)*2+1] // 全倒獎勵下兩球分數

			if isStrike(points, round+1) { // 全倒計為 0, 10 但只有一球 所以再往下一回合找分數
				total += points[(round-1+2)*2]

				if isStrike(points, round+2) { // 因為全倒為0, 10所以上面那行會是+0 補上分數
					total += 10
				}
			}

		}

		if isSpare(points, round) { // 補全倒獎勵一球分數
			total += points[(round-1+1)*2]

			if isStrike(points, round+1) { // 全倒為0, 10所以上面那行會是+0 補上分數
				total += 10
			}
		}
		fmt.Printf("第 %d 局 累計分數: %d \n", round, total)
		round++
	}

	// 第十局有全倒或補全倒 +上最後一球分數
	if points[(round-1)*2] == 10 || points[(round-1)*2]+points[(round-1)*2+1] == 10 {
		return total + points[round*2]
	}

	fmt.Printf("第 %d 局 累計分數: %d \n", round, total)

	return total
}

func isStrike(points []int, round int) bool {
	return points[(round-1)*2] == 0 && points[(round-1)*2+1] == 10
}

func isSpare(points []int, round int) bool {
	return points[(round-1)*2] > 0 && points[(round-1)*2]+points[(round-1)*2+1] == 10
}
