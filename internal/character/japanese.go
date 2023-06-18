package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsJapanese(r rune) bool {

	for _, jpRange := range codepoints.JapaneseRanges {
		if r >= jpRange[0] && r <= jpRange[1] {
			return true
		}
	}

	return false
}
