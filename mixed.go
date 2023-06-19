package wanakana

import "github.com/deelawn/wanakana/internal/character"

func IsMixed(s string, passKanji bool) bool {

	if len(s) == 0 {
		return false
	}

	var hasKanji, hasHiragana, hasKatakana, hasRomaji bool
	for _, r := range []rune(s) {

		// Check if this character is a kanji character if `passKanji` is set and we haven't previously
		// identified a kanji character.
		if !passKanji && !hasKanji {
			if character.IsKanji(r) {
				return false
			}
		}

		if !hasHiragana && character.IsHiragana(r) {
			hasHiragana = true
			continue
		}

		if !hasKatakana && character.IsKatakana(r) {
			hasKatakana = true
			continue
		}

		if !hasRomaji && character.IsRomaji(r) {
			hasRomaji = true
			continue
		}
	}

	return (hasHiragana || hasKatakana) && hasRomaji
}
