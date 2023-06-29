package transform

import (
	"strings"
	"sync"

	"github.com/deelawn/wanakana/internal/tree"
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
	rtkTreeMap   *tree.Map
	rtkTreeMapMu sync.Mutex
)

func GetRomajiToKanaTree() *tree.Map {

	rtkTreeMapMu.Lock()
	defer rtkTreeMapMu.Unlock()

	if rtkTreeMap == nil {
		rtkTreeMap = createRomajiToKanaTree()
	}

	return rtkTreeMap
}

func createRomajiToKanaTree() *tree.Map {

	treeMap := new(tree.Map)

	addBasicKunreiToRTKTree(treeMap)
	addSpecialSymbolsToRTKTree(treeMap)

	// Add the consonants + small y kana.
	for consonant, yKana := range rtkConsonants {
		for roma, kana := range rtkSmallY {
			treeMap.PutValue(append([]rune{consonant}, []rune(roma)...), string(yKana)+string(kana))
		}
	}

	// Add things like うぃ, くぃ, etc.
	for consonant, aiueoKana := range rtkAIUEOConstructions {
		for vowel, kana := range rtkSmallVowels {
			treeMap.PutValue(append([]rune(consonant), []rune(vowel)...), string(aiueoKana)+string(kana))
		}
	}

	// Add different ways to write ん.
	treeMap.PutValue([]rune("n"), "ん")
	treeMap.PutValue([]rune("n'"), "ん")
	treeMap.PutValue([]rune("xn"), "ん")

	// 'c' is nearly identical to 'k', so make a copy of it to start.
	treeMap.PutMap([]rune{'c'}, treeMap.GetMap([]rune{'k'}).Copy())

	// Apply the aliases by inserted a copied version of the target node into a
	// new branch path in the tree.
	for _, pair := range rtkAliases {
		treeMap.PutMap([]rune(pair.from), treeMap.GetMap([]rune(pair.to)).Copy())
	}

	for kunreiRoma, kana := range rtkSmallLetters {
		treeMap.PutValue(append([]rune{'x'}, []rune(kunreiRoma)...), string(kana))
		treeMap.PutValue(append([]rune{'l'}, []rune(kunreiRoma)...), string(kana))

		if altRoma := getAlternative(kunreiRoma); altRoma != "" {
			treeMap.PutValue(append([]rune{'x'}, []rune(altRoma)...), string(kana))
			treeMap.PutValue(append([]rune{'l'}, []rune(altRoma)...), string(kana))
		}
	}

	// TODO: Make this static.
	for roma, kana := range specialCases {
		treeMap.PutValue([]rune(roma), kana)
	}

	// Add the little tsu when typing two of the same consonant.
	for roma := range rtkConsonants {
		treeMap.PutValue([]rune{roma, roma}, "っ")
	}

	// And do the same for these letters as well.
	treeMap.PutValue([]rune{'c', 'c'}, "っ")
	treeMap.PutValue([]rune{'y', 'y'}, "っ")
	treeMap.PutValue([]rune{'w', 'w'}, "っ")
	treeMap.PutValue([]rune{'j', 'j'}, "っ")
	treeMap.PutValue([]rune{'t', 't'}, "っ")

	return treeMap
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

func addBasicKunreiToRTKTree(treeMap *tree.Map) {
	treeMap.PutValue([]rune("a"), "あ")
	treeMap.PutValue([]rune("i"), "い")
	treeMap.PutValue([]rune("u"), "う")
	treeMap.PutValue([]rune("e"), "え")
	treeMap.PutValue([]rune("o"), "お")
	treeMap.PutValue([]rune("ka"), "か")
	treeMap.PutValue([]rune("ki"), "き")
	treeMap.PutValue([]rune("ku"), "く")
	treeMap.PutValue([]rune("ke"), "け")
	treeMap.PutValue([]rune("ko"), "こ")
	treeMap.PutValue([]rune("sa"), "さ")
	treeMap.PutValue([]rune("si"), "し")
	treeMap.PutValue([]rune("su"), "す")
	treeMap.PutValue([]rune("se"), "せ")
	treeMap.PutValue([]rune("so"), "そ")
	treeMap.PutValue([]rune("ta"), "た")
	treeMap.PutValue([]rune("ti"), "ち")
	treeMap.PutValue([]rune("tu"), "つ")
	treeMap.PutValue([]rune("te"), "て")
	treeMap.PutValue([]rune("to"), "と")
	treeMap.PutValue([]rune("na"), "な")
	treeMap.PutValue([]rune("ni"), "に")
	treeMap.PutValue([]rune("nu"), "ぬ")
	treeMap.PutValue([]rune("ne"), "ね")
	treeMap.PutValue([]rune("no"), "の")
	treeMap.PutValue([]rune("ha"), "は")
	treeMap.PutValue([]rune("hi"), "ひ")
	treeMap.PutValue([]rune("hu"), "ふ")
	treeMap.PutValue([]rune("he"), "へ")
	treeMap.PutValue([]rune("ho"), "ほ")
	treeMap.PutValue([]rune("ma"), "ま")
	treeMap.PutValue([]rune("mi"), "み")
	treeMap.PutValue([]rune("mu"), "む")
	treeMap.PutValue([]rune("me"), "め")
	treeMap.PutValue([]rune("mo"), "も")
	treeMap.PutValue([]rune("ya"), "や")
	treeMap.PutValue([]rune("yu"), "ゆ")
	treeMap.PutValue([]rune("yo"), "よ")
	treeMap.PutValue([]rune("ra"), "ら")
	treeMap.PutValue([]rune("ri"), "り")
	treeMap.PutValue([]rune("ru"), "る")
	treeMap.PutValue([]rune("re"), "れ")
	treeMap.PutValue([]rune("ro"), "ろ")
	treeMap.PutValue([]rune("wa"), "わ")
	treeMap.PutValue([]rune("wi"), "ゐ")
	treeMap.PutValue([]rune("wu"), "う")
	treeMap.PutValue([]rune("we"), "ゑ")
	treeMap.PutValue([]rune("wo"), "を")
	treeMap.PutValue([]rune("ga"), "が")
	treeMap.PutValue([]rune("gi"), "ぎ")
	treeMap.PutValue([]rune("gu"), "ぐ")
	treeMap.PutValue([]rune("ge"), "げ")
	treeMap.PutValue([]rune("go"), "ご")
	treeMap.PutValue([]rune("za"), "ざ")
	treeMap.PutValue([]rune("zi"), "じ")
	treeMap.PutValue([]rune("zu"), "ず")
	treeMap.PutValue([]rune("ze"), "ぜ")
	treeMap.PutValue([]rune("zo"), "ぞ")
	treeMap.PutValue([]rune("da"), "だ")
	treeMap.PutValue([]rune("di"), "ぢ")
	treeMap.PutValue([]rune("du"), "づ")
	treeMap.PutValue([]rune("de"), "で")
	treeMap.PutValue([]rune("do"), "ど")
	treeMap.PutValue([]rune("ba"), "ば")
	treeMap.PutValue([]rune("bi"), "び")
	treeMap.PutValue([]rune("bu"), "ぶ")
	treeMap.PutValue([]rune("be"), "べ")
	treeMap.PutValue([]rune("bo"), "ぼ")
	treeMap.PutValue([]rune("pa"), "ぱ")
	treeMap.PutValue([]rune("pi"), "ぴ")
	treeMap.PutValue([]rune("pu"), "ぷ")
	treeMap.PutValue([]rune("pe"), "ぺ")
	treeMap.PutValue([]rune("po"), "ぽ")
	treeMap.PutValue([]rune("va"), "ゔぁ")
	treeMap.PutValue([]rune("vi"), "ゔぃ")
	treeMap.PutValue([]rune("vu"), "ゔ")
	treeMap.PutValue([]rune("ve"), "ゔぇ")
	treeMap.PutValue([]rune("vo"), "ゔぉ")
}

func addSpecialSymbolsToRTKTree(treeMap *tree.Map) {
	treeMap.PutValue([]rune{'.'}, "。")
	treeMap.PutValue([]rune{','}, "、")
	treeMap.PutValue([]rune{':'}, "：")
	treeMap.PutValue([]rune{'/'}, "・")
	treeMap.PutValue([]rune{'!'}, "！")
	treeMap.PutValue([]rune{'?'}, "？")
	treeMap.PutValue([]rune{'~'}, "〜")
	treeMap.PutValue([]rune{'-'}, "ー")
	treeMap.PutValue([]rune{'‘'}, "「")
	treeMap.PutValue([]rune{'’'}, "」")
	treeMap.PutValue([]rune{'“'}, "『")
	treeMap.PutValue([]rune{'”'}, "』")
	treeMap.PutValue([]rune{'['}, "［")
	treeMap.PutValue([]rune{']'}, "］")
	treeMap.PutValue([]rune{'('}, "（")
	treeMap.PutValue([]rune{')'}, "）")
	treeMap.PutValue([]rune{'{'}, "｛")
	treeMap.PutValue([]rune{'}'}, "｝")
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

func ToIMEModeTree(treeMap *tree.Map) *tree.Map {

	if treeMap == nil {
		return nil
	}

	treeMapCopy := treeMap.Copy()
	treeMapCopy.PutValue([]rune("nn"), "ん")
	treeMapCopy.PutValue([]rune("n "), "ん")

	return treeMapCopy
}

func ToTreeWithObsoleteKana(treeMap *tree.Map) *tree.Map {

	treeMapCopy := treeMap.Copy()
	treeMapCopy.PutValue([]rune("wi"), "ゐ")
	treeMapCopy.PutValue([]rune("we"), "ゑ")

	return treeMapCopy
}
