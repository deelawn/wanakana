package wanakana

import (
	"regexp"
	"strings"

	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/transform"
	"github.com/deelawn/wanakana/tree"
)

// IsRomaji returns true if all characters in the string are Romaji or match
// the optional regular expression.
func IsRomaji(input string, regex *regexp.Regexp) bool {

	if len(input) == 0 {
		return false
	}

	for _, r := range []rune(input) {
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

// ToRomaji converts input to romaji with the option to uppercase katakana.
func ToRomaji(input string, options Options, treeMap *tree.Map) string {

	if treeMap == nil {
		treeMap = createKanaToRomajiTree(options.Romanization, options.CustomKanaMapping)
	}

	inputRunes := []rune(input)
	hiraganaInput := transform.KatakanaToHiragana(input, treeMap, true, !options.IgnoreLongVowelMark)
	tokens := transform.ToKanaToken([]rune(hiraganaInput), treeMap, false)

	var result string
	for _, token := range tokens {
		if options.UppercaseKatakana && IsKatakana(string(inputRunes[token.Start:token.End])) {
			token.Value = strings.ToUpper(token.Value)
		}

		result += token.Value
	}

	return result
}

func createKanaToRomajiTree(romanization Romanization, customMapping CustomMapping) *tree.Map {

	treeMap := transform.GetKanaToRomajiTreeMap(string(romanization))
	if customMapping != nil {
		treeMap = treeMap.Copy()
		customMapping.Apply(treeMap)
	}

	return treeMap
}
