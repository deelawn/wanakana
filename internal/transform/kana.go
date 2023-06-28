package transform

import "github.com/deelawn/wanakana/internal/tree"

var (
	basicKanaToRomajiMapping = map[string]string{
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

	kanaToRomajiSpecialSymbolsMapping = map[rune]rune{
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

	ambiguousVowels = []rune{'あ', 'い', 'う', 'え', 'お', 'や', 'ゆ', 'よ'}
	smallY          = map[rune]string{'ゃ': "ya", 'ゅ': "yu", 'ょ': "yo"}
	smallYExtra     = map[rune]string{'ぃ': "yi", 'ぇ': "ye"}
	smallAIUEO      = map[rune]rune{
		'ぁ': 'a',
		'ぃ': 'i',
		'ぅ': 'u',
		'ぇ': 'e',
		'ぉ': 'o',
	}

	yoonKana       = []rune{'き', 'に', 'ひ', 'み', 'り', 'ぎ', 'び', 'ぴ', 'ゔ', 'く', 'ふ'}
	yoonExceptions = map[rune]string{
		'し': "sh",
		'ち': "ch",
		'じ': "j",
		'ぢ': "j",
	}

	smallKana = map[rune]string{
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

type ConvertedKanaToken struct {
	Start int
	End   int
	Value string
}

func ToKanaToken(input []rune, treeMap *tree.Map, convertEnding bool) (result []ConvertedKanaToken) {

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
		// parseKanaToken is called recusrively, so each time it returns a non-nil coverted token,
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
