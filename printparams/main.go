package piscine

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	count := 1
	for x := range os.Args {
		count = x
	}
	i := 1
	for i < count {
		out := []rune(os.Args[i])
		for _, j := range out {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
		i++
	}

}
