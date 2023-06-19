package wanakana_test

import (
	"testing"

	"github.com/deelawn/wanakana"
)

func TestIsKanji(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name: "empty string",
		},
		{
			name:     "sword",
			input:    "刀",
			expected: true,
		},
		{
			name:     "seppuku",
			input:    "切腹",
			expected: true,
		},
		{
			name:  "force",
			input: "勢い",
		},
		{
			name:  "mixed kana latin",
			input: "あAア",
		},
		{
			name:  "frog emoji",
			input: "🐸",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wanakana.IsKanji(tt.input)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}
