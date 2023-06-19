package wanakana

import "github.com/deelawn/wanakana/internal/character"

func IsKanji(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, r := range []rune(s) {
		if !character.IsKanji(r) {
			return false
		}
	}

	return true
}
