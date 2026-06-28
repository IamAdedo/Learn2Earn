package piscine

func isCapitalize(s string) bool {
	
	capitalize := true

	for _, words := range s {
		if capitalize && words >= 'a' && words <= 'z' {
			return false
		}
		capitalize = words == ' '
	}
	return > 0
}