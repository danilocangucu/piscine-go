package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := make([]string, len(os.Args))
	copy(arguments, os.Args)

	var argrunes []rune

	for i := len(arguments) - 1; i > 0; i-- {
		argrunes = []rune(arguments[i])
		if i < len(arguments) {
			for j := 0; j < len(argrunes); j++ {
				z01.PrintRune(argrunes[j])
			}
			z01.PrintRune('\n')
		}
	}
}
