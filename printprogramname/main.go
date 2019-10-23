package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	out := []rune(os.Args[0])
	for i := range out {
		z01.PrintRune(out[i])
	}
	z01.PrintRune(10)
}
