package character

import "github.com/deelawn/wanakana/internal/codepoints"

func IsSlashDot(r rune) bool {
	return r == codepoints.KanaSlashDot
}
