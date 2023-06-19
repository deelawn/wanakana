package wanakana_test

import (
	"regexp"
	"testing"

	"github.com/deelawn/wanakana"
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
