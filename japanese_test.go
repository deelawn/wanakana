package wanakana_test

import (
	"regexp"
	"testing"

	"github.com/deelawn/wanakana"
)

func TestIsJapanese(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		regex    *regexp.Regexp
		expected bool
	}{
		{
			name:     "crybaby",
			input:    "泣き虫",
			expected: true,
		},
		{
			name:     "aa mixed hiragana katakana",
			input:    "あア",
			expected: true,
		},
		{
			name:     "february with zenkaku number",
			input:    "２月",
			expected: true,
		},
		{
			name:     "crybaby with zenkaku punctuation",
			input:    "泣き虫。！〜＄",
			expected: true,
		},
		{
			name:  "crybaby with latin punctuation",
			input: "泣き虫.!~$",
		},
		{
			name:  "crybaby latin mix",
			input: "A泣き虫",
		},
		{
			name:     "weird punctuation with regex",
			input:    "≪偽括弧≫",
			regex:    regexp.MustCompile(`[≪≫]`),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.IsJapanese(tt.input, tt.regex)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}
