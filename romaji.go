package wanakana

import (
	"regexp"

	"github.com/deelawn/wanakana/internal/character"
)

func IsRomaji(s string, regex *regexp.Regexp) bool {

	if len(s) == 0 {
		return false
	}

	for _, r := range []rune(s) {
		if character.IsRomaji(r) {
			// This character is Romaji; keep going.
			continue
		}

		if regex != nil && regex.MatchString(string(r)) {
			// This character isn't Romaji but matches the regex; keep going.
			continue
		}

		return false
	}

	return true
}
