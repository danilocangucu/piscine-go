package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(NRune("String", 1))
}

func NRune(s string, n int) rune {
	if n > len(s) {
		return '0'
	} else {
		if n < 0 {
			return '0'
		} else {
			letter := []rune(s)
			return letter[n-1]
		}
	}
	return '0'
}
