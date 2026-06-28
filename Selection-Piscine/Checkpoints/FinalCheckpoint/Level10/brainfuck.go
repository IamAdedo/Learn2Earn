package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}
	program := os.Args[1]
	memory := make([]byte, 2048)
	ptr := 0
	matching := make(map[int]int)
	stack := []int{}
	for i, cmd := range program {
		if cmd == '[' {
			stack = append(stack, i)
		} else if cmd == ']' {
			if len(stack) == 0 {
				return
			}
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			matching[open] = i
			matching[i] = open
		}
	}
	if len(stack) > 0 {
		return
	}
	pc := 0
	for pc < len(program) {
		cmd := program[pc]
		switch cmd {
		case '>':
			ptr++
			if ptr >= len(memory) {
				return
			}
		case '<':
			ptr--
			if ptr < 0 {
				return
			}
		case '+':
			memory[ptr]++
		case '-':
			memory[ptr]--
		case '.':
			z01.PrintRune(rune(memory[ptr]))
		case '[':
			if memory[ptr] == 0 {
				pc = matching[pc]
			}
		case ']':
			if memory[ptr] != 0 {
				pc = matching[pc]
			}
		}
		pc++
	}
}
