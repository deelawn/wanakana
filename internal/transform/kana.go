package transform

import (
	"sync"

	"github.com/deelawn/wanakana/config"
	"github.com/deelawn/wanakana/tree"
)

// KTR = kana to romaji

var (
	basicKTRMapping = map[string]string{
		"あ": "a", "い": "i", "う": "u", "え": "e", "お": "o",
		"か": "ka", "き": "ki", "く": "ku", "け": "ke", "こ": "ko",
		"さ": "sa", "し": "shi", "す": "su", "せ": "se", "そ": "so",
		"た": "ta", "ち": "chi", "つ": "tsu", "て": "te", "と": "to",
		"な": "na", "に": "ni", "ぬ": "nu", "ね": "ne", "の": "no",
		"は": "ha", "ひ": "hi", "ふ": "fu", "へ": "he", "ほ": "ho",
		"ま": "ma", "み": "mi", "む": "mu", "め": "me", "も": "mo",
		"ら": "ra", "り": "ri", "る": "ru", "れ": "re", "ろ": "ro",
		"や": "ya", "ゆ": "yu", "よ": "yo",
		"わ": "wa", "ゐ": "wi", "ゑ": "we", "を": "wo",
		"ん": "n",
		"が": "ga", "ぎ": "gi", "ぐ": "gu", "げ": "ge", "ご": "go",
		"ざ": "za", "じ": "ji", "ず": "zu", "ぜ": "ze", "ぞ": "zo",
		"だ": "da", "ぢ": "ji", "づ": "zu", "で": "de", "ど": "do",
		"ば": "ba", "び": "bi", "ぶ": "bu", "べ": "be", "ぼ": "bo",
		"ぱ": "pa", "ぴ": "pi", "ぷ": "pu", "ぺ": "pe", "ぽ": "po",
		"ゔぁ": "va", "ゔぃ": "vi", "ゔ": "vu", "ゔぇ": "ve", "ゔぉ": "vo",
	}

	ktrSpecialSymbolsMapping = map[rune]rune{
		'。': '.',
		'、': ',',
		'：': ':',
		'・': '/',
		'！': '!',
		'？': '?',
		'〜': '~',
		'ー': '-',
		'「': '‘',
		'」': '’',
		'『': '“',
		'』': '”',
		'［': '[',
		'］': ']',
		'（': '(',
		'）': ')',
		'｛': '{',
		'｝': '}',
		'　': ' ',
	}

	ktrAmbiguousVowels = []rune{'あ', 'い', 'う', 'え', 'お', 'や', 'ゆ', 'よ'}
	ktrSmallY          = map[rune]string{'ゃ': "ya", 'ゅ': "yu", 'ょ': "yo"}
	ktrSmallYExtra     = map[rune]string{'ぃ': "yi", 'ぇ': "ye"}
	ktrSmallAIUEO      = map[rune]rune{
		'ぁ': 'a',
		'ぃ': 'i',
		'ぅ': 'u',
		'ぇ': 'e',
		'ぉ': 'o',
	}

	ktrYoonKana       = []rune{'き', 'に', 'ひ', 'み', 'り', 'ぎ', 'び', 'ぴ', 'ゔ', 'く', 'ふ'}
	ktrYoonExceptions = map[rune]string{
		'し': "sh",
		'ち': "ch",
		'じ': "j",
		'ぢ': "j",
	}

	ktrSmallKana = map[rune]string{
		'っ': "",
		'ゃ': "ya",
		'ゅ': "yu",
		'ょ': "yo",
		'ぁ': "a",
		'ぃ': "i",
		'ぅ': "u",
		'ぇ': "e",
		'ぉ': "o",
	}

	// sukuonWhiteList needs to be a rune to rune map because 'c' and 't' both map to 't'.
	// The following is a comment from the original JS source:
	// going with the intuitive (yet incorrect) solution where っや -> yya and っぃ -> ii
	// in other words, just assume the sokuon could have been applied to anything
	sokuonWhitelist = map[rune]rune{
		'b': 'b',
		'c': 't',
		'd': 'd',
		'f': 'f',
		'g': 'g',
		'h': 'h',
		'j': 'j',
		'k': 'k',
		'm': 'm',
		'p': 'p',
		'q': 'q',
		'r': 'r',
		's': 's',
		't': 't',
		'v': 'v',
		'w': 'w',
		'x': 'x',
		'z': 'z',
	}
)

var (
	ktrTreeMap   *tree.Map
	ktrTreeMapMu sync.Mutex
)

func GetKanaToRomajiTreeMap(romanization config.Romanization) *tree.Map {

	ktrTreeMapMu.Lock()
	defer ktrTreeMapMu.Unlock()

	if ktrTreeMap == nil {
		ktrTreeMap = createKanaToRomajiTreeMap(romanization)
	}

	return ktrTreeMap
}

