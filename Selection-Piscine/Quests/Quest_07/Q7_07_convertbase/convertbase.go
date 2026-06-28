package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num := 0
	baseFromLen := len(baseFrom)
	for _, c := range nbr {
		digit := findIndex(baseFrom, c)
		if digit == -1 {
			return "" // invalid digit for baseFrom
		}
		num = num*baseFromLen + digit
	}
	if num == 0 {
		return string(baseTo[0])
	}
	baseToLen := len(baseTo)
	result := []rune{}
	for num > 0 {
		remainder := num % baseToLen
		result = append([]rune{rune(baseTo[remainder])}, result...)
		num = num / baseToLen
	}
	return string(result)
}

// Helper function to find the index of a rune in a string
func findIndex(s string, r rune) int {
	for i, c := range s {
		if c == r {
			return i
		}
	}
	return -1
}
