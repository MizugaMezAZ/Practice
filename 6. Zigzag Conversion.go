package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1 {
		return s
	}

	result := make([]string, numRows)
	cross := numRows + numRows - 2

	result[0] += string(s[0])

	for i := 1; i < len(s); i++ {
		index := i % cross
		if index < numRows {
			result[index] += string(s[i])
			continue
		}

		result[numRows-index%numRows-2] += string(s[i])
	}

	fmt.Println(result)

	var answer string
	for i := 0; i < numRows; i++ {
		answer += result[i]
	}
	return answer
}

// func convert(s string, numRows int) string {
// 	if len(s) <= numRows {
// 		return s
// 	}

// 	up := false
// 	n := 0

// 	result := make(map[int][]rune)

// 	b := []rune(s)

// 	for _, v := range b {
// 		result[n] = append(result[n], v)

// 		if n == 0 {
// 			up = true
// 		} else if n == numRows-1 {
// 			up = false
// 		}

// 		if up {
// 			n++
// 		} else {
// 			n--
// 		}
// 	}

// 	var answer string
// 	fmt.Println(result)
// 	for i := 0; i < numRows; i++ {
// 		fmt.Println(result[i])
// 		answer += string(result[i])
// 	}

// 	return answer
// }
