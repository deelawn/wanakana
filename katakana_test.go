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

func TestToKatakana(t *testing.T) {

	var tests = []struct {
		name     string
		input    string
		options  wanakana.Options
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
			options:  wanakana.Options{PassRomaji: true},
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
			options:  wanakana.Options{UseObsoleteKana: true},
			expected: "ヰ",
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
