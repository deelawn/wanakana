package wanakana

import "testing"

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
			result := IsKana(tt.input)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}
