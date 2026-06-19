package piscine

func CountAlpha(s string) int {
	count := 0
	for _, alpha := range s {
		if (alpha >= 'a' && alpha <= 'z') || (alpha >= 'A' && alpha <= 'Z') {
			count++
		}
	}
	return count
}
