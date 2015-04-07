package sigilize

import (
	"regexp"
	"strings"
)

func removeVowels(text []string) string {
	r := regexp.MustCompile("[aeiou]")
	return r.ReplaceAllString(strings.Join(text, ""), "")
}

func removeRepeatedChars(text string) string {
	chars := map[string]bool{}

	for _, c := range text {
		char := string(c)
		if _, found := chars[char]; found {
			r := regexp.MustCompile(char)
			text = r.ReplaceAllString(text, "")
		}
		chars[char] = true
	}
	return text
}

func Sigilize(text []string) string {
	return removeRepeatedChars(removeVowels(text))
}