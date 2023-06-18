package wanakana

import (
	"regexp"

	"github.com/deelawn/wanakana/internal/character"
)

// IsJapanese returns true if all characters in the string are Japanese or match
// the optional regular expression.
func IsJapanese(s string, regex *regexp.Regexp) bool {

	for _, r := range []rune(s) {
		isJP := character.IsJapanese(r)
		if isJP {
			// This character is Japanese; keep going.
			continue
		}

		if regex != nil && regex.MatchString(string(r)) {
			// This character isn't Japanese but matches the regex; keep going.
			continue
		}

		return false
	}

	return true
}
