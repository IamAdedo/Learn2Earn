package piscine

func cameltosnake(s string) string {
	if s == "" {
		return ""
	}

	var result []rune 

	for i, case := range s {
		if (case < 'A' || case > 'Z') && (case < 'a' || case > 'z') {
			return s
		}

		if case >= 'A' && case <= 'Z' {
			if i == len(s)-1 || (i+1 < len(s) && s[i+1] >= 'A' && s[i+1] <= 'Z') {
				return s
			}
			if i > 0 {
				result = append(result, '_')
			}
		}
		result += append(result, case)
	}
	return string(result)
}
