package transform

import (
	"strings"
	"sync"
)

// TODO:
// - obsolete kana map
// - ime mode map

// rtk = Romaji To Kana
var (
	rtkConsonants = map[rune]rune{
		'k': 'き',
		's': 'し',
		't': 'ち',
		'n': 'に',
		'h': 'ひ',
		'm': 'み',
		'r': 'り',
		'g': 'ぎ',
		'z': 'じ',
		'd': 'ぢ',
		'b': 'び',
		'p': 'ぴ',
		'v': 'ゔ',
		'q': 'く',
		'f': 'ふ',
	}

	rtkSmallY = map[string]rune{
		"ya": 'ゃ',
		"yi": 'ぃ',
		"yu": 'ゅ',
		"ye": 'ぇ',
		"yo": 'ょ',
	}

	rtkSmallVowels = map[string]rune{
		"a": 'ぁ',
		"i": 'ぃ',
		"u": 'ぅ',
		"e": 'ぇ',
		"o": 'ぉ',
	}

	rtkSmallLetters = mergeMaps(
		rtkSmallY,
		rtkSmallVowels,
		map[string]rune{
			"tu": 'っ',
			"wa": 'ゎ',
			"ka": 'ヵ',
			"ke": 'ヶ',
		},
	)

	rtkAIUEOConstructions = map[string]rune{
		"wh": 'う',
		"kw": 'く',
		"qw": 'く',
		"q":  'く',
		"gw": 'ぐ',
		"sw": 'す',
		"ts": 'つ',
		"th": 'て',
		"tw": 'と',
		"dh": 'で',
		"dw": 'ど',
		"fw": 'ふ',
		"f":  'ふ',
	}

	// Typing one should be the same as having typed the other instead.

	rtkAliases = []aliasPair{
		{from: "sh", to: "sy"},  // sha -> sya
		{from: "ch", to: "ty"},  // cho -> tyo
		{from: "cy", to: "ty"},  // cyo -> tyo
		{from: "chy", to: "ty"}, // chyu -> tyu
		{from: "shy", to: "sy"}, // shya -> sya
		{from: "j", to: "zy"},   // ja -> zya
		{from: "jy", to: "zy"},  // jye -> zye

		// Exceptions to above rules.
		{from: "shi", to: "si"},
		{from: "chi", to: "ti"},
		{from: "tsu", to: "tu"},
		{from: "ji", to: "zi"},
		{from: "fu", to: "hu"},
	}

	// These don't follow any noticeable patters.
	specialCases = map[string]string{
		"yi":  "い",
		"wu":  "う",
		"ye":  "いぇ",
		"wi":  "うぃ",
		"we":  "うぇ",
		"kwa": "くぁ",
		"whu": "う",
		// because it's not thya for てゃ but tha
		// and tha is not てぁ, but てゃ
		"tha": "てゃ",
		"thu": "てゅ",
		"tho": "てょ",
		"dha": "でゃ",
		"dhu": "でゅ",
		"dho": "でょ",
	}
)

type aliasPair struct {
	from string
	to   string
}

var (
	rtkTree   *node
	rtkTreeMu sync.Mutex
)

func getRomajiToKanaTree() *node {

	rtkTreeMu.Lock()
	defer rtkTreeMu.Unlock()

	if rtkTree == nil {
		rtkTree = createRomajiToKanaTree()
	}

	return rtkTree
}

