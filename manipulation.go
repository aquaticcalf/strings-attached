package stringsattached

import (
	"strings"
	"unicode"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j -1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
} // returns the reverse of a string ( abcdefgh -> hgfedcba )

func Is_palindrome(s string) bool {
	var cleaned strings.Builder
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			cleaned.WriteRune(r)
		}
	}
	clean_string := cleaned.String()
	return clean_string == Reverse(clean_string)
} // gives out a boolean, true if palindrome

func Truncate(s string, max_length int) string {
	if len(s) <= max_length {
		return s
	}
	return s[:max_length - 3] + "..."
} // cuts a string to a specified length and adds ellipsis ( if necessary )

func Word_count(s string) int {
	words := strings.Fields(strings.TrimSpace(s))
	return len(words)
} // counts the number of words in a string