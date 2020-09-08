package strings

func Reverse(text string) string {
	characters := []rune(text)
	lastIndex := len(text) - 1
	middle := lastIndex / 2

	for index := 0; index < middle; index++ {
		characters[index], characters[lastIndex-index] = characters[lastIndex-index], characters[index]
	}

	return string(characters)
}
