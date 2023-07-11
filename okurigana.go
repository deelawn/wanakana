package wanakana

import (
	"regexp"

	"github.com/deelawn/wanakana/internal/character"
)

func isLeadingWithoutInitialKana(input string, leading bool) bool {
	return leading && !IsKana(string([]rune(input)[0]))
}

func isTrailingWithoutFinalKana(input string, leading bool) bool {
	runes := []rune(input)
	return !leading && !IsKana(string(runes[len(runes)-1]))
}

func isValidMatcher(input string, matchKanji string) bool {

	if matchKanji != "" {
		for _, r := range []rune(matchKanji) {
			if character.IsKanji(r) {
				return true
			}
		}
		return false
	}

	return !IsKana(input)
}

// StripOkurigana removes leading or trailing kana from a string.
func StripOkurigana(input string, leading bool, matchKanji string) string {
	if !IsJapanese(input, nil) ||
		isLeadingWithoutInitialKana(input, leading) ||
		isTrailingWithoutFinalKana(input, leading) ||
		!isValidMatcher(input, matchKanji) {
		return input
	}

	modifiedInput := input
	if matchKanji != "" {
		modifiedInput = matchKanji
	}

	tokens := Tokenize(modifiedInput, false, false)
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

	return regex.ReplaceAllString(input, "")
}
