package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for h := j + 1; h <= 9; h++ {
				z01.PrintRune(rune(i + 48))
				z01.PrintRune(rune(j + 48))
				z01.PrintRune(rune(h + 48))
				if h == 9 && j == 8 && i == 7 {
					break
				} else {
					z01.PrintRune(',')
					z01.PrintRune(32)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
