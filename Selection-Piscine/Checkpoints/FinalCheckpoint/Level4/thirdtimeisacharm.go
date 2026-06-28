package main

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	charm := []rune(str)
	result := ""

	for i, thirdtime := range charm {
		if (i+1)%3 == 0 {
			result += string(thirdtime)
		}
	}
	return result + "\n"
}
