package transform

import "github.com/deelawn/wanakana/internal/codepoints"

func isCharInitialLongDash(r rune, index int) bool {
	return r == codepoints.ProlongedSoundMark && index < 1
}

func isCharInnerLongDash(r rune, index int) bool {
	return r == codepoints.ProlongedSoundMark && index > 0
}

func KatakanaToHiragana(s string, isDestinationRomaji, convertLongVowelMark bool) string {

	var (
		previousKana rune
		result       string
	)

	for i, r := range []rune(s) {
		if r == codepoints.KanaSlashDot || isCharInitialLongDash(r, i) || r == 'ヶ' || r == 'ヵ' {
			result += string(r)
			continue
		}

		if convertLongVowelMark && previousKana != 0 && isCharInnerLongDash(r, i) {

		}
	}

	return result
}
