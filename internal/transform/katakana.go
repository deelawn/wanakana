package transform

import (
	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/codepoints"
	"github.com/deelawn/wanakana/internal/tree"
)

var longVowels = map[rune]rune{
	'a': 'あ',
	'i': 'い',
	'u': 'う',
	'e': 'え',
	'o': 'う',
}

func isCharInitialLongDash(r rune, index int) bool {
	return r == codepoints.ProlongedSoundMark && index < 1
}

func isCharInnerLongDash(r rune, index int) bool {
	return r == codepoints.ProlongedSoundMark && index > 0
}

func KatakanaToHiragana(s string, treeMap *tree.Map, isDestinationRomaji, convertLongVowelMark bool) string {

	var (
		previousKana rune
		result       string
	)

	runes := []rune(s)
	for i, r := range runes {
		if r == codepoints.KanaSlashDot || isCharInitialLongDash(r, i) || r == 'ヶ' || r == 'ヵ' {
			result += string(r)
			continue
		}

		if convertLongVowelMark && previousKana != 0 && isCharInnerLongDash(r, i) {
			previousKanaRomaji := treeMap.GetValue(string(previousKana))
			romajiVowel := rune(previousKanaRomaji[len(previousKanaRomaji)-1])

			if character.IsKatakana(runes[i-1]) && romajiVowel == 'o' && isDestinationRomaji {
				result += "お"
				continue
			}

			result += string(longVowels[romajiVowel])
			continue
		}

		if r != codepoints.ProlongedSoundMark && character.IsKatakana(r) {
			hiraganaChar := r + (codepoints.HiraganaStart - codepoints.KatakanaStart)
			previousKana = hiraganaChar
			result += string(hiraganaChar)
			continue
		}

		// Pass non katakana chars through.
		previousKana = 0
		result += string(r)
	}

	return result
}
