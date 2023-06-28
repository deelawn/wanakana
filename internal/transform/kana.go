package transform

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
