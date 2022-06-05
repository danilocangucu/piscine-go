package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sortascii := len(args)

	for i := 0; i < sortascii; i++ {
		for j := 0; j < (sortascii - 1 - i); j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}
	for _, sr := range args {
		for _, j := range sr {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
