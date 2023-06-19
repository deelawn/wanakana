package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
)

func TestIsMixed(t *testing.T) {

	var tests = []struct {
		name      string
		input     string
		passKanji bool
		expected  bool
	}{
		{
			name: "empty string",
		},
		{
			name:     "mixed scripts no kanji",
			input:    "Abあア",
			expected: true,
		},
		{
			name:      "mixed scripts with kanji",
			input:     "お腹A",
			passKanji: true,
			expected:  true,
		},
		{
			name:  "mixed scripts with kanji no pass",
			input: "お腹A",
		},
		{
			name:  "latin only",
			input: "ab",
		},
		{
			name:  "kana only",
			input: "あア",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wanakana.IsMixed(tt.input, tt.passKanji)
			if actual != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
