package main

func isAlienSorted(words []string, order string) bool {
	orderIndex := make(map[byte]int)

	for i := 0; i < len(order); i++ {
		orderIndex[order[i]] = i
	}

	for i := 1; i < len(words); i++ {
		index := 0
		word1 := words[i-1]
		word2 := words[i]
		samePrefix := true

		for index < len(word1) && index < len(word2) {
			if word1[index] != word2[index] {
				if orderIndex[word1[index]] > orderIndex[word2[index]] {
					return false
				}

				samePrefix = false

				break
			}

			index++
		}

		if samePrefix && len(word1) > len(word2) {
			return false
		}
	}

	return true
}
