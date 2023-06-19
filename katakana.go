package wanakana

import "github.com/deelawn/wanakana/internal/character"

func IsKatakana(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, r := range []rune(s) {
		if !character.IsKatakana(r) {
			return false
		}
	}

	return true
}
