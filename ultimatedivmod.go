package piscine

func UltimateDivMod(a *int, b *int) {
	temp := *a / *b
	temp2 := *a % *b
	*a = temp
	*b = temp2
}
