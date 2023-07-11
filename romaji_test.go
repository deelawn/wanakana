package wanakana_test

import (
	"regexp"
	"testing"

	"github.com/deelawn/wanakana"
	"github.com/deelawn/wanakana/config"
)

func TestIsRomaji(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		regex    *regexp.Regexp
		expected bool
	}{
		{
			name: "empty string",
		},
		{
			name:     "romaji characters",
			input:    "Tōkyō and Ōsaka",
			expected: true,
		},
		{
			name:     "romaji with symbols",
			input:    "12a*b&c-d",
			expected: true,
		},
		{
			name:  "mixed with kana",
			input: "あアA",
		},
		{
			name:  "no romaji",
			input: "お願い",
		},
		{
			name:  "mixed with zenkaku punctuation",
			input: "a！b&cーd",
		},
		{
			name:     "mixed with zenkaku punctuation with regex",
			input:    "a！b&cーd",
			regex:    regexp.MustCompile(`[！ー]`),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wanakana.IsRomaji(tt.input, tt.regex)
			if actual != tt.expected {
				t.Errorf("wanakana.IsRomaji(%v, %v): expected %v, actual %v", tt.input, tt.regex, tt.expected, actual)
			}
		})
	}
}

func TestToRomaji(t *testing.T) {

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
			input:    "ひらがな　カタカナ",
			expected: "hiragana katakana",
		},
		{
			name:     "hiragana katakana with long dashes",
			input:    "げーむ　ゲーム",
			expected: "ge-mu geemu",
		},
		{
			name:     "upcase katakana",
			input:    "ひらがな　カタカナ",
			options:  config.Options{UppercaseKatakana: true},
			expected: "hiragana KATAKANA",
		},
		{
			name:  "custom mapping",
			input: "つじぎり",
			options: config.Options{
				CustomKanaMapping: config.NewCustomMappingKeyValue(
					map[string]string{
						"じ": "zi",
						"つ": "tu",
						"り": "li",
					},
				),
			},
			expected: "tuzigili",
		},
		{
			name:     "osaka",
			input:    "オーサカ おおさか オオサカ",
			expected: "oosaka oosaka oosaka",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wanakana.ToRomaji(tt.input, tt.options, nil)
			if actual != tt.expected {
				t.Errorf("wanted %s, got %s", tt.expected, actual)
			}
		})
	}
}
