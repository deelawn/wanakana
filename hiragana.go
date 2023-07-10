package wanakana

import "github.com/deelawn/wanakana/internal/character"

// IsHiragana returns true if all characters in the string are Hiragana.
func IsHiragana(input string) bool {

	if len(input) == 0 {
		return false
	}

	for _, r := range []rune(input) {
		if !character.IsHiragana(r) {
			return false
		}
	}

	return true
}
