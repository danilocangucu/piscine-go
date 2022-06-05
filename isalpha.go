package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if v >= 'a' && v <= 'z' ||
			v >= 'A' && v <= 'Z' ||
			v >= '0' && v <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
