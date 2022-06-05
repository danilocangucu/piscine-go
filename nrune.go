package piscine

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	} else if n > 0 && n <= len(s) {
		for i, v := range s {
			if i == n-1 {
				return v
			}
		}
	}
	return 0
}
