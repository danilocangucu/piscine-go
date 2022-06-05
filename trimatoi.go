package piscine

func TrimAtoi(s string) int {
	nb := 0
	minus := false
	for _, v := range s {
		lastdigit := 0

		if v <= '9' && v >= '0' {
			for i := '1'; i <= v; i++ {
				lastdigit++
			}
			nb = nb*10 + lastdigit

		} else if nb == 0 && v == '-' {
			minus = true
			continue
		} else {
			continue
		}
	}
	if minus == true {
		nb = nb * -1
	}
	return nb
}
