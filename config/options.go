package config

// Romanization is an enum that determines the romanization to use when converting to/from romaji.
// Hepburn is the only supported romanization at this time.
type Romanization int

const (
	RomanizationHepburn Romanization = iota
)

// ToKanaMethod is an enum that determines the method to use when converting to kana.
type ToKanaMethod int

const (
	ToKanaMethodNone ToKanaMethod = iota
	ToKanaMethodHiragana
	ToKanaMethodKatakana
)

// Options is a struct that contains options for the various methods. Note that not all methods
// will use all options.
type Options struct {
	// UseObsoleteKana determines whether to use obsolete characters such as ゐ and ゑ.
	UseObsoleteKana bool
	// PassRomaji determines whether or not to pass romaji through untransformed when using mixed syllabaries.
	PassRomaji bool
	// IgnoreLongVowelMark determines whether or not to ignore long vowel marks when converting to hiragana.
	IgnoreLongVowelMark bool
	// UppercaseKatakana determines whether to use uppercase characters when converting from katakana to romaji.
	UppercaseKatakana bool
	// Romanization determines the romanization to use when converting to/from romaji. Hepburn is the default and only
	// supported romanization at this time. Providing any other value will result in an empty tree being generated
	// and no transliteration taking place.
	Romanization Romanization
	// IMEMode determines whether to use IME mode as well as the to kana method
	IMEMode ToKanaMethod
	// CustomKanaMapping is a custom mapping to use when converting to/from romaji.
	CustomKanaMapping CustomMapping
}
