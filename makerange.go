package piscine

func MakeRange(min, max int) []int {
	var slcint []int
	if max > min {
		slcint = make([]int, max-min)
	}
	for i := 0; i < (max - min); i++ {
		slcint[i] = min + i
	}
	return slcint
}
