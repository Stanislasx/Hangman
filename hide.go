package hangman

func HideWord(word string, revealed []bool) string {

	// hide word

	hidden := ""
	for i, c := range word {
		if revealed[i] {
			hidden += string(c)
		} else {
			hidden += "_"
		}
	}
	return hidden
}
