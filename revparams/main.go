package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	count := 0
	for range os.Args {
		count++
	}
	i := count - 1
	for i > 0 {
		out := []rune(os.Args[i])
		for _, j := range out {
			z01.PrintRune(j)
		}
		z01.PrintRune(10)
		i++
	}

}
