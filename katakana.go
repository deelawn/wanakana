package wanakana

import (
	"strings"

	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/transform"
)

func IsKatakana(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, r := range []rune(s) {
		if !character.IsKatakana(r) {
			return false
		}
	}

	return true
}

func ToKatakana(input string, options Options) string {

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
