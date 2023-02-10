package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	rns := []rune(s)
	st := []rune{}
	for i := 0; i <= len(rns)/2; i++ {
		st = append(st, rns[i])

		if len(rns)%len(st) != 0 {
			continue
		}

		tmp := s
		count := 0

	tmploop:
		ok, tmp := findPrefix(tmp, string(st))
		if !ok {
			continue
		}
		count++
		if tmp == "" {
			if count > 1 {
				return true
			}
			return false
		}

		if !strings.Contains(tmp, string(st)) {
			continue
		}

		goto tmploop
	}
	return false
}

func findPrefix(s, pre string) (b bool, sub string) {
	ok := strings.HasPrefix(s, pre)
	if !ok {
		return false, ""
	}
	return true, strings.TrimPrefix(s, pre)

	//-------------------------------------------
	// str := s + s
	// str = str[1 : len(str)-1]
	// fmt.Println(str)
	// if strings.Contains(str, s) {
	// 	return true
	// }

	// return false

	//-------------------------------------------
	// 	s1 := []rune(s)

	// con:
	// 	for i := 1; i <= len(s1)/2; i++ {
	// 		if len(s1)%i == 0 {
	// 			for j := i; j < len(s1); j += i {
	// 				for k := 0; k < i; k++ {
	// 					if s1[k] != s1[k+j] {
	// 						continue con
	// 					}
	// 				}
	// 			}

	// 			return true
	// 		}
	// 	}

	// 	return false

	//-------------------------------------------

	// s1 := []rune(s)
	// n := make([]rune, 0)
	// n = append(n, s1[0])

	// for i := 1; i < len(s1); i++ {

	// 	if s1[i] != n[0] {
	// 		n = append(n, s1[i])
	// 	} else {
	// 		if (len(s1)-i)%len(n) == 0 {

	// 			index := i

	// 			for index != len(s1) {

	// 				for _, v := range n {
	// 					if v == s1[index] {
	// 						index++
	// 					} else {
	// 						goto AAA
	// 					}
	// 				}
	// 			}

	// 			return true

	// 		}
	// 	AAA:
	// 		n = append(n, s1[i])

	// 	}

	// }

	// return false
}
