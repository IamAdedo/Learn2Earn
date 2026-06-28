package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	pattern := os.Args[1]
	text := os.Args[2]
	if text == "" {
		return
	}
	patterns := []string{}
	current := ""
	for _, r := range pattern {
		if r == '|' {
			if current == "" {
				return
			}
			patterns = append(patterns, current)
			current = ""
		} else if r == '(' || r == ')' {
			continue
		} else {
			current += string(r)
		}
	}
	if current == "" {
		return
	}
	patterns = append(patterns, current)
	words := []string{}
	word := ""
	for _, r := range text {
		if r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == ':' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		words = append(words, word)
	}
	matches := []string{}
	for _, word := range words {
		for _, pat := range patterns {
			if containsAll(word, pat) {
				matches = append(matches, word)
				break
			}
		}
	}
	for i, match := range matches {
		printInt(i + 1)
		printStr(": ")
		printStr(match)
		z01.PrintRune('\n')
	}
}

func containsAll(word, pat string) bool {
	seen := make(map[rune]bool)
	for _, r := range word {
		seen[r] = true
	}
	for _, r := range pat {
		if !seen[r] {
			return false
		}
	}
	return true
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := []rune{}
	for n > 0 {
		digits = append(digits, rune('0'+n%10))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
