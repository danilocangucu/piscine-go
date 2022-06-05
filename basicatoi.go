package piscine

func BasicAtoi(s string) int {
	nb := 0

	for _, v := range s {
		lastdigit := 0
		for i := '1'; i <= v; i++ {
			lastdigit++
		}
		nb = nb*10 + lastdigit
	}
	return nb
}
