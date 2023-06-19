package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsRomaji(r rune) bool {

	for _, romajiRange := range codepoints.RomajiRanges {
		if r >= romajiRange[0] && r <= romajiRange[1] {
			return true
		}
	}

	return false
}
