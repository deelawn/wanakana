package wanakana

import "github.com/deelawn/wanakana/internal/character"

// IsHiragana returns true if all characters in the string are Hiragana.
func IsHiragana(s string) bool {

	for _, r := range []rune(s) {
		if !character.IsHiragana(r) {
			return false
		}
	}

	return true
}
