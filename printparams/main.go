package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := make([]string, len(os.Args))
	copy(arguments, os.Args)

	for i := 1; i < len(arguments); i++ {
		argrunes := []rune(arguments[i])
		for j := 0; j < len(argrunes); j++ {
			z01.PrintRune(argrunes[j])
		}
		z01.PrintRune('\n')
	}
}
