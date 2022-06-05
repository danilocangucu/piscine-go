package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	order := [10]int{}

	if n >= 0 && n <= 9 {
		z01.PrintRune(rune(n + 48))
	} else {
		for {
			if n == 0 {
				break
			}
			order[(n%10)]++
			n = n / 10
		}

		for index, num := range order {
			if num != 0 {
				for i := 0; i < num; i++ {
					z01.PrintRune(rune(index + 48))
				}
			}
		}
	}
}
