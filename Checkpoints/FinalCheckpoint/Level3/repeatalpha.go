package piscine

func RepeatAlpha(s string) string {
	result := ""
	for _, repeat := range s {
		n := 1
		if repeat >= 'a' && repeat <= 'z' {
			n = int(repeat - 'a' + 1)
		} else if repeat >= 'A' && repeat <= 'Z' {
			n = int(repeat - 'A' + 1)
		}
		for i := 0; i < n; i++ {
			result += string(repeat)
		}
	}
	return result
}
