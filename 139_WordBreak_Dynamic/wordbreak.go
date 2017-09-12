package wordbreak

func wordbreak(str string) bool {
	// str = "heyaman"
	dict := [2]string{"hey", "man"}

	//temp := [6]bool{}

	wb := [...]bool{}

	for i := 1; i <= len(str); i++ {
		if !wb[i] && dictionaryContains(dict, str[0:i]) {
			wb[i] = true
		}

		if i == len(str) {
			return true
		}

		if wb[i] {
			for j := i + 1; j <= len(str); j++ {
				if !wb[i] && dictionaryContains(dict, str[i:j-i]) {
					wb[i] = true
				}

				if j == len(str) && wb[i] {
					return true
				}
			}
		}

	}

	return false
}

func dictionaryContains(dict [2]string, word string) bool {
	for i := 0; i < len(dict); i++ {
		if dict[i] == word {
			return true
		}
	}
	return false
}
