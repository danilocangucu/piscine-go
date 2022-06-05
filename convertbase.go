package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	result1 := AtoiBase(nbr, baseFrom)

	result2 := ConvertNbrBase(result1, baseTo)

	return (result2)
}

func AtoiBase(s string, base string) int {
	counter := 0
	sCounter := 0
	for i, v := range base {
		if v == '-' || v == '+' {
			return 0
		}
		counter++
		for j, v2 := range base {
			if v == v2 && i != j {
				return 0
			}
		}
	}
	for _, v := range s {
		if v == '-' || v == '+' {
			return 0
		}
		sCounter++
	}
	if counter < 2 {
		return 0
	}
	runes := []rune(s)
	result := 0
	j := 0
	for i := sCounter - 1; i >= 0; i-- {
		r := -1
		for k, v := range base {
			if v == runes[i] {
				r = k
			}
		}
		if r == (-1) {
			return 0
		}
		if j == 0 {
			result += r
		} else {
			ss := 1
			for l := 0; l < j; l++ {
				ss *= counter
			}
			result += ss * r
		}
		j++
	}
	return result
}

func ConvertNbrBase(n int, base string) string {
	var result string
	length := len(base)

	for n >= length {
		result = string(base[(n%length)]) + result
		n = n / length
	}
	result = string(base[n]) + result

	return result
}
