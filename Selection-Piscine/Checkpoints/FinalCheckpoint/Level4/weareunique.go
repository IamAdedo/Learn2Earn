package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}
	inString1 := make(map[rune]bool)
	inString2 := make(map[rune]bool)
	for _, r := range str1 {
		inString1[r] = true
	}
	for _, r := range str2 {
		inString2[r] = true
	}
	count := 0
	for r := range inString1 {
		if !inString2[r] {
			count++
		}
	}
	for r := range inString2 {
		if !inString1[r] {
			count++
		}
	}
	return count
}
