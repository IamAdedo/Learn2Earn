package piscine

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	count := 1

	for i := 1; i <= len(s); i++ {
		if i < len(s) && s[i] == s[i-1] {
			count++
		} else {
			result += string(rune(count + '0'))
			result += string(s[i-1])
			count = 1
		}
	}
	return result
}
