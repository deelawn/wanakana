package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
)

func TestStripOkurigana(t *testing.T) {

	var tests = []struct {
		name       string
		input      string
		leading    bool
		matchKanji string
		expected   string
	}{
		{
			name: "empty string",
		},
		{
			name:     "trailing 1",
			input:    "踏み込む",
			expected: "踏み込",
		},
		{
			name:     "trailing 2",
			input:    "お祝い",
			expected: "お祝",
		},
		{
			name:     "leading",
			input:    "お腹",
			leading:  true,
			expected: "腹",
		},
		{
			name:       "match kanji trailing",
			input:      "ふみこむ",
			matchKanji: "踏み込む",
			expected:   "ふみこ",
		},
		{
			name:       "match kanji leading",
			input:      "おみまい",
			leading:    true,
			matchKanji: "お祝い",
			expected:   "みまい",
		},
		{
			// Invalid match kanji results in this spitting out the input.
			name:       "invalid match kanji",
			input:      "おみまい",
			matchKanji: "おみまい",
			expected:   "おみまい",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.StripOkurigana(tt.input, tt.leading, tt.matchKanji)
			if result != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, result)
			}
		})
	}
}
