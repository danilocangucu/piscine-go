package piscine

func AppendRange(min, max int) []int {
	var slcint []int
	for i := min; i < max; i++ {
		slcint = append(slcint, i)
	}
	return slcint
}
