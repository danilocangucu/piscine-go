package piscine

func BasicJoin(elems []string) string {
	var strjoin string
	for _, v := range elems {
		strjoin += string(v)
	}
	return strjoin
}
