package wanakana

type Romanization string

const (
	RomanizationHepburn Romanization = "hepburn"
)

type ToKanaMethod int

const (
	ToKanaMethodNone ToKanaMethod = iota
	ToKanaMethodHiragana
	ToKanaMethodKatakana
)

type Options struct {
	// UseObsoleteKana determines whether to use obsolete characters such as ゐ and ゑ.
	UseObsoleteKana bool
	// PassRomaji determines whether or not to pass romaji through untransformed when using mixed syllabaries.
	PassRomaji bool
	// IgnoreLongVowelMark determines whether or not to ignore long vowel marks when converting to hiragana.
	IgnoreLongVowelMark bool
	// UppercaseKatakana determines whether to use uppercase characters when converting from katakana to romaji.
	UppercaseKatakana bool
	// Romanization determines the romanization to use when converting to/from romaji.
	Romanization Romanization
	// IMEMode determines whether to use IME mode as well as the to kana method
	IMEMode ToKanaMethod
	// CustomKanaMapping is a custom mapping to use when converting to/from romaji.
	CustomKanaMapping CustomMapping
}
