package piscine

func SplitWhiteSpaces(str string) []string {
	counter := 0
	for _, v := range str {
		if v == ' ' || v == '\n' || v == '\t' {
			counter++
		}
	}
	res := make([]string, counter+1)
	j := 0
	last := 0
	for i, v := range str {
		if v == ' ' || v == '\n' || v == '\t' {
			if str[last:i] != "" {
				res[j] = str[last:i]
				j++
			}
			last = i + 1
		}
	}
	res[j] = str[last:]
	c := 0
	for _, v := range res {
		if v != "" {
			c++
		}
	}
	res2 := make([]string, c)
	r := 0
	for _, v := range res {
		if v != "" {
			res2[r] = v
			r++
		}
	}
	return res2
}
