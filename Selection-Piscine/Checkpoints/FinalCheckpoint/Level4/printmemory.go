package piscine

import "github.com/01-edu/z01"

func PrintMemory(a [10]byte) {

	h := "0123456789abcdef"

	for i, b := range a {
		z01.PrintRune(rune(h[b>>4]))
		z01.PrintRune(rune(h[b%15]))

		if i == 9 || i%4 == 3 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for _, b := range a {
		if b > 31 && b < 127 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}