func createRomajiToKanaTree() *node {

	tree := new(node)

	addBasicKunreiToRTKTree(tree)
	addSpecialSymbolsToRTKTree(tree)

	// Add the consonants + small y kana.
	for consonant, yKana := range rtkConsonants {
		for roma, kana := range rtkSmallY {
			tree.putValue(append([]rune{consonant}, []rune(roma)...), string(yKana)+string(kana))
		}
	}

	// Add things like うぃ, くぃ, etc.
	for consonant, aiueoKana := range rtkAIUEOConstructions {
		for vowel, kana := range rtkSmallVowels {
			tree.putValue(append([]rune(consonant), []rune(vowel)...), string(aiueoKana)+string(kana))
		}
	}

	// Add different ways to write ん.
	tree.putValue([]rune("n"), "ん")
	tree.putValue([]rune("n'"), "ん")
	tree.putValue([]rune("xn"), "ん")

	// 'c' is nearly identical to 'k', so make a copy of it to start.
	tree.branches['c'] = tree.branches['k'].copy()

	// Apply the aliases by inserted a copied version of the target node into a
	// new branch path in the tree.
	for _, pair := range rtkAliases {
		tree.putNode([]rune(pair.from), tree.getNode([]rune(pair.to)).copy())
	}

	for kunreiRoma, kana := range rtkSmallLetters {
		tree.putValue(append([]rune{'x'}, []rune(kunreiRoma)...), string(kana))
		tree.putValue(append([]rune{'l'}, []rune(kunreiRoma)...), string(kana))

		if altRoma := getAlternative(kunreiRoma); altRoma != "" {
			tree.putValue(append([]rune{'x'}, []rune(altRoma)...), string(kana))
			tree.putValue(append([]rune{'l'}, []rune(altRoma)...), string(kana))
		}
	}

	// TODO: Make this static.
	for roma, kana := range specialCases {
		tree.putValue([]rune(roma), kana)
	}

	// Add the little tsu when typing two of the same consonant.
	for roma, kana := range rtkConsonants {
		tree.putValue([]rune{roma, roma}, "っ"+string(kana))
	}

	// And do the same for these letters as well.
	tree.putValue([]rune{'c', 'c'}, "っc")
	tree.putValue([]rune{'y', 'y'}, "っy")
	tree.putValue([]rune{'w', 'w'}, "っw")
	tree.putValue([]rune{'j', 'j'}, "っj")

	return tree
}

func getAlternative(s string) string {

	if len(s) == 0 {
		return ""
	}

	// Check for c -> k. Put this check first because it is pretty common.
	if strings.HasPrefix(s, "k") {
		return "c" + s[1:]
	}

	for _, pair := range rtkAliases {
		if strings.HasPrefix(s, pair.to) {
			return strings.Replace(s, pair.to, pair.from, 1)
		}
	}

	return ""
}

