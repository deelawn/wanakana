package codepoints

const (
	LatinLowercaseStart rune = 0x61
	LatinLowercaseEnd   rune = 0x7a
	LatinUppercaseStart rune = 0x41
	LatinUppercaseEnd   rune = 0x5a

	LowercaseZenkakuStart rune = 0xff41
	LowercaseZenkakuEnd   rune = 0xff5a
	UppercaseZenkakuStart rune = 0xff21
	UppercaseZenkakuEnd   rune = 0xff3a

	HiraganaStart rune = 0x3041
	HiraganaEnd   rune = 0x3096
	KatakanaStart rune = 0x30a1
	KatakanaEnd   rune = 0x30fc

	KanjiStart rune = 0x4e00
	KanjiEnd   rune = 0x9faf

	ProlongedSoundMark rune = 0x30fc
	KanaSlashDot       rune = 0x30fb
)

type Range [2]rune

var (
	ZenkakuNumbersRange            = Range{0xff10, 0xff19}
	ZenkakuUppercaseRange          = Range{UppercaseZenkakuStart, UppercaseZenkakuEnd}
	ZenkakuLowercaseRange          = Range{LowercaseZenkakuStart, LowercaseZenkakuEnd}
	ZenkakuePunctuationRange1      = Range{0xff01, 0xff0f}
	ZenkakuePunctuationRange2      = Range{0xff1a, 0xff1f}
	ZenkakuePunctuationRange3      = Range{0xff3b, 0xff3f}
	ZenkakuePunctuationRange4      = Range{0xff5b, 0xff60}
	ZenkakuSymbolsAndCurrencyRange = Range{0xffe0, 0xffee}

	HiraganaCharsRange            = Range{0x3040, 0x309f}
	KatakanaCharsRange            = Range{0x30a0, 0x30ff}
	HankakuKatakanaRange          = Range{0xff66, 0xff9f}
	KatakanaPunctuationRange      = Range{0x30fb, 0x30fc}
	KanaPunctuationRange          = Range{0xff61, 0xff65}
	CJKSymbolsAndPunctuationRange = Range{0x3000, 0x303f}
	CommonCJKRange                = Range{0x4e00, 0x9fff}
	RareCJKRange                  = Range{0x3400, 0x4dbf}

	KanaRanges = []Range{
		HiraganaCharsRange,
		KatakanaCharsRange,
		KanaPunctuationRange,
		HankakuKatakanaRange,
	}

	JAPunctuationRanges = []Range{
		CJKSymbolsAndPunctuationRange,
		KanaPunctuationRange,
		KatakanaPunctuationRange,
		ZenkakuePunctuationRange1,
		ZenkakuePunctuationRange2,
		ZenkakuePunctuationRange3,
		ZenkakuePunctuationRange4,
		ZenkakuSymbolsAndCurrencyRange,
	}

	ModernEnglishRange = Range{0x0000, 0x007f}

	HepburnMacronRanges = []Range{
		{0x0100, 0x0101},
		{0x0112, 0x0113},
		{0x012a, 0x012b},
		{0x014c, 0x014d},
		{0x016a, 0x016b},
	}

	SmartQuoteRanges = []Range{
		{0x2018, 0x2019},
		{0x201c, 0x201d},
	}

	// The slices below are initialized separately in init() because they make use of other
	// slices that are initialized above.
	JapaneseRanges      []Range
	RomajiRanges        []Range
	ENPunctuationRanges []Range
)

func init() {

	JapaneseRanges = append(JapaneseRanges, KanaRanges...)
	JapaneseRanges = append(JapaneseRanges, JAPunctuationRanges...)
	JapaneseRanges = append(
		JapaneseRanges,
		ZenkakuUppercaseRange,
		ZenkakuLowercaseRange,
		ZenkakuNumbersRange,
		CommonCJKRange,
		RareCJKRange,
	)

	RomajiRanges = append([]Range{ModernEnglishRange}, HepburnMacronRanges...)

	ENPunctuationRanges = []Range{
		{0x20, 0x2f},
		{0x3a, 0x3f},
		{0x5b, 0x60},
		{0x7b, 0x7e},
	}
	ENPunctuationRanges = append(ENPunctuationRanges, SmartQuoteRanges...)
}
