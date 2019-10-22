package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune(0)
	} else if 1 <= n && n <= 9 {
		x := rune(n)
		z01.PrintRune(n)
	} else {
		i := 0
		ch := n
		var a []int
		for 1 < ch {
			a = append(a, ch%10)
			ch -= (ch % 10)
			ch /= 10
			i++
		}
		for j := 0; j < i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j+1]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}

		for j := 0; j < i; j++ {
			ai := rune(a[j])
			z01.PrintRune(ai)
		}

	}
}
