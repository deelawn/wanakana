package wanakana

import "github.com/deelawn/wanakana/internal/character"

// IsKanji returns true if all runes in the string are Kanji.
func IsKanji(input string) bool {

	if len(input) == 0 {
		return false
	}

	for _, r := range []rune(input) {
		if !character.IsKanji(r) {
			return false
		}
	}

	return true
}
