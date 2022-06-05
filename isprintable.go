package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)

	for _, v := range runes {
		if !(v >= 32 && v <= 127) {
			return false
		}
	}
	return true
}
