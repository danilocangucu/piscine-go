package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var args int

	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"

	if len(os.Args) > 1 {
		args = len(os.Args[1:])
	}

	if isEven(args) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}
