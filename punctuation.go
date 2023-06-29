package wanakana

import "github.com/deelawn/wanakana/internal/character"

func hasEnglishPunctuation(input string) bool {

	for _, char := range []rune(input) {
		if character.IsEnglishPunctuation(char) {
			return true
		}
	}

	return false
}
