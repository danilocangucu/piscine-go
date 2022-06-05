package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a {
		for _, ch := range []rune(str) {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}
