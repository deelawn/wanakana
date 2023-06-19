package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsKatakana(r rune) bool {
	return r >= codepoints.KatakanaStart && r <= codepoints.KatakanaEnd
}
