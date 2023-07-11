package wanakana

import (
	"strings"

	"github.com/deelawn/wanakana/config"
	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/transform"
)

// IsKatana returns true if all runes in the string are katakana.
func IsKatakana(input string) bool {

	if len(input) == 0 {
		return false
	}

	for _, r := range []rune(input) {
		if !character.IsKatakana(r) {
			return false
		}
	}

	return true
}

// ToKatakana converts input to katakana with the option to pass romaji runes through untransformed.
func ToKatakana(input string, options config.Options) string {

	if len(input) == 0 {
		return ""
	}

	if options.PassRomaji {
		return string(transform.HiraganaToKatakana([]rune(input)))
	}

	if IsMixed(input, true) || IsRomaji(input, nil) || hasEnglishPunctuation(input) {
		input = ToKana(strings.ToLower(input), options, nil)
	}

	return string(transform.HiraganaToKatakana([]rune(input)))
}
