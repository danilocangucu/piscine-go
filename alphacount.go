package piscine

func AlphaCount(s string) int {
	var count int

	for _, v := range s {
		if v >= 'a' && v <= 'z' ||
			v >= 'A' && v <= 'Z' {
			count++
		}
	}
	return count
}
