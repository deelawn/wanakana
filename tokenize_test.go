package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
)

func TestTokenize(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		compact  bool
		detailed bool
		expected []string
	}{
		{
			name: "empty string",
		},
		{
			name:     "fufufufu",
			input:    "ふふフフ",
			expected: []string{"ふふ", "フフ"},
		},
		{
			name:     "kanji",
			input:    "感じ",
			expected: []string{"感", "じ"},
		},
		{
			name:     "i'm truly miserable",
			input:    "truly 私は悲しい",
			expected: []string{"truly", " ", "私", "は", "悲", "しい"},
		},
		{
			name:     "i'm truly miserable compact",
			input:    "truly 私は悲しい",
			compact:  true,
			expected: []string{"truly ", "私は悲しい"},
		},
		{
			name:     "long mixed",
			input:    "5romaji here...!?漢字ひらがなカタ　カナ４「ＳＨＩＯ」。！",
			expected: []string{"5", "romaji", " ", "here", "...!?", "漢字", "ひらがな", "カタ", "　", "カナ", "４", "「", "ＳＨＩＯ", "」。！"},
		},
		{
			name:     "long mixed compact",
			input:    "5romaji here...!?漢字ひらがなカタ　カナ４「ＳＨＩＯ」。！",
			compact:  true,
			expected: []string{"5", "romaji here", "...!?", "漢字ひらがなカタ　カナ", "４「", "ＳＨＩＯ", "」。！"},
		},
		{
			name:     "long mixed detailed",
			input:    "5romaji here...!?漢字ひらがなカタ　カナ４「ＳＨＩＯ」。！ لنذهب",
			detailed: true,
			expected: []string{
				"{ type: 'englishNumeral', value: '5' }",
				"{ type: 'en', value: 'romaji' }",
				"{ type: 'space', value: ' ' }",
				"{ type: 'en', value: 'here' }",
				"{ type: 'englishPunctuation', value: '...!?' }",
				"{ type: 'kanji', value: '漢字' }",
				"{ type: 'hiragana', value: 'ひらがな' }",
				"{ type: 'katakana', value: 'カタ' }",
				"{ type: 'space', value: '　' }",
				"{ type: 'katakana', value: 'カナ' }",
				"{ type: 'japaneseNumeral', value: '４' }",
				"{ type: 'japanesePunctuation', value: '「' }",
				"{ type: 'ja', value: 'ＳＨＩＯ' }",
				"{ type: 'japanesePunctuation', value: '」。！' }",
				"{ type: 'space', value: ' ' }",
				"{ type: 'other', value: 'لنذهب' }",
			},
		},
		{
			name:     "long mixed detailed compact",
			input:    "5romaji here...!?漢字ひらがなカタ　カナ４「ＳＨＩＯ」。！ لنذهب",
			detailed: true,
			compact:  true,
			expected: []string{
				"{ type: 'other', value: '5' }",
				"{ type: 'en', value: 'romaji here' }",
				"{ type: 'other', value: '...!?' }",
				"{ type: 'ja', value: '漢字ひらがなカタ　カナ' }",
				"{ type: 'other', value: '４「' }",
				"{ type: 'ja', value: 'ＳＨＩＯ' }",
				"{ type: 'other', value: '」。！' }",
				"{ type: 'en', value: ' ' }",
				"{ type: 'other', value: 'لنذهب' }",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.Tokenize(tt.input, tt.compact, tt.detailed)
			if len(result) != len(tt.expected) {
				t.Fatalf("expected results of length %d, got %d: %+v", len(tt.expected), len(result), result)
			}

			for i, token := range result {
				if token.String() != tt.expected[i] {
					t.Errorf("at index %d expected %s, got %s", i, tt.expected[i], token.String())
				}
			}
		})
	}
}
