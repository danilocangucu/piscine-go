package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if (runes[i] >= 'A' && runes[i] <= 'L') || (runes[i] >= 'a' && runes[i] <= 'l') {
			runes[i] = runes[i] + 14
		} else if (runes[i] >= 'M' && runes[i] <= 'Z') || (runes[i] >= 'm' && runes[i] <= 'z') {
			runes[i] = runes[i] - 12
		}
	}
	return string(runes)
}
