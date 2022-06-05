package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}

	factorial := 1

	if nb > 0 && nb < 21 {
		for i := 1; i <= nb; i++ {
			factorial *= i
		}
	} else {
		return 0
	}

	return factorial
}
