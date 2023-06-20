package wanakana

import (
	"regexp"

	"github.com/deelawn/wanakana/internal/character"
)

func isLeadingWithoutInitialKana(s string, leading bool) bool {
	return leading && !IsKana(string([]rune(s)[0]))
}

func isTrailingWithoutFinalKana(s string, leading bool) bool {
	runes := []rune(s)
	return !leading && !IsKana(string(runes[len(runes)-1]))
}

func isInvalidMatcher(s string, matchKanji string) bool {

	if matchKanji != "" {
		for _, r := range []rune(matchKanji) {
			if character.IsKanji(r) {
				return false
			}
		}
		return true
	}

	return IsKana(s)
}

func StripOkurigana(s string, leading bool, matchKanji string) string {
	if !IsJapanese(s, nil) ||
		isLeadingWithoutInitialKana(s, leading) ||
		isTrailingWithoutFinalKana(s, leading) ||
		isInvalidMatcher(s, matchKanji) {
		return s
	}

	input := s
	if matchKanji != "" {
		input = matchKanji
	}

	tokens := Tokenize(input, false, false)
	var removeToken Token
	if leading {
		removeToken = tokens[0]
	} else {
		removeToken = tokens[len(tokens)-1]
	}

	var result string
	for _, token := range tokens {
		result += token.Value
	}

	regexString := removeToken.Value + "$"
	if leading {
		regexString = "^" + removeToken.Value
	}

	regex, err := regexp.Compile(regexString)
	if err != nil {
		return input
	}

	return regex.ReplaceAllString(s, "")
}
