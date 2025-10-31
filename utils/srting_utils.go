package utils

import "strings"

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CountVowels(s string) int {
	vowels := "aeiouаеёиоуыэюяAEIOUАЕЁИОУЫЭЮЯ"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func IsPalindrome(s string) bool {
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == Reverse(cleaned)
}
