package piscine

func HashCode(dec string) string {
	if len(dec) == 0 {
		return ""
	}

	size := len(dec)

	var result string

	for _, hCode := range dec {
		hashed := (int(hCode) + size) % 127
		if hashed < 33 {
			hashed += 33
		}
		result += string(rune(hashed))
	}
	return result
}
