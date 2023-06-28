package wanakana

import (
	"strings"

	"github.com/deelawn/wanakana/internal/character"
	"github.com/deelawn/wanakana/internal/transform"
	"github.com/deelawn/wanakana/internal/tree"
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

	if treeMap == nil {
		treeMap = createRomajiToKanaTree(options.IMEMode, options.UseObsoleteKana, options.CustomKanaMapping)
	}

	tokens := transform.ToKanaToken([]rune(strings.ToLower(input)), treeMap, false)

	var result string
	for _, token := range tokens {
		result += token.Value
	}

	return result
}

func createRomajiToKanaTree(imeMode, useObsoleteKana bool, customKanaMapping CustomMapping) *tree.Map {

	treeMap := transform.GetRomajiToKanaTree()
	if imeMode {
		treeMap = transform.ToIMEModeTree(treeMap)
	}
	if useObsoleteKana {
		treeMap = transform.ToTreeWithObsoleteKana(treeMap)
	}

	if customKanaMapping != nil {
		customKanaMapping.Apply(treeMap)
	}

	return treeMap
}