func addBasicKunreiToRTKTree(tree *node) {
	tree.putValue([]rune("a"), "あ")
	tree.putValue([]rune("i"), "い")
	tree.putValue([]rune("u"), "う")
	tree.putValue([]rune("e"), "え")
	tree.putValue([]rune("o"), "お")
	tree.putValue([]rune("ka"), "か")
	tree.putValue([]rune("ki"), "き")
	tree.putValue([]rune("ku"), "く")
	tree.putValue([]rune("ke"), "け")
	tree.putValue([]rune("ko"), "こ")
	tree.putValue([]rune("sa"), "さ")
	tree.putValue([]rune("si"), "し")
	tree.putValue([]rune("su"), "す")
	tree.putValue([]rune("se"), "せ")
	tree.putValue([]rune("so"), "そ")
	tree.putValue([]rune("ta"), "た")
	tree.putValue([]rune("ti"), "ち")
	tree.putValue([]rune("tu"), "つ")
	tree.putValue([]rune("te"), "て")
	tree.putValue([]rune("to"), "と")
	tree.putValue([]rune("na"), "な")
	tree.putValue([]rune("ni"), "に")
	tree.putValue([]rune("nu"), "ぬ")
	tree.putValue([]rune("ne"), "ね")
	tree.putValue([]rune("no"), "の")
	tree.putValue([]rune("ha"), "は")
	tree.putValue([]rune("hi"), "ひ")
	tree.putValue([]rune("hu"), "ふ")
	tree.putValue([]rune("he"), "へ")
	tree.putValue([]rune("ho"), "ほ")
	tree.putValue([]rune("ma"), "ま")
	tree.putValue([]rune("mi"), "み")
	tree.putValue([]rune("mu"), "む")
	tree.putValue([]rune("me"), "め")
	tree.putValue([]rune("mo"), "も")
	tree.putValue([]rune("ya"), "や")
	tree.putValue([]rune("yu"), "ゆ")
	tree.putValue([]rune("yo"), "よ")
	tree.putValue([]rune("ra"), "ら")
	tree.putValue([]rune("ri"), "り")
	tree.putValue([]rune("ru"), "る")
	tree.putValue([]rune("re"), "れ")
	tree.putValue([]rune("ro"), "ろ")
	tree.putValue([]rune("wa"), "わ")
	tree.putValue([]rune("wi"), "ゐ")
	tree.putValue([]rune("wu"), "う")
	tree.putValue([]rune("we"), "ゑ")
	tree.putValue([]rune("wo"), "を")
	tree.putValue([]rune("ga"), "が")
	tree.putValue([]rune("gi"), "ぎ")
	tree.putValue([]rune("gu"), "ぐ")
	tree.putValue([]rune("ge"), "げ")
	tree.putValue([]rune("go"), "ご")
	tree.putValue([]rune("za"), "ざ")
	tree.putValue([]rune("zi"), "じ")
	tree.putValue([]rune("zu"), "ず")
	tree.putValue([]rune("ze"), "ぜ")
	tree.putValue([]rune("zo"), "ぞ")
	tree.putValue([]rune("da"), "だ")
	tree.putValue([]rune("di"), "ぢ")
	tree.putValue([]rune("du"), "づ")
	tree.putValue([]rune("de"), "で")
	tree.putValue([]rune("do"), "ど")
	tree.putValue([]rune("ba"), "ば")
	tree.putValue([]rune("bi"), "び")
	tree.putValue([]rune("bu"), "ぶ")
	tree.putValue([]rune("be"), "べ")
	tree.putValue([]rune("bo"), "ぼ")
	tree.putValue([]rune("pa"), "ぱ")
	tree.putValue([]rune("pi"), "ぴ")
	tree.putValue([]rune("pu"), "ぷ")
	tree.putValue([]rune("pe"), "ぺ")
	tree.putValue([]rune("po"), "ぽ")
	tree.putValue([]rune("va"), "ゔぁ")
	tree.putValue([]rune("vi"), "ゔぃ")
	tree.putValue([]rune("vu"), "ゔ")
	tree.putValue([]rune("ve"), "ゔぇ")
	tree.putValue([]rune("vo"), "ゔぉ")
}

func addSpecialSymbolsToRTKTree(tree *node) {
	tree.putValue([]rune{'.'}, "。")
	tree.putValue([]rune{','}, "、")
	tree.putValue([]rune{':'}, "：")
	tree.putValue([]rune{'/'}, "・")
	tree.putValue([]rune{'!'}, "！")
	tree.putValue([]rune{'?'}, "？")
	tree.putValue([]rune{'~'}, "〜")
	tree.putValue([]rune{'-'}, "ー")
	tree.putValue([]rune{'‘'}, "「")
	tree.putValue([]rune{'’'}, "」")
	tree.putValue([]rune{'“'}, "『")
	tree.putValue([]rune{'”'}, "』")
	tree.putValue([]rune{'['}, "［")
	tree.putValue([]rune{']'}, "］")
	tree.putValue([]rune{'('}, "（")
	tree.putValue([]rune{')'}, "）")
	tree.putValue([]rune{'{'}, "｛")
	tree.putValue([]rune{'}'}, "｝")
}

func mergeMaps(maps ...map[string]rune) map[string]rune {

	var result = make(map[string]rune)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}

	return result
}

func imeModeTree(tree *node) *node {

	if tree == nil {
		return nil
	}

	treeCopy := tree.copy()
	treeCopy.putValue([]rune("nn"), "ん")
	treeCopy.putValue([]rune("n "), "ん")

	return treeCopy
}
