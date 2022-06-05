package piscine

func BasicAtoi2(s string) int {
	nb := 0

	for _, v := range s {
		lastdigit := 0
		if v <= '9' && v >= '0' {
			for i := '1'; i <= v; i++ {
				lastdigit++
			}
			nb = nb*10 + lastdigit
		} else {
			return 0
		}
	}
	return nb
}
