package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsKanji(r rune) bool {
	return r >= codepoints.KanjiStart && r <= codepoints.KanjiEnd
}
