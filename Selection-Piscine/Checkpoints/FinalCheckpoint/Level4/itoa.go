package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	if n < 0 {
		return "-" + Itoa(-n)
	}
	var result string
	for ; n > 0; n /= 10 {
		result = string('0'+n%10) + result
	}
	return result
}
