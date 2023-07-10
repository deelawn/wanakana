package wanakana

import (
	"strings"

	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/transform"
	"github.com/deelawn/wanakana/tree"
)

func IsKana(s string) bool {

	if len(s) == 0 {
		return false
	}

	for _, r := range []rune(s) {
		if character.IsHiragana(r) || character.IsKatakana(r) {
			continue
		}

		return false
	}

	return true
}

func ToKana(input string, options Options, treeMap *tree.Map) string {

	// TODO:
	// treemap can't be an internal type if it is being used as input.
	// Or use an interface instead.
	if treeMap == nil {
		treeMap = createRomajiToKanaTree(options.IMEMode, options.UseObsoleteKana, options.CustomKanaMapping)
	}

	tokens := transform.ToKanaToken([]rune(strings.ToLower(input)), treeMap, !(options.IMEMode == ToKanaMethodNone))

	var result string
	for _, token := range tokens {
		enforceHiragana := options.IMEMode == ToKanaMethodHiragana
		// Katakana preference can be delineated by making an intended token all uppercase.
		enforceKatakana := options.IMEMode == ToKanaMethodKatakana ||
			strings.ToUpper(string(input[token.Start:token.End])) == string(input[token.Start:token.End])

		if enforceHiragana || !enforceKatakana {
			result += token.Value
			continue
		}

		result += string(transform.HiraganaToKatakana([]rune(token.Value)))
	}

	return result
}

func createRomajiToKanaTree(imeMode ToKanaMethod, useObsoleteKana bool, customKanaMapping CustomMapping) *tree.Map {

	treeMap := transform.GetRomajiToKanaTree()
	if imeMode != ToKanaMethodNone {
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
