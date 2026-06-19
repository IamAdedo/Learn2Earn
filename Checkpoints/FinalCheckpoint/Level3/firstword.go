package piscine

func FirstWord(s string) string {
	word := ""
	for _, fword := range s {
		if fword == ' ' && word == "" {
			continue
		}
		if fword == ' ' {
			break
		}
		word += string(fword)
	}
	return word + "\n"
}
