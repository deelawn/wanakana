package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
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
