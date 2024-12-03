func isPrefixOfWord(sentence string, searchWord string) int {
	// 1. split the sentence into array, separator: space
	words := strings.Split(sentence, " ")

	// 2. loop for each word
	for index, word := range words {
		// 2.1. check if the word contains searchWord at the first index
		if strings.Index(word, searchWord) == 0 {
			// 2.1.1 return index+1
			return index + 1
		}
	}

	return -1
}