package isogram

import (
	"unicode"
)

func IsIsogram(word string) bool {
	var doesItExist = map[rune]bool{}
	for _, letter := range word {
		if !unicode.IsLetter(letter){
			continue
		}
		lowerCharacter := unicode.ToLower(letter)

		if doesItExist[lowerCharacter]{
			return false
		}
		doesItExist[lowerCharacter] = true
	}
	return true
}