func createKanaToRomajiTreeMap(romanization config.Romanization) *tree.Map {

	// They are requesting an unsupported romanization system.
	if romanization != config.RomanizationHepburn {
		return nil
	}

	treeMap := new(tree.Map)
	for kana, romaji := range basicKTRMapping {
		treeMap.PutValue([]rune(kana), romaji)
	}

	for kana, romaji := range ktrSpecialSymbolsMapping {
		treeMap.PutValue([]rune{kana}, string(romaji))
	}

	for kana, romaji := range ktrSmallY {
		treeMap.PutValue([]rune{kana}, romaji)
	}

	for kana, romaji := range ktrSmallAIUEO {
		treeMap.PutValue([]rune{kana}, string(romaji))
	}

	for _, kana := range ktrYoonKana {

		// Get the first romaji character that corresponds to the kana that already exists in the tree.
		kanaMapping := treeMap.GetValue([]rune{kana})
		if kanaMappingRunes := []rune(kanaMapping); len(kanaMappingRunes) > 0 {
			kanaMapping = string(kanaMappingRunes[0])
		}

		// Add entries to the tree for that existing first character plus the small y kana.
		for yKana, yRomaji := range ktrSmallY {
			treeMap.PutValue([]rune{kana, yKana}, kanaMapping+yRomaji)
		}

		for yKana, yRomaji := range ktrSmallYExtra {
			treeMap.PutValue([]rune{kana, yKana}, kanaMapping+yRomaji)
		}
	}

	for kana, romaji := range ktrYoonExceptions {
		for yKana, yRomaji := range ktrSmallY {
			treeMap.PutValue([]rune{kana, yKana}, romaji+yRomaji[1:])
		}

		treeMap.PutValue([]rune{kana, 'ぃ'}, romaji+"yi")
		treeMap.PutValue([]rune{kana, 'ぇ'}, romaji+"e")
	}

	treeMap.Branches['っ'] = resolveTsu(treeMap)

	for kana, romaji := range ktrSmallKana {
		treeMap.PutValue([]rune{kana}, romaji)
	}

	for _, vowel := range ktrAmbiguousVowels {
		treeMap.PutValue([]rune{'ん', vowel}, "n"+treeMap.GetValue([]rune{vowel}))
	}

	return treeMap
}

func resolveTsu(treeMap *tree.Map) *tree.Map {

	newTreeMap := new(tree.Map)
	if treeMap.Branches == nil {
		if treeMap.Value == nil {
			return nil
		}

		consonant := rune((*treeMap.Value)[0])
		if sokuon, ok := sokuonWhitelist[consonant]; ok {
			newValue := string(sokuon) + *treeMap.Value
			newTreeMap.Value = &newValue
			return newTreeMap
		}
	}

	newTreeMap.Branches = make(map[rune]*tree.Map)
	for k, v := range treeMap.Branches {
		newTreeMap.Branches[k] = resolveTsu(v)
	}

	return newTreeMap
}

type ConvertedKanaToken struct {
	Start int
	End   int
	Value string
}

func ToKanaToken(input []rune, treeMap *tree.Map, convertEnding bool) (result []ConvertedKanaToken) {

	// TODO: Figure out how convertEnding is supposed to work.

	if treeMap == nil {
		return
	}

	// Iterate over the input characters while calling parseKanaToken to try to build the
	// longest chain of characters possible that matches an entry in the tree map.
	for i := 0; i < len(input); i++ {
		nextConvertedToken := parseKanaToken(input[i:], treeMap, convertEnding)

		// A nil value indicates that the value wasn't found in the tree, so pass it through.
		if nextConvertedToken == nil {
			result = append(
				result,
				ConvertedKanaToken{
					Start: i,
					End:   i + 1,
					Value: string(input[i]),
				},
			)
			continue
		}

		// The start and end values always need to be offset by i because the converted token start
		// and end indexes are always relative to the character slice that is passed to parseKanaToken.
		nextConvertedToken.Start += i
		nextConvertedToken.End += i
		result = append(result, *nextConvertedToken)

		// Subtract one because the loop increases by one each iteration.
		i = nextConvertedToken.End - 1

		// The last two characters of input would have been a duplicate such as "tt", so if
		// the full string was "hattsu", then we need to process the second "t" again because
		// it is part of the next token.
		if nextConvertedToken.Value == "っ" && nextConvertedToken.End-nextConvertedToken.Start == 2 &&
			input[nextConvertedToken.Start] == input[nextConvertedToken.Start+1] {
			i--
		}

		// // The last two characters of input would have been a duplicate such as "tt", so if
		// // the full string was "hattsu", then we need to process the second "t" again because
		// // it is part of the next token.
		// if nextConvertedToken.Value == "っ" {
		// 	i--
		// }
	}

	return result
}

func parseKanaToken(input []rune, treeMap *tree.Map, convertEnding bool) *ConvertedKanaToken {

	nextTreeMap := treeMap.GetMap([]rune{input[0]})
	if nextTreeMap == nil {
		return nil
	}

	// We are at the end of the input, so there is either a match or it returns nil.
	if len(input) == 1 {
		if value := nextTreeMap.Value; value != nil {
			return &ConvertedKanaToken{
				Start: 0,
				End:   1,
				Value: *value,
			}
		}

		return nil
	}

	convertedToken := parseKanaToken(input[1:], nextTreeMap, convertEnding)
	if convertedToken != nil {
		// parseKanaToken is called recusrively, so each time it returns a non-nil converted token,
		// the end index needs to be incremented by one because the returning invocation has no
		// knowledge of the number of invocations that preceded it.
		convertedToken.End += 1
		return convertedToken
	}

	// The previous invocation tried to find a match on the next character but didn't.
	// If the current character has a match then return it. Otherwise nil will be returned again
	// and it will try to match backwards until one or no value matches are found.
	if convertedToken == nil && nextTreeMap.Value != nil {
		convertedToken = &ConvertedKanaToken{
			Start: 0,
			End:   1,
			Value: *nextTreeMap.Value,
		}
	}

	return convertedToken
}
