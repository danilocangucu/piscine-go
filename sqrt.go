package piscine

func Sqrt(nb int) int {
	oldnb := nb

	var sqrt int

	for odd := 0; nb > 0; odd++ {
		odd += 1
		nb = nb - odd
		sqrt++

	}

	check := sqrt * sqrt

	if check == oldnb {
		return sqrt
	}

	return 0
}
