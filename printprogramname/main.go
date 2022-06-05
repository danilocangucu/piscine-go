package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argue := os.Args[0]
	for i, prog := range argue {
		if i > 1 {
			z01.PrintRune(rune(prog))
		}
	}
	z01.PrintRune('\n')
}
