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
}