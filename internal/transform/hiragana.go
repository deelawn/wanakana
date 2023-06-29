package transform

import (
	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/codepoints"
)

func HiraganaToKatakana(input []rune) []rune {

	var result []rune
	for _, char := range input {
		if char == codepoints.ProlongedSoundMark || char == codepoints.KanaSlashDot {
			result = append(result, char)
			continue
		}

		if character.IsHiragana(char) {
			katakanaChar := char + (codepoints.KatakanaStart - codepoints.HiraganaStart)
			result = append(result, katakanaChar)
			continue
		}

		// Non-hiragana characters are passed through.
		result = append(result, char)
	}

	return result
}
