package wanakana

import "github.com/deelawn/wanakana/internal/character"

func IsKana(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, r := range []rune(s) {
		if character.IsHiragana(r) || character.IsKatakana(r) {
			continue
		}

		return false
	}

	return true
}
