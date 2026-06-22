package piscine

func CheckNumber(arg string) bool {
	for _, num := range arg {
		if num >= '0' && num <= '9' {
			return true
		}
	}
	return false
}
