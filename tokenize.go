package wanakana

import (
	"fmt"
	"regexp"

	"github.com/deelawn/wanakana/internal/character"
)

// TokenType indicates the type of token produced by calling Tokenize.
type TokenType string

const (
	TokenTypeEN            TokenType = "en"
	TokenTypeJA            TokenType = "ja"
	TokenTypeENNumeral     TokenType = "englishNumeral"
	TokenTypeJANumeral     TokenType = "japaneseNumeral"
	TokenTypeENPunctuation TokenType = "englishPunctuation"
	TokenTypeJAPunctuation TokenType = "japanesePunctuation"
	TokenTypeKanji         TokenType = "kanji"
	TokenTypeHiragana      TokenType = "hiragana"
	TokenTypeKatakana      TokenType = "katakana"
	TokenTypeSpace         TokenType = "space"
	TokenTypeOther         TokenType = "other"

	enSpace rune = ' '
	jaSpace rune = '　'
)

var (
	jaNumRegex = regexp.MustCompile(`[０-９]`)
	enNumRegex = regexp.MustCompile(`[0-9]`)
)

func getTokenType(r rune, compact bool) TokenType {

	if compact {
		switch {
		case jaNumRegex.MatchString(string(r)):
			return TokenTypeOther

		case enNumRegex.MatchString(string(r)):
			return TokenTypeOther

		case r == enSpace:
			return TokenTypeEN

		case character.IsEnglishPunctuation(r):
			return TokenTypeOther

		case r == jaSpace:
			return TokenTypeJA

		case character.IsJapanesePunctuation(r):
			return TokenTypeOther

		case character.IsJapanese(r):
			return TokenTypeJA

		case character.IsRomaji(r):
			return TokenTypeEN

		default:
			return TokenTypeOther
		}
	}

	switch {
	case r == jaSpace:
		return TokenTypeSpace

	case r == enSpace:
		return TokenTypeSpace

	case jaNumRegex.MatchString(string(r)):
		return TokenTypeJANumeral

	case enNumRegex.MatchString(string(r)):
		return TokenTypeENNumeral

	case character.IsJapanesePunctuation(r):
		return TokenTypeJAPunctuation

	case character.IsEnglishPunctuation(r):
		return TokenTypeENPunctuation

	case character.IsKanji(r):
		return TokenTypeKanji

	case character.IsHiragana(r):
		return TokenTypeHiragana

	case character.IsKatakana(r):
		return TokenTypeKatakana

	case character.IsJapanese(r):
		return TokenTypeJA

	case character.IsRomaji(r):
		return TokenTypeEN
	}

	return TokenTypeOther
}

// Token represents a token produced by calling Tokenize.
type Token struct {
	Type  TokenType
	Value string
}

func newToken(tokenType TokenType, value string, detailed bool) Token {

	if detailed {
		return Token{Type: tokenType, Value: value}
	}

	return Token{Value: value}
}

// String returns the string representation of a Token as a JSON object with two string fields:
// `type` and `value`. This is the same way it is done in the original javaScript implementation.
func (t Token) String() string {

	if t.Type == "" {
		return t.Value
	}

	return fmt.Sprintf(`{ type: '%s', value: '%s' }`, t.Type, t.Value)
}

// Tokenize returns an array of Token objects representing the contents of the input string.
func Tokenize(input string, compact, detailed bool) []Token {

	if len(input) == 0 {
		return []Token{}
	}

	var (
		newValue string
		prevType TokenType
		result   []Token
	)

	for _, r := range []rune(input) {
		currType := getTokenType(r, compact)

		if currType == prevType || prevType == "" {
			newValue += string(r)
		} else {
			result = append(result, newToken(prevType, newValue, detailed))
			newValue = string(r)
		}

		prevType = currType
	}

	// Add the last accumulated token if there is one.
	if newValue != "" {
		result = append(result, newToken(prevType, newValue, detailed))
	}

	return result
}
