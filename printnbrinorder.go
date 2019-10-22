package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune(48)
	} else if 1 <= n && n <= 9 {
		x := rune(n + 48)
		z01.PrintRune(x)
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
			for k := j+1; k < i-1; k++ {
				if a[j] > a[k] {
					temp := a[j]
					a[j] = a[k]
					a[k] = temp
				}
			}

		}

		for j := 0; j < i; j++ {
			x := rune(a[j])
			z01.PrintRune(x + 48)
		}

	}
}
