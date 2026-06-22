package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	for _, arg := range args {
		if isBalanced(arg) {
			printStr("OK")
		} else {
			printStr("Error")
		}
		z01.PrintRune('\n')
	}
}

func isBalanced(s string) bool {
	stack := []rune{}
	closingToOpening := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if closingToOpening[r] != top {
				return false
			}
		}
	}
	return len(stack) == 0
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
