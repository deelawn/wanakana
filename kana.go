package wanakana

import (
	"strings"

	"github.com/deelawn/wanakana/config"
	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/transform"
	"github.com/deelawn/wanakana/tree"
)

// IsKana returns true if every rune in the input string is hiragana or katakana.
func IsKana(input string) bool {

	if len(input) == 0 {
		return false
	}

	for _, r := range []rune(input) {
		if character.IsHiragana(r) || character.IsKatakana(r) {
			continue
		}

		return false
	}

	return true
}

// ToKana converts input to kana (hiragana and/or katakana).
func ToKana(input string, options config.Options, treeMap *tree.Map) string {

	if treeMap == nil {
		treeMap = createRomajiToKanaTree(options.IMEMode, options.UseObsoleteKana, options.CustomKanaMapping)
	}

	tokens := transform.ToKanaToken(
		[]rune(strings.ToLower(input)),
		treeMap,
		options.IMEMode != config.ToKanaMethodNone,
	)

	var result string
	for _, token := range tokens {
		enforceHiragana := options.IMEMode == config.ToKanaMethodHiragana
		// Katakana preference can be delineated by making an intended token all uppercase.
		enforceKatakana := options.IMEMode == config.ToKanaMethodKatakana ||
			strings.ToUpper(string(input[token.Start:token.End])) == string(input[token.Start:token.End])

		if enforceHiragana || !enforceKatakana {
			result += token.Value
			continue
		}

		result += string(transform.HiraganaToKatakana([]rune(token.Value)))
	}

	return result
}

func createRomajiToKanaTree(
	imeMode config.ToKanaMethod,
	useObsoleteKana bool,
	customKanaMapping config.CustomMapping,
) *tree.Map {

	treeMap := transform.GetRomajiToKanaTree()
	if imeMode != config.ToKanaMethodNone {
		treeMap = transform.ToIMEModeTree(treeMap)
	}
	if useObsoleteKana {
		treeMap = transform.ToTreeWithObsoleteKana(treeMap)
	}

	if customKanaMapping != nil {
		treeMap = treeMap.Copy()
		customKanaMapping.Apply(treeMap)
	}

	return treeMap
}
