package piscine

func IsLower(s string) bool {
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
