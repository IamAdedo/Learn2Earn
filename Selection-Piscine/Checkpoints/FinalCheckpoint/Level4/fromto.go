package main

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "invalid\n"
	}
	var result string
	for from != to {
		result += fmt.Sprintf("%02d, ", from)

		if from < to {
			from++
		} else {
			from--
		}
	}
	result += fmt.Sprintf("%02d\n", from)
	return result
}
