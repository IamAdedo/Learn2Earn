package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, charCount := range str {
		if charCount == c {
			count++
		}
	}
	return count
}
