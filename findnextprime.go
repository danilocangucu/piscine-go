package piscine

func FindNextPrime(nb int) int {
	if nb == 2 || nb == 3 {
		return nb
	}

	if nb <= 1 {
		return 2
	} else if nb%2 == 0 || nb%3 == 0 {
		return FindNextPrime(nb + 1)
	}

	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return FindNextPrime(nb + 1)
		}
	}
	return nb
}
