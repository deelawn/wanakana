package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
)

func TestIsHiragana(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			expected: true,
		},
		{
			name:     "game",
			input:    "げーむ",
			expected: true,
		},
		{
			name:     "letter A",
			input:    "A",
			expected: false,
		},
		{
			name:     "mixed with katakana",
			input:    "あア",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.IsHiragana(tt.input)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}
