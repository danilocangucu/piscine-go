package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 1 {
		arguments := make([][]rune, len(os.Args)-1)
		shift := ' '
		if os.Args[1] == "--upper" {
			arguments = arguments[:len(arguments)-1]
		} else {
			shift -= ' '
		}
		for i := 0; i < len(arguments); i++ {
			if os.Args[1] == "--upper" {
				arguments[i] = []rune(os.Args[i+2])
			} else {
				arguments[i] = []rune(os.Args[i+1])
			}
		}
		output := make([]rune, len(arguments))
		for i := 0; i < len(arguments); i++ {
			if len(arguments[i]) > 2 || (len(arguments[i]) == 2 && (arguments[i][0] < '1' || arguments[i][0] > '2' || (arguments[i][0] == '2' && (arguments[i][1] > '6' || arguments[i][1] < '0')) || (arguments[i][0] == '1' && (arguments[i][1] > '9' || arguments[i][1] < '0')))) || (len(arguments[i]) == 1 && (arguments[i][0] < '1' || arguments[i][0] > '9')) {
				output[i] = ' '
			} else if len(arguments[i]) == 1 {
				output[i] = arguments[i][0] + 48 - shift
			} else if len(arguments[i]) == 2 && arguments[i][0] == '1' {
				output[i] = arguments[i][1] + 58 - shift
			} else if len(arguments[i]) == 2 && arguments[i][0] == '2' {
				output[i] = arguments[i][1] + 68 - shift
			}
			z01.PrintRune(output[i])
		}
		z01.PrintRune('\n')
	}
}
