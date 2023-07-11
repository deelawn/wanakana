package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
	"github.com/deelawn/wanakana/config"
)

func TestIsKatakana(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "empty string",
		},
		{
			name:     "game",
			input:    "ゲーム",
			expected: true,
		},
		{
			name:  "a hiragana",
			input: "あ",
		},
		{
			name:  "letter A",
			input: "A",
		},
		{
			name:  "mixed with hiragana",
			input: "あア",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.IsKatakana(tt.input)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestToKatakana(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		options  config.Options
		expected string
	}{
		{
			name: "empty string",
		},
		{
			name:     "romaji hiragana",
			input:    "toukyou, おおさか",
			expected: "トウキョウ、 オオサカ",
		},
		{
			name:     "pass romaji with hiragana",
			input:    "only かな",
			options:  config.Options{PassRomaji: true},
			expected: "only カナ",
		},
		{
			name:     "wi",
			input:    "wi",
			expected: "ウィ",
		},
		{
			name:     "obsolete wi",
			input:    "wi",
			options:  config.Options{UseObsoleteKana: true},
			expected: "ヰ",
		},
		{
			name:     "english punctuation",
			input:    string([]rune{0x2018, 0x2019}),
			expected: "「」",
		},
		{
			name:     "unsupported",
			input:    string([]rune{0x0600}),
			expected: "؀",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.ToKatakana(tt.input, tt.options)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
