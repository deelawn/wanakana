package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsHiragana(r rune) bool {

	if r == codepoints.ProlongedSoundMark {
		return true
	}

	return r >= codepoints.HiraganaStart && r <= codepoints.HiraganaEnd
}
