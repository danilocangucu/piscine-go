package piscine

func Swap(a *int, b *int) {
	temp := *a
	temp2 := *b
	*b = temp
	*a = temp2
}
