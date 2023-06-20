package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsEnglishPunctuation(r rune) bool {

	for _, punctuationRange := range codepoints.ENPunctuationRanges {
		if r >= punctuationRange[0] && r <= punctuationRange[1] {
			return true
		}
	}

	return false
}

func IsJapanesePunctuation(r rune) bool {

	for _, punctuationRange := range codepoints.JAPunctuationRanges {
		if r >= punctuationRange[0] && r <= punctuationRange[1] {
			return true
		}
	}

	return false
}
