package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
	"github.com/deelawn/wanakana/config"
)

func TestIsKana(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "empty string",
		},
		{
			name:     "a hiragana",
			input:    "あ",
			expected: true,
		},
		{
			name:     "a katakana",
			input:    "ア",
			expected: true,
		},
		{
			name:     "mixed a dash a",
			input:    "あーア",
			expected: true,
		},
		{
			name:  "latin a",
			input: "A",
		},
		{
			name:  "mixed with latin a",
			input: "あAア",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.IsKana(tt.input)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

func TestToKana(t *testing.T) {

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
			name:     "hiragana katakana",
			input:    "onaji BUTTSUUJI",
			expected: "おなじ ブッツウジ",
		},
		{
			name:     "katakana hiragana",
			input:    "ONAJI buttsuuji",
			expected: "オナジ ぶっつうじ",
		},
		{
			name:     "mixed with kanji",
			input:    "座禅‘zazen’スタイル",
			expected: "座禅「ざぜん」スタイル",
		},
		{
			name:     "with dash",
			input:    "batsuge-mu",
			expected: "ばつげーむ",
		},
		{
			name:     "punctuation",
			input:    "!?.:/,~-‘’“”[](){}",
			expected: "！？。：・、〜ー「」『』［］（）｛｝",
		},
		{
			name:     "obsolete we",
			input:    "we",
			options:  config.Options{UseObsoleteKana: true},
			expected: "ゑ",
		},
		{
			name:  "custom mapping",
			input: "wanakana",
			options: config.Options{
				CustomKanaMapping: config.NewCustomMappingKeyValue(
					map[string]string{
						"na": "に",
						"ka": "bana",
					},
				),
			},
			expected: "わにbanaに",
		},
		{
			name:     "ch tsu",
			input:    "hacchi",
			expected: "はっち",
		},
		{
			name:     "hahha",
			input:    "hahha",
			expected: "はっは",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.ToKana(tt.input, tt.options, nil)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
