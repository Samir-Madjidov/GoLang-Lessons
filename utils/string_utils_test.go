// string_utils_test.go
package utils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"single character", "a", "a"},
		{"ascii string", "hello", "olleh"},
		{"unicode string", "привет", "тевирп"},
		{"with spaces", "hello world", "dlrow olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; expected %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"no vowels", "bcdfg", 0},
		{"english vowels", "hello world", 3},
		{"russian vowels", "привет мир", 4},
		{"mixed case", "Hello World", 3},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountVowels(tt.input)
			if result != tt.expected {
				t.Errorf("CountVowels(%q) = %d; expected %d",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"palindrome ascii", "racecar", true},
		{"palindrome russian", "казак", true},
		{"not palindrome", "hello", false},
		{"with spaces", "а роза упала на лапу азора", true},
		{"case insensitive", "Racecar", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindrome(%q) = %t; expected %t",
					tt.input, result, tt.expected)
			}
		})
	}
}
