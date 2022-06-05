package piscine

func LastRune(s string) rune {
	var firstrune rune

	for i, v := range s {
		if i == len(s)-1 {
			return v
		}
	}

	return firstrune
}
