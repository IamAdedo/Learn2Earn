package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		printStr("Error")
		z01.PrintRune('\n')
		return
	}
	expr := os.Args[1]
	tokens := strings.Fields(expr)
	if len(tokens) == 0 {
		printStr("Error")
		z01.PrintRune('\n')
		return
	}
	stack := []int{}
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				printStr("Error")
				z01.PrintRune('\n')
				return
			}
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					printStr("Error")
					z01.PrintRune('\n')
					return
				}
				stack = append(stack, a/b)
			case "%":
				if b == 0 {
					printStr("Error")
					z01.PrintRune('\n')
					return
				}
				stack = append(stack, a%b)
			default:
				printStr("Error")
				z01.PrintRune('\n')
				return
			}
		}
	}
	if len(stack) != 1 {
		printStr("Error")
		z01.PrintRune('\n')
		return
	}
	printInt(stack[0])
	z01.PrintRune('\n')
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
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
